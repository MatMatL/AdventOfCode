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

	Part1(input)
}

func ReadFile(filename string) string {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result string

	for scanner.Scan() {
		result += scanner.Text()
	}
	fmt.Println(result)
	return result
}

//####################################################################################################################################
//####################################################################################################################################
//########################################                  PART 1                   #################################################
//####################################################################################################################################
//####################################################################################################################################

func Part1(input string) {
	mul := GetAllMul(input)

	result := calculateMul(mul)
	fmt.Println("Resultat part 1 : ", result)
}

func GetAllMul(input string) []string {
	var result []string = []string{}
	var current string = ""
	var isRecording bool = false
	var do bool = true

	for i := 0; i < len(input); i++ {
		fmt.Println(".", current, ".")
		if input[i] == 'd' {
			if input[i:i+4] == "do()" {
				do = true
			} else if input[i:i+7] == "don't()" {
				do = false
			}
		}
		if do {
			if input[i] == ')' && isRecording {
				isRecording = false
				result = append(result, current)
				current = ""
			}
			if isRecording && (input[i] < '0' || input[i] > '9') && (input[i] != ',') {
				current = ""
				isRecording = false
			}
			if input[i] == 'm' {
				if input[i:i+4] == "mul(" {
					i += 4
					isRecording = true
				}
			}

			if isRecording {
				current += string(input[i])
			}
		}
	}
	return result
}

func calculateMul(mul []string) int {
	result := 0
	for _, i := range mul {
		fmt.Println("spliting ", i)
		numbs := strings.Split(i, ",")
		if len(numbs) == 2 {
			num1, _ := strconv.Atoi(numbs[0])
			num2, _ := strconv.Atoi(numbs[1])
			result += num1 * num2
		}
	}
	return result
}

//####################################################################################################################################
//####################################################################################################################################
//########################################                  PART 2                   #################################################
//####################################################################################################################################
//####################################################################################################################################
