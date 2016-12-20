package main

import (
	"fmt"

	"github.com/junneyang/xcloud/common"
)

func main() {
	fmt.Println(common.SayHello("junneyang"))
	sum, err := common.Sum()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("sum: %v\n", sum)
}
