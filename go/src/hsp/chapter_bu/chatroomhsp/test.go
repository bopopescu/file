package main

import (
	"fmt"
	"regexp"
)

func main() {
	var r  = regexp.MustCompile(`"https://photo.zastatic.com/(.*)\);"></div>`)


	str :="url(\"https://photo.zastatic.com/phot079270683&amp;cpo;h=2);\"></div>"
	match := r.FindSubmatch([]byte(str))

	fmt.Print(string(match[1]))
}
