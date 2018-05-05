package main

import (
	"testing"
)

func TestFloatEquals(t *testing.T) {
	// tc1: equal floats
	float1 := 0.00000001
	float2 := 0.00000000
	if !floatEquals(float1, float2) {
		t.Errorf("%v should be equal to %v", float1, float2)
	}

	// tc2: non-equal floats
	float1 = 0.0000001
	float2 = 0.0000000
	if floatEquals(float1, float2) {
		t.Errorf("%v should not be equal to %v", float1, float2)
	}
}

func TestDegreessToRadians(t *testing.T) {
	// tc1: lat to radians
	homeLatDegrees := 52.3706233
	homeExpectedLatRadians := 0.9140398079044
	homeLatRadians := degreesToRadians(homeLatDegrees)
	if !floatEquals(homeLatRadians, homeExpectedLatRadians) {
		t.Errorf("Failed to convert degress to radians got %v, expected %v", homeLatRadians, homeExpectedLatRadians)
	}

	// tc2: lon to radians
	homeLonDegrees := 4.9284594
	homeExpectedLonRadians := 0.0860178435807
	homeLonRadians := degreesToRadians(homeLonDegrees)
	if !floatEquals(homeLonRadians, homeExpectedLonRadians) {
		t.Errorf("Failed to convert degress to radians got %v, expected %v", homeLatRadians, homeExpectedLatRadians)
	}
}

func TestPointToRadians(t *testing.T) {
	// tc1: convert given Point to PointInRadians
	intercomDublinRadians := IntercomDublin.toPointRadians()
	intercomDublinExpectedRadians := &PointInRadians{LatitudeRadian: 0.930948639728, LongitudeRadian: -0.10921684028}
	if intercomDublinRadians == nil {
		t.Errorf("Failed to convert Point to PointInRadians")
	}
	if !floatEquals(intercomDublinRadians.LatitudeRadian, intercomDublinRadians.LatitudeRadian) {
		t.Errorf("Failed to convert degress to radians got %v, expected %v", intercomDublinRadians.LatitudeRadian, intercomDublinExpectedRadians.LatitudeRadian)
	}
	if !floatEquals(intercomDublinRadians.LongitudeRadian, intercomDublinRadians.LongitudeRadian) {
		t.Errorf("Failed to convert degress to radians got %v, expected %v", intercomDublinRadians.LongitudeRadian, intercomDublinExpectedRadians.LongitudeRadian)
	}
}

var distanceTests = []struct {
	point1             *Point
	point2             *Point
	expectedDistanceKM float64
}{
	{
		&Point{LatitudeDegree: 52.3706233, LongitudeDegree: 4.9284594}, // Amsterdam, Netherlands
		IntercomDublin, // Intercom Dublin, Ireland
		757.95284943639736,
	},
	{
		&Point{LatitudeDegree: 22.55, LongitudeDegree: 43.12},  // Rio de Janeiro, Brazil
		&Point{LatitudeDegree: 13.45, LongitudeDegree: 100.28}, // Bangkok, Thailand
		6094.544408786774,
	},
	{
		&Point{LatitudeDegree: 20.10, LongitudeDegree: 57.30}, // Port Louis, Mauritius
		&Point{LatitudeDegree: 0.57, LongitudeDegree: 100.21}, // Padang, Indonesia
		5145.525771394785,
	},
	{
		&Point{LatitudeDegree: 51.45, LongitudeDegree: 1.15},  // Oxford, United Kingdom
		&Point{LatitudeDegree: 41.54, LongitudeDegree: 12.27}, // Vatican, City Vatican City
		1389.1793118293067,
	},
	{
		&Point{LatitudeDegree: 22.34, LongitudeDegree: 17.05}, // Windhoek, Namibia
		&Point{LatitudeDegree: 51.56, LongitudeDegree: 4.29},  // Rotterdam, Netherlands
		3429.89310043882,
	},
	{
		&Point{LatitudeDegree: 63.24, LongitudeDegree: 56.59}, // Esperanza, Argentina
		&Point{LatitudeDegree: 8.50, LongitudeDegree: 13.14},  // Luanda, Angola
		6996.18595539861,
	},
	{
		&Point{LatitudeDegree: 90.00, LongitudeDegree: 0.00}, // North/South Poles
		&Point{LatitudeDegree: 48.51, LongitudeDegree: 2.21}, // Paris,  France
		4613.477506482742,
	},
	{
		&Point{LatitudeDegree: 45.04, LongitudeDegree: 7.42},  // Turin, Italy
		&Point{LatitudeDegree: 3.09, LongitudeDegree: 101.42}, // Kuala Lumpur, Malaysia
		10078.111954385415,
	},
}

// Given distance calculator test it with test cases defined above
func testGivenCalculator(t *testing.T, distanceCalculator DistanceCalculator) {
	for _, input := range distanceTests {
		km := distanceCalculator.GreatCircleDistanceInKm(input.point1, input.point2)
		if !floatEquals(input.expectedDistanceKM, km) {
			t.Errorf("points: %v, %v, want %v, got %v", input.point1, input.point2, input.expectedDistanceKM, km)
		}
	}
}

func TestSphericalLawofCosinesCalculator(t *testing.T) {
	testGivenCalculator(t, &DistanceCalculatorWithSphericalLawofCosines{})
}

func TestHaversineCalculator(t *testing.T) {
	testGivenCalculator(t, &DistanceCalculatorWithHaversine{})
}

func TestVincentyCalculator(t *testing.T) {
	testGivenCalculator(t, &DistanceCalculatorWithVincenty{})
}
