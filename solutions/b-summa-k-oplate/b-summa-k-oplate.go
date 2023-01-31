package b_summa_k_oplate

import (
	"bufio"
	"fmt"
	"os"
)

func sum() {
	in := bufio.NewReader(os.Stdin)
	res := bufio.NewWriter(os.Stdout)
	defer res.Flush()

	var n map[int]int

	fmt.Fscan(in, &n)

	for i := 0; i < len(n); i++ {
	}
}

func testing_sum(a int, b int, res []int) int {
	return 12
}
