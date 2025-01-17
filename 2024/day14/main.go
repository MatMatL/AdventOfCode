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

	//fmt.Println(input)

	Part1(input)
	//Part2(input)
}

type robot struct {
	x  int
	y  int
	vx int
	vy int
}

func ReadFile(filename string) []robot {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result []robot = []robot{}

	for scanner.Scan() {
		tempo := robot{}
		parts := strings.Split(scanner.Text(), " ")
		pos, velocity := strings.Split(parts[0], ","), strings.Split(parts[1], ",")
		tempo.x, _ = strconv.Atoi(pos[0][2:])
		tempo.y, _ = strconv.Atoi(pos[1])
		tempo.vx, _ = strconv.Atoi(velocity[0][2:])
		tempo.vy, _ = strconv.Atoi(velocity[1])
		result = append(result, tempo)
	}

	return result
}

// ####################################################################################################################################
// ####################################################################################################################################
// ########################################                  PART 1                   #################################################
// ####################################################################################################################################
// ####################################################################################################################################

func Part1(input []robot) {
	var NorthEst, NorthWest, SouthEst, SouthWest int = 0, 0, 0, 0

	var mapWide, mapHeight int = 101, 103

	for t := 0; t < 100; t++ {
		for i := 0; i < len(input); i++ {
			//fmt.Println(input[i])
			input[i].x += input[i].vx
			input[i].y += input[i].vy
			//fmt.Println(input[i])
			if input[i].x >= mapWide {
				input[i].x -= mapWide
			} else if input[i].x < 0 {
				input[i].x += mapWide
			}
			if input[i].y >= mapHeight {
				input[i].y -= mapHeight
			} else if input[i].y < 0 {
				input[i].y += mapHeight
			}
			//fmt.Println(input[i], "\n")
		}
		PrintRobots(input, mapWide, mapHeight)
	}

	for _, robot := range input {
		if robot.x < int(mapWide/2) && robot.y < int(mapHeight/2) {
			NorthWest++
		} else if robot.x < int(mapWide/2) && robot.y > int(mapHeight/2) {
			SouthWest++
		} else if robot.x > int(mapWide/2) && robot.y < int(mapHeight/2) {
			NorthEst++
		} else if robot.x > int(mapWide/2) && robot.y > int(mapHeight/2) {
			SouthEst++
		}
	}

	fmt.Println(NorthEst, NorthWest, SouthEst, SouthWest)
	result := NorthEst * NorthWest * SouthEst * SouthWest

	fmt.Println("Resultat part 1 : ", result)
}

func PrintRobots(input []robot, mapWide int, mapHeight int) {
	var Map [][]int = make([][]int, mapHeight)
	for i := 0; i < mapHeight; i++ {
		Map[i] = make([]int, mapWide)
		for j := 0; j < mapWide; j++ {
			Map[i][j] = 0
		}
	}

	for _, robot := range input {
		Map[robot.y][robot.x]++
	}

	fmt.Println()
	fmt.Println()
	for i := 0; i < mapHeight; i++ {
		for j := 0; j < mapWide; j++ {
			fmt.Print(Map[i][j], " ")
		}
		fmt.Println()
	}
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
