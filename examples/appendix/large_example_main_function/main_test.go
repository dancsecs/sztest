package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"testing"

	"github.com/dancsecs/sztest"
)

func Test_PASS_Main_No_Args(t *testing.T) {
	chk := sztest.CaptureLogWithStderrAndStdout(t)
	defer chk.Release()

	log.Println("Testing missing angle")
	chk.SetArgs("progname")
	chk.Panic(
		main,
		"angle required",
	)

	log.Println("Testing invalid angle")
	chk.SetArgs("progname", "notANumber")
	chk.Panic(
		main,
		"invalid angle: notANumber",
	)

	fmt.Println("Testing angle 0 no flags")
	chk.SetArgs("progname", "0")
	chk.NoPanic(main)

	fmt.Println("Testing angle 0 with verbose flag")
	chk.SetArgs("progname", "-v", "0")
	chk.NoPanic(main)

	twoPi := strconv.FormatFloat(math.Pi*2, 'f', -1, 64)
	fmt.Println("Testing angle 2Pi with radian flag")
	chk.SetArgs("progname", "-r", twoPi)
	chk.NoPanic(main)

	fmt.Println("Testing angle 2Pi with radian and verbose flag")
	chk.SetArgs("progname", "-v", "-r", twoPi)
	chk.NoPanic(main)

	chk.Stdout(
		"Testing angle 0 no flags",
		"0.000000",
		"1.000000",
		"Testing angle 0 with verbose flag",
		"Report on 0.000000 degrees (0.000000 radians)",
		"Sin(0.000000°) = 0.000000",
		"Cos(0.000000°) = 1.000000",
		"Testing angle 2Pi with radian flag",
		"-0.000000",
		"1.000000",
		"Testing angle 2Pi with radian and verbose flag",
		"Report on 6.283185 radians (360.000000 degrees)",
		"Sin(6.283185) = -0.000000",
		"Cos(6.283185) = 1.000000",
	)
	chk.Log(
		"Testing missing angle",
		"angle required",
		//
		"Testing invalid angle",
		"invalid angle: notANumber",
		//
	)
}

func Test_FAIL_Main_No_Args(t *testing.T) {
	chk := sztest.CaptureLogWithStderrAndStdout(t)
	defer chk.Release()

	chk.FailFast(false) // Do not terminate function on first error.

	log.Println("Testing missing angle")
	chk.SetArgs("progname")
	chk.Panic(
		main,
		"angle is required",
	)

	log.Println("Testing invalid angle")
	chk.SetArgs("progname", "notANumber")
	chk.Panic(
		main,
		"invalid angle: not A Number",
	)

	fmt.Println("Testing angle 0 no flags")
	chk.SetArgs("progname", "0")
	chk.NoPanic(main)

	fmt.Println("Testing angle 0 with verbose flag")
	chk.SetArgs("progname", "-v", "0")
	chk.NoPanic(main)

	twoPi := strconv.FormatFloat(math.Pi*2, 'f', -1, 64)
	fmt.Println("Testing angle 2Pi with radian flag")
	chk.SetArgs("progname", "-r", twoPi)
	chk.NoPanic(main)

	fmt.Println("Testing angle 2Pi with radian and verbose flag")
	chk.SetArgs("progname", "-v", "-r", twoPi)
	chk.NoPanic(main)

	chk.Stdout(
		"Testing angle 0 no flags",
		"0.000000",
		"1.000000",
		"Testing angle 0 with verbose flag",
		"Report on 0.000000 degrees (0.000000 radians)",
		"Sin(0.000000°) = 0.000000",
		"Cos(0.000000°) = 1.000000",
		"Testing angle 2PI with radian flag",
		"-0.000000",
		"1.000000",
		"Testing angle 2Pi with radian and verbose flag",
		"Report on 6.283185 radians (360.000000 degrees)",
		"Sin(6.283185) = -0.000000",
		"Cos(6.283185) = 1.000000",
	)
	chk.Log(
		"Testing missing angle",
		"angle is required",
		//
		"Testing invalid angle",
		"invalid angle: not A Number",
		//
	)
}
