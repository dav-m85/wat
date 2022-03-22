package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/dav-m85/wat/internal"
)

var color = struct{ Reset, Red, Green, Yellow, Blue, Purple, Cyan, Gray, White string }{
	Reset:  "\033[0m",
	Red:    "\033[31m",
	Green:  "\033[32m",
	Yellow: "\033[33m",
	Blue:   "\033[34m",
	Purple: "\033[35m",
	Cyan:   "\033[36m",
	Gray:   "\033[37m",
	White:  "\033[97m",
}

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	cfg := struct {
		depth int
		wd    string
	}{}

	flag.StringVar(&cfg.wd, "wd", wd, "Working directory")
	flag.IntVar(&cfg.depth, "depth", 0, "How deep wat should look for project READMEs, -1 for unlimited depth (default: 0).")
	flag.Parse()

	fs := os.DirFS(cfg.wd)

	printer := func(e *internal.Entry, err error) {
		if err != nil {
			fmt.Println(color.Red + err.Error() + color.Reset)
			return
		}
		var sb strings.Builder
		sb.WriteString(color.Green + e.Dir + color.Reset)
		sb.WriteString("\n")

		switch true {
		case strings.Contains(e.Git, "ahead"):
			sb.WriteString(color.Red + e.Git + color.Reset)
			sb.WriteString("\n")
		case e.Git != "":
			sb.WriteString(e.Git)
			sb.WriteString("\n")
		}

		if e.Excerpt != "" {
			sb.WriteString(e.Excerpt)
			sb.WriteString("\n")
		}

		fmt.Println(sb.String())
	}

	internal.Walk(fs, cfg.depth, printer)
}
