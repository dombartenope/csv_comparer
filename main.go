package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/dombartenope/csv_comparer.git/promptui"
)

func main() {

	/* DYNAMICALLY FIND FILES */

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fileStore := []string{}
	for _, file := range files {
		if !file.IsDir() {
			if filepath.Ext(file.Name()) == ".csv" && file.Name() != "out.csv" {
				fileStore = append(fileStore, file.Name())
			}
		}
	}

	/*CSV INPUT*/
	if len(fileStore) > 2 {
		log.Fatalf("\nToo many input files\nPlease delete the extra file and try again")
		//TODO Allow for more inputs
	}
	input_one, err := os.Open(fileStore[0])
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer input_one.Close()

	input_two, err := os.Open(fileStore[1])
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer input_two.Close()

	/* SLICE FOR STORAGE */
	var slice_one []string
	var slice_two []string

	/*READERS*/
	reader_one := csv.NewReader(input_one)
	// reader_two := csv.NewReader(input_two)
	//Column names and index
	columnAndIndexOne := make(map[string]int)
	column_names_one, err := reader_one.Read()
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	reader_two := csv.NewReader(input_two)
	columnAndIndexTwo := make(map[string]int)
	column_names_two, err := reader_two.Read()
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	/* GET COLUMN NAMES */
	for i, v := range column_names_one {
		columnAndIndexOne[v] = i
	}
	for i, v := range column_names_two {
		columnAndIndexTwo[v] = i
	}

	/* PROMPT USER */
	inputOneColumnName, inputTwoColumnName := promptui.PromptUser(columnAndIndexOne, columnAndIndexTwo)

	/* READ ALL FROM BOTH FILES */
	rows, err := reader_one.ReadAll()
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	for _, v := range rows {
		if v[columnAndIndexOne[inputOneColumnName]] != "" {
			slice_one = append(slice_one, v[columnAndIndexOne[inputOneColumnName]])
		}
	}

	rows, err = reader_two.ReadAll()
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	for _, v := range rows {
		if v[columnAndIndexTwo[inputTwoColumnName]] != "" {
			slice_two = append(slice_two, v[columnAndIndexTwo[inputTwoColumnName]])
		}
	}

	out, err := os.Create("out.csv")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	writer := csv.NewWriter(out)
	defer writer.Flush()

	var count int
	for _, v := range slice_one {
		for _, j := range slice_two {
			if v == j {
				yes_slice := []string{v}
				writer.Write(yes_slice)
				count++
			}
		}
	}
	fmt.Println(count, "matches found")
}
