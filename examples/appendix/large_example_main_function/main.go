package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

const radiansPerDegree = math.Pi / 180.0
const degreesPerRadian = 180.0 / math.Pi
const bits64 = 64
const numDecimals = 6

func makeReport(degrees, radians float64, useRadians, verbose bool) []string {
	var rep []string
	addToReport := func(v float64, name string) {
		line := strconv.FormatFloat(v, 'f', numDecimals, bits64)
		if verbose {
			arg := ""
			if !useRadians {
				arg = strconv.FormatFloat(degrees, 'f', numDecimals, bits64) + "°"
			} else {
				arg = strconv.FormatFloat(radians, 'f', numDecimals, bits64)
			}
			line = name + "(" + arg + ") = " + line
		}
		rep = append(rep, line)
	}
	addToReport(math.Sin(radians), "Sin")
	addToReport(math.Cos(radians), "Cos")
	return rep
}

// Program takes an angle and reports its Sin and Cos values.
// -v causes a more detailed response.
// -r cause the angle input to be interrupted as radians.
func main() {
	var degrees, radians float64

	verbose := flag.Bool("v", false, "More detailed information.")
	useRadians := flag.Bool("r", false, "Value is in radians.")

	flag.Parse()

	if flag.NArg() != 1 {
		log.Panic("angle required")
	}
	v, err := strconv.ParseFloat(flag.Args()[0], bits64)

	if err != nil {
		log.Panicf("invalid angle: %s", flag.Args()[0])
	}

	if *useRadians {
		degrees = degreesPerRadian * v
		radians = v
	} else {
		degrees = v
		radians = radiansPerDegree * v
	}

	if *verbose {
		if *useRadians {
			fmt.Printf("Report on %f radians (%f degrees)\n", radians, degrees)
		} else {
			fmt.Printf("Report on %f degrees (%f radians)\n", degrees, radians)
		}
	}

	fmt.Print(
		strings.Join(makeReport(degrees, radians, *useRadians, *verbose), "\n"),
		"\n",
	)
}
