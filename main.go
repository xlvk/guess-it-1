package main

import (
	"bufio"
	"fmt"
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
		if ha < 100 || ha > 200 {
			arr = append(arr, (int(math.Round(sum / float64(len(arr))))))
			sum = sum + float64((int(math.Round(sum / float64(len(arr))))))
		} else {
			arr = append(arr, ha)
			sum = sum + float64(ha)
		}
		sort.Ints(arr)
		avr = math.Round(sum / float64(len(arr)))
		for j := 0; j < len(arr); j++ {
			sd += math.Pow((float64(arr[j]) - avr), 2)
		}
		sd = math.Round(math.Sqrt(sd / float64(len(arr))))
		y := avr 
		low := int(y - (sd/4))
		up := int(y + (sd/4))
		fmt.Printf("%d %d\n", low, up)
	}
}