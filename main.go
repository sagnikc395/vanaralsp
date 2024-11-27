package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sagnikc395/vanaralsp/rpc"
)

func main() {
	fmt.Println("LSP START >>>>> ")

	scanner := bufio.NewScanner(os.Stdin)

	//using our custom split method
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(msg)
	}
}

func handleMessage(_ any) {

}
