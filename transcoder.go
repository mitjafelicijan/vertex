package main

import (
	"strings"
)

var rules = map[string]string{
	"localStorage.getItem":    "localStorage.GetItem",
	"localStorage.setItem":    "localStorage.SetItem",
	"localStorage.removeItem": "localStorage.RemoveItem",
	"localStorage.clear":      "localStorage.Clear",
}

func transcode(source string) string {
	for k, v := range rules {
		source = strings.Replace(source, k, v, -1)
	}
	return source
}
