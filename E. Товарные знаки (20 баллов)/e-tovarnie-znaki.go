package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var k int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &k)

	builds := make([][]string, k)
	for i := 0; i < k; i++ {
		var l int
		fmt.Fscanln(reader, &l)
		var build = make([]string, l)
		for j := 0; j < l; j++ {
			fmt.Fscanln(reader, &(build[j]))
		}
		builds[i] = build
	}
	for i := range builds {
		for j := range builds[i] {
			bytestr := []byte(builds[i][j])
			newstr := ``
			var oldlrt byte
			for s := range bytestr {
				if s != 0 {
					if bytestr[s] == bytestr[s-1] {
						if oldlrt != bytestr[s] {
							oldlrt = bytestr[s]
						} else {
							continue
						}
					} else {
						oldlrt = 0
					}
				}
				newstr += string(bytestr[s])
			}
			builds[i][j] = newstr
		}
		fmt.Println(Unique(builds[i]))
	}
}

func Unique(arr []string) int {
	var uniq []string
	for _, s := range arr {
		if !Contains(uniq, s) {
			uniq = append(uniq, s)
		}
	}
	return len(uniq)
}

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
