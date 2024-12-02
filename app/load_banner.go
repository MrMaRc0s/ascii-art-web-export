package main

import (
	"fmt"
	"os"
	"strings"
)

// loadBanner loads the selected banner file into a slice of strings
func loadBanner(banner string) ([]string, error) {
	bannerFile, err := os.ReadFile("banners/" + banner + ".txt")
	if err != nil {
		return nil, fmt.Errorf("banner file not found")
	}
	line := strings.Replace(string(bannerFile), "\r\n", "\n", -1)
	return strings.Split(line, "\n"), nil
}
