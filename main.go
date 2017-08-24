package main

import (
	"flag"
	"log"
	"os/exec"
	"time"

	tm "github.com/buger/goterm"
)

func main() {
	flag.Parse()

	args := flag.Args()

	tm.Clear()

	var output []byte
	var err error

	tm.MoveCursor(1, 1)
	tm.Print("")
	tm.Flush()

	for {
		tm.MoveCursor(1, 1)

		output, err = exec.Command(args[0], args[1:]...).CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}

		tm.Print(time.Now().Format(time.RFC1123) + "\n" + string(output))

		tm.Flush()

		time.Sleep(time.Second * 10)
	}
}
