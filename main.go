package main


import (
	"os"
	"fmt"
	"log"
	"time"
	"sync"
	"strings"
	"io/ioutil"

	db "github.com/carlozamagni/geolab/storage"
	geojson "github.com/carlozamagni/geolab/geojson"
	gpx "github.com/carlozamagni/geolab/gpxToGeoJson"
)

type GeoTrack struct{
	Name string			`json:"name"`
	Path geojson.LineString		`json:"path"`
}

func parseGpx(basePath string, file os.FileInfo, wg *sync.WaitGroup) {

	fullFileName := strings.Join([]string{basePath, file.Name()}, "/")

	if(! strings.HasSuffix(strings.ToLower(file.Name()), ".gpx")){
		return
	}

	f, err := ioutil.ReadFile(fullFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	parsed, err := gpx.ParseGpxFile(f)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	trackName := strings.Split(strings.ToLower(file.Name()), ".gpx")[0]
	lineString, err := gpx.CreateLineString(parsed)

	if err == nil {

		geoTrack := GeoTrack{Name:trackName, Path:lineString}

		mongo := db.MongoSession("local")
		geoDataCollection := mongo.DB("geolab").C("paths")
		geoDataCollection.Insert(geoTrack)
	}

	wg.Done()
}


func checkParams()(string){
	arguments := os.Args[1:]
	if(len(arguments) < 1){
		fmt.Println("Use as follows:")
		fmt.Println("$ geolab <gpx source dir>")
		os.Exit(0)
	}

	return os.Args[1]
}


func main() {

	basePath := checkParams()
	fmt.Printf("\n search for GPX tracks in %s \n", basePath)

	files, err := ioutil.ReadDir(basePath)
	if err != nil{
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	var wg sync.WaitGroup
	processedFiles := 0

	startTime := time.Now()

	for _, file := range files{
		// read GPX
		if(strings.HasSuffix(strings.ToLower(file.Name()), ".gpx")){
			//fmt.Println(file.Name())
			wg.Add(1)
			processedFiles ++
			go parseGpx(basePath, file, &wg);
		}
	}

	wg.Wait()

	endTime := time.Now()

	fmt.Printf("process completed \n")
	fmt.Printf("converted and imported %d GPX tracks \n", processedFiles)

	fmt.Printf("execution time: %d sec.", (endTime.Sub(startTime) / 1000000000))
}
