package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/isaias-dgr/go_test/src/mapper"
)

func CapitalizeEveryThirdAlphanumericChar(s string) string {
	runas := []rune(strings.ToLower(s))
	for i := 3; i <= len(runas); i += 3 {
		runas[i-1] = unicode.To(unicode.UpperCase, runas[i-1])
	}
	return string(runas)
}

func main() {
	s, e := mapper.NewSkipString(3, "Aspiration.com")
	if e != nil {
		fmt.Printf("skipString fail: %s", e.Error())
	}
	mapper.MapString(s)
	fmt.Println(s)
}
