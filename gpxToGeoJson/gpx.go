package gpxToGeoJson

type(
	Trkpt struct{
		Lat float32	`xml:"lat,attr"`
		Lon float32	`xml:"lon,attr"`
		Time string 	`xml:"time"`
		Ele float32 	`xml:"ele"`
	}

	Trkseg struct {
		Trkpt []Trkpt	`xml:"trkpt"`
	}

	Trk struct {
		Name string	`xml:"name"`
		Trkseg []Trkseg `xml:"trkseg"`
	}

	Gpx struct{
		Trk *Trk `xml:"trk"`
	}
)