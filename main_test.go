package main

import (
	"strings"
	"testing"
)

func TestDoRequests(t *testing.T) {

	buf := strings.Builder{}

	cases := []map[string]interface{}{
		{
			"input": []string{
				"novalid",
			},
			"output": 1,
		},
		{
			"input": []string{
				"adjust.com",
				"google.com",
				"facebook.com",
				"yahoo.com",
				"yandex.com",
				"twitter.com",
				"reddit.com/r/funny",
				"reddit.com/r/notfunny",
				"baroquemusiclibrary.com",
			},
			"output": 10,
		},
	}

	for _, c := range cases {
		input := c["input"].([]string)
		output := c["output"].(int)
		doRequests(input, 10, &buf)
		resStr := buf.String()
		lines := strings.Split(resStr, "\n")
		linesCount := len(lines)
		if linesCount != output {
			t.Fatalf("Expected %d entries, returned %d", output, linesCount)
		}
		if linesCount > 1 {
			for _, str := range input {
				if !strings.Contains(resStr, str) {
					t.Fatalf("Expected output for url %s but not returned", str)
				}
			}
		}
		buf.Reset()
	}
}
