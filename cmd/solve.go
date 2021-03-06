/*
Copyright © 2021 Deepak Kumar

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"github.com/kumadee/mazesolver/pkg/mazesolver"
	"github.com/spf13/cobra"
)

var InImg, OutImg string

// solveCmd represents the solve command
var solveCmd = &cobra.Command{
	Use:   "solve",
	Short: "Solve the given image with the choice of solving method",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		opt := mazesolver.Options{
			Method:         mazesolver.Bfs,
			InImg:          InImg,
			OutImg:         OutImg,
			HighlightNodes: true,
		}
		opt.Solve()
	},
}

func init() {
	rootCmd.AddCommand(solveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// solveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// solveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	solveCmd.Flags().StringVarP(&InImg, "input", "i", "images/tiny.png", "Maze input image")
	solveCmd.Flags().StringVarP(&OutImg, "output", "o", "solutions/tiny.png", "Maze solution image")
}
