package plans

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type dataLoader struct {
	// Any dependencies would go here
}

type DataLoader interface {
	BuildPlanMap(fileName string) PlanMap
	BuildUserMap(fileName string) UserMap
}

func NewDataLoader() DataLoader {
	return &dataLoader{}
}

func (r *dataLoader) BuildPlanMap(fileName string) PlanMap {
	csvFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Unable to load Plans data file: %v", err))
	}

	zipToPlanIdMap := PlanMap{}
	emptyStruct := struct{}{}

	count := 0
	reader := csv.NewReader(csvFile)
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}

		if len(line) != 2 {
			log.Printf("Invalid Plan at row [%d]: %v", count, line)
			continue
		}

		zips := strings.Split(line[1], ", ")
		for _, zip := range zips {
			if zipToPlanIdMap[zip] == nil {
				zipToPlanIdMap[zip] = map[string]struct{}{}
			}
			planId := line[0]
			zipToPlanIdMap[zip][planId] = emptyStruct
		}

		count++
	}

	return zipToPlanIdMap
}

func (r *dataLoader) BuildUserMap(fileName string) UserMap {
	csvFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Unable to load Users data file: %v", err))
	}

	zipToUserMap := UserMap{}

	count := 0
	reader := csv.NewReader(csvFile)
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}

		if len(line) != 3 {
			log.Printf("Invalid User at row [%d]: %v", count, line)
			continue
		}

		zip := GetZipFromAddress(line[2])
		zipToUserMap[zip] = append(zipToUserMap[zip], User{
			UserId:  line[0],
			Name:    line[1],
			Address: line[2],
			ZipCode: zip,
		})

		count++
	}

	return zipToUserMap
}
