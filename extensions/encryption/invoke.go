package encryption

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"
	"github.com/takeshisean/cckit/convert"
	"github.com/takeshisean/cckit/state"
)

// InvokeChaincode decrypts received payload
func InvokeChaincode(
	stub shim.ChaincodeStubInterface, encKey []byte, chaincodeName string,
	args []interface{}, channel string, target interface{}) (interface{}, error) {

	// args are not encrypted cause we cannot pass encryption key in transient map while invoking cc from cc
	// thus target cc cannot decrypt args
	aa, err := convert.ArgsToBytes(args...)
	if err != nil {
		return nil, errors.Wrap(err, `encrypt args`)
	}

	response := stub.InvokeChaincode(chaincodeName, aa, channel)
	if response.Status != shim.OK {
		return nil, errors.New(response.Message)
	}

	if len(response.Payload) == 0 {
		return nil, state.ErrEmptyChaincodeResponsePayload
	}

	decrypted, err := Decrypt(encKey, response.Payload)
	if err != nil {
		return nil, errors.Wrap(err, `decrypt payload`)
	}
	return convert.FromBytes(decrypted, target)
}
