package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"reflect"
	"regexp"
	"strings"
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

func colorize(color Color, message string) {
	fmt.Println(string(color), message, string(ColorReset))
}

func main() {
	camelCaseInput := flag.String("camel", "", "convert json string to camel case")
	flag.Parse()

	if camelCaseInput != nil && *camelCaseInput != "" {
		camelCaseResult, err := toCamelCase(*camelCaseInput)
		if err != nil {
			fmt.Println(err)
			return
		}

		colorize(ColorRed, "Output:")
		colorize(ColorBlue, camelCaseResult)
	}
}

// convert json string to camel case
func toCamelCase(input string) (string, error) {
	var obj interface{}
	err := json.Unmarshal([]byte(input), &obj)
	if err != nil {
		return "", err
	}

	obj = convertToCamelCase(obj)

	output, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func convertToCamelCase(obj interface{}) interface{} {
	switch reflect.TypeOf(obj).Kind() {
	case reflect.Map:
		m := obj.(map[string]interface{})
		for k, v := range m {
			delete(m, k)
			k = toCamelCaseString(k)
			// check v is nil
			if v != nil {
				m[k] = convertToCamelCase(v)
			} else {
				m[k] = v
			}
		}
		return m
	case reflect.Slice:
		s := obj.([]interface{})
		for i, v := range s {
			s[i] = convertToCamelCase(v)
		}
		return s
	default:
		return obj
	}
}

func toCamelCaseString(input string) string {
	re := regexp.MustCompile("[^a-zA-Z0-9]+")
	input = re.ReplaceAllString(input, " ")
	input = strings.Title(input)
	re = regexp.MustCompile("[^a-zA-Z0-9]")
	input = re.ReplaceAllString(input, "")
	input = strings.ToLower(input[:1]) + input[1:]
	return input
}
