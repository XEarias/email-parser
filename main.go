package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type emailRecord []string

type emailRecords [][]string

func main() {

	file, err := os.Open("src/email_record.csv")
	if err != nil {
		fmt.Println("Error", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	record, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error", err)
	}

	include := flag.String("include", "", "Crear una lista de dominios unicos")
	exclude := flag.String("exclude", "", "Excluir un dominio de la lista")
	flag.Parse()

	if *include == "" && *exclude == "" {

		fmt.Println("Debes espificar una exclusion o una inclusion")
		return
	}

	var name string

	if *exclude != "" {
		name = "exclude"
	} else if *include != "" {
		name = "include"
	}

	var emailRecords emailRecords

	for _, value := range record {

		email := value[0]
		name := value[1]
		singleRecord := emailRecord{email, name}

		if *exclude != "" {
			if !strings.Contains(email, "@"+*exclude) {

				emailRecords = append(emailRecords, singleRecord)
			}

		} else if *include != "" {
			if strings.Contains(email, "@"+*include) {
				emailRecords = append(emailRecords, singleRecord)
			}
		}

	}

	newFile, err := os.Create("output/new_records_" + name + ".csv")

	if err != nil {
		fmt.Println(err)
	}
	defer newFile.Close()

	w := csv.NewWriter(newFile)
	defer w.Flush()

	w.WriteAll(emailRecords) // calls Flush internally

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}
