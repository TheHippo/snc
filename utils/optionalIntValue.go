package utils

import (
	"fmt"
	"strconv"

	flag "github.com/ogier/pflag"
)

type OptionalIntValue struct {
	Init   bool
	Parsed bool
	Value  uint32
}

// compile time check if interface is implemented
var _ flag.Value = &OptionalIntValue{}

func (lv *OptionalIntValue) String() string {
	if lv.Init == false {
		return "8888"
	}
	return fmt.Sprintf("%t %d", lv.Parsed, lv.Value)
}

func (lv *OptionalIntValue) Set(input string) error {
	lv.Init = true
	if input == "" {
		lv.Parsed = false
		return nil
	}
	intVal, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		return err
	}
	lv.Parsed = true
	lv.Value = uint32(intVal)
	return nil
}
