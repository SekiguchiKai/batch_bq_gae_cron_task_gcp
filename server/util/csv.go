package util

import (
	"encoding/csv"
	"os"
)

// ファイルを作成して、CSVに書き込みを行う
func WriteCsv(fileName string, target []string) error {
	// ファイルを作成する
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Write(target)
	writer.Flush()

	return nil
}
