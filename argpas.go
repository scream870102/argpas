package argpas

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func parse(arg, name string) (string, error) {
	reg, _ := regexp.Compile("--\\s*(\\w+)\\s*=*\\s*(.*)")
	res := reg.FindStringSubmatch(arg)
	if len(res) > 0 && strings.EqualFold(res[1], name) {
		return res[2], nil
	}
	return "", fmt.Errorf("no arg \"%s\" exists", name)
}

// Parse Check if specific arg exist
func Parse(arg, name string) (string, error) {
	if _, err := parse(arg, name); err == nil {
		return name, nil
	}
	return "", fmt.Errorf("no arg \"%s\" exists", name)
}

// ParseToTrigger Check if specific arg exist but return as Trigger
func ParseToTrigger(arg, name string) (*Trigger, error) {
	if _, err := parse(arg, name); err == nil {
		return &Trigger{IsTrigger: true}, nil
	}
	return nil, fmt.Errorf("no arg \"%s\" exists", name)
}

// ParseToInt parse the argument and return as Integer
func ParseToInt(arg, name string) (int, error) {
	res, err := parse(arg, name)
	if err != nil {
		return 0, fmt.Errorf("Parse %s failed", name)
	}
	result, conErr := strconv.Atoi(res)
	if conErr != nil {
		return 0, fmt.Errorf("Can't parse \"%s\" to Integer", res)
	}
	return result, nil
}

// ParseToStr parse the argument and return as String
func ParseToStr(arg, name string) (string, error) {
	res, err := parse(arg, name)
	if err != nil {
		return "", fmt.Errorf("Parse %s failed", name)
	}
	return res, nil
}

// ParseToBool parse the argument and return as bool
func ParseToBool(arg, name string) (bool, error) {
	res, err := parse(arg, name)
	if err != nil {
		return false, fmt.Errorf("Parse %s failed", name)
	}
	result, conErr := strconv.ParseBool(res)
	if conErr != nil {
		return false, fmt.Errorf("Can't parse %s to bool Error: %s", res, conErr.Error())
	}
	return result, nil
}

// ParseToFloat32 parse the argument to float32
func ParseToFloat32(arg, name string) (float32, error) {
	res, err := parse(arg, name)
	if err != nil {
		return 0, fmt.Errorf("Parse %s failed", name)
	}
	result, conErr := strconv.ParseFloat(res, 32)
	if conErr != nil {
		return 0, fmt.Errorf("Can't parse \"%s\" to Float32", res)
	}
	return float32(result), nil
}

// ParseToFloat64 parse the argument to float64
func ParseToFloat64(arg, name string) (float64, error) {
	res, err := parse(arg, name)
	if err != nil {
		return 0, fmt.Errorf("Parse %s failed", name)
	}
	result, conErr := strconv.ParseFloat(res, 64)
	if conErr != nil {
		return 0, fmt.Errorf("Can't parse \"%s\" to Float32", res)
	}
	return result, nil
}

// ParseToInterface pass an pointer and check all argument
func ParseToInterface(pointer interface{}, args []string) error {

	if reflect.ValueOf(pointer).Kind() != reflect.Ptr {
		msg := fmt.Sprintln(reflect.TypeOf(pointer), " is not a pointer")
		return errors.New(msg)
	}

	errMsg := ""

	var argpairs []argPair
	rValue := reflect.ValueOf(pointer).Elem()
	rType := reflect.TypeOf(pointer).Elem()

	// get all argpairs from target type
	for i := 0; i < rValue.NumField(); i++ {
		field := rValue.Field(i)
		fieldName := rType.Field(i).Name
		argpairs = append(argpairs, argPair{field, fieldName})
	}

	// iterate all os args
	for _, arg := range args {
		for _, info := range argpairs {
			if _, err := parse(arg, info.name); err == nil {
				kind := info.field.Type().Kind()
				if kind == reflect.String {
					if pv, err := ParseToStr(arg, info.name); err != nil {
						errMsg += err.Error() + "\n"
					} else {
						info.field.SetString(pv)
					}
				} else if kind == reflect.Bool {
					if pv, err := ParseToBool(arg, info.name); err != nil {
						errMsg += err.Error() + "\n"
					} else {
						info.field.SetBool(pv)
					}
				} else if kind == reflect.Int {
					if pv, err := ParseToInt(arg, info.name); err != nil {
						errMsg += err.Error() + "\n"
					} else {
						info.field.SetInt(int64(pv))
					}
				} else if kind == reflect.Float64 {
					if pv, err := ParseToFloat64(arg, info.name); err != nil {
						errMsg += err.Error() + "\n"
					} else {
						info.field.SetFloat(pv)
					}
				} else if info.field.Type() == reflect.TypeOf(Trigger{}) {
					info.field.Field(0).SetBool(true)
				} else {
					errMsg += fmt.Sprintln(info.field.Type(), " doesn't support")
				}
			}
		}

	}

	if errMsg != "" {
		return errors.New(errMsg)
	}
	return nil
}
