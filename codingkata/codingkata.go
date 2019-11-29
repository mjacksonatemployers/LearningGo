package codingkata

import (
	"fmt"
	"strings"
)

/*
	Print Smythe kata result
*/
func PrintSmytheExercise() {

	var counter int

	for i := 0; i <= 22222; i++ {
		int2Str := PadNumberWithZero(i)
		if !strings.ContainsAny(int2Str, "3456789") {
			if int2Str[0:2] != "00" && int2Str[1:3] != "00" && int2Str[2:4] != "00" && int2Str[3:5] != "00" {
				counter++
			}
		}
	}
	println("Count: ", counter)
}

//Pad the with zeros number
func PadNumberWithZero(value int) string {
	return fmt.Sprintf("%05d", value)
}

//Print the spiral solution
func PrintGridClockwiseSpiralTraversal(gridXY int) {

	// Create the grid
	intGrid := make([][]int, gridXY)
	for i := 0; i < gridXY; i++ {
		intGrid[i] = make([]int, gridXY)
	}

	// Fill the grid
	total :=1
	for i:=0;i<gridXY;i++ {
		for j:=0;j<gridXY;j++ {
			intGrid[i][j] = total
			total++
		}
	}

	// Print the Grid
	fmt.Println("Grid:")
	for i :=0; i<len(intGrid[0]);i++ {
		fmt.Printf("%-2v\n", intGrid[i])
	}

	//Initialize
	var trow = 0
	var brow = len(intGrid) - 1
	var lcol = 0
	var rcol = len(intGrid[0]) - 1
	var i = 0

	fmt.Println("Result:")

	// While true
	for {

		if lcol > rcol {
			break
		}

		// Traverse top row
		for i := lcol; i <= rcol; i++ {
			fmt.Print(intGrid[trow][i])
			fmt.Print(" ")
		}

		trow++
		if trow > brow {
			break
		}

		// Traverse right row
		for i := trow; i <= brow; i++ {
			fmt.Print(intGrid[i][rcol])
			fmt.Print(" ")
		}

		rcol--
		if lcol > rcol {
			break
		}

		// Traverse bottom row
		for i = rcol; i >= lcol; i-- {
			fmt.Print(intGrid[brow][i])
			fmt.Print(" ")
		}

		brow--
		if trow > brow {
			break
		}

		// Traverse left row
		for i = brow; i >= trow; i-- {
			fmt.Print(intGrid[i][lcol])
			fmt.Print(" ")
		}

		lcol++
	}

	//new line
	fmt.Println("")
}
