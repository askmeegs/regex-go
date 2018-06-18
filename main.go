package main

import (
	"fmt"

	"github.com/m-okeefe/regex-go/helpers"
)

func main() {
	fmt.Println(" ---------- go regular expressions --------------")

	//runTest(helpers.Ares, []string{"archeobotanies", "arbitrates", "arches", "arranges", "arficies", "ares", "are", "", "_ares", "bubbles", "123abc", "cares"})

	// runTest(helpers.PhoneNumber, []string{"(123)-456-7890", "abcde-999", "+1 555.555.4444", "1234512345", ""})

	runTest(helpers.Email, []string{"abc@efgh.com", "my.cool.email@gmail.com", "rain@verizon.net", "hello world", "@aol.com", "$runner$@yahoo.com", "coffee@", "@", ".com"})

}

// pretty-prints regex helper output for multiple inputs
func runTest(fn func(string) bool, input []string) {
	for _, x := range input {
		result := fn(x)
		fmt.Printf("%s:  %v\n", x, result)
	}
}
