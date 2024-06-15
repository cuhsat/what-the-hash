// What the hash?
//
// Usage:
//
//	wth HASH
package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/cuhsat/what-the-hash/pkg/wth"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "usage: wth HASH")
		os.Exit(2)
	}

	hash := []byte(strings.ToLower(os.Args[1]))

	for _, e := range wth.DB {
		re := regexp.MustCompile(e.Regex)

		if re.Match(hash) {
			fmt.Println(strings.Join(e.Algos, "\n"))
		}
	}
}
