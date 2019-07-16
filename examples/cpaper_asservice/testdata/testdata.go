package testdata

import (
	"io/ioutil"

	"github.com/takeshisean/cckit/testing"
)

func GetTestIdentity(msp, file string) *testing.Identity {
	identity, err := testing.IdentityFromFile(msp, file, ioutil.ReadFile)
	if err != nil {
		panic(err)
	}

	return identity
}
