package dataprovider

import (
	"bufio"
	"os"
	"strings"

	"employee-csv-parser/utils"
)

func RegisterCsv(employeeLine []string, fileStatus string) {

	csv, err := os.OpenFile(fileStatus+"Data.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	utils.PrintError(err)
	reader := bufio.NewReader(csv)
	line, _ := reader.ReadString('\n')
	if !strings.Contains(line, "Id,Full Name,E-mail,Salary") &&
		fileStatus == "valid" {
		csv.WriteString("Id,Full Name,E-mail,Salary\n")
	}

	csv.WriteString(strings.Join(employeeLine, ",") + "\n")

	csv.Close()

}
