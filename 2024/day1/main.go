package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	tab1, tab2 := ReadFile("input.txt")

	Part1(tab1, tab2)
	Part2(tab1, tab2)

}

func ReadFile(filename string) ([]int, []int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return nil, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var tab1, tab2 []int = []int{}, []int{}

	for scanner.Scan() {
		numbs := strings.Split(scanner.Text(), "   ")

		num1, _ := strconv.Atoi(numbs[0])
		num2, _ := strconv.Atoi(numbs[1])

		tab1 = append(tab1, num1)
		tab2 = append(tab2, num2)
	}
	return tab1, tab2
}

//####################################################################################################################################
//####################################################################################################################################
//########################################                  PART 1                   #################################################
//####################################################################################################################################
//####################################################################################################################################

func Part1(tab1 []int, tab2 []int) {
	tab1, tab2 = Sort(tab1), Sort(tab2)

	for i := 0; i < len(tab1); i++ {
		fmt.Println(tab1[i], "  ", tab2[i])
	}

	result := FindDiff(tab1, tab2)

	fmt.Println("Resultat part 1 : ", result)
}

func Sort(tab []int) []int {
	for i := 0; i < len(tab)-1; i++ {
		if tab[i] > tab[i+1] {
			tab[i], tab[i+1] = tab[i+1], tab[i]
			i = 0
		}
	}
	for i := 0; i < len(tab)-1; i++ {
		if tab[i] > tab[i+1] {
			tab[i], tab[i+1] = tab[i+1], tab[i]
			i = 0
		}
	}
	for i := 0; i < len(tab)-1; i++ {
		if tab[i] > tab[i+1] {
			tab[i], tab[i+1] = tab[i+1], tab[i]
			i = 0
		}
	}
	for i := 0; i < len(tab)-1; i++ {
		if tab[i] > tab[i+1] {
			tab[i], tab[i+1] = tab[i+1], tab[i]
			i = 0
		}
	}
	return tab
}

func FindDiff(tab1 []int, tab2 []int) int {
	result := 0

	for i := 0; i < len(tab1); i++ {
		//println(tab1[i], "  ", tab2[i])
		if tab1[i] > tab2[i] {
			result += tab1[i] - tab2[i]
			//println("result : ", result)
		} else {
			result += tab2[i] - tab1[i]
			//println("result : ", result)
		}
	}

	return result
}

//####################################################################################################################################
//####################################################################################################################################
//########################################                  PART 2                   #################################################
//####################################################################################################################################
//####################################################################################################################################

func Part2(tab1 []int, tab2 []int) {
	result := 0
	for i := 0; i < len(tab1); i++ {
		result += tab1[i] * HowManyIn(tab1[i], tab2)
		println("Tempo result : ", result)
	}

	println("Resultat part 2 : ", result)
}

func HowManyIn(n int, tab []int) int {
	result := 0
	for i := 0; i < len(tab); i++ {
		if tab[i] == n {
			result++
		}
	}
	return result
}
