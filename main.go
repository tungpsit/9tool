package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/atotto/clipboard"
	"github.com/tungpsit/9tool/pkg/datetime"
	"github.com/tungpsit/9tool/pkg/json"
)

type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset        = "\u001b[0m"
)

// colorize string
func colorize(color Color, message string) {
	fmt.Println(string(color), message, string(ColorReset))
}

// write string into clipboard
func writeToClipboard(s string) {
	clipboard.WriteAll(s)
}

func output(o string) {
	colorize(ColorRed, "Output (Copied to clipboard!):")
	colorize(ColorBlue, o)
	writeToClipboard(o)
	colorize(ColorReset, "\nCopied to clipboard!")
}

func main() {
	camelCaseInput := flag.String("camel", "", "convert json string to camel case")
	iso8601Input := flag.String("iso", "", "convert datetime string to ISO8601 format")

	flag.Parse()

	// handle convert json string to camel case
	if camelCaseInput != nil && *camelCaseInput != "" {
		camelCaseResult, err := json.ToCamelCase(*camelCaseInput)
		if err != nil {
			fmt.Println(err)
			return
		}

		output(camelCaseResult)
	}

	// handle convert datetime string to ISO8601 format
	if iso8601Input != nil && *iso8601Input != "" {

		if iVal, err := strconv.ParseInt(*iso8601Input, 10, 64); err == nil {
			output(datetime.ToISO8601String(datetime.ParseUnixTime(iVal)))
		} else if timeVal, err := time.Parse("Mon Jan 02 2006 15:04:05 GMT-0700 (MST)", *iso8601Input); err == nil {
			output(datetime.ToISO8601String(timeVal))
		} else {
			output("Something went wrong!")
		}
	}
}
