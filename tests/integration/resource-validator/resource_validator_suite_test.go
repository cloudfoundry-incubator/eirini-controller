package resource_validator_test

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"testing"

	eirinictrl "code.cloudfoundry.org/eirini-controller"
	"code.cloudfoundry.org/eirini-controller/tests"
	"code.cloudfoundry.org/eirini-controller/tests/integration"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	arv1 "k8s.io/api/admissionregistration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func TestResourceValidator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ResourceValidator Suite")
}

var (
	fixture        *tests.Fixture
	eiriniBins     integration.EiriniBinaries
	config         *eirinictrl.ResourceValidatorConfig
	configFilePath string
	fingerprint    string
	certDir        string
	hookSession    *gexec.Session
)

var _ = SynchronizedBeforeSuite(func() []byte {
	eiriniBins = integration.NewEiriniBinaries()

	fixture = tests.NewFixture(GinkgoWriter)

	port := int32(tests.GetTelepresencePort())
	lrpsPath := "/lrps"
	telepresenceService := tests.GetTelepresenceServiceName()
	telepresenceDomain := fmt.Sprintf("%s.default.svc", telepresenceService)
	fingerprint = "lrp-" + tests.GenerateGUID()[:8]
	var caBundle []byte
	certDir, caBundle = tests.GenerateKeyPairDir("tls", telepresenceDomain)
	sideEffects := arv1.SideEffectClassNone
	scope := arv1.NamespacedScope

	_, err := fixture.Clientset.AdmissionregistrationV1().ValidatingWebhookConfigurations().Create(context.Background(),
		&arv1.ValidatingWebhookConfiguration{
			ObjectMeta: metav1.ObjectMeta{
				Name: fingerprint + "-validating-hook",
			},
			Webhooks: []arv1.ValidatingWebhook{
				{
					Name: fingerprint + "-validating-webhook.cloudfoundry.org",
					ClientConfig: arv1.WebhookClientConfig{
						Service: &arv1.ServiceReference{
							Namespace: "default",
							Name:      telepresenceService,
							Port:      &port,
							Path:      &lrpsPath,
						},
						CABundle: caBundle,
					},
					SideEffects:             &sideEffects,
					AdmissionReviewVersions: []string{"v1beta1"},
					Rules: []arv1.RuleWithOperations{
						{
							Operations: []arv1.OperationType{arv1.Update},
							Rule: arv1.Rule{
								APIGroups:   []string{"eirini.cloudfoundry.org"},
								APIVersions: []string{"v1"},
								Resources:   []string{"lrps"},
								Scope:       &scope,
							},
						},
					},
				},
			},
		}, metav1.CreateOptions{})
	Expect(err).NotTo(HaveOccurred())

	config = &eirinictrl.ResourceValidatorConfig{
		Port: port,
		KubeConfig: eirinictrl.KubeConfig{
			ConfigPath: fixture.KubeConfigPath,
		},
	}
	env := fmt.Sprintf("%s=%s", eirinictrl.EnvResourceValidatorCertDir, certDir)
	hookSession, configFilePath = eiriniBins.ResourceValidator.Run(config, env)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	Eventually(func() error {
		fmt.Fprintf(GinkgoWriter, "CHECKING: %s\n", telepresenceDomain)
		resp, err := client.Get(fmt.Sprintf("https://%s:%d/lrps", telepresenceDomain, port))
		if err != nil {
			fmt.Fprintf(GinkgoWriter, "failed to call telepresence: %s\n", err.Error())

			return err
		}
		resp.Body.Close()

		fmt.Fprintf(GinkgoWriter, "SUCCESS: %#v\n", resp)

		return nil
	}, "2m", "500ms").Should(Succeed())

	return []byte{}
}, func(data []byte) {
	fixture = tests.NewFixture(GinkgoWriter)
})

var _ = SynchronizedAfterSuite(func() {
	fixture.Destroy()
}, func() {
	if configFilePath != "" {
		Expect(os.Remove(configFilePath)).To(Succeed())
	}
	if hookSession != nil {
		Eventually(hookSession.Kill()).Should(gexec.Exit())
	}
	err := fixture.Clientset.AdmissionregistrationV1().ValidatingWebhookConfigurations().Delete(context.Background(), fingerprint+"-validating-hook", metav1.DeleteOptions{})
	Expect(err).NotTo(HaveOccurred())

	Expect(os.RemoveAll(certDir)).To(Succeed())

	eiriniBins.TearDown()

	fixture.Destroy()
})

var _ = BeforeEach(func() {
	fixture.SetUp()
})

var _ = AfterEach(func() {
	fixture.TearDown()
})
