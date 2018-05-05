package main

import (
	"encoding/json"
	"testing"
)

func TestCustomerSerialization(t *testing.T) {
	// tc1: smoke test, initing a new Customer programmatically
	var customer = &Customer{UserID: 99, Name: "Baran Kucukguzel", Latitude: 52.3706233, Longitude: 4.9284594}
	if customer == nil {
		t.Errorf("cannot init new Customer programatically")
	}

	// tc2: building a Customer from valid JSON
	customerJsonStr := `{"user_id":99,"name":"Baran Kucukguzel","latitude":"52.3706233","longitude":"4.9284594"}`
	customerFromJson, err := NewCustomerFromJSON(customerJsonStr)
	if err != nil {
		t.Errorf("failed unmarshalling valid json err: %v", err)
	}
	if customerFromJson == nil {
		t.Errorf("cannot init new customer from JSON")
	}
	if customer.UserID != customerFromJson.UserID {
		t.Errorf("UserID is not same for programmatic and json one")
	}
	if customer.Name != customerFromJson.Name {
		t.Errorf("Name is not same for programmatic and json one")
	}
	if customer.Latitude != customerFromJson.Latitude {
		t.Errorf("Latitude is not same for programmatic and json one")
	}
	if customer.Longitude != customerFromJson.Longitude {
		t.Errorf("Longitude is not same for programmatic and json one")
	}

	// tc3: marshaling a Customer
	customerMarshaledBytes, err := json.Marshal(customer)
	if err != nil {
		t.Errorf("cannot marshall customer instance %v", err)
	}
	customerMarshaslledString := string(customerMarshaledBytes)
	if customerMarshaslledString != customerJsonStr {
		t.Errorf("marshalling and programmatic customer doesn't match expected JSON")
	}

	// tc4: trying to build a Customer from invalid JSON
	invalidCustomerJsonStr := `{user_id":99,"name":"Baran Kucukguzel","latitude":52.3706233,"longitude":4.9284594}`
	customerFromJson, err = NewCustomerFromJSON(invalidCustomerJsonStr)
	if err == nil || customerFromJson != nil {
		t.Errorf("unmarshalling invalid json should fail")
	}
}

func TestCustomerLocation(t *testing.T) {
	// tc1: smoke test, initing a new Customer and location matches it's latitude and longitude
	var customer = &Customer{UserID: 99, Name: "Baran Kucukguzel", Latitude: 52.3706233, Longitude: 4.9284594}
	customerLocation := customer.Location()
	if customerLocation.LatitudeDegree != customer.Latitude {
		t.Errorf("customer location latitude is not correct")
	}
	if customerLocation.LongitudeDegree != customer.Longitude {
		t.Errorf("customer location longitude is not correct")
	}
}
