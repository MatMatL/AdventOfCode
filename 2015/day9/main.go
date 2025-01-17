package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := ReadFile("input2.txt")

	fmt.Print(input)

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
		for _, numb := range scanner.Text() {
			num, _ := strconv.Atoi(string(numb))
			result = append(result, num)
		}
	}

	return result
}

//####################################################################################################################################
//####################################################################################################################################
//########################################                  PART 1                   #################################################
//####################################################################################################################################
//####################################################################################################################################

func Part1(input []int) {
	fmt.Println(input)

	DiskData := ToData(input)

	//fmt.Println(DiskData)

	OrderedDisk := OrderDisk(DiskData)

	//fmt.Println(OrderedDisk)

	result := CalCulSomme(OrderedDisk)

	fmt.Println("Resultat part 1 : ", result)
}

func ToData(input []int) []int {
	var result []int = []int{}
	var FileNumber int = 0
	var IsBlank bool = false

	for i := 0; i < len(input); i++ {
		if IsBlank {
			for j := 0; j < input[i]; j++ {
				result = append(result, -1)
			}
			IsBlank = false
		} else {
			for j := 0; j < input[i]; j++ {
				result = append(result, FileNumber)
			}
			FileNumber++
			IsBlank = true
		}
	}

	return result
}

func OrderDisk(DiskData []int) []int {
	for i := len(DiskData) - 1; i >= 0; i-- {
		if DiskData[i] != -1 {
			for j := 0; j <= i; j++ {
				if DiskData[j] == -1 {
					DiskData[i], DiskData[j] = DiskData[j], DiskData[i]
					break
				}
			}
		}
	}

	return DiskData
}

func CalCulSomme(OrderedDisk []int) int {
	result := 0

	for i := 0; i < len(OrderedDisk); i++ {
		if OrderedDisk[i] != -1 {
			result += OrderedDisk[i] * i
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
	fmt.Println(input)

	DiskData := ToData(input)

	//fmt.Println(DiskData)

	OrderedDisk := OrderDisk2(DiskData)

	//fmt.Println(OrderedDisk)

	result := CalCulSomme(OrderedDisk)

	fmt.Println("Resultat part 2 : ", result)
}

func OrderDisk2(DiskData []int) []int {
	var ActualFileLenth int = 0
	var Recording bool = false

	for i := len(DiskData) - 1; i >= 0; i-- {
		//fmt.Println("Have ", DiskData[i], " at ", i, " with ", ActualFileLenth, " and ", Recording)
		if DiskData[i] != -1 && !Recording {
			Recording = true
			ActualFileLenth++
		} else if DiskData[i] != DiskData[i+1] && Recording {
			var ActualBlankLenth int = 0
			for j := 0; j < i+1; j++ {
				//fmt.Println("found numb ", DiskData[i+1], "of lenth ", ActualFileLenth, " surching ", DiskData[j], " at ", j, " with ", ActualBlankLenth)
				if DiskData[j] == -1 {
					ActualBlankLenth++
				} else {
					ActualBlankLenth = 0
				}
				//fmt.Println("found numb ", DiskData[i+1], "of lenth ", ActualFileLenth, " surching ", DiskData[j], " at ", j, " with ", ActualBlankLenth)
				if ActualBlankLenth == ActualFileLenth {
					//fmt.Println("found place for ", DiskData[i+1], " at ", j)
					//fmt.Println(DiskData)
					for k := 0; k < ActualFileLenth; k++ {
						DiskData[i+ActualFileLenth-k], DiskData[j-k] = DiskData[j-k], DiskData[i+ActualFileLenth-k]
						//fmt.Println(DiskData)
					}
					break
				}
			}
			if DiskData[i] == -1 {
				Recording = false
				ActualFileLenth = 0
			} else {
				ActualFileLenth = 1
			}
		} else if Recording && DiskData[i] != -1 {
			ActualFileLenth++
		} else {
			//fmt.Println("prout")
		}
	}
	return DiskData
}
