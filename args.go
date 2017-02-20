package snc

import (
	"errors"
	"strconv"
)

func ParseArgs(args []string) (string, uint32, error) {
	if len(args) < 3 {
		return "", 0, errors.New("Need at least hostname and port")
	}
	intVal, err := strconv.ParseUint(args[2], 10, 32)
	if err != nil {
		return "", 0, err
	}
	return args[1], uint32(intVal), nil
}
