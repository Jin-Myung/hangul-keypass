// main.go
package main

import (
	"crypto/rand"
	_ "embed"
	"math/big"
	"strings"
	"syscall/js"
)

//go:embed internal/words.txt
var wordText string

//go:embed internal/words_strong.txt
var strongText string

var wordList []string
var strongList []string

func init() {
	wordList = strings.Split(strings.TrimSpace(wordText), "\n")
	strongList = strings.Split(strings.TrimSpace(strongText), "\n")
}

// 확률 p (0.0 ~ 1.0)로 true 반환
func secureBool(p float64) bool {
	nBig, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		return false
	}
	return float64(nBig.Int64()) < p*100
}

func secureIndex(max int) int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0 // fallback
	}
	return int(nBig.Int64())
}

func secureShuffle(slice []string) {
	n := len(slice)
	for i := n - 1; i > 0; i-- {
		j := secureIndex(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func secureSample(charset string, n int) string {
	var out strings.Builder
	max := big.NewInt(int64(len(charset)))

	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, max)
		if err != nil {
			// fallback: panic or return partial
			break
		}
		out.WriteByte(charset[num.Int64()])
	}
	return out.String()
}

func decomposeHangul(r rune) (rune, rune, rune) {
	const (
		base     = 0xAC00
		chosung  = 588
		jungsung = 28
	)

	offset := r - base
	if offset < 0 || offset > 11171 {
		return 0, 0, 0
	}

	cho := offset / chosung
	jung := (offset % chosung) / jungsung
	jong := offset % jungsung

	return cho, jung, jong
}

var chosungKey = []string{
	"r", "R", "s", "e", "E", "f", "a", "q", "Q",
	"t", "T", "d", "w", "W", "c", "z", "x", "v", "g",
}
var jungsungKey = []string{
	"k", "o", "i", "O", "j", "p", "u", "P", "h", "hk",
	"ho", "hl", "y", "n", "nj", "np", "nl", "b", "m",
	"ml", "l",
}
var jongsungKey = []string{
	"", "r", "R", "rt", "s", "sw", "sg", "e", "f", "fr",
	"fa", "fq", "ft", "fx", "fv", "fg", "a", "q", "qt",
	"t", "T", "d", "w", "c", "z", "x", "v", "g",
}

func hangulToQwerty(input string) string {
	var result strings.Builder

	for _, r := range input {
		if r >= 0xAC00 && r <= 0xD7A3 {
			cho, jung, jong := decomposeHangul(r)
			result.WriteString(chosungKey[cho])
			result.WriteString(jungsungKey[jung])
			if jong != 0 {
				result.WriteString(jongsungKey[jong])
			}
		} else {
			// 알파벳, 숫자 등은 그대로
			result.WriteRune(r)
		}
	}

	return result.String()
}

func generate(this js.Value, args []js.Value) interface{} {
	useNum := args[0].Bool()
	useSym := args[1].Bool()

	// 단어 구성: 강한 단어 + 일반 단어
	parts := []string{
		strongList[secureIndex(len(strongList))],
		wordList[secureIndex(len(wordList))],
	}

	if useNum {
		parts = append(parts, secureSample("0123456789", 1))
		for i := 0; i < 3; i++ {
			if secureBool(0.3) {
				parts = append(parts, secureSample("0123456789", 1))
			}
		}
	}

	if useSym {
		parts = append(parts, secureSample("!@#$%^&*()-_=+[]{}", 1))
		for i := 0; i < 2; i++ {
			if secureBool(0.2) {
				parts = append(parts, secureSample("!@#$%^&*()-_=+[]{}", 1))
			}
		}
	}

	secureShuffle(parts)

	origin := strings.Join(parts, "")
	password := hangulToQwerty(origin)

	js.Global().Get("document").Call("getElementById", "result").Set("innerText", origin)
	js.Global().Get("document").Call("getElementById", "origin-word").Set("innerText", password)
	return nil
}

func main() {
	js.Global().Set("generatePassword", js.FuncOf(generate))
	select {}
}
