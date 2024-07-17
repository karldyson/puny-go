package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/idna"
)

var (
	terseOutput = flag.Bool("terse", false, "Terse output")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <string>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nOptions:\n\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	switch len(flag.Args()) {
	case 1:
		inputString := flag.Args()[0]
		if strings.HasPrefix(inputString, "xn--") {
			u, _ := idna.ToUnicode(inputString)
			if *terseOutput {
				fmt.Println(u)
			} else {
				fmt.Printf("ASCII..: %s\n", inputString)
				fmt.Printf("Unicode: %s\n", u)
			}
		} else {
			a, _ := idna.ToASCII(inputString)
			if *terseOutput {
				fmt.Println(a)
			} else {
				fmt.Printf("Unicode: %s\n", inputString)
				fmt.Printf("ASCII..: %s\n", a)
			}
		}
	default:
		fmt.Fprintf(os.Stderr, "ERROR: invalid number of CLI parameters\n")
		flag.Usage()
	}
}
