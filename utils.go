package main

import (
	"strconv"
	"strings"
)

// ParseURL is function for get id of schema and return
func ParseURL(url string) int {
	if url != "" {
		urlSplited := strings.Split(url, "/")
		ID, errParse := strconv.Atoi(urlSplited[len(urlSplited)-1])
		if errParse != nil {
			panic(errParse)
		}
		return ID
	}
	return 0
}
