package cmd

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"code.cloudfoundry.org/eirini"
	"code.cloudfoundry.org/eirini/bifrost"
	cmdcommons "code.cloudfoundry.org/eirini/cmd"
	"code.cloudfoundry.org/eirini/events"
	"code.cloudfoundry.org/eirini/handler"
	"code.cloudfoundry.org/eirini/k8s"
	k8sevent "code.cloudfoundry.org/eirini/k8s/informers/event"
	"code.cloudfoundry.org/eirini/k8s/kubelet"
	"code.cloudfoundry.org/eirini/metrics"
	"code.cloudfoundry.org/eirini/stager"
	"code.cloudfoundry.org/eirini/util"
	loggregator "code.cloudfoundry.org/go-loggregator"
	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/tps/cc_client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
	"k8s.io/client-go/kubernetes"
	metricsclientset "k8s.io/metrics/pkg/client/clientset/versioned"

	// For gcp and oidc authentication
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connects CloudFoundry with Kubernetes",
	Run:   connect,
}

func connect(cmd *cobra.Command, args []string) {
	path, err := cmd.Flags().GetString("config")
	cmdcommons.ExitWithError(err)
	if path == "" {
		cmdcommons.ExitWithError(errors.New("--config is missing"))
	}

	cfg := setConfigFromFile(path)

	stager := initStager(cfg)
	bifrost := initBifrost(cfg)
	clientset := cmdcommons.CreateKubeClient(cfg.Properties.ConfigPath)
	metricsClient := cmdcommons.CreateMetricsClient(cfg.Properties.ConfigPath)

	tlsConfig, err := loggregator.NewIngressTLSConfig(
		cfg.Properties.LoggregatorCAPath,
		cfg.Properties.LoggregatorCertPath,
		cfg.Properties.LoggregatorKeyPath,
	)
	cmdcommons.ExitWithError(err)

	loggregatorClient, err := loggregator.NewIngressClient(
		tlsConfig,
		loggregator.WithAddr(cfg.Properties.LoggregatorAddress),
	)
	cmdcommons.ExitWithError(err)
	defer func() {
		if err = loggregatorClient.CloseSend(); err != nil {
			cmdcommons.ExitWithError(err)
		}
	}()
	launchMetricsEmitter(
		clientset,
		metricsClient,
		loggregatorClient,
		cfg.Properties.Namespace,
		cfg.Properties.AppMetricsEmissionIntervalInSecs,
	)

	launchEventReporter(
		clientset,
		cfg.Properties.CcInternalAPI,
		cfg.Properties.CCCAPath,
		cfg.Properties.CCCertPath,
		cfg.Properties.CCKeyPath,
		cfg.Properties.Namespace,
	)

	handlerLogger := lager.NewLogger("handler")
	handlerLogger.RegisterSink(lager.NewPrettySink(os.Stdout, lager.DEBUG))
	handler := handler.New(bifrost, stager, handlerLogger)

	var server *http.Server
	handlerLogger.Info("opi-connected")
	server = &http.Server{
		Addr:      fmt.Sprintf("0.0.0.0:%d", cfg.Properties.TLSPort),
		Handler:   handler,
		TLSConfig: serverTLSConfig(cfg),
	}
	handlerLogger.Fatal("opi-crashed",
		server.ListenAndServeTLS(cfg.Properties.ServerCertPath, cfg.Properties.ServerKeyPath))
}

func serverTLSConfig(cfg *eirini.Config) *tls.Config {
	bs, err := ioutil.ReadFile(cfg.Properties.ClientCAPath)
	cmdcommons.ExitWithError(err)

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(bs) {
		panic("invalid client CA cert data")
	}

	return &tls.Config{
		ClientCAs:  certPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
}

func initStager(cfg *eirini.Config) eirini.Stager {
	clientset := cmdcommons.CreateKubeClient(cfg.Properties.ConfigPath)
	taskDesirer := &k8s.TaskDesirer{
		Namespace:       cfg.Properties.Namespace,
		CertsSecretName: cfg.Properties.CCCertsSecretName,
		Client:          clientset,
	}

	stagerCfg := eirini.StagerConfig{
		EiriniAddress:   cfg.Properties.EiriniAddress,
		DownloaderImage: cfg.Properties.DownloaderImage,
		UploaderImage:   cfg.Properties.UploaderImage,
		ExecutorImage:   cfg.Properties.ExecutorImage,
	}

	httpClient, err := util.CreateTLSHTTPClient(
		[]util.CertPaths{
			{
				Crt: cfg.Properties.CCCertPath,
				Key: cfg.Properties.CCKeyPath,
				Ca:  cfg.Properties.CCCAPath,
			},
		},
	)
	if err != nil {
		panic(errors.Wrap(err, "failed to create stager http client"))
	}

	stagerLogger := lager.NewLogger("stager")
	stagerLogger.RegisterSink(lager.NewPrettySink(os.Stdout, lager.DEBUG))
	return stager.New(taskDesirer, httpClient, stagerCfg, stagerLogger)
}

func initBifrost(cfg *eirini.Config) eirini.Bifrost {
	kubeNamespace := cfg.Properties.Namespace
	clientset := cmdcommons.CreateKubeClient(cfg.Properties.ConfigPath)
	desireLogger := lager.NewLogger("desirer")
	desireLogger.RegisterSink(lager.NewPrettySink(os.Stdout, lager.DEBUG))
	desirer := k8s.NewStatefulSetDesirer(clientset, kubeNamespace, cfg.Properties.RegistrySecretName, cfg.Properties.RootfsVersion, desireLogger)
	convertLogger := lager.NewLogger("convert")
	convertLogger.RegisterSink(lager.NewPrettySink(os.Stdout, lager.DEBUG))
	registryIP := cfg.Properties.RegistryAddress
	converter := bifrost.NewConverter(convertLogger, registryIP, cfg.Properties.DiskLimitMB)

	return &bifrost.Bifrost{
		Converter: converter,
		Desirer:   desirer,
	}
}

func setConfigFromFile(path string) *eirini.Config {
	fileBytes, err := ioutil.ReadFile(filepath.Clean(path))
	cmdcommons.ExitWithError(err)

	var conf eirini.Config
	conf.Properties.DiskLimitMB = 2048
	err = yaml.Unmarshal(fileBytes, &conf)
	cmdcommons.ExitWithError(err)

	return &conf
}

func initConnect() {
	connectCmd.Flags().StringP("config", "c", "", "Path to the Eirini config file")
}

func launchMetricsEmitter(
	clientset kubernetes.Interface,
	metricsClient metricsclientset.Interface,
	loggregatorClient metrics.LoggregatorClient,
	namespace string,
	metricsEmissionInterval int,
) {
	work := make(chan []metrics.Message, 20)
	podClient := clientset.CoreV1().Pods(namespace)

	podMetricsClient := metricsClient.MetricsV1beta1().PodMetricses(namespace)
	metricsLogger := lager.NewLogger("metrics")
	metricsLogger.RegisterSink(lager.NewPrettySink(os.Stdout, lager.DEBUG))

	tickerInterval := eirini.AppMetricsEmissionIntervalInSecs
	if metricsEmissionInterval > 0 {
		tickerInterval = metricsEmissionInterval
	}
	collectorScheduler := &util.TickerTaskScheduler{
		Ticker: time.NewTicker(time.Duration(tickerInterval) * time.Second),
		Logger: metricsLogger.Session("collector.scheduler"),
	}

	metricsCollectorLogger := metricsLogger.Session("metrics-collector", lager.Data{})
	diskClientLogger := metricsCollectorLogger.Session("disk-metrics-client", lager.Data{})
	kubeletClient := kubelet.NewClient(clientset.CoreV1().RESTClient())
	diskClient := kubelet.NewDiskMetricsClient(clientset.CoreV1().Nodes(),
		kubeletClient,
		namespace,
		diskClientLogger)
	collector := k8s.NewMetricsCollector(podMetricsClient, podClient, diskClient, metricsCollectorLogger)

	forwarder := metrics.NewLoggregatorForwarder(loggregatorClient)
	emitterScheduler := &util.SimpleLoopScheduler{
		CancelChan: make(chan struct{}, 1),
		Logger:     metricsLogger.Session("emitter.scheduler"),
	}
	emitter := metrics.NewEmitter(work, emitterScheduler, forwarder)

	go collectorScheduler.Schedule(func() error {
		return k8s.ForwardMetricsToChannel(collector, work)
	})
	go emitter.Start()
}

func launchEventReporter(clientset kubernetes.Interface, uri, ca, cert, key, namespace string) {
	work := make(chan events.CrashReport, 1)
	tlsConf, err := cc_client.NewTLSConfig(cert, key, ca)
	cmdcommons.ExitWithError(err)

	client := cc_client.NewCcClient(uri, tlsConf)
	crashReporterLogger := lager.NewLogger("instance-crash-reporter")
	crashReporterLogger.RegisterSink(lager.NewPrettySink(os.Stdout, lager.DEBUG))

	scheduler := &util.SimpleLoopScheduler{
		CancelChan: make(chan struct{}, 1),
		Logger:     crashReporterLogger.Session("scheduler"),
	}
	reporter := events.NewCrashReporter(work, scheduler, client, crashReporterLogger)

	crashLogger := lager.NewLogger("instance-crash-informer")
	crashLogger.RegisterSink(lager.NewPrettySink(os.Stdout, lager.DEBUG))
	crashInformer := k8sevent.NewCrashInformer(clientset, 0, namespace, work, make(chan struct{}), crashLogger, k8sevent.DefaultCrashReportGenerator{})

	go crashInformer.Start()
	go reporter.Run()
}
