package main

import (
	"encoding/json"
	"fmt"
)

// Customer represents customer we have in the input spec.
type Customer struct {
	UserID    int32   `json:"user_id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude,string"`
	Longitude float64 `json:"longitude,string"`
}

// initalizes new Customer instances from given JSON string
func NewCustomerFromJSON(customerJSON string) (*Customer, error) {
	customer := &Customer{}
	err := json.Unmarshal([]byte(customerJSON), customer)

	if err != nil {
		return nil, err
	}
	return customer, nil
}

// returns customer's location as a Point
func (customer *Customer) Location() *Point {
	location := &Point{}

	location.LatitudeDegree = customer.Latitude
	location.LongitudeDegree = customer.Longitude

	return location
}

func (customer *Customer) Print() {
	fmt.Printf("Customer\tid: %v\tname: %v\n", customer.UserID, customer.Name)
}
