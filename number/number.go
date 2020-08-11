package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func countDigits(number int) (count int) {
	for number != 0 {
		number /= 10
		count++
	}
	return count
}

func main() {
	defer writer.Flush()
	var number int
	scanf("%d", &number)
	digits := countDigits(number)
	for i := digits - 1; i >= 0; i-- {
		digit := number / int(math.Pow(10, float64(i)))
		result := digit * int(math.Pow(10, float64(i)))
		number -= result
		printf("%d\n", int(result))
	}
}
