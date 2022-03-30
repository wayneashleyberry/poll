package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/gosuri/uilive"
)

func main() {
	i := flag.String("i", "1s", "polling interval")
	flag.Parse()

	duration, err := time.ParseDuration(*i)
	if err != nil {
		log.Fatal(err)
	}

	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Usage: poll [flags] [command]")

		return
	}

	var output []byte

	writer := uilive.New()
	writer.Start()

	for {
		output, err = exec.Command(args[0], args[1:]...).CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprint(writer, string(output))
		time.Sleep(duration)
	}
}
