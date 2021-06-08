package entrypoint

import (
	"bufio"
	"employee-csv-parser/configuration"
	"employee-csv-parser/usecase"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"

	"employee-csv-parser/utils"
)

func ProcessFiles() {

	fmt.Println("Starting the .csv files import.")

	files, err := ioutil.ReadDir("csv")
	utils.PrintError(err)
	fmt.Println("Files successfully imported.")
	configuration.DbConnect()
	for _, f := range files {
		fmt.Println("Starting " + f.Name() + " file validation...")
		employeeTable := parseTable(readFile(f))
		verifier := usecase.VerifyTable(employeeTable)

		if !verifier {
			fmt.Println(f.Name() + " file didn't match the minimal required column validations.")
		} else {
			fmt.Println("Csv parsing job finished with success for the file: " + f.Name())
		}
	}

}

func readFile(f fs.FileInfo) []string {

	file, err := os.Open("csv/" + f.Name())
	utils.PrintError(err)
	lines := readLines(file)
	file.Close()
	return lines

}

func readLines(file *os.File) []string {
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func parseTable(lines []string) [][]string {

	employeeTable := make([][]string, len(lines))
	for i := 0; i < len(lines); i++ {
		s := strings.Split(lines[i], ",")
		employeeTable[i] = make([]string, 0)
		for j := 0; j < len(s); j++ {
			employeeTable[i] = append(employeeTable[i], strings.TrimSpace(s[j]))
		}

	}
	return employeeTable

}
