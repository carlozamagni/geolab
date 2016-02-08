package gpxToGeoJson


import(
	"encoding/xml"
	"log"
	"fmt"
)


// maybe using this:  http://kpawlik.github.io/geojson/
func ConvertToGeoJson(gpxStructure Gpx) (bool, error){

	segments := len(gpxStructure.Trk.Trkseg)
	points := len(gpxStructure.Trk.Trkseg[0].Trkpt)

	if(segments < 1 || points < 1){
		return true, nil
	}

	var coordinates [][3]float32

	for i, point := range gpxStructure.Trk.Trkseg[0].Trkpt{
		//fmt.Printf("%f - %f \n", point.Lat, point.Lon)
		fmt.Print(i)

		coordinates = append(coordinates, [3]float32{
			point.Lon,
			point.Lat,
			point.Ele,
		})
	}

	return true, nil
}

func ParseGpxFile(gpxData []byte) (Gpx, error) {
	// parsing xml:
	// https://www.socketloop.com/tutorials/golang-xml-to-json-example
	// https://github.com/revh/gpxjson/blob/master/gpxjson.go
	var g Gpx
	err := xml.Unmarshal(gpxData, &g)

	if err != nil{
		log.Fatal(err)
	}

	return g, err
}
/*
<?xml version='1.0' encoding='UTF-8'?>
<gpx xmlns:gpxdata="http://www.cluetrust.com/XML/GPXDATA/1/0" xmlns:gpxtpx="http://www.garmin.com/xmlschemas/TrackPointExtension/v1" xmlns:gpxext="http://www.garmin.com/xmlschemas/GpxExtensions/v3" xmlns="http://www.topografix.com/GPX/1/1" creator="tapiriik-sync">
  <metadata>
    <name>Corsa serale</name>
  </metadata>
  <trk>
    <name>Corsa serale</name>
    <trkseg>
      <trkpt lat="44.066427" lon="12.453816">
        <time>2014-04-18T17:31:15+00:00</time>
        <ele>37.3</ele>
      </trkpt>
      <trkpt lat="44.066499" lon="12.453673">
        <time>2014-04-18T17:31:17+00:00</time>
        <ele>37.6</ele>
      </trkpt>
      <trkpt lat="44.06649" lon="12.453497">
        <time>2014-04-18T17:31:18+00:00</time>
        <ele>37.5</ele>
      </trkpt>
*/