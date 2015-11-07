package maps

import (
	"math"
	"fmt"
)

// Coordinate is a point on the earth, specified in Latitude and Longitude, 
// where positive values denote North and East, and negative values South and West.
type Coordinate struct {
	Latitude float64
	Longitude float64
}

// R is the radius of the earth, in kilometers, for distance calculations
const (
	R = 6367
)

// String implements the Stringer interface for Coordinate. It generates a string in the format
// 00.00 North, 00.00 East
func (c Coordinate) String() string {
	var northsouth string
	var eastwest string
	var lat float64
	var long float64
	if c.Latitude < 0 {
		northsouth = "South"
		lat = c.Latitude*-1
	} else {
		northsouth = "North"
		lat = c.Latitude
	}
	if c.Longitude < 0 {
		eastwest = "West"
		long = c.Longitude*-1
	} else {
		eastwest = "East"
		long = c.Longitude
	}
	return fmt.Sprintf("%.2f %s, %.2f %s",lat,northsouth,long,eastwest)
}

// radians is a helper function to convert from degrees to radians
func radians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// degrees is a helper function to convert from radians to degrees
func degrees(radians float64) float64 {
	return radians * 180 / math.Pi
}

// Distance calculates the distance between two _Coordinate_s on a great circle.
// It returns a float64 understood to be in kilometers.
func Distance (origin Coordinate, destination Coordinate) float64 {
	originLat := radians(origin.Latitude)
	originLon := radians(origin.Longitude)
	destLat := radians(destination.Latitude)
	destLon := radians(destination.Longitude)
	deltaLat := destLat - originLat
	deltaLon := destLon - originLon

	a := math.Pow(math.Sin(deltaLat/2), 2) + math.Cos(originLat)*math.Cos(destLat)*math.Pow(math.Sin(deltaLon/2), 2)
	c := 2 * math.Asin(math.Sqrt(a))
	return c * R
}

// PointOnGreatCircle returns the coordinates of a point that is _distance_ (in kilometers) from the origin.
// if the distance provided is larger than the actual distance, it returns the coordinates of the destination.
func PointOnGreatCircle (origin Coordinate, destination Coordinate, distance float64) (waypoint Coordinate) {

	totalDistance := Distance(origin, destination)

	// if the distance provided by the user is larger than the actual distance, return the destination coordinates
	if distance > totalDistance {
		waypoint.Latitude=destination.Latitude
		waypoint.Longitude=destination.Longitude
 		return
	}

	// from http://williams.best.vwh.net/avform.htm#Intermediate
	// this works, but I'll happily admit I don't understand the math
	fraction := distance / totalDistance
	originLatitude := radians(origin.Latitude)
	originLongitude := radians(origin.Longitude * -1) // quirk of the formula
	destinationLatitude := radians(destination.Latitude)
	destinationLongitude := radians(destination.Longitude * -1) // quirk of the formula

	d := math.Acos(math.Sin(originLatitude)*math.Sin(destinationLatitude) + math.Cos(originLatitude)*math.Cos(destinationLatitude)*math.Cos(originLatitude-destinationLatitude))
	A := math.Sin((1-fraction)*d) / math.Sin(d)
	B := math.Sin(fraction*d) / math.Sin(d)
	x := A*math.Cos(originLatitude)*math.Cos(originLongitude) + B*math.Cos(destinationLatitude)*math.Cos(destinationLongitude)
	y := A*math.Cos(originLatitude)*math.Sin(originLongitude) + B*math.Cos(destinationLatitude)*math.Sin(destinationLongitude)
	z := A*math.Sin(originLatitude) + B*math.Sin(destinationLatitude)
	waypoint.Latitude = degrees(math.Atan2(z, math.Sqrt(x*x+y*y)))
	waypoint.Longitude = degrees(-1 * math.Atan2(y, x))
	return
}
