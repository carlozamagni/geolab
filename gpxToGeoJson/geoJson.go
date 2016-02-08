package gpxToGeoJson


type LineString struct{
	Type string 		`json:"type"`
	Coordinates [][]float32 `json:"coordinates"`
}

func New() *LineString{
	return &LineString{
		Type: "LineString",
	}
}