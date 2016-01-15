package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/yuya-takeyama/argf"
)

var (
	name  = "ncat"
	usage = fmt.Sprintf("usage: %s N [FILE]...", name)
)

func printErr(err error) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", name, err)
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "--help" {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(2)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		printErr(err)
		os.Exit(2)
	}

	r, err := argf.From(os.Args[2:])
	if err != nil {
		printErr(err)
		os.Exit(2)
	}

	src, err := ioutil.ReadAll(r)
	if err != nil {
		printErr(err)
		os.Exit(1)
	}

	switch {
	case n < 0:
		for {
			os.Stdout.Write(src)
		}
	default:
		for i := 0; i < n; i++ {
			os.Stdout.Write(src)
		}
	}
}
