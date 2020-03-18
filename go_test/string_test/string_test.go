package string_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestStringConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}

}

func TestStringMap(t *testing.T) {
	s1 := "how do you do"
	s2 := strings.Split(s1, " ")
	m1 := make(map[string]int, 10)
	for _, w := range s2 {
		m1[w]++
	}

	for key, value := range m1 {
		fmt.Println(key, value)
	}
}

func isPalindrome(str string) bool {
	length := utf8.RuneCountInString(str)
	r := make([]rune, 0, length)
	for _, c := range str {
		r = append(r, c)
	}

	j := length - 1
	for i := 0; i < length/2; i++ {
		if r[i] != r[j] {
			return false
		}
		j--
	}
	return true
}

func TestPalindrome(t *testing.T) {
	str := "上海自来水来自海上"
	if isPalindrome(str) {
		fmt.Println("str是回文")
	} else {
		fmt.Println("str不是回文")
	}
}
