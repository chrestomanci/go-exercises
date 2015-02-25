package main

import (
	"fmt"
)

func mkOddGenerator() func() uint {

	i := uint(1)

	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {

	nextodd := mkOddGenerator()

	for i := 1; i <= 10; i++ {
		fmt.Println( nextodd())
	}
}
