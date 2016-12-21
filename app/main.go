package main

import (
	//	"bufio"
	"fmt"
	//	"log"
	//	"os"
	//	"time"

	"github.com/junneyang/xcloud/common"
)

func main() {
	defer func() {
		fmt.Println("DEFER START...")
		if r := recover(); r != nil {
			fmt.Print("DONT PANIC! ")
			fmt.Println("JUST SOME LITTLE BUG: ", r)
		}
		fmt.Println("DEFER END...")
	}()

	fmt.Println(common.SayHello("junneyang"))
	sum, err := common.Sum()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("sum: %v\n", sum)
	fmt.Println(common.CatStr_New("ABC", "-", "DEF", "-", "GHI"))

	//	log.Println("Press any key to continue...")
	//	reader := bufio.NewReader(os.Stdin)
	//	data, _, _ := reader.ReadLine()
	//	command := string(data)
	//	log.Println("Your input: ", command)
	//	log.Println("Exit after 5 seconds...")
	//	time.Sleep(time.Second * 5)
}
