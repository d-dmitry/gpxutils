package cmd

import (
	"fmt"
	"github.com/tkrajina/gpxgo/gpx"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

func LoadGPXs(filesDir, regex string) []string {
	gpxes := make([]string, 0)
	dirs, err := ioutil.ReadDir(filesDir)
	if err != nil {
		log.Fatal(err)
	}
	
	for _, fileInfo := range dirs {
		match, err := regexp.MatchString(regex, fileInfo.Name())

		if err != nil {
			panic("Your regex is faulty")
		}

		if strings.HasSuffix(fileInfo.Name(), ".gpx") && match {
			gpxes = append(gpxes, fmt.Sprintf("%s/%s", filesDir, fileInfo.Name()))
		}
	}
	
	if len(gpxes) == 0 {
		panic("No GPX files found in dir")
	}

	log.Printf("Found %d gpx files\n", len(gpxes))
	return gpxes
}

func MergeGpx() *gpx.GPX {
	NewGPX := new(gpx.GPX)
	for _, gpxfile := range LoadGPXs(dirMerge, regexMerge) {
		gpxFile, err := gpx.ParseFile(gpxfile)
		
		filename := filepath.Base(gpxfile)
		
		if err != nil {
			log.Fatal(err)
		}

		if gpxFile.Waypoints != nil {
			log.Printf("Found %d waypoint(s) in %s\n", len(gpxFile.Waypoints), filename)
			for _, point := range gpxFile.Waypoints {
				NewGPX.AppendWaypoint(&point)
			}
		}else{log.Printf("File without waypoints %s\n", gpxfile)}


		for _, track := range gpxFile.Tracks {
			track.Name = filename
			NewGPX.AppendTrack(&track)
		}
	}
	return NewGPX
}

func Simplify (gpxtrack *gpx.GPX, maxdistance float64) {
	// Упрощаем трек по алгоритму - Ramer-Douglas-Peucker - https://habr.com/ru/post/448618/
	gpxtrack.SimplifyTracks(maxdistance)
}

func Output(gpxtrack *gpx.GPX, filename string) {
	xmlBytes, _ := gpxtrack.ToXml(gpx.ToXmlParams{Version: "1.1", Indent: true})
    err := ioutil.WriteFile(filename, xmlBytes, 0644)

    if err != nil {
        log.Fatal(err)
    }

	log.Printf("Done. Result file = %s\n", filename)
	//fmt.Println(string(xmlBytes))
}