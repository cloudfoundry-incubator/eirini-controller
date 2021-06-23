package main

import (
	"fmt"
	"os"

	eirinictrl "code.cloudfoundry.org/eirini-controller"
	cmdcommons "code.cloudfoundry.org/eirini-controller/cmd"
	"code.cloudfoundry.org/eirini-controller/k8s"
	"code.cloudfoundry.org/eirini-controller/k8s/client"
	"code.cloudfoundry.org/eirini-controller/k8s/crclient"
	eirinievent "code.cloudfoundry.org/eirini-controller/k8s/informers/event"
	"code.cloudfoundry.org/eirini-controller/k8s/jobs"
	"code.cloudfoundry.org/eirini-controller/k8s/pdb"
	"code.cloudfoundry.org/eirini-controller/k8s/reconciler"
	"code.cloudfoundry.org/eirini-controller/k8s/stset"
	eiriniv1 "code.cloudfoundry.org/eirini-controller/pkg/apis/eirini/v1"
	eirinischeme "code.cloudfoundry.org/eirini-controller/pkg/generated/clientset/versioned/scheme"
	"code.cloudfoundry.org/eirini-controller/prometheus"
	"code.cloudfoundry.org/eirini-controller/util"
	"code.cloudfoundry.org/lager"
	"github.com/jessevdk/go-flags"
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/clock"
	"k8s.io/client-go/kubernetes"
	kscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

type options struct {
	ConfigFile string `short:"c" long:"config" description:"Config for running eirini-controller"`
}

func main() {
	if err := kscheme.AddToScheme(eirinischeme.Scheme); err != nil {
		cmdcommons.Exitf("failed to add the k8s scheme to the LRP CRD scheme: %v", err)
	}

	var opts options
	_, err := flags.ParseArgs(&opts, os.Args)
	cmdcommons.ExitfIfError(err, "Failed to parse args")

	var cfg eirinictrl.ControllerConfig
	err = cmdcommons.ReadConfigFile(opts.ConfigFile, &cfg)
	cmdcommons.ExitfIfError(err, "Failed to read config file")

	kubeConfig, err := clientcmd.BuildConfigFromFlags("", cfg.ConfigPath)
	cmdcommons.ExitfIfError(err, "Failed to build kubeconfig")

	controllerClient, err := ctrlruntimeclient.New(kubeConfig, ctrlruntimeclient.Options{Scheme: eirinischeme.Scheme})
	cmdcommons.ExitfIfError(err, "Failed to create k8s runtime client")

	clientset, err := kubernetes.NewForConfig(kubeConfig)
	cmdcommons.ExitfIfError(err, "Failed to create k8s clientset")

	logger := lager.NewLogger("eirini-controller")
	logger.RegisterSink(lager.NewPrettySink(os.Stdout, lager.DEBUG))

	managerOptions := manager.Options{
		MetricsBindAddress: "0",
		Scheme:             eirinischeme.Scheme,
		Namespace:          cfg.WorkloadsNamespace,
		Logger:             util.NewLagerLogr(logger),
		LeaderElection:     true,
		LeaderElectionID:   "eirini-controller-leader",
	}

	if cfg.PrometheusPort > 0 {
		managerOptions.MetricsBindAddress = fmt.Sprintf(":%d", cfg.PrometheusPort)
	}

	if cfg.LeaderElectionID != "" {
		managerOptions.LeaderElectionNamespace = cfg.LeaderElectionNamespace
		managerOptions.LeaderElectionID = cfg.LeaderElectionID
	}

	mgr, err := manager.New(kubeConfig, managerOptions)
	cmdcommons.ExitfIfError(err, "Failed to create k8s controller runtime manager")

	lrpReconciler, err := createLRPReconciler(logger, controllerClient, clientset, cfg, mgr.GetScheme())
	cmdcommons.ExitfIfError(err, "Failed to create LRP reconciler")

	taskReconciler := createTaskReconciler(logger, controllerClient, clientset, cfg, mgr.GetScheme())
	podCrashReconciler := createPodCrashReconciler(logger, cfg.WorkloadsNamespace, controllerClient, clientset)

	err = builder.
		ControllerManagedBy(mgr).
		For(&eiriniv1.LRP{}).
		Owns(&appsv1.StatefulSet{}).
		Complete(lrpReconciler)
	cmdcommons.ExitfIfError(err, "Failed to build LRP reconciler")

	err = builder.
		ControllerManagedBy(mgr).
		For(&eiriniv1.Task{}).
		Owns(&batchv1.Job{}).
		Complete(taskReconciler)
	cmdcommons.ExitfIfError(err, "Failed to build Task reconciler")

	predicates := []predicate.Predicate{reconciler.NewSourceTypeUpdatePredicate(stset.AppSourceType)}
	err = builder.
		ControllerManagedBy(mgr).
		For(&corev1.Pod{}, builder.WithPredicates(predicates...)).
		Complete(podCrashReconciler)
	cmdcommons.ExitfIfError(err, "Failed to build Pod Crash reconciler")

	err = mgr.Start(ctrl.SetupSignalHandler())
	cmdcommons.ExitfIfError(err, "Failed to start manager")
}

func createLRPReconciler(
	logger lager.Logger,
	controllerClient ctrlruntimeclient.Client,
	clientset kubernetes.Interface,
	cfg eirinictrl.ControllerConfig,
	scheme *runtime.Scheme,
) (*reconciler.LRP, error) {
	logger = logger.Session("lrp-reconciler")
	lrpToStatefulSetConverter := stset.NewLRPToStatefulSetConverter(
		cfg.ApplicationServiceAccount,
		cfg.RegistrySecretName,
		cfg.UnsafeAllowAutomountServiceAccountToken,
		cfg.AllowRunImageAsRoot,
		k8s.CreateLivenessProbe,
		k8s.CreateReadinessProbe,
	)
	workloadClient := k8s.NewLRPClient(
		logger.Session("stateful-set-desirer"),
		client.NewSecret(clientset),
		client.NewStatefulSet(clientset, cfg.WorkloadsNamespace),
		client.NewPod(clientset, cfg.WorkloadsNamespace),
		pdb.NewUpdater(client.NewPodDisruptionBudget(clientset)),
		client.NewEvent(clientset),
		lrpToStatefulSetConverter,
		eirinischeme.Scheme,
	)

	decoratedWorkloadClient, err := prometheus.NewLRPClientDecorator(logger.Session("prometheus-decorator"), workloadClient, metrics.Registry, clock.RealClock{})
	if err != nil {
		return nil, err
	}

	lrpsCrClient := crclient.NewLRPs(controllerClient)

	return reconciler.NewLRP(
		logger,
		lrpsCrClient,
		decoratedWorkloadClient,
		client.NewStatefulSet(clientset, cfg.WorkloadsNamespace),
	), nil
}

func createTaskReconciler(
	logger lager.Logger,
	controllerClient ctrlruntimeclient.Client,
	clientset kubernetes.Interface,
	cfg eirinictrl.ControllerConfig,
	scheme *runtime.Scheme,
) *reconciler.Task {
	taskToJobConverter := jobs.NewTaskToJobConverter(
		cfg.ApplicationServiceAccount,
		cfg.RegistrySecretName,
		cfg.UnsafeAllowAutomountServiceAccountToken,
	)
	workloadClient := k8s.NewTaskClient(
		logger,
		client.NewJob(clientset, cfg.WorkloadsNamespace),
		client.NewSecret(clientset),
		taskToJobConverter,
		scheme,
	)
	tasksCrClient := crclient.NewTasks(controllerClient)

	return reconciler.NewTask(logger, tasksCrClient, workloadClient, cfg.TaskTTLSeconds)
}

func createPodCrashReconciler(
	logger lager.Logger,
	workloadsNamespace string,
	controllerClient ctrlruntimeclient.Client,
	clientset kubernetes.Interface) *reconciler.PodCrash {
	eventsClient := client.NewEvent(clientset)
	statefulSetClient := client.NewStatefulSet(clientset, workloadsNamespace)
	crashEventGenerator := eirinievent.NewDefaultCrashEventGenerator(eventsClient)

	return reconciler.NewPodCrash(logger, controllerClient, crashEventGenerator, eventsClient, statefulSetClient)
}
