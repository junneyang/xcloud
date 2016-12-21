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
	fmt.Println(common.SayHello("junneyang"))
	sum, err := common.Sum()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("sum: %v\n", sum)

	//	log.Println("Press any key to continue...")
	//	reader := bufio.NewReader(os.Stdin)
	//	data, _, _ := reader.ReadLine()
	//	command := string(data)
	//	log.Println("Your input: ", command)
	//	log.Println("Exit after 5 seconds...")
	//	time.Sleep(time.Second * 5)
}
