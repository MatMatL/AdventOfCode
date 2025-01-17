package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := ReadFile("input.txt")

	Part2(input)
}

type data struct {
	value int
	numbs []int
}

func ReadFile(filename string) []data {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result []data = []data{}

	for scanner.Scan() {
		var tempo data = data{}
		datas := strings.Split(scanner.Text(), ": ")
		tempo.value, _ = strconv.Atoi(datas[0])
		numbs := strings.Split(datas[1], " ")

		for _, numb := range numbs {
			num, _ := strconv.Atoi(numb)
			tempo.numbs = append(tempo.numbs, num)
		}
		result = append(result, tempo)
	}

	return result
}

//####################################################################################################################################
//####################################################################################################################################
//########################################                  PART 1                   #################################################
//####################################################################################################################################
//####################################################################################################################################

func Part1(input []data) {
	var result int = 0

	for _, data := range input {
		fmt.Println("######### data : ", data)
		result += SolveData(data)
	}

	fmt.Println("Resultat part 1 : ", result)
}

func SolveData(data data) int {
	var ranging []bool = []bool{}

	for i := 0; i < len(data.numbs)-1; i++ {
		ranging = append(ranging, false)
	}

	//fmt.Println("starting ranging : ", ranging)

	result := data.numbs[0]

	for i := 0; i < len(ranging); i++ {
		if !ranging[i] {
			result += data.numbs[i+1]
		} else {
			result *= data.numbs[i+1]
		}
		//fmt.Println("result : ", result)
	}
	if result == data.value {
		//fmt.Println("found, result : ", result)
		return result
	}

	for !Finished(ranging) {

		result := data.numbs[0]

		ranging = NextRange(ranging)
		//fmt.Println("new ranging : ", ranging)

		for i := 0; i < len(ranging); i++ {
			if !ranging[i] {
				result += data.numbs[i+1]
			} else {
				result *= data.numbs[i+1]
			}
			//fmt.Println("result : ", result)
		}
		if result == data.value {
			//fmt.Println("found, result : ", result)
			return result
		}
	}

	return 0
}

func NextRange(ranging []bool) []bool {
	for i := len(ranging) - 1; i >= 0; i-- {
		if !ranging[i] {
			ranging[i] = true
			for j := i + 1; j < len(ranging); j++ {
				ranging[j] = false
			}
			return ranging
		}
	}
	return ranging
}

func Finished(ranging []bool) bool {
	for _, value := range ranging {
		if !value {
			return false
		}
	}
	return true
}

//####################################################################################################################################
//####################################################################################################################################
//########################################                  PART 2                   #################################################
//####################################################################################################################################
//####################################################################################################################################

func Part2(input []data) {
	var result int = 0

	for _, data := range input {
		fmt.Println(" Trying to solve : ", data.value)
		result += SolveData2(data)
	}

	fmt.Println("Resultat part 2 : ", result)
}

func SolveData2(data data) int {
	var ranging []int = []int{}

	for i := 0; i < len(data.numbs)-1; i++ {
		ranging = append(ranging, 0)
	}

	//fmt.Println("starting ranging : ", ranging)

	result := data.numbs[0]

	for i := 0; i < len(ranging); i++ {
		if ranging[i] == 0 {
			result += data.numbs[i+1]
		} else if ranging[i] == 1 {
			result *= data.numbs[i+1]
		} else if ranging[i] == 2 {
			result = Concat(result, data.numbs[i+1])
		}
		//fmt.Println("result : ", result)
	}
	FormatedPrint(result, ranging, data.numbs)
	if result == data.value {
		fmt.Println("found, result : ", result)
		//fmt.Println("found, result : ", result)
		return result
	}

	for !Finished2(ranging) {

		result = data.numbs[0]

		ranging = NextRange2(ranging)
		//fmt.Println("new ranging : ", ranging)

		for i := 0; i < len(ranging); i++ {
			if ranging[i] == 0 {
				result += data.numbs[i+1]
			} else if ranging[i] == 1 {
				result *= data.numbs[i+1]
			} else if ranging[i] == 2 {
				result = Concat(result, data.numbs[i+1])
			}
			//fmt.Println("result : ", result)
		}
		//FormatedPrint(result, ranging, data.numbs)
		if result == data.value {
			// FormatedPrint(result, ranging, data.numbs)
			fmt.Println("found, result : ", result)
			return result
		}
	}

	//FormatedPrint(result, ranging, data.numbs)
	fmt.Println("end", result)
	return 0
}

func NextRange2(ranging []int) []int {
	for i := len(ranging) - 1; i >= 0; i-- {
		if ranging[i] < 2 {
			ranging[i]++
			for j := i + 1; j < len(ranging); j++ {
				ranging[j] = 0
			}
			return ranging
		}
	}
	return ranging
}

func Finished2(ranging []int) bool {
	for _, value := range ranging {
		if value != 2 {
			return false
		}
	}
	return true
}

func Concat(a, b int) int {
	bah, _ := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	return bah
}

func FormatedPrint(result int, ranging []int, numbs []int) {
	fmt.Print("Result : ", result, " = ", numbs[0])
	for i := 0; i < len(ranging); i++ {
		if ranging[i] == 0 {
			fmt.Print(" + ", numbs[i+1])
		} else if ranging[i] == 1 {
			fmt.Print(" * ", numbs[i+1])
		} else if ranging[i] == 2 {
			fmt.Print(" || ", numbs[i+1])
		}
	}
	fmt.Println()
}
