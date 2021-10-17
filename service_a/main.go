package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/smith-30/monorepo/service_a/model"
)

var t = "a"

func main() {
	u, _ := model.NewUser()
	fmt.Printf("%#v\n", u)

	spew.Dump(u)

	p := []string{}
	fmt.Printf("%#v\n", p)
}
