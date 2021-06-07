package entrypoint

import (
	"bufio"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"

	"github.com/guimsmendes/csv-parser/usecase"
	"github.com/guimsmendes/csv-parser/utils"
)

func ProcessFiles() {

	files, err := ioutil.ReadDir("csv")
	utils.PrintError(err)
	for _, f := range files {
		employeeTable := parseTable(readFile(f))
		verification := usecase.VerifyTable(employeeTable)

		if verification {
			registerCsv(f.Name(), "valid")
		} else {
			registerCsv(f.Name(), "invalid")
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
		employeeTable[i] = make([]string, len(s))
		for j := 0; j < len(s); j++ {
			employeeTable[i] = append(employeeTable[i], strings.TrimSpace(s[j]))

		}

	}
	return employeeTable

}
