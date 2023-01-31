package a_summator

import (
	"bufio"
	"fmt"
	"os"
)

func adder() {
	in := bufio.NewReader(os.Stdin)
	res := bufio.NewWriter(os.Stdout)
	defer res.Flush()

	var count int
	fmt.Fscan(in, &count)

	for i := 0; i < count; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		fmt.Fprintln(res, a+b)
	}
}
