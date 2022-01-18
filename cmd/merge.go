package cmd

import (
	"github.com/spf13/cobra"
	"time"
)

var dirMerge string
var regexMerge string
var nameMerge string
var maxDistance float64

var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge tracks to one",

	Run: func(cmd *cobra.Command, args []string) {
		if dirMerge != "" {
			newgpx := MergeGpx()
			Simplify(newgpx, maxDistance)
			Output(newgpx, nameMerge)
		}
	},
}

func init() {
	rootCmd.AddCommand(mergeCmd)
	mergeCmd.Flags().StringVarP (&dirMerge, "dir", "d","","Dir with gpx tracks")
	mergeCmd.Flags().StringVarP (&regexMerge, "regex", "r",".*","Regex for search gpx tracs")
	mergeCmd.Flags().StringVarP (&nameMerge, "name", "n",time.Now().Format("2006_01_02-150405"+".gpx"),"Result filename")
	mergeCmd.Flags().Float64Var (&maxDistance, "maxdistance", 0.5, "Max distance - https://habr.com/ru/post/448618/" )
}
