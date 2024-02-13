package main

import (
	"strconv"
	"os"
	"math"
	"fmt"
	"sort"
)

func main() {
	if len(os.Args) == 2 {
		content, err := os.Open(os.Args[1])
		if err != nil {
			os.Stdout.Write([]byte("ERROR: open " + os.Args[1] + ": no such file or directory\n"))
			os.Exit(1)
		} else {
			contSize, _ := content.Stat()
			arr := make([]byte, contSize.Size())
			content.Read(arr)
			//séparation avec \n
			intTable := []int{}
			temp := ""
			for _, v := range arr {
				if v == '\n' {
					nmb, err := strconv.Atoi(temp)
					if err != nil {
						fmt.Println("Error during conversion")
					} else {
						intTable = append(intTable, nmb)
					}
					temp = ""
				} else {
					temp += string(v)
				}
			}
			if len(temp) > 0 {
				nmb, err := strconv.Atoi(temp)
				if err != nil {
					fmt.Println("Error during conversion")
				} else {
					intTable = append(intTable, nmb)
				}
			}
			content.Close()
			/*
			Format:
			Average: 35
			Median: 4
			Variance: 5
			Standard Deviation: 65
			*/
			// Average calculation
			average := 0
			for i, v := range intTable {
				average += v
				if i+1 == len(intTable) {
					temp := 0.0
					temp = float64(average) / float64((i+1))
					average = int(math.Round(temp))
				}
			}
			// Deviations calculation
			deviation := []int{}
			for _, v := range intTable {
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
			// Standard deviation calculation
			standDev := int(math.Round(math.Sqrt(float64(variance))))
			// Median calculation
			median := 0
			//sort the table
			sort.Ints(intTable)
			if len(intTable) % 2 == 0 {
				median = int(math.Round((float64(intTable[len(intTable)/2-1]) + float64(intTable[len(intTable)/2]))/2))
			} else {
				median = intTable[len(intTable)/2]
			}
			fmt.Println("Average: " + strconv.Itoa(average) + "\nMedian: " + strconv.Itoa(median) + "\nVariance: " + strconv.Itoa(variance) + "\nStandard Deviation: " + strconv.Itoa(standDev))
		}
	} else {
		os.Stdout.Write([]byte("Not enough or too many arguments\n"))
		os.Exit(2)
	}
}
