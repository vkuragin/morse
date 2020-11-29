package morse

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

type Morse struct {
	charMap map[rune]string
	codeMap map[string]rune
}

func New() (*Morse, error) {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ "
	codes := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....",
		"..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.",
		"--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-",
		"-.--", "--..", " "}
	if len(letters) != len(codes) {
		return nil, fmt.Errorf("len(letters) != len(codes): %d, %d\n", len(letters), len(codes))
	}

	// populate letter-to-code mappings
	charMap := make(map[rune]string, 27)
	for i, r := range letters {
		charMap[r] = codes[i]
	}

	// populate code-to-letter mappings
	codeMap := make(map[string]rune, 27)
	for i, r := range letters {
		codeMap[codes[i]] = r
	}

	return &Morse{charMap: charMap, codeMap: codeMap}, nil
}

func (m *Morse) Encode(src string) string {
	src = strings.ToUpper(src)
	b := strings.Builder{}
	for i, r := range src {
		if i != 0 {
			b.WriteRune(' ')
		}
		b.WriteString(m.charMap[r])
	}
	return b.String()
}

func (m *Morse) Decode(src string) string {

	var res []rune
	seq := bytes.Buffer{}
	prevSpace := true
	for i, r := range src {
		isSpace := unicode.IsSpace(r)
		if isSpace && prevSpace {
			res = append(res, ' ')
		} else if isSpace {
			code := string(seq.Bytes())
			res = append(res, m.codeMap[code])
			seq.Reset()
		} else if i == len(src)-1 {
			seq.WriteRune(r)
			code := string(seq.Bytes())
			res = append(res, m.codeMap[code])
			seq.Reset()
		} else if prevSpace {
			seq.Reset()
			seq.WriteRune(r)
		} else {
			seq.WriteRune(r)
		}
		prevSpace = isSpace
	}

	return string(res)
}
