// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	tab := ReadFile("input.txt")

// 	Part2(tab)
// }

// func ReadFile(filename string) [][]int {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
// 		return nil
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)

// 	var tab [][]int = [][]int{}

// 	for scanner.Scan() {
// 		var tempTab []int = []int{}
// 		numbs := strings.Split(scanner.Text(), " ")

// 		for _, numb := range numbs {
// 			num, _ := strconv.Atoi(numb)
// 			tempTab = append(tempTab, num)
// 		}
// 		tab = append(tab, tempTab)
// 	}
// 	return tab
// }

// func PrintTab(tab [][]int) {
// 	for i := 0; i < len(tab); i++ {
// 		for j := 0; j < len(tab[i]); j++ {
// 			fmt.Print(tab[i][j], " ")
// 		}
// 		fmt.Println()
// 	}
// }

// //####################################################################################################################################
// //####################################################################################################################################
// //########################################                  PART 1                   #################################################
// //####################################################################################################################################
// //####################################################################################################################################

// func Part1(tab [][]int) {
// 	result := 0
// 	for j, i := range tab {
// 		fmt.Println(j, i)
// 		result += CheckPart1(i)
// 	}

// 	fmt.Println("Resultat part 1 : ", result)
// }

// func CheckPart1(tab []int) int {
// 	IsAscending := false

// 	if tab[0] < tab[1] {
// 		IsAscending = true
// 	}

// 	if IsAscending {
// 		for i := 0; i < len(tab)-1; i++ {
// 			if tab[i] > tab[i+1] {
// 				return 0
// 			}
// 			if tab[i+1]-tab[i] > 3 {
// 				return 0
// 			}
// 			if tab[i] == tab[i+1] {
// 				return 0
// 			}
// 		}
// 	} else {
// 		for i := 0; i < len(tab)-1; i++ {
// 			if tab[i] < tab[i+1] {
// 				return 0
// 			}
// 			if tab[i]-tab[i+1] > 3 {
// 				return 0
// 			}
// 			if tab[i] == tab[i+1] {
// 				return 0
// 			}
// 		}
// 	}
// 	return 1
// }

// //####################################################################################################################################
// //####################################################################################################################################
// //########################################                  PART 2                   #################################################
// //####################################################################################################################################
// //####################################################################################################################################

// func Part2(tab [][]int) {
// 	result := 0
// 	for _, i := range tab {
// 		//fmt.Println(j)
// 		result += CheckPart2(i)
// 	}

// 	fmt.Println("Resultat part 2 : ", result)
// }

// func CheckPart2(tab []int) int {
// 	// count mistakes
// 	mistake := 0

// 	//bool to know if it works
// 	AscendingWorks := true
// 	DescendingWorks := true

// 	// copy the tab
// 	tab1 := make([]int, len(tab))
// 	copy(tab1, tab)

// 	// check if possible in ascending order
// 	for i := 0; i < len(tab1)-1; i++ {
// 		//  if before sup to after    if diff more than 3         if same number
// 		if (tab1[i] > tab1[i+1]) || (tab1[i+1]-tab1[i] > 3) || (tab1[i] == tab1[i+1]) {
// 			// add a chance
// 			mistake++
// 			// remove the 2nd number
// 			tab1 = UpdateTabRemove(tab1, i+1)
// 			// recheck the number with the new next
// 			i--
// 		}
// 		// if 2 or more it's not possible
// 		if mistake > 1 {
// 			AscendingWorks = false
// 		}
// 	}

// 	// copy the tab
// 	tab2 := make([]int, len(tab))
// 	copy(tab2, tab)

// 	// reset mistake
// 	mistake = 0

// 	// check if possible in descending order
// 	for i := 0; i < len(tab2)-1; i++ {
// 		//  if before inf to after    if diff more than 3         if same number
// 		if (tab2[i] < tab2[i+1]) || (tab2[i]-tab2[i+1] > 3) || (tab2[i] == tab2[i+1]) {
// 			// add a chance
// 			mistake++
// 			// remove the 2nd number
// 			tab2 = UpdateTabRemove(tab2, i+1)
// 			// recheck the number with the new next
// 			i--
// 		}
// 		// if 2 or more it's not possible
// 		if mistake > 1 {
// 			DescendingWorks = false
// 		}
// 	}

// 	// try without the first number (without first because it's the one that can be removed)
// 	TryRemovingFirst := false
// 	if CheckPart1(tab[1:]) == 1 {
// 		TryRemovingFirst = true
// 	}

// 	if DescendingWorks || AscendingWorks || TryRemovingFirst {
// 		// if CheckPart1(tab) == 0 {
// 		// 	fmt.Println(tab)
// 		// 	fmt.Println("with ", DescendingWorks, tab2, AscendingWorks, tab1, TryRemovingFirst)
// 		// }
// 		return 1
// 	}
// 	if CheckDouble(tab) {
// 		fmt.Println(tab, "    \t\t     with ", DescendingWorks, tab2, "    \t\t    ", AscendingWorks, tab1, "     \t\t    ", TryRemovingFirst)
// 	}
// 	return 0
// }

// func UpdateTabRemove(tab []int, index int) []int {
// 	tab = append(tab[:index], tab[index+1:]...)
// 	return tab
// }

// func CheckDouble(tab []int) bool {
// 	for i := 0; i < len(tab)-1; i++ {
// 		if tab[i] == tab[i+1] {
// 			return false
// 		}
// 	}
// 	return true
// }

// //####################################################################################################################################
// //####################################################################################################################################
// //####################################################################################################################################
// //########################################                    V2                     #################################################
// //####################################################################################################################################
// //####################################################################################################################################
// //####################################################################################################################################

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	tab := ReadFile("input.txt")

	Part2(tab)
}

func ReadFile(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var tab [][]int = [][]int{}

	for scanner.Scan() {
		var tempTab []int = []int{}
		numbs := strings.Split(scanner.Text(), " ")

		for _, numb := range numbs {
			num, _ := strconv.Atoi(numb)
			tempTab = append(tempTab, num)
		}
		tab = append(tab, tempTab)
	}
	return tab
}

func PrintTab(tab [][]int) {
	for i := 0; i < len(tab); i++ {
		for j := 0; j < len(tab[i]); j++ {
			fmt.Print(tab[i][j], " ")
		}
		fmt.Println()
	}
}

//####################################################################################################################################
//####################################################################################################################################
//########################################                  PART 1                   #################################################
//####################################################################################################################################
//####################################################################################################################################

func Part1(tab [][]int) {
	result := 0
	for j, i := range tab {
		fmt.Println(j, i)
		result += CheckPart1(i)
	}

	fmt.Println("Resultat part 1 : ", result)
}

func CheckPart1(tab []int) int {
	IsAscending := false

	if tab[0] < tab[1] {
		IsAscending = true
	}

	if IsAscending {
		for i := 0; i < len(tab)-1; i++ {
			if tab[i] > tab[i+1] {
				return 0
			}
			if tab[i+1]-tab[i] > 3 {
				return 0
			}
			if tab[i] == tab[i+1] {
				return 0
			}
		}
	} else {
		for i := 0; i < len(tab)-1; i++ {
			if tab[i] < tab[i+1] {
				return 0
			}
			if tab[i]-tab[i+1] > 3 {
				return 0
			}
			if tab[i] == tab[i+1] {
				return 0
			}
		}
	}
	return 1
}

//####################################################################################################################################
//####################################################################################################################################
//########################################                  PART 2                   #################################################
//####################################################################################################################################
//####################################################################################################################################

func Part2(tab [][]int) {
	result := 0
	for _, i := range tab {
		result += CheckPart2(i)
	}

	fmt.Println("Resultat part 2 : ", result)
}

func CheckPart2(tab []int) int {
	for i := 0; i < len(tab); i++ {
		temp := make([]int, len(tab[:i]))
		copy(temp, tab[:i])
		temp = append(temp, tab[i+1:]...)
		fmt.Println("Trying : ", temp, " from ", tab, " returned ", CheckPart1(temp))
		if CheckPart1(temp) == 1 {
			return 1
		}
	}
	return 0
}
