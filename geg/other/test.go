package main

import (
	"fmt"

	"github.com/gogf/gf/g/os/gtime"
)

func main() {
	fmt.Println(gtime.Now().ISO8601())
}
