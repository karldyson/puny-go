package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"
	"strings"

	"golang.org/x/net/idna"
)

var (
	terseOutput          = flag.Bool("terse", false, "Terse output")
	version              = flag.Bool("version", false, "the code version")
	revision             = flag.Bool("revision", false, "revision and build information")
	versionString string = "devel"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <string>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nOptions:\n\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	// output revision info
	if *revision {
		bi, ok := debug.ReadBuildInfo()
		if !ok {
			panic("not ok reading build info!")
		}
		fmt.Printf("%s version information:\ncommit %s\n%+v\n", os.Args[0], versionString, bi)
		return
	}

	// output version info
	if *version {
		fmt.Printf("%s version %s\n", os.Args[0], versionString)
		return
	}

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
