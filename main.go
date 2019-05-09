// Example app for trying to figure out how packages work in Go

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/bobfirestone/go-packages-in-project-example/double"
	"github.com/bobfirestone/go-packages-in-project-example/triple"
)

func main() {
	if len(os.Args) < 3 {
		return
	}
	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])

	doStuffWithNumbers(a, b)
}

func doStuffWithNumbers(x int, y int) {
	pkg1 := double.Number(x)
	pkg2 := triple.Number(y)
	add := pkg1 + pkg2
	fmt.Printf("(%d * 2) + (%d * 3) = %d\n", x, y, add)
}
