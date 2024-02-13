package calc

import (
	"fmt"
	"math"
	"strconv"
)

// var x float64
// var y float64
// var xx float64
// var xy float64
//var yy float64


func LinearR(arr []string) {

	// calc Linear Regression Line
	sumX := 0.0
	sumY := 0.0
	sumXY := 0.0
	sumXX := 0.0
	for i := 0; i < len(arr); i++ {
		if len(arr[i]) > 0 {
			vf, _ := strconv.Atoi(arr[i])
			sumX += float64(i)
			sumXX += float64(i) * float64(i)
			sumY += float64(vf)
			sumXY += float64(vf) * float64(i)
		}
	}
	m := (float64(len(arr)-1) * sumXY-sumX * sumY) / (float64(len(arr)-1) * sumXX - sumX * sumX)
	b := (sumY - m * sumX) / float64(len(arr)-1)

	// calc Pearson Correlation Coefficient
	moyX := sumX / float64(len(arr)-1)
	moyY := sumY / float64(len(arr)-1)
	n := 1.0 / float64(len(arr)-1)
	covXY := 0.0
	varX := 0.0
	varY := 0.0
	for i := 0; i < len(arr); i++ {
		if len(arr[i]) > 0 {
			vf, _ := strconv.Atoi(arr[i])
			covXY += (float64(i)-moyX)*(float64(vf)-moyY)
			varX += (float64(i)-moyX)*(float64(i)-moyX)
			varY += (float64(vf)-moyY)*(float64(vf)-moyY)
		}
	}
	covXY = n * covXY
	varX = n * varX
	varY = n * varY
	r := covXY / math.Sqrt(varX*varY)

	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\nPearson Correlation Coefficient: %.10f\n", m, b, r)
}