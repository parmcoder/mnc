package main

import (
	"github.com/parmcoder/transaction-tracker/core"
)

func main() {
	e := core.Setup()

	e.Logger.Fatal(e.Start(":1323"))
}
