package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rules, lines := ReadFile("input.txt")

	Part2(rules, lines)
}

func ReadFile(filename string) ([][]int, [][]int) {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return nil, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rules [][]int = [][]int{}
	var partUnDone = false
	var lines [][]int = [][]int{}

	for scanner.Scan() {
		if scanner.Text() == "" {
			partUnDone = true
		} else {
			if !partUnDone {
				numbs := strings.Split(scanner.Text(), "|")
				num1, _ := strconv.Atoi(numbs[0])
				num2, _ := strconv.Atoi(numbs[1])

				rules = append(rules, []int{num1, num2})
			} else {
				var temp []int = []int{}
				numbs := strings.Split(scanner.Text(), ",")
				for _, numb := range numbs {
					num, _ := strconv.Atoi(numb)
					temp = append(temp, num)
				}
				lines = append(lines, temp)
			}
		}
	}

	return rules, lines
}

//####################################################################################################################################
//####################################################################################################################################
//########################################                  PART 1                   #################################################
//####################################################################################################################################
//####################################################################################################################################

func Part1(rules [][]int, lines [][]int) {
	var result int = 0

	for _, line := range lines {
		result += CheckLine(rules, line)
	}

	fmt.Println("Resultat part 1 : ", result)
}

func CheckLine(rules [][]int, line []int) int {
	for i, numb := range line {
		for _, rule := range rules {
			if rule[1] == numb {
				if CheckThen(line, i, rule[0]) {
					return 0
				}
			}
			if rule[0] == numb {
				if CheckBack(line, i, rule[1]) {
					return 0
				}
			}
		}
	}
	return line[len(line)/2]
}

func CheckThen(line []int, i int, numb int) bool {
	for j := i + 1; j < len(line)-1; j++ {
		if line[j] == numb {
			return true
		}
	}
	return false
}

func CheckBack(line []int, i int, numb int) bool {
	for j := i - 1; j >= 0; j-- {
		if line[j] == numb {
			return true
		}
	}
	return false
}

//####################################################################################################################################
//####################################################################################################################################
//########################################                  PART 2                   #################################################
//####################################################################################################################################
//####################################################################################################################################

func Part2(rules [][]int, lines [][]int) {
	var result int = 0

	for _, line := range lines {
		if CheckLine(rules, line) == 0 {
			fmt.Println("###### Resolving line : ", line)
			result += ResolvLine(rules, line)
		}
	}

	fmt.Println("Resultat part 2 : ", result)
}

func ResolvLine(rules [][]int, line []int) int {
	var result []int = []int{}

	for i := 0; i < len(line); i++ {
		result = Append(result, line[i])
		var done = false
		for j := 0; j < len(result)-1 && !done; j++ {
			fmt.Println("### Current result : ", result, " of line ", line)
			if CheckNumb(result[j], result[j+1], rules) {
				result[j], result[j+1] = result[j+1], result[j]
				fmt.Println("## changes result to : ", result)
			} else {
				done = true
			}
		}
	}

	fmt.Println("Resolved line : ", result)
	return result[len(result)/2]
}

func Append(array []int, numb int) []int {
	var result []int = []int{numb}
	for _, n := range array {
		result = append(result, n)
	}

	return result
}

func CheckNumb(numb1 int, numb2 int, rules [][]int) bool {
	fmt.Println("Check rule with : ", numb1, numb2)
	for i := 0; i < len(rules); i++ {
		if rules[i][0] == numb1 && rules[i][1] == numb2 {
			fmt.Println("Rule found : ", rules[i])
			return false
		}
	}

	fmt.Println(" # changing ")
	return true
}
