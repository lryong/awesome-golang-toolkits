package main

import (
	"log"
	"strconv"
	"strings"
)

func MustBool(s string, defaultVal ...bool) (bool, error) {
	getDefault := func() bool {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return false
	}

	if s == "" {
		return getDefault()
	}
	b, err := strconv.ParseBool(strings.TrimSpace(s))
	if err != nil {
		log.Println("strconv.ParseBool error:", err)
		return getDefault()
	}
	return b

}
