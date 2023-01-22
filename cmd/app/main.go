package main

import (
	"log"
	"os"
	"time"
	"yadroTest/internal/usecase/csvParser"
)

// dataPath - константа для передачи пути CSV-файла
const dataPath = "../../data/"

func main() {
	if len(os.Args) != 2 {
		log.Fatal("please add filename in subcommand")
	}
	filename := os.Args[1]

	start := time.Now()
	go func() {
		for {
			if time.Since(start) > time.Second*3 {
				log.Fatal("end by timeout, maybe self-pointer value in data, check you CSV-file")
			}
		}
	}()
	csvParser.HashTableOfCsvFile(dataPath + filename)
	csvParser.PrintResult()
}
