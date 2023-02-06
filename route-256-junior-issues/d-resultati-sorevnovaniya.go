package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCounter int
	fmt.Fscan(in, &testCounter)

	for i := 0; i < testCounter; i++ {
		var k int
		fmt.Fscan(in, &k)

		input := []int{}
		ans := []int{}
		times := map[int]int{}
		used := map[int]int{}

		for j := 0; j < k; j++ {
			var n int
			fmt.Fscan(in, &n)
			input = append(input, n)
			ans = append(ans, n)
			times[n]++
			used[n] = -1
		}

		input = unique(input)
		place := 1
		sort.Ints(input)

		for _, n := range input {
			if used[n] != -1 {
				continue
			}
			place += checkNearest(n, used, times, place)
		}

		for _, n := range ans {
			fmt.Fprint(out, used[n], " ")
		}
		fmt.Fprintln(out)
	}
}

func checkNearest(n int, used map[int]int, times map[int]int, place int) int {
	used[n] = place
	if val, ok := used[n+1]; ok && val == -1 {
		return checkNearest(n+1, used, times, place) + times[n]
	}
	return times[n]
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
