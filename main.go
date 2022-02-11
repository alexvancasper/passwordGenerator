package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
)

const (
	alphabetCaptial string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabetSmall   string = "abcdefghijklmnopqrstuvwxyz"
	digits          string = "1234567890"
	symbols         string = "{}[]()/,;:.<>"
	specSymbols     string = "@#$%"
)

var alphabet string
var num int

func randStringNum(password []string, alphabet string, num int) []string {
	// var output []string
	var leftInt int
	randomInt := rand.Intn(len(alphabet))
	leftInt = randomInt - 1
	if (randomInt - 1) < 1 {
		leftInt = randomInt - 1
	}
	if (randomInt + 1) > len(alphabet) {
		leftInt = randomInt - 1
	}

	password = append(password, alphabet[leftInt:randomInt])
	if len(password) != num {
		password = randStringNum(password, alphabet, num)
	}

	return password
}

func randString(num int) (password []string) {

	//var password []string
	var leftInt int
	randomInt := rand.Intn(len(alphabet))
	leftInt = randomInt - 1
	if (randomInt - 1) < 1 {
		leftInt = randomInt - 1
	}
	if (randomInt + 1) > len(alphabet) {
		leftInt = randomInt - 1
	}

	password = append(password, alphabet[leftInt:randomInt])
	if len(password) != num {
		password = randString(num)
	}

	return password
}

func init() {
	lengthPtr := flag.Int("n", 16, "length of password")
	captialFlagPtr := flag.Bool("C", true, "use alphabet \"ABCDEFGHIJKLMNOPQRSTUVWXYZ\" ")
	smallFlagPtr := flag.Bool("c", true, "use alphabet \"abcdefghijklmnopqrstuvwxyz\" ")
	digitsFlagPtr := flag.Bool("d", true, "use alphabet \"1234567890\" ")
	symbolsFlagPtr := flag.Bool("s", false, "use alphabet \"{}[]()/,;:.<>\" ")
	specSymbolsFlagPtr := flag.Bool("a", false, "use alphabet \"@#$%\" ")
	customerAlphabetPtr := flag.String("custom", "", "use your own alphabet ")
	flag.Parse()
	num = *lengthPtr

	if *captialFlagPtr && *smallFlagPtr && *digitsFlagPtr && *symbolsFlagPtr && *specSymbolsFlagPtr {
		alphabet = fmt.Sprintf("%s%s%s%s%s", alphabetCaptial, alphabetSmall, digits, symbols, specSymbols)
	}
	if *captialFlagPtr && *smallFlagPtr && *digitsFlagPtr && *symbolsFlagPtr {
		alphabet = fmt.Sprintf("%s%s%s%s", alphabetCaptial, alphabetSmall, digits, symbols)
	}
	if *captialFlagPtr && *smallFlagPtr && *digitsFlagPtr {
		alphabet = fmt.Sprintf("%s%s%s", alphabetCaptial, alphabetSmall, digits)
	}
	if *captialFlagPtr && *smallFlagPtr {
		alphabet = fmt.Sprintf("%s%s", alphabetCaptial, alphabetSmall)
	}
	if *captialFlagPtr {
		alphabet = fmt.Sprintf("%s", alphabetCaptial)
	}

	if *customerAlphabetPtr == "" {
		alphabet = fmt.Sprintf("%s", customerAlphabetPtr)
	}

}

func main() {

	pass := randString(num)
	fmt.Println(strings.Join(pass[:], ""))
}
