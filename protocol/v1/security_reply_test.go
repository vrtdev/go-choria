package v1

import (
	"crypto/sha256"
	"encoding/base64"

	"github.com/choria-io/go-protocol/protocol"
	"github.com/tidwall/gjson"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SecureReply", func() {
	It("Should create valid replies", func() {
		request, _ := NewRequest("test", "go.tests", "rip.mcollective", 120, "a2f0ca717c694f2086cfa81b6c494648", "mcollective")
		request.SetMessage(`{"test":1}`)

		reply, err := NewReply(request, "testing")
		Expect(err).ToNot(HaveOccurred())

		rj, err := reply.JSON()
		Expect(err).ToNot(HaveOccurred())

		sha := sha256.Sum256([]byte(rj))

		sreply, _ := NewSecureReply(reply)
		sj, err := sreply.JSON()
		Expect(err).ToNot(HaveOccurred())

		Expect(gjson.Get(sj, "protocol").String()).To(Equal(protocol.SecureReplyV1))
		Expect(gjson.Get(sj, "message").String()).To(Equal(rj))
		Expect(gjson.Get(sj, "hash").String()).To(Equal(base64.StdEncoding.EncodeToString(sha[:])))
		Expect(sreply.Valid()).To(BeTrue())
	})
})
