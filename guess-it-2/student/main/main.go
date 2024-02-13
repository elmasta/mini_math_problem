// package main

// import (
// 	"fmt"
// 	"math"
// 	"strconv"
// )

// func main() {
// 	var input int
// 	y := []float64{}
// 	i := 0
// 	for i <= 12500 {
// 		fmt.Scan(&input)
// 		y = append(y, float64(input))
// 		x := []float64{}
// 		for subI := range y {
// 			x = append(x, float64(subI))
// 		}
// 		if i == 0 {
// 			fmt.Println(strconv.Itoa(input-5) + " " + strconv.Itoa(input+5))
// 		} else {
// 			m, b := linearRegression(x, y)
// 			// Calcul du coefficient de corrélation de Pearson
// 			//r := pearsonCorrelationCoefficient(x, y)

// 			// Prédiction pour le prochain nombre
// 			nextX := float64(i)+1.0
// 			nextY := m*nextX + b
// 			fmt.Println(strconv.Itoa(int(nextY-4.0)) + " " + strconv.Itoa(int(nextY+4.0)))
// 		}
// 		i++
// 	}
// }

// // Fonction pour calculer la régression linéaire
// func linearRegression(x, y []float64) (float64, float64) {
// 	//n := len(x)
// 	meanX := calculateMean(x)
// 	meanY := calculateMean(y)

// 	covarianceXY := calculateCovariance(x, y, meanX, meanY)
// 	varianceX := calculateVariance(x, meanX)

// 	m := covarianceXY / varianceX
// 	b := meanY - m*meanX

// 	return m, b
// }

// // Fonction pour calculer le coefficient de corrélation de Pearson
// func pearsonCorrelationCoefficient(x, y []float64) float64 {
// 	n := len(x)
// 	meanX := calculateMean(x)
// 	meanY := calculateMean(y)

// 	numerator := 0.0
// 	denominatorX := 0.0
// 	denominatorY := 0.0

// 	for i := 0; i < n; i++ {
// 		numerator += (x[i] - meanX) * (y[i] - meanY)
// 		denominatorX += math.Pow(x[i]-meanX, 2)
// 		denominatorY += math.Pow(y[i]-meanY, 2)
// 	}

// 	r := numerator / (math.Sqrt(denominatorX) * math.Sqrt(denominatorY))
// 	return r
// }

// // Fonction pour calculer la moyenne
// func calculateMean(data []float64) float64 {
// 	sum := 0.0
// 	for _, value := range data {
// 		sum += value
// 	}
// 	return sum / float64(len(data))
// }

// // Fonction pour calculer la covariance
// func calculateCovariance(x, y []float64, meanX, meanY float64) float64 {
// 	sum := 0.0
// 	for i := 0; i < len(x); i++ {
// 		sum += (x[i] - meanX) * (y[i] - meanY)
// 	}
// 	return sum / float64(len(x))
// }

// // Fonction pour calculer la variance
// func calculateVariance(data []float64, mean float64) float64 {
// 	sum := 0.0
// 	for _, value := range data {
// 		sum += math.Pow(value-mean, 2)
// 	}
// 	return sum / float64(len(data))
// }

package main

import (
	"fmt"
	//"math"
	"strconv"
)

func main() {
	var input int
	nmbList := []int{}
	i := 0
	for i <= 12500 {
		fmt.Scan(&input)
		nmbList = append(nmbList, input)
		if i == 0 {
			fmt.Println(input-5, input+5)
		} else {
			sumX := 0.0
			sumY := 0.0
			sumXY := 0.0
			sumXX := 0.0
			for i := 0; i < len(nmbList); i++ {
				vf := nmbList[i]
				sumX += float64(i)
				sumXX += float64(i) * float64(i)
				sumY += float64(vf)
				sumXY += float64(vf) * float64(i)
			}
			m := (float64(len(nmbList)) * sumXY-sumX * sumY) / (float64(len(nmbList)) * sumXX - sumX * sumX)
			b := (sumY - m * sumX) / float64(len(nmbList))


			// calc Pearson Correlation Coefficient
			moyX := sumX / float64(len(nmbList))
			moyY := sumY / float64(len(nmbList))
			n := 1.0 / float64(len(nmbList))
			covXY := 0.0
			varX := 0.0
			varY := 0.0
			for i := 0; i < len(nmbList); i++ {
				vf := nmbList[i]
				covXY += (float64(i)-moyX)*(float64(vf)-moyY)
				varX += (float64(i)-moyX)*(float64(i)-moyX)
				varY += (float64(vf)-moyY)*(float64(vf)-moyY)
			}
			covXY = n * covXY
			varX = n * varX
			varY = n * varY
			//mm := covXY / varX
			//r := covXY / math.Sqrt(varX*varY)
			//bb := moyY - mm  * moyX*r
			nextX := float64(i)+1.0
			nextY := m*nextX + b
			//r := covXY / math.Sqrt(varX*varY)
			//ran := nextY - nextY*r
			fmt.Println(strconv.Itoa(int(nextY-5.5)) + " " + strconv.Itoa(int(nextY+5.5)))

			//fmt.Println(strconv.Itoa(int((bb)-5.0)) + " " + strconv.Itoa(int((bb)+5.0)))
		}
		i++
	}
}
