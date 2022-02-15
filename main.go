package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"strings"
)

type Bitmask uint

const (
	FLAG_CAPITAL Bitmask = 1 << iota
	FLAG_SMALL
	FLAG_DIGIT
	FLAG_SYMBOL
	FLAG_SPECSUMBOL
	FLAG_MYSUMBOL
)

var (
	passwdLen int
	mainFlag  Bitmask
)

// based on unicode table
//  symbols from 65 - 90  -> ABCDEFGHIJKLMNOPQRSTUVWXYZ
//  symbols from 97 - 122 -> abcdefghijklmnopqrstuvwxyz
//  symbols from 48 - 57  -> 1234567890
//  symbols from 33 - 47, 58 - 63, 91 - 96, 123 - 126 -> special symbols

func (f Bitmask) HasFlag(flag Bitmask) bool { return f&flag != 0 }
func (f *Bitmask) AddFlag(flag Bitmask)     { *f |= flag }

// func (f *Bitmask) ClearFlag(flag Bitmask)   { *f &= ^flag }
// func (f *Bitmask) ToggleFlag(flag Bitmask)  { *f ^= flag }

func initAlphabet(alphabetFlag Bitmask) (dict map[int]int) {
	var l int = 0
	dict = make(map[int]int, 93) //26+26+10+15+6+6+4
	if alphabetFlag.HasFlag(FLAG_CAPITAL) {
		for i := 65; i <= 90; i++ { // ABCDEFGHIJKLMNOPQRSTUVWXYZ -> 0..25
			dict[l] = i
			l++
		}
	}
	if alphabetFlag.HasFlag(FLAG_SMALL) {
		for i := 97; i <= 122; i++ { //abcdefghijklmnopqrstuvwxyz -> 26 .. 51
			dict[l] = i
			l++
		}
	}
	if alphabetFlag.HasFlag(FLAG_DIGIT) {
		for i := 48; i <= 57; i++ { //1234567890 -> 10 ->  52..61
			dict[l] = i
			l++
		}
	}
	if alphabetFlag.HasFlag(FLAG_SYMBOL) {
		for i := 33; i <= 47; i++ { // 62..76  !#%&()*+,-./
			if i == 34 || i == 36 || i == 39 {
				continue
			}
			dict[l] = i
			l++
		}
		for i := 58; i <= 63; i++ { // 77..82 :;< = >?
			dict[l] = i
			l++
		}
		for i := 91; i <= 95; i++ { // 83..88 [\]^_
			dict[l] = i
			l++
		}
		for i := 123; i <= 125; i++ { // 89..92 {|}
			dict[l] = i
			l++
		}
	}
	if alphabetFlag.HasFlag(FLAG_SPECSUMBOL) { // " ' ` $ ~
		dict[l] = 34 // "
		l++
		dict[l] = 39 // '
		l++
		dict[l] = 96 // `
		l++
		dict[l] = 36 // $
		l++
		dict[l] = 126 // ~
	}
	if alphabetFlag.HasFlag(FLAG_MYSUMBOL) { // !@# $%^ &
		dict[l] = 33 // !
		l++
		dict[l] = 64 // @
		l++
		dict[l] = 35 // #
		l++
		dict[l] = 36 // $
		l++
		dict[l] = 37 // %
		l++
		dict[l] = 94 // ^
		l++
		dict[l] = 38 // &
	}

	return dict
}

func passwordCryptGenerate(length int, dictionary map[int]int) string {
	if len(dictionary) == 0 {
		return "error to initialaze the alphabet; use -h for helping"
	}
	var passInt []int
	for i := 0; i < length; i++ {
		nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(dictionary))))
		if err != nil {
			panic(err)
		}
		n := nBig.Int64()
		passInt = append(passInt, dictionary[int(n)])
		if len(passInt) == length {
			break
		}
	}
	output := make([]string, length)
	for key, value := range passInt {
		output[key] = fmt.Sprintf("%c", value)
	}

	return strings.Join(output[:], "")
}

func init() {
	lengthPtr := flag.Int("n", 16, "length of password")
	captialFlagPtr := flag.Bool("C", false, "use Capital lettter ABCDEFGHIJKLMNOPQRSTUVWXYZ ")
	smallFlagPtr := flag.Bool("c", false, "use small letter abcdefghijklmnopqrstuvwxyz ")
	digitsFlagPtr := flag.Bool("d", false, "use digits 1234567890 ")
	symbolsFlagPtr := flag.Bool("s", false, "use symbols !#%&()*+,-./:;< = >?[\\]^_{|} ")
	specSymbolsFlagPtr := flag.Bool("a", false, "use special symbols \"'`$~ ")
	myOwnSymbolsFlagPtr := flag.Bool("m", false, "use special symbols !@#$%^& ")
	// customerAlphabetPtr := flag.String("custom", "", "use your own alphabet ")

	flag.Parse()
	passwdLen = *lengthPtr

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
	if *myOwnSymbolsFlagPtr {
		mainFlag.AddFlag(FLAG_MYSUMBOL)
	}

	// if *customerAlphabetPtr != "" {
	// 	alphabet = fmt.Sprintf("%s", *customerAlphabetPtr)
	// }
}

func main() {
	resultDict := initAlphabet(mainFlag)
	pass := passwordCryptGenerate(passwdLen, resultDict)
	fmt.Println(pass)
}
