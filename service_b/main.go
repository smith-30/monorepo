package main

import (
	"fmt"

	"github.com/smith-30/monorepo/service_b/model"
)

func main() {
	u, _ := model.NewUser()
	fmt.Printf("%#v\n", u)
}
