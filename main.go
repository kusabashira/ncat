package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

var (
	name  = "catn"
	usage = fmt.Sprintf("usage: %s N [FILE]...", name)
)

func printErr(err error) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", name, err)
}

func toARGF(args []string) (r io.Reader, err error) {
	if len(args) < 1 {
		return os.Stdin, nil
	} else {
		rs := make([]io.Reader, len(args))
		for i := 0; i < len(args); i++ {
			f, err := os.Open(args[i])
			if err != nil {
				return nil, err
			}
			rs[i] = f
		}
		return io.MultiReader(rs...), nil
	}
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

	r, err := toARGF(os.Args[2:])
	if err != nil {
		printErr(err)
		os.Exit(2)
	}

	src, err := ioutil.ReadAll(r)
	if err != nil {
		printErr(err)
		os.Exit(1)
	}

	if n < 0 {
		for {
			os.Stdout.Write(src)
		}
	} else {
		for i := 0; i < n; i++ {
			os.Stdout.Write(src)
		}
	}
}
