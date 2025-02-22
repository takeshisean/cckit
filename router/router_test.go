package router_test

import (
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"

	"github.com/hyperledger/fabric/protos/peer"

	// testcc "github.com/takeshisean/cckit/testing"
	testcc "github.com/takeshisean/cckit/testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	// "github.com/takeshisean/cckit/router"
	"github.com/takeshisean/cckit/router"
)

func TestRouter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Router suite")
}

func New() *router.Chaincode {
	r := router.New(`router`).
		Init(router.EmptyContextHandler).
		Invoke(`empty`, func(c router.Context) (interface{}, error) {
			return nil, nil
		})

	return router.NewChaincode(r)
}

var cc *testcc.MockStub

var _ = Describe(`Router`, func() {

	BeforeSuite(func() {
		cc = testcc.NewMockStub(`Router`, New())
	})

	It(`Allow empty response`, func() {

		Expect(cc.Invoke(`empty`)).To(BeEquivalentTo(peer.Response{
			Status:  shim.OK,
			Payload: nil,
			Message: ``,
		}))
	})

})
