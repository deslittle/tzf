// CLI tool to preindex location shape.
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/deslittle/pinpoint/pb"
	"github.com/deslittle/pinpoint/preindex"
	"github.com/paulmach/orb/maptile"
	"google.golang.org/protobuf/proto"
)

var (
	idxZoom            = 13
	aggZoom            = 3
	maxZoomLevelToKeep = 10
	layerDrop          = 2
)

func main() {
	originalProbufPath := os.Args[1]
	rawFile, err := os.ReadFile(originalProbufPath)
	if err != nil {
		panic(err)
	}
	input := &pb.Locations{}
	if err := proto.Unmarshal(rawFile, input); err != nil {
		panic(err)
	}

	output := preindex.PreIndexLocations(input, maptile.Zoom(idxZoom), maptile.Zoom(aggZoom), maptile.Zoom(maxZoomLevelToKeep), layerDrop)

	// file := preindex.PreIndexLocationsToGeoJSON(output)
	// err = os.WriteFile("preindex_tiles.geojson", file, 0644)
	// if err != nil {
	// 	panic(err)
	// }

	outputPath := strings.Replace(originalProbufPath, ".pb", ".preindex.pb", 1)
	outputBin, _ := proto.Marshal(output)
	f, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	_, _ = f.Write(outputBin)
	fmt.Println(outputPath)
}
