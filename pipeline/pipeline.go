package pipeline

import (
	"strings"
	"unicode"
)

func step1(in <-chan string, out chan<- string) {
	for strIn := range in {
		if strIn == "" {
			continue
		}
		strOut := []rune{}
		str := strings.TrimSpace(strIn)

		var prev rune
		for _, current := range str {
			if unicode.IsSpace(current) && unicode.IsSpace(prev) {
				continue
			}
			strOut = append(strOut, current)
			prev = current
		}
		out <- string(strOut)
	}
	close(out)
}

func step2(in <-chan string, out chan<- string) {
	for s := range in {
		sentances := strings.Split(strings.Trim(s, "."), ".")
		for _, sentance := range sentances {
			if len(sentance) > 0 {
				out <- strings.TrimSpace(sentance)
			}
		}
	}
	close(out)
}

// Обратите внимание, что step3 должен вернуть канал, в который будет записывать.
// Это значит, что внутри функции нужно запустить отдельную горутину, читающую in.
func step3(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for v := range in {
			str := []rune(v)
			str[0] = unicode.ToUpper(str[0])
			out <- string(str)
		}
		close(out)
	}()
	return out
}
