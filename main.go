package main

import (
	"github.com/YouSangSon/Practice_Blockchain/explorer"
	"github.com/YouSangSon/Practice_Blockchain/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
