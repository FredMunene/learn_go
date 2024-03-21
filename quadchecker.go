package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func getPreviousOutput() string {
	input, _ := ioutil.ReadAll(os.Stdin) //reads data from stdin. _, handles errors
	str := string(input)                 //data read is converted to string. //takes symbols o--o
	return str
}
func execCommand(quad, x, y string) string { //creates a cmd structure that executes quad exec file with x,y
	cmd := exec.Command(quad, x, y)
	output, _ := cmd.Output()
	return string(output) //returns the output from quad file.. goes back to find same results
}
func getDimensions(str string) (int, int, bool) { //
	err := false
	lines := strings.Split(str, "\n") //splits input into rows
	y := len(lines) - 1               //takes no. of row
	if y == 0 {                       // if there is no row, returns error
		return 0, 0, true
	}
	for _, char := range lines[0] { //iteration over first row,
		if char == ' ' { //if character is empty, row is empty
			return 0, 0, true
		}
	}
	for i := 0; i < len(lines)-1; i++ { //
		if len(lines[0]) != len(lines[i]) { // check rows have same length>>no of values
			return 0, 0, true
		}
	}
	x := len(lines[0]) //determines x
	return x, y, err
}
func findSameResults(x int, y int, resultOfGiven string) []string { // checks quad with same pattern/results
	files := []string{"./quadA", "./quadB", "./quadC", "./quadD", "./quadE"} // initiates an array that records quad file executables
	sameQuads := []string{}                                                  //intiate array to store squads with similar results
	strX := strconv.Itoa(x)                                                  //integer x conversion to string
	strY := strconv.Itoa(y)                                                  //integer y conversion to string
	for _, elem := range files {                                             //iterate through files with executable quad files
		output := execCommand(elem, strX, strY) // calls execCommand func
		if resultOfGiven == output {            //if results are same, appends to that samequads
			sameQuads = append(sameQuads, elem[2:])
		}
	}
	return sameQuads
}
func quadChecker(previousOutput string) { //arg are input of figures from stdin
	x, y, err := getDimensions(previousOutput) // initiate variables of dimensions and error handling
	if !err {                                  //if there is no error
		sameQuads := findSameResults(x, y, previousOutput) //intiate variable that records same squads//calls find same results func
		for i, elem := range sameQuads {                   // iliterate through arry
			fmt.Printf("[%s] [%d] [%d]", elem, x, y) //prints
			if i != len(sameQuads)-1 {               //print || if it is not the last quad
				fmt.Printf(" || ")
			}
		}
		fmt.Printf("\n")
	} else {
		fmt.Printf("Not a quad function\n")
	}
}
func main() {
	previousOutput := getPreviousOutput() //get string of data from stdio
	if len(previousOutput) > 0 {          //if there is valid output, continue
		quadChecker(previousOutput) //initiate quadchecker func
	} else {
		fmt.Printf("Not a quad function\n")
	}
}
