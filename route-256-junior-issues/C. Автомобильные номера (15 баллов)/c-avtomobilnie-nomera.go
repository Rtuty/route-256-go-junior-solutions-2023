package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	var k int
	reader := bufio.NewReader(os.Stdin)
	fmt.Scanln(&k)
	numbers := make([]string, k)
	for i := 0; i < k; i++ {
		fmt.Fscanln(reader, &(numbers[i]))

	}
	for i := 0; i < k; i++ {
		rule := regexp.MustCompile("([a-zA-Z]\\d\\d[a-zA-Z][a-zA-Z])|([a-zA-Z]\\d[a-zA-Z][a-zA-Z])")

		res := rule.FindAllString(numbers[i], 999)

		if pass := rule.ReplaceAllString(numbers[i], ""); len(pass) == 0 && len(res) != 0 {
			fmt.Println(strings.Join(res, " "))
		} else {
			fmt.Println("-")
		}

	}
}
