package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := ReadFile("input2.txt")

	fmt.Print(input)

	//Part1(input)
	Part2(input)
}

func ReadFile(filename string) []int {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result []int = []int{}

	for scanner.Scan() {
		numbs := strings.Split(scanner.Text(), " ")
		for _, numb := range numbs {
			num, _ := strconv.Atoi(numb)
			result = append(result, num)
		}
	}

	return result
}

// ####################################################################################################################################
// ####################################################################################################################################
// ########################################                  PART 1                   #################################################
// ####################################################################################################################################
// ####################################################################################################################################

func Part1(input []int) {
	for i := 0; i < 25; i++ {
		fmt.Println(i)
		input = Blink(input)
	}

	fmt.Println("Resultat part 1 : ", len(input))
}

func Blink(input []int) []int {
	limit := len(input)
	for i := 0; i < limit; i++ {
		if input[i] == 0 {
			input[i] = 1
		} else if BumberOfDigits(input[i])%2 == 0 {
			input = append(input, input[i]%int(math.Pow(10, float64(BumberOfDigits(input[i])/2))))
			input[i] = input[i] / int(math.Pow(10, float64(BumberOfDigits(input[i])/2)))

		} else {
			input[i] *= 2024
		}
	}
	//fmt.Println("END     Blink with ", input)

	return input
}

func BumberOfDigits(n int) int {
	result := 0
	for n != 0 {
		n /= 10
		result++
	}
	return result
}

func SlitAtIndex(index int, input []int) []int {
	var result []int = []int{}

	for i := 0; i < len(input); i++ {
		if i == index {
			result = append(result, input[i]/int(math.Pow(10, float64(BumberOfDigits(input[i])/2))))
			result = append(result, input[i]%int(math.Pow(10, float64(BumberOfDigits(input[i])/2))))
		} else {
			result = append(result, input[i])
		}
	}
	return result
}

//####################################################################################################################################
//####################################################################################################################################
//########################################                  PART 2                   #################################################
//####################################################################################################################################
//####################################################################################################################################

func Part2(input []int) {
	var result int = 0

	numbers, quantity := Init(input)

	for i := 0; i < 75; i++ {
		fmt.Println(i)
		fmt.Println("\n numbers : ", numbers)
		fmt.Println("\n quantity : ", quantity)
		numbers, quantity = Blink2(numbers, quantity)
	}

	for i := 0; i < len(numbers); i++ {
		result += quantity[i]
	}

	fmt.Println("Resultat part 2 : ", result)
}

func Init(input []int) ([]int, []int) {
	var numbers []int = []int{}
	var quantity []int = []int{}

	for i := 0; i < len(input); i++ {
		index := NumberIn(numbers, input[i])
		if index == -1 {
			numbers = append(numbers, input[i])
			quantity = append(quantity, 1)
		} else {
			quantity[index]++
		}
	}

	return numbers, quantity
}

func NumberIn(numbers []int, number int) int {
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == number {
			return i
		}
	}
	return -1
}

func Blink2(numbers []int, quantity []int) ([]int, []int) {
	var newNumbers []int = []int{}
	var newQuantity []int = []int{}

	for i := 0; i < len(numbers); i++ {
		var NumbOfDigits int = BumberOfDigits(numbers[i])
		if numbers[i] == 0 {
			newNumbers = append(newNumbers, 1)
			newQuantity = append(newQuantity, quantity[i])
		} else if NumbOfDigits%2 == 0 {
			index1 := NumberIn(newNumbers, numbers[i]%int(math.Pow(10, float64(NumbOfDigits/2))))
			if index1 == -1 {
				newNumbers = append(newNumbers, numbers[i]%int(math.Pow(10, float64(NumbOfDigits/2))))
				newQuantity = append(newQuantity, quantity[i])
			} else {
				newQuantity[index1] += quantity[i]
			}
			index2 := NumberIn(newNumbers, numbers[i]/int(math.Pow(10, float64(NumbOfDigits/2))))
			if index2 == -1 {
				newNumbers = append(newNumbers, numbers[i]/int(math.Pow(10, float64(NumbOfDigits/2))))
				newQuantity = append(newQuantity, quantity[i])
			} else {
				newQuantity[index2] += quantity[i]
			}
		} else {
			newNumbers = append(newNumbers, numbers[i]*2024)
			newQuantity = append(newQuantity, quantity[i])
		}
	}
	return newNumbers, newQuantity
}
