package argpas

import (
	"fmt"
	"regexp"
	"strconv"
)

func parseArg(arg, name string) (string, error) {
	reg, _ := regexp.Compile("--(\\w+)=(.+)")
	res := reg.FindStringSubmatch(arg)
	if len(res) > 0 && res[1] == name {
		return res[2], nil
	}
	return "", fmt.Errorf("no arg\"%s\" exists", name)
}

// ParseArgToInt parse the argument and return as Integer
func ParseArgToInt(arg, name string) (int, error) {
	res, err := parseArg(arg, name)
	if err != nil {
		return 0, fmt.Errorf("Parse %s failed", name)
	}
	result, conErr := strconv.Atoi(res)
	if conErr != nil {
		return 0, fmt.Errorf("Can't parse \"%s\" to Integer", res)
	}
	return result, nil
}

// ParseArgToStr parse the argument and return as String
func ParseArgToStr(arg, name string) (string, error) {
	res, err := parseArg(arg, name)
	if err != nil {
		return "", fmt.Errorf("Parse %s failed", name)
	}
	return res, nil
}

// ParseArg Check if specific arg exist
func ParseArg(arg, name string) (string, error) {
	reg, _ := regexp.Compile("--(\\w+)")
	res := reg.FindStringSubmatch(arg)
	if len(res) > 0 && res[1] == name {
		return res[1], nil
	}
	return "", fmt.Errorf("no arg\"%s\" exists", name)
}
