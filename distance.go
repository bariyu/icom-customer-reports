package main

import (
	"math"
)

// Point is a geographic coordinate in degrees.
type Point struct {
	LatitudeDegree  float64
	LongitudeDegree float64
}

// PointInRadians is a geographic coordinate in radians.
type PointInRadians struct {
	LatitudeRadian  float64
	LongitudeRadian float64
}

// DistanceCalcuator defines an interface that calculates distance between points.
type DistanceCalculator interface {
	// Given the points calculate distance between them in km
	GreatCircleDistanceInKm(point1, point2 *Point) float64
}

// converts given degrees to radians
func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// given to float number checks they are close enough
func floatEquals(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}

// converts given point to PointInRadians
func (p *Point) toPointRadians() *PointInRadians {
	pointInRadians := &PointInRadians{}

	pointInRadians.LatitudeRadian = degreesToRadians(p.LatitudeDegree)
	pointInRadians.LongitudeRadian = degreesToRadians(p.LongitudeDegree)

	return pointInRadians
}

// DistanceCalculatorWithSphericalLawofCosines implements formula one in wikipage: https://en.wikipedia.org/wiki/Great-circle_distance which uses spherical law of cosines
type DistanceCalculatorWithSphericalLawofCosines struct{}

func (f *DistanceCalculatorWithSphericalLawofCosines) GreatCircleDistanceInKm(point1, point2 *Point) float64 {
	point1InRadians := point1.toPointRadians()
	point2InRadians := point2.toPointRadians()

	deltaAlpha := math.Abs(point1InRadians.LongitudeRadian - point2InRadians.LongitudeRadian)
	deltaSigma := math.Acos(math.Sin(point1InRadians.LatitudeRadian)*math.Sin(point2InRadians.LatitudeRadian) + math.Cos(point1InRadians.LatitudeRadian)*math.Cos(point2InRadians.LatitudeRadian)*math.Cos(deltaAlpha))

	return earthRadiusKm * deltaSigma
}

// DistanceCalculatorWithHaversine implements formula one in wikipage: https://en.wikipedia.org/wiki/Great-circle_distance which uses haversine formula
type DistanceCalculatorWithHaversine struct{}

func (f *DistanceCalculatorWithHaversine) GreatCircleDistanceInKm(point1, point2 *Point) float64 {
	point1InRadians := point1.toPointRadians()
	point2InRadians := point2.toPointRadians()

	deltaTheta := math.Abs(point1InRadians.LatitudeRadian - point2InRadians.LatitudeRadian)
	deltaAlpha := math.Abs(point1InRadians.LongitudeRadian - point2InRadians.LongitudeRadian)

	inside := math.Pow(math.Sin(deltaTheta/2.0), 2) + math.Cos(point1InRadians.LatitudeRadian)*math.Cos(point2InRadians.LatitudeRadian)*math.Pow(math.Sin(deltaAlpha/2.0), 2)
	deltaSigma := 2 * math.Asin(math.Sqrt(inside))

	return earthRadiusKm * deltaSigma
}

// DistanceCalculatorWithHaversine implements formula one in wikipage: https://en.wikipedia.org/wiki/Great-circle_distance which uses vincenty formula
type DistanceCalculatorWithVincenty struct{}

func (f *DistanceCalculatorWithVincenty) GreatCircleDistanceInKm(point1, point2 *Point) float64 {
	point1InRadians := point1.toPointRadians()
	point2InRadians := point2.toPointRadians()

	// deltaTheta := math.Abs(point1InRadians.LatitudeRadian - point2InRadians.LatitudeRadian)
	deltaAlpha := math.Abs(point1InRadians.LongitudeRadian - point2InRadians.LongitudeRadian)

	nominatorFirstPart := math.Pow(math.Cos(point2InRadians.LatitudeRadian)*math.Sin(deltaAlpha), 2)
	nominatorSecondPart := math.Pow(math.Cos(point1InRadians.LatitudeRadian)*math.Sin(point2InRadians.LatitudeRadian)-math.Sin(point1InRadians.LatitudeRadian)*math.Cos(point2InRadians.LatitudeRadian)*math.Cos(deltaAlpha), 2)
	nominator := math.Sqrt(nominatorFirstPart + nominatorSecondPart)

	denomiatorFirstPart := math.Sin(point1InRadians.LatitudeRadian) * math.Sin(point2InRadians.LatitudeRadian)
	denomiatorSecondPart := math.Cos(point1InRadians.LatitudeRadian) * math.Cos(point2InRadians.LatitudeRadian) * math.Cos(deltaAlpha)
	denominator := denomiatorFirstPart + denomiatorSecondPart

	deltaSigma := math.Atan2(nominator, denominator)

	return earthRadiusKm * deltaSigma
}
