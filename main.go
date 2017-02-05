package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/hikhvar/zkillbeat/beater"
)

func main() {
	err := beat.Run("zkillbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
