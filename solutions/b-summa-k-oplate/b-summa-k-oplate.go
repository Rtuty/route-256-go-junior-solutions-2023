package b_summa_k_oplate

import (
	"bufio"
	"fmt"
	"os"
)

func summaKoplate() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ {
		var productCount, sum int
		var products = make(map[int]int)
		fmt.Fscan(in, &productCount)
		for i := 0; i < productCount; i++ {
			var productPrice int
			fmt.Fscan(in, &productPrice)
			products[productPrice]++
		}
		for k, v := range products {
			sum += (k * ((v / 3 * 2) + v%3))
		}
		fmt.Fprintln(out, sum)
	}

}
