package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	dir := "C:\\Users\\dsp_lab\\OneDrive - Florida International University\\subjects_data_processed\\1M\\MK\\"
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

		record := []string{strconv.FormatFloat(l, 'f', -1, 64), strconv.FormatFloat(r, 'f', -1, 64)}
		csvWriter.Write(record)
		csvWriter.Flush()
	}
}

func ReadCSV(csvLocation *os.File) (l float64, r float64) {
	reader := csv.NewReader((csvLocation))
	rows, err := reader.ReadAll()

	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	MK_time_left := rows[3][1]
	MK_time_right := rows[0][1]

	KF_time_left := rows[7][1]
	KF_time_right := rows[5][1]

	// last_left, _ := strconv.Atoi(MK_time_left)
	// first_left, _ := strconv.Atoi(MK_time_right)

	// last_right, _ := strconv.Atoi(KF_time_left)
	// first_right, _ := strconv.Atoi(KF_time_right)

	last_left, err := time.Parse("20060102150405", MK_time_left[:14])
	first_left, err := time.Parse("20060102150405", MK_time_right[:14])

	last_right, _ := time.Parse("20060102150405", KF_time_left[:14])
	first_right, _ := time.Parse("20060102150405", KF_time_right[:14])

	// if
	// da, err := time.Parse(time.StampMilli, MK_time_left)
	millisec_last_left, _ := strconv.Atoi(MK_time_left[14:])
	millisec_first_left, _ := strconv.Atoi(MK_time_right[14:])

	KF_millisec_last_left, _ := strconv.Atoi(KF_time_left[14:])
	KF_millisec_first_left, _ := strconv.Atoi(KF_time_right[14:])

	ll := last_left.Sub(first_left).Milliseconds() // + millisec_first_left - millisec_last_left
	rr := last_right.Sub(first_right).Milliseconds()

	fmt.Println(csvLocation.Name(), float64(ll)+float64(millisec_last_left-millisec_first_left))

	left := float64(ll*10) + float64(millisec_last_left-millisec_first_left)
	right := float64(rr*10) + float64(KF_millisec_last_left-KF_millisec_first_left)
	return left, right
}
