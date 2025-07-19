package main

import (
	"strconv"
)

const DEFAULTLIMIT = 2
const MAXLIMIT = 1000

func postLimit(cmd command) int32 {
	if len(cmd.arguments) == 0 {
		return DEFAULTLIMIT
	}

	i, err := strconv.Atoi(cmd.arguments[0])

	if err != nil || i > MAXLIMIT || i < 0 {
		return DEFAULTLIMIT
	}

	return int32(i)
}
