package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var numbers []float64
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		word := scanner.Text()
		num, err := strconv.Atoi(word)
		if err != nil {
			fmt.Println("It's not a number")
			continue
		}
		if len(numbers) > 2 && (num > int(average(numbers))*2 || num < int(average(numbers)*0.3)) {
			num = int(average(numbers))
		}
		numbers = append(numbers, float64(num))
		guessIt(numbers)
	}
}

func standardDev(arr []float64) float64 {
	var sum, mean, sd, count float64
	for i := 1; i <= len(arr); i++ {
		sum += arr[i-1]
	}
	count = float64(len(arr))
	mean = sum / count
	for j := 0; j < len(arr); j++ {
		sd += math.Pow(arr[j]-mean, 2)
	}
	sd = math.Sqrt(sd / count)
	return sd
}

func average(arr []float64) float64 {
	var sum float64
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum / float64(len(arr))
}

func guessIt(arr []float64) {
	fmt.Println(average(arr)-standardDev(arr)*1.28, average(arr)+standardDev(arr)*1.28)
}
