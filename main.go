package main

import (
	"Practice_Blockchain/cli"
	"Practice_Blockchain/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
