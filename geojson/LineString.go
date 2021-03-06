package geojson

import (
	"bytes"
	"fmt"
	"log"
)

type LineString struct {
	Type        string      `json:"type"`
	Coordinates [][]float32 `json:"coordinates"`
}

func NewLineString() *LineString {
	return &LineString{
		Type: "LineString",
	}
}

func ToString(ls LineString) string {
	var buffer bytes.Buffer

	for _, p := range ls.Coordinates {
		pointString := fmt.Sprintf("[%g,%g,%g]", p[0], p[1], p[2])
		_, err := buffer.WriteString(pointString)
		if err != nil {
			log.Fatal(err)
		}
	}

	return fmt.Sprintf("{ 'type':%s, 'coordinates':[%s] ", ls.Type, buffer.String())

}

// re define using: http://mholt.github.io/json-to-go/
