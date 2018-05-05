package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

// CLI flags
var (
	flagFileName        = flag.String("f", "input/customers.txt", "input file name")
	flagCircleRadius    = flag.Float64("r", 100, "radius of the circle we want to invite our customers")
	flagOfficeLatitude  = flag.Float64("lat", IntercomDublin.LatitudeDegree, "optional Latitude of our office")
	flagOfficeLongitude = flag.Float64("lon", IntercomDublin.LongitudeDegree, "optional Longitude of our office")
)

// reads customers from given file and returns
func readCustomers(inputFileName string) (customers []*Customer) {
	customers = make([]*Customer, 0)

	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		customer, err := NewCustomerFromJSON(line)
		if err != nil {
			log.Printf("cannot read customer ignoring, error :%v\n", err)
		} else {
			customers = append(customers, customer)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return customers
}

// Given customer list, our office location and radius finds invitee list
func getInvitedCustomerList(customers []*Customer, officeLocation *Point, radius float64) (invitedCustomers []*Customer) {
	invitedCustomers = make([]*Customer, 0)

	// helper functions decides whether customer is close to the office or not
	customerInRange := func(officeLoc, customerLoc *Point) bool {
		var distanceCalculator DistanceCalculator
		distanceCalculator = &DistanceCalculatorWithVincenty{}

		return distanceCalculator.GreatCircleDistanceInKm(officeLoc, customerLoc) < radius
	}

	// filter invited customers
	for _, customer := range customers {
		if customerInRange(officeLocation, customer.Location()) {
			invitedCustomers = append(invitedCustomers, customer)
		}
	}

	// sort invited customers
	sort.Slice(invitedCustomers[:], func(i, j int) bool {
		return invitedCustomers[i].UserID < invitedCustomers[j].UserID
	})

	return invitedCustomers
}

func main() {
	flag.Parse()

	// Get params from flags
	inputFileName := *flagFileName
	officeLoc := &Point{LatitudeDegree: *flagOfficeLatitude, LongitudeDegree: *flagOfficeLongitude}
	radius := *flagCircleRadius

	// Read customers from the file
	customers := readCustomers(inputFileName)
	// fmt.Println("***ALL CUSTOMERS***")
	// for _, customer := range customers {
	// 	customer.Print()
	// }

	// get list of invited customers
	invitedCustomers := getInvitedCustomerList(customers, officeLoc, radius)

	// print list of invited customers
	fmt.Println("***INVITED CUSTOMERS***")
	for _, customer := range invitedCustomers {
		customer.Print()
	}

}
