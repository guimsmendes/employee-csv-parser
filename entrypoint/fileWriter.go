package entrypoint

import (
	"os"

	"github.com/guimsmendes/csv-parser/utils"
)

func registerCsv(fileName string, fileStatus string) {

	csv, err := os.OpenFile(fileStatus+"Csv.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	utils.PrintError(err)
	csv.WriteString(fileName + "is a/an " + fileStatus + " CSV file \n")
	csv.Close()

}
