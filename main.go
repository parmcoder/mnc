package main

import (
	"fmt"
	"github.com/parmcoder/transaction-tracker/core"
)

func main() {
	fmt.Println("hello world!")

	e := core.Setup()

	e.Logger.Fatal(e.Start(":1323"))
}
