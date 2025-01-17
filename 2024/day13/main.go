package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := ReadFile("exemple.txt")

	fmt.Print(input)

	Part1(input)
	//Part2(input)
}

type Case struct {
	xA          int
	yA          int
	xB          int
	yB          int
	xP          int
	yP          int
	currentCost int
	currentA    int
	currentB    int
}

func ReadFile(filename string) []Case {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result []Case = []Case{}
	current := Case{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			result = append(result, current)
			current = Case{}
		} else {
			parts := strings.Split(scanner.Text(), ": ")
			coord := strings.Split(parts[1], ", ")
			if parts[0] == "Button A" {
				current.xA, _ = strconv.Atoi(coord[0][2:])
				current.yA, _ = strconv.Atoi(coord[1][2:])
			} else if parts[0] == "Button B" {
				current.xB, _ = strconv.Atoi(coord[0][2:])
				current.yB, _ = strconv.Atoi(coord[1][2:])
			} else if parts[0] == "Prize" {
				current.xP, _ = strconv.Atoi(coord[0][2:])
				current.yP, _ = strconv.Atoi(coord[1][2:])
			}
		}
	}

	result = append(result, current)

	return result
}

// ####################################################################################################################################
// ####################################################################################################################################
// ########################################                  PART 1                   #################################################
// ####################################################################################################################################
// ####################################################################################################################################

func Part1(input []Case) {
	var result int = 0

	for i, c := range input {
		fmt.Println("########## Solving for : ", c)
		result += Solve(i, c.xP, c.yP, c.xA, c.xB, c.yA, c.yB, 0, 0, 0)
		fmt.Println("########## EENNDD for : ", c)
	}

	fmt.Println("Resultat part 1 : ", result)
}

func Solve(i int, xP int, yP int, xA int, xB int, yA int, yB int, currentCost int, currentA int, currentB int) int {
	//fmt.Println(i, currentA, currentB, xP, yP, currentCost)
	//fmt.Println("Solving ", i, " with : ", xP, yP, xA, xB, yA, yB, currentCost, currentA, currentB)
	//fmt.Println(" ################# tac ", yP%yB, yP%yA, xP%xA, xP%xB)
	if yP == 0 && xP == 0 {
		fmt.Println("#### FOUND #### : ", currentCost, currentA, currentB)
		return currentCost
	}

	if yP < 0 || xP < 0 {
		//fmt.Println("#### DEAD #### : ")
		return -1
	}

	if currentA > 100 || currentB > 100 {
		return -1
	}

	case1 := Solve(i, xP-xA, yP-yA, xA, xB, yA, yB, currentCost+3, currentA+1, currentB)
	case2 := Solve(i, xP-xB, yP-yB, xA, xB, yA, yB, currentCost+1, currentA, currentB+1)

	if (case1 < case2 && case1 != -1) || (case2 == -1) {
		return case1
	} else if case2 != -1 {
		return case2
	}

	//fmt.Println("#### DEAD #### : ")
	return -1
}

//####################################################################################################################################
//####################################################################################################################################
//########################################                  PART 2                   #################################################
//####################################################################################################################################
//####################################################################################################################################

func Part2(input [][]rune) {
	var result int = 0

	fmt.Println("Resultat part 2 : ", result)
}
