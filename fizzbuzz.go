package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getStdin() string {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	return stdin.Text()
}

func fizz(x int) string {
	if x % 3 == 0 {
		return "Fizz"
	}
	return ""
}

func buzz(x int) string {
	if x % 5 == 0 {
		return "Buzz"
	}
	return ""
}

func main()  {
	n, _ := strconv.Atoi(getStdin())
	fmt.Printf(fizz(n)+buzz(n))
}

