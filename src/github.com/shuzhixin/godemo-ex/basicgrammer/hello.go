package main

import (
	"fmt"
	myfunc "github.com/shuzhixin/godemo-ex/src/github.com/shuzhixin/godemo-ex/basicgrammer/function"
)

func main() {
	fmt.Println("Hello, World!")

	var sum int
	sum = myfunc.Add(2, 3)
	fmt.Printf("%d + %d = %d", 2, 3, sum)

}
