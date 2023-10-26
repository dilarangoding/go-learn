package main

import (
	"fmt"
	"regexp"
)

func main() {

	var regex *regexp.Regexp = regexp.MustCompile("e([a-z])o")

	fmt.Println(regex.MatchString("eto"))
	fmt.Println(regex.MatchString("euo"))
	fmt.Println(regex.MatchString("eDo"))

	search := regex.FindAllString("budi buda badi bubi badu eto", 1)
	fmt.Println(search)

}
