package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tkrajina/gpxgo/gpx"
	"log"
)

var dirDistance string
var regexDistance string

// distanceCmd represents the distance command
var distanceCmd = &cobra.Command{
	Use:   "distance",
	Short: "Return distance summ by tracks",
	Run: func(cmd *cobra.Command, args []string) {
		if dirDistance != "" {
			var DistSum float64 = 0

			for _, gpxfile := range LoadGPXs(dirDistance, regexDistance) {
				gpxFile, err := gpx.ParseFile(gpxfile)

				if err != nil {
					log.Fatal(err)
				}

				TrackDistance := gpxFile.MovingData().MovingDistance
				fmt.Printf("track_name: %s, distance: %f km\n", gpxfile, TrackDistance/1000)
				DistSum += TrackDistance
			}

			fmt.Printf("Summ = %f km", DistSum/1000)
		}
	},
}

func init() {
	rootCmd.AddCommand(distanceCmd)
	distanceCmd.Flags().StringVarP (&dirDistance, "dir", "d","","Dir with gpx tracks")
	distanceCmd.Flags().StringVarP (&regexDistance, "regex", "r",".*","Regex for search gpx tracs")
}
