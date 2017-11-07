/*first GO Program CSV reader*/
package main

import "fmt"
import "io"
import "os"
import "bufio"

var firstname string
var lastname string
var filename string


func main() {

	fmt.Println("Please enter your full name:")
	fmt.Scanln(&firstname, &lastname)
	fmt.Printf("Hi %s %s", firstname, lastname)

	fmt.Println("Enter file name:")
	fmt.Scanln(&filename)
	inputFile, inputError := os.Open(filename)

	if inputError != nil {
		fmt.Printf("An Error Occurred opening the file")
		return
		}

	inputReader := bufio.NewReader(inputFile)

	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			return
		}

		fmt.Printf("the input was: %s", inputString)

	}

}
