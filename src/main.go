package main

import (
	"flag"
	"fmt"
	"log"
	"plans"
)

func main() {
	zipFlag := flag.String("zip", "-1", "a five-digit zip code")
	plansCsv := flag.String("plans", "plans.csv", "path to CSV containing Plan information")
	usersCsv := flag.String("users", "users.csv", "path to CSV containing User information")

	flag.Parse()

	zip := *zipFlag
	if zip == "-1" {
		log.Fatalln("Please specify a ZipCode when running!\n" +
			"usage: ./bin/plans -zip=12345")
	}

	loader := plans.NewDataLoader()
	plansMap := loader.BuildPlanMap(*plansCsv)
	usersMap := loader.BuildUserMap(*usersCsv)

	fmt.Printf("ZipCode: %s\n", zip)
	fmt.Println(plansMap.GetPlansFromZip(zip))
	fmt.Println(usersMap.GetUsersFromZip(zip))
}
