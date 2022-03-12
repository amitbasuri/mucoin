package main

import (
	"mucoin"
)

func main() {

	chain := mucoin.InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

}
