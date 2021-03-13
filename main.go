package main

import (
	"fmt"
	"os"

	"github.com/kolikons/label-watch/cmd"
	"github.com/kolikons/label-watch/pkg/label"
)

func main() {
	opts, e := cmd.ParseFlags()
	if e != nil {
		fmt.Print(e)
		os.Exit(0)
	}
	fmt.Println("Running label-watch")
	// Run Lable pkg to set labels
	if opts.Interval == "" {
		label.RunLabel(opts)
	}

	// Run Ticker daemon
	if opts.Interval != "" {
		label.RunTimerLabel(opts)
	}

	os.Exit(0)
}
