package main

import (
	"flag"
	"fmt"
	"os"
    // "log"
	"github.com/nvbn/shell_logger/client/shell"
)

func configure() {
	clientPath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	sh, err := shell.Get()
	if err != nil {
		panic(err)
	}

	if shell.InWrapper() {
		fmt.Println(sh.SetupHooks(clientPath))
	} else {
		fmt.Println(sh.SetupWrapper(clientPath))
	}
}

func main() {
	mode := flag.String("mode", "", "configure|wrapper|submit")

	flag.Parse()

	switch *mode {
	case "configure":
		configure()
	case "wrapper":
		fmt.Println("wrapper")
	case "submit":
        var successfulCommand string = os.Getenv(shell.CommandEnv)
        var failedCommand string = os.Getenv(shell.FailedCommandEnv)
        fmt.Println("successful command: " + successfulCommand)
        fmt.Println("failed command: " + failedCommand)
/*
        err := Insert([]byte(successfulCommand), []byte(failedCommand))
        if err != nil {
            log.Fatal(err)
        }
*/
	default:
		flag.Usage()
		os.Exit(2)
	}

}

