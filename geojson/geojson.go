package geojson


/*
Point coordinates are in x, y order (easting, northing for projected coordinates,
longitude, latitude for geographic coordinates)
 */
type Point struct {
	Type string 		`json:"type"`
	Coordinates []float64	`json:"coordinates"`
}

func NewPoint() *Point{
	return &Point{
		Type: "Point",
	}
}

/*
Coordinates of LineString are an array of positions.
A position is represented by an array of numbers. There must be at least two elements,
and may be more. The order of elements must follow x, y, z order (easting, northing,
altitude for coordinates in a projected coordinate reference system, or longitude,
latitude, altitude for coordinates in a geographic coordinate reference system).
Any number of additional elements are allowed.
 */

/*
type LineString struct {
	Type string 		`json:"type"`
	Coordinates []struct {
		Lon float32 	`json:"0"`
		Lat float32 	`json:"1"`
		Ele float32	`json:"2"`
	} 			`json:"coordinates"`
}

func NewLineString() *LineString{
	return &LineString{
		Type: "LineString",
	}
}
*/

/*
For type "MultiLineString", the "coordinates" member must be
an array of LineString coordinate arrays.
 */
type MultiLineString struct {
	Type string 		`json:"type"`
	Coordinates []struct {
		Lon []float64	`json:"0"`
		Lat []float64	`json:"1"`
	} 			`json:"coordinates"`
}

func NewMultiLineString() *MultiLineString{
	return &MultiLineString{
		Type: "MultiLineString",
	}
}