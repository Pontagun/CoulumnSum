package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func main() {
	dir := ""
	files, err := os.ReadDir(dir)

	if err != nil {
		log.Fatal("Error while reading the directory", err)
	}

	for _, f := range files {
		file, err := os.Open(dir + f.Name())

		if err != nil {
			log.Fatal("Error while reading the file", err)
		}

		// l is a total time of the left boxes in the experiment
		// r is a total time of the right boxes in the experiment
		l, r := ReadCSV(file)

		defer file.Close()

		csvFile, _ := os.Create(dir + "diff_" + f.Name())
		defer csvFile.Close()
		csvWriter := csv.NewWriter(csvFile)

		record := []string{strconv.Itoa(l), strconv.Itoa(r)}
		csvWriter.Write(record)
		csvWriter.Flush()
	}
}

func ReadCSV(csvLocation *os.File) (l int, r int) {
	reader := csv.NewReader((csvLocation))
	rows, err := reader.ReadAll()

	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	last_left, _ := strconv.Atoi(rows[3][1])
	first_left, _ := strconv.Atoi(rows[1][1])

	last_right, _ := strconv.Atoi(rows[7][1])
	first_right, _ := strconv.Atoi(rows[5][1])

	l = last_left - first_left
	r = last_right - first_right

	return l, r
}
