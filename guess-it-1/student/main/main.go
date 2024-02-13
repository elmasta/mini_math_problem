package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var input int
	nmbList := []int{}
	i := 0
	for i <= 12500 {
		fmt.Scan(&input)
		if input <= 200 && input >= 100 {
			nmbList = append(nmbList, input)

			// Average calculation
			average := 0
			for i, v := range nmbList {
				average += v
				if i+1 == len(nmbList) {
					temp := 0.0
					temp = float64(average) / float64((i+1))
					average = int(math.Round(temp))
				}
			}

			// Deviations calculation
			deviation := []int{}
			for _, v := range nmbList {
				deviation = append(deviation, (v-average)*(v-average))
			}

			// Variance calculation
			variance := 0
			for i, v := range deviation {
				variance += v
				if i+1 == len(deviation) {
					temp := 0.0
					temp = float64(variance) / float64((i+1))
					variance = int(math.Round(temp))
				}
			}

			standDev := int(math.Round(math.Sqrt(float64(variance))))

			fmt.Println(strconv.Itoa(average- standDev) + " " + strconv.Itoa(average + standDev))
		} else {
			fmt.Println("100 200")
		}
		i++
	}
}
