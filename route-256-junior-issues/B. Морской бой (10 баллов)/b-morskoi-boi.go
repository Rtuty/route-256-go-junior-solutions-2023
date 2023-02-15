package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var k int
	fmt.Scanln(&k)

	scanner := bufio.NewScanner(os.Stdin)

	var testCounter = make([]string, k)
	for i := 0; i < k; i++ {
		scanner.Scan()
		testCounter[i] = scanner.Text()
	}

	for i := range testCounter {
		mass := strings.Split(testCounter[i], " ")

		massInt := make([]int, len(mass))
		for i := range mass {
			v, _ := strconv.Atoi(mass[i])
			massInt[i] = v
		}

		var ship = make(map[int]int)

		ship[1] = 0
		ship[2] = 0
		ship[3] = 0
		ship[4] = 0

		for _, v := range massInt {
			switch v {
			case 1:
				ship[1] += 1
			case 2:
				ship[2] += 1
			case 3:
				ship[3] += 1
			case 4:
				ship[4] += 1
			}
		}

		if ship[1] == 4 && ship[2] == 3 && ship[3] == 2 && ship[4] == 1 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
