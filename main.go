package main

import (
	"bufio"
	"fmt"
	// "log"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	var arr []int
	sum := 0.0
	sd := 0.0
	avr := 0.0
	fileScanner := bufio.NewScanner(os.Stdin)
	for fileScanner.Scan() {
		ha, e := strconv.Atoi(fileScanner.Text())
		if e != nil {
			fmt.Printf("%T \n %v", ha, ha)
		}
		arr = append(arr, ha)
		sum = sum + float64(ha)
		sort.Ints(arr)
		avr = math.Round(sum / float64(len(arr)))
		for j := 0; j < len(arr); j++ {
			sd += math.Pow((float64(arr[j]) - avr), 2)
		}
		sd = math.Round(math.Sqrt(sd / float64(len(arr))))
		y := avr - sd
		low := int(y - 3)
		up := int(y + 2)

		fmt.Printf("%d %d\n" ,low, up)
	}

}
