package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/sagnikc395/vanaralsp/rpc"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	logger := getLogger(fmt.Sprintf("%s/log.txt", dir))

	logger.Println("Starting vanaralsp logger ...")
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

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("file format is not writable for logging")
	}

	return log.New(logfile, "[vanaralsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
