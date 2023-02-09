package json

import (
	"encoding/json"
	"reflect"
	"regexp"
	"strings"
)

func ToCamelCase(input string) (string, error) {
	var obj interface{}
	err := json.Unmarshal([]byte(input), &obj)
	if err != nil {
		return "", err
	}

	obj = convertToCamelCase(obj)

	output, err := json.MarshalIndent(obj, "", "  ")
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
