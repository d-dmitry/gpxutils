package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tkrajina/gpxgo/gpx"
	"path/filepath"
)

var file string

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Print info about gpx track",
	Run: func(cmd *cobra.Command, args []string) {
		if file != "" {
			gpxFile, err := gpx.ParseFile(file)
			if err != nil {
				fmt.Println("Error opening gpx file: ", err)
				return
			}
			gpxPath, _ := filepath.Abs(file)
			fmt.Print("File: ", gpxPath, "\n")
			fmt.Println(gpxFile.GetGpxInfo())
		}
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
	infoCmd.Flags().StringVarP (&file, "file", "f","","Gpx track filename")
}
