package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	// "github.com/s7techlab/cckit/examples/erc20"
	"github.com/takeshisean/cckit/examples/erc20"
)

func main() {
	err := shim.Start(erc20.NewErc20FixedSupply())
	if err != nil {
		fmt.Printf("Error starting ERC-20 chaincode: %s", err)
	}
}
