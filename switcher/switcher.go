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

func main() {
	defer writer.Flush()
	var nLamps int
	scanf("%d", &nLamps)
	nIter := int(math.Sqrt(float64(nLamps)))
	printf("%d lamps\n", nIter)
	printf("Lamp's number: ")
	for i := 1; i <= nIter; i++ {
		result := i * i
		if i == nIter {
			printf("%d\n", result)
		} else {
			printf("%d ", result)
		}
	}
}
