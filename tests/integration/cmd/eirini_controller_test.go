package cmd_test

import (
	"os"

	eirinictrl "code.cloudfoundry.org/eirini-controller"
	"code.cloudfoundry.org/eirini-controller/tests/integration"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("EiriniController", func() {
	var (
		config         *eirinictrl.ControllerConfig
		configFilePath string
		session        *gexec.Session
	)

	BeforeEach(func() {
		config = integration.DefaultControllerConfig(fixture.Namespace)
	})

	JustBeforeEach(func() {
		session, configFilePath = eiriniBins.EiriniController.Run(config)
	})

	AfterEach(func() {
		if configFilePath != "" {
			Expect(os.Remove(configFilePath)).To(Succeed())
		}
		if session != nil {
			Eventually(session.Kill()).Should(gexec.Exit())
		}
	})

	It("should be able to start properly", func() {
		Consistently(session).ShouldNot(gexec.Exit())
	})
})
