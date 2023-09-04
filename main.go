package main

import (
	
	"bufio"
	"fmt"
	"log"
	"os"
	"math"
	"strconv"
	"sort"
)

func main() {

	var arr []int
var ForY []int
	sum := 0.0
	sd := 0.0
	avr := 0.0
	readFile, err := os.Open("data.txt")
	defer readFile.Close()

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		ha, e := strconv.Atoi(fileScanner.Text())
    	if e != nil {
        	fmt.Printf("%T \n %v", ha, ha)
		}
		arr = append(arr, ha)
		ForY = append(ForY, ha)
		sum = sum + float64(ha)
	}
	
	sort.Ints(arr)
	x, e := strconv.Atoi(os.Args[1])
    	if e != nil {
        	fmt.Printf("%T \n %v", x, x)
		}
		if x > len(arr)-1 {
			fmt.Println("Error: NO Y VALUE GIVEN ON THE", x , "VALUE.")
			return
		}
	c := ForY[0]
	m := float64(ForY[x+1]- ForY[0])/float64(x+1)
	y := (m*float64(x+1)) + float64(c)
	avr = math.Round(sum/float64(len(arr)))
	for j := 0; j < len(arr); j++ {
		sd += math.Pow((float64(arr[j])- avr), 2)
	}
	sd = math.Round(math.Sqrt(sd / float64(len(arr))))
	low := (y-(sd/8))
	up := (y+(sd/8))
	fmt.Println(low, "~" , up)
	fmt.Println(ForY[x+1])

}
