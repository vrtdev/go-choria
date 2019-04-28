package machine

import (
	gomock "github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Aagent/Machine/Notifications", func() {
	var (
		mockctl  *gomock.Controller
		service1 *MockNotificationService
		service2 *MockNotificationService
		manager  *MockWatcherManager
		machine  *Machine
	)

	BeforeEach(func() {
		mockctl = gomock.NewController(GinkgoT())
		service1 = NewMockNotificationService(mockctl)
		service2 = NewMockNotificationService(mockctl)
		manager = NewMockWatcherManager(mockctl)
		machine = &Machine{
			notifiers:   []NotificationService{},
			manager:     manager,
			MachineName: "ginkgo",
		}
	})

	AfterEach(func() {
		mockctl.Finish()
	})

	Describe("RegisterNotifier", func() {
		It("Should add the notifier to the list", func() {
			Expect(machine.notifiers).To(BeEmpty())
			machine.RegisterNotifier(service1)
			Expect(machine.notifiers[0]).To(Equal(service1))
			Expect(machine.notifiers).To(HaveLen(1))
		})
	})

	Describe("Notifications", func() {
		BeforeEach(func() {
			machine.RegisterNotifier(service1, service2)
		})

		It("Should support notifying state", func() {
			state := map[string]interface{}{}

			service1.EXPECT().NotifyWatcherState("ginkgo", "w1", state)
			service2.EXPECT().NotifyWatcherState("ginkgo", "w1", state)

			machine.NotifyWatcherState("w1", state)
		})

		It("Should support common loggers", func() {
			service1.EXPECT().Debugf("ginkgo", "w1", "format", "debugarg")
			service2.EXPECT().Debugf("ginkgo", "w1", "format", "debugarg")
			service1.EXPECT().Infof("ginkgo", "w1", "format", "infoarg")
			service2.EXPECT().Infof("ginkgo", "w1", "format", "infoarg")
			service1.EXPECT().Warnf("ginkgo", "w1", "format", "warnarg")
			service2.EXPECT().Warnf("ginkgo", "w1", "format", "warnarg")
			service1.EXPECT().Errorf("ginkgo", "w1", "format", "errorarg")
			service2.EXPECT().Errorf("ginkgo", "w1", "format", "errorarg")

			machine.Debugf("w1", "format", "debugarg")
			machine.Infof("w1", "format", "infoarg")
			machine.Warnf("w1", "format", "warnarg")
			machine.Errorf("w1", "format", "errorarg")
		})
	})
})
