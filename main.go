package main


import (
	"os"
	"fmt"
	"log"
	"strings"
	"io/ioutil"

	gpx "github.com/carlozamagni/geolab/gpxToGeoJson"
)

func parseGpx(basePath string, file os.FileInfo){
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
	}

	gpx.ConvertToGeoJson(parsed)

	fmt.Println(len(parsed.Trk.Trkseg[0].Trkpt))
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

	files, err := ioutil.ReadDir(basePath)
	if err != nil{
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	for _, file := range files{
		// read GPX
		if(strings.HasSuffix(strings.ToLower(file.Name()), ".gpx")){
			fmt.Println(file.Name())
			parseGpx(basePath, file);
		}
	}

	fmt.Printf("import completed\n")
}
