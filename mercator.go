package maps

import (
	"fmt"
	"math"
)

/*
MaxLatitude specified the maximum latitude (above or below the equator) that will be converted
into a PlotPoint, given the distortion of the Mercator projection close to the poles.

The ProjectionMeridian of -Pi ensures that the Greenwich meridian is in the center of the projection.
*/
const (
	MaxLatitude = 70
	ProjectionMeridian = -1 * math.Pi
)

/*
CoordinateToMercator converts a Coordinate into a PlotPoint and scales the result.

For scale = 1, values for x will range 0...2*Pi, and for y from (roughly) 0...4.
The equator will be at y=2.

A scale < 1 will be adjusted to 1.

Returns an error for latitudes higher than 70 degrees, as the Mercator projection
becomes practically unusable towards the poles.
*/
func CoordinateToMercator(c Coordinate, scale float64) (p PlotPoint, err error) {
	if c.isValid() != true {
		err = fmt.Errorf("Invalid coordinate.")
		return
	}

	if scale < 1 {
		scale = 1
	}
	// error for latitudes higher than 70 degrees
	if (c.Latitude > MaxLatitude) || (c.Latitude < MaxLatitude*-1) {
		err = fmt.Errorf("Mercator projection should not be used for latitudes higher than 70 degrees.")
		return
	}

	lat := radians(c.Latitude)
	lon := radians(c.Longitude)
	p.X = (lon - projectionMeridian) * scale
	p.Y = (math.Log(math.Tan(lat) + (1 / math.Cos(lat)))) * scale
	p.Y += scale * 2 // ensure all values are positive
	return
}

// Not yet implemented
func MercatorToCoordinate(p PlotPoint, scale float64) (c Coordinate, err error) {
	return
}
