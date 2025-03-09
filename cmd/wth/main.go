// What The Hash!? is a simple hash reverse lookup.
//
// It searches a database of 270+ hash algorithms for the possible source of the given hash sum
// and outputs all found matches in hashcat notation to STDOUT.
//
// Usage:
//
//	wth hashsum
//
// The arguments are:
//
//	hashsum
//		Hash sum to find all possible sources for (required).
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/cuhsat/what-the-hash/pkg/wth"
)

const banner = `
░█░█░█░█░█▀█░▀█▀░░░▀█▀░█░█░█▀▀░░░█░█░█▀█░█▀▀░█░█░█░▀▀█
░█▄█░█▀█░█▀█░░█░░░░░█░░█▀█░█▀▀░░░█▀█░█▀█░▀▀█░█▀█░▀░░▀░
░▀░▀░▀░▀░▀░▀░░▀░░░░░▀░░▀░▀░▀▀▀░░░▀░▀░▀░▀░▀▀▀░▀░▀░▀░░▀░
`

func main() {
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Fprintln(os.Stderr, "usage: wth HASHSUM")
		os.Exit(2)
	}

	fmt.Println(banner)

	ch := make(chan string, len(wth.DB))

	go wth.Search([]byte(strings.ToLower(os.Args[1])), ch)

	for m := range ch {
		fmt.Println(m)
	}
}
