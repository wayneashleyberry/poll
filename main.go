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

	if len(args) == 0 {
		return
	}

	tm.Clear()

	var output []byte
	var err error

	tm.MoveCursor(0, 0)
	tm.Print("")
	tm.Flush()

	for {
		output, err = exec.Command(args[0], args[1:]...).CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}

		// clear screen
		tm.MoveCursor(0, 0)
		tm.Print("")
		tm.Flush()

		// print output
		tm.Print(string(output))
		tm.MoveCursor(0, 0)
		tm.Flush()

		time.Sleep(time.Second)
	}
}
