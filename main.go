package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"math/bits"
	"math/rand"
	// "crypto/rand"
	// "math/big"
)

type Bitmask uint

const (
	alphabetCaptial string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabetSmall   string = "abcdefghijklmnopqrstuvwxyz"
	digits          string = "1234567890"
	symbols         string = "{}[]()/,;:.<>"
	specSymbols     string = "@#$%"

	FLAG_CAPITAL Bitmask = 1 << iota
	FLAG_SMALL
	FLAG_DIGIT
	FLAG_SYMBOL
	FLAG_SPECSUMBOL
)

var (
	alphabet string
	num      int
	password []string
)

// func (f Bitmask) HasFlag(flag Bitmask) bool { return f&flag != 0 }
func (f *Bitmask) AddFlag(flag Bitmask) { *f |= flag }

// func (f *Bitmask) ClearFlag(flag Bitmask)   { *f &= ^flag }
// func (f *Bitmask) ToggleFlag(flag Bitmask)  { *f ^= flag }

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

func randString(num int) []string {

	//var password []string
	var leftInt int
	rand.Seed(time.Now().UnixNano())
	if alphabet == "" {
		return []string{"error"}
	}
	randomInt := rand.Intn(len(alphabet))
	time.Sleep(time.Duration(10 * time.Millisecond))
	// randomInt, err := rand.Int(rand.Reader, big.NewInt(6))
	leftInt = randomInt - 1
	if randomInt == 0 {
		leftInt = randomInt
		randomInt += 1
	}
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
	captialFlagPtr := flag.Bool("C", false, "use alphabet \"ABCDEFGHIJKLMNOPQRSTUVWXYZ\" ")
	smallFlagPtr := flag.Bool("c", false, "use alphabet \"abcdefghijklmnopqrstuvwxyz\" ")
	digitsFlagPtr := flag.Bool("d", false, "use alphabet \"1234567890\" ")
	symbolsFlagPtr := flag.Bool("s", false, "use alphabet \"{}[]()/,;:.<>\" ")
	specSymbolsFlagPtr := flag.Bool("a", false, "use alphabet \"@#$%\" ")
	customerAlphabetPtr := flag.String("custom", "", "use your own alphabet ")

	flag.Parse()
	num = *lengthPtr

	var mainFlag Bitmask

	if *captialFlagPtr {
		mainFlag.AddFlag(FLAG_CAPITAL)
	}
	if *smallFlagPtr {
		mainFlag.AddFlag(FLAG_SMALL)
	}
	if *digitsFlagPtr {
		mainFlag.AddFlag(FLAG_DIGIT)
	}
	if *symbolsFlagPtr {
		mainFlag.AddFlag(FLAG_SYMBOL)
	}
	if *specSymbolsFlagPtr {
		mainFlag.AddFlag(FLAG_SPECSUMBOL)
	}
	var number uint = bits.RotateLeft(uint(mainFlag), -5)

	alphabet = fmt.Sprintf("%s%s%s", digits, alphabetCaptial, alphabetSmall)

	if number == 31 {
		alphabet = fmt.Sprintf("%s%s%s%s%s", alphabetCaptial, alphabetSmall, digits, symbols, specSymbols)
	}
	if number == 30 {
		alphabet = fmt.Sprintf("%s%s%s%s", alphabetSmall, digits, symbols, specSymbols)
	}
	if number == 29 {
		alphabet = fmt.Sprintf("%s%s%s%s", alphabetCaptial, digits, symbols, specSymbols)
	}
	if number == 28 {
		alphabet = fmt.Sprintf("%s%s%s", digits, symbols, specSymbols)
	}
	if number == 27 {
		alphabet = fmt.Sprintf("%s%s%s%s", alphabetCaptial, alphabetSmall, symbols, specSymbols)
	}
	if number == 26 {
		alphabet = fmt.Sprintf("%s%s%s", alphabetSmall, symbols, specSymbols)
	}
	if number == 25 {
		alphabet = fmt.Sprintf("%s%s%s", alphabetCaptial, symbols, specSymbols)
	}
	if number == 24 {
		alphabet = fmt.Sprintf("%s%s", symbols, specSymbols)
	}
	if number == 23 {
		alphabet = fmt.Sprintf("%s%s%s%s", alphabetCaptial, alphabetSmall, digits, specSymbols)
	}
	if number == 22 {
		alphabet = fmt.Sprintf("%s%s%s", alphabetSmall, digits, specSymbols)
	}
	if number == 21 {
		alphabet = fmt.Sprintf("%s%s%s", alphabetCaptial, digits, specSymbols)
	}
	if number == 20 {
		alphabet = fmt.Sprintf("%s%s", digits, specSymbols)
	}
	if number == 19 {
		alphabet = fmt.Sprintf("%s%s%s", alphabetCaptial, alphabetSmall, specSymbols)
	}
	if number == 18 {
		alphabet = fmt.Sprintf("%s%s", alphabetSmall, specSymbols)
	}
	if number == 17 {
		alphabet = fmt.Sprintf("%s%s", alphabetCaptial, specSymbols)
	}
	if number == 16 {
		alphabet = fmt.Sprintf("%s", specSymbols)
	}
	if number == 15 {
		alphabet = fmt.Sprintf("%s%s%s%s", alphabetCaptial, alphabetSmall, digits, symbols)
	}
	if number == 14 {
		alphabet = fmt.Sprintf("%s%s%s", alphabetSmall, digits, symbols)
	}
	if number == 13 {
		alphabet = fmt.Sprintf("%s%s%s", alphabetCaptial, digits, symbols)
	}
	if number == 12 {
		alphabet = fmt.Sprintf("%s%s", digits, symbols)
	}
	if number == 11 {
		alphabet = fmt.Sprintf("%s%s%s", alphabetCaptial, alphabetSmall, symbols)
	}
	if number == 10 {
		alphabet = fmt.Sprintf("%s%s", alphabetSmall, symbols)
	}
	if number == 9 {
		alphabet = fmt.Sprintf("%s%s", alphabetCaptial, symbols)
	}
	if number == 8 {
		alphabet = fmt.Sprintf("%s", symbols)
	}
	if number == 7 {
		alphabet = fmt.Sprintf("%s%s%s", alphabetCaptial, alphabetSmall, digits)
	}
	if number == 6 {
		alphabet = fmt.Sprintf("%s%s", alphabetSmall, digits)
	}
	if number == 5 {
		alphabet = fmt.Sprintf("%s%s", alphabetCaptial, digits)
	}
	if number == 4 {
		alphabet = fmt.Sprintf("%s", digits)
	}
	if number == 3 {
		alphabet = fmt.Sprintf("%s%s", alphabetCaptial, alphabetSmall)
	}
	if number == 2 {
		alphabet = fmt.Sprintf("%s", alphabetSmall)
	}
	if number == 1 {
		alphabet = fmt.Sprintf("%s", alphabetCaptial)
	}

	if *customerAlphabetPtr != "" {
		alphabet = fmt.Sprintf("%s", *customerAlphabetPtr)
	}
	fmt.Printf("%d\n", number)
}

func main() {
	pass := randString(num)
	fmt.Println(strings.Join(pass[:], ""))
}
