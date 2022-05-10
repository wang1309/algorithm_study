package search

import (
	"fmt"
	"testing"
)

func TestBfSearch(t *testing.T)  {
	m := "abcdaa"
	s1 := "dad"
	s2 := "a"
	s3 := "ab"
	s4 := "aa"
	s5 := "bd"
	s6 := "abcdaa"

	fmt.Println(bfSearch(m, s1))
	fmt.Println(bfSearch(m, s2))
	fmt.Println(bfSearch(m, s3))
	fmt.Println(bfSearch(m, s4))
	fmt.Println(bfSearch(m, s5))
	fmt.Println(bfSearch(m, s6))
}
