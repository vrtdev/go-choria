package filewatcher

import (
	"encoding/json"
	"path/filepath"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/choria-io/go-choria/aagent/watchers/watcher"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AAgent/Watchers/ExecWatcher")
}

var _ = Describe("ExecWatcher", func() {
	var (
		mockctl     *gomock.Controller
		mockMachine *watcher.MockMachine
		watch       *Watcher
		now         time.Time
	)

	BeforeEach(func() {
		mockctl = gomock.NewController(GinkgoT())
		mockMachine = watcher.NewMockMachine(mockctl)

		mockMachine.EXPECT().Name().Return("file").AnyTimes()
		mockMachine.EXPECT().Identity().Return("ginkgo").AnyTimes()
		mockMachine.EXPECT().InstanceID().Return("1234567890").AnyTimes()
		mockMachine.EXPECT().Version().Return("1.0.0").AnyTimes()
		mockMachine.EXPECT().TimeStampSeconds().Return(now.Unix()).AnyTimes()
		mockMachine.EXPECT().Directory().Return(".").AnyTimes()

		now = time.Unix(1606924953, 0)

		wi, err := New(mockMachine, "ginkgo", []string{"always"}, "fail", "success", "2m", time.Second, map[string]interface{}{
			"path": filepath.Join("bin", "sh"),
		})
		Expect(err).ToNot(HaveOccurred())
		watch = wi.(*Watcher)
	})

	AfterEach(func() {
		mockctl.Finish()
	})

	Describe("setProperties", func() {
		It("Should parse valid properties", func() {
			prop := map[string]interface{}{
				"path":                 "cmd",
				"gather_initial_state": "t",
			}
			Expect(watch.setProperties(prop)).ToNot(HaveOccurred())
			Expect(watch.properties.Path).To(Equal("cmd"))
			Expect(watch.properties.Initial).To(BeTrue())
		})

		It("Should handle errors", func() {
			watch.properties = &Properties{}
			err := watch.setProperties(map[string]interface{}{})
			Expect(err).To(MatchError("path is required"))
		})
	})

	Describe("CurrentState", func() {
		It("Should be a valid state", func() {
			watch.previous = Changed
			cs := watch.CurrentState()
			csj, err := cs.(*StateNotification).JSON()
			Expect(err).ToNot(HaveOccurred())

			event := map[string]interface{}{}
			err = json.Unmarshal(csj, &event)
			Expect(err).ToNot(HaveOccurred())
			delete(event, "id")

			Expect(event).To(Equal(map[string]interface{}{
				"time":            "2020-12-02T16:02:33Z",
				"type":            "io.choria.machine.watcher.file.v1.state",
				"subject":         "ginkgo",
				"specversion":     "1.0",
				"source":          "io.choria.machine",
				"datacontenttype": "application/json",
				"data": map[string]interface{}{
					"id":               "1234567890",
					"identity":         "ginkgo",
					"machine":          "file",
					"name":             "ginkgo",
					"protocol":         "io.choria.machine.watcher.file.v1.state",
					"type":             "file",
					"version":          "1.0.0",
					"timestamp":        float64(now.Unix()),
					"previous_outcome": "changed",
					"path":             filepath.Join("bin", "sh"),
				},
			}))
		})
	})
})
