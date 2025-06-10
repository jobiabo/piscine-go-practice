package main

import (
	"fmt"
	"strings"
)

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	fmt.Println(ToUpper("hello how are you"))

}
