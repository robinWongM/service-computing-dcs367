package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
)

type selpgParams struct {
	startPage      int
	endPage        int
	linesPerPage   int
	isFormFeedMode bool
	destination    string
}

var startPage = pflag.IntP("start", "s", -1, "Start Page Num")
var endPage = pflag.IntP("end", "e", -1, "End Page Num")
var linesPerPage = pflag.IntP("lines-per-page", "l", 72, "Lines per Page")
var isFormFeedMode = pflag.BoolP("form-feed", "f", false, "f mode")
var destination = pflag.StringP("destination", "d", "-", "Destination")

var programName string
var inFile io.Reader
var outFile io.Writer

func main() {
	programName = os.Args[0]

	processFlags()
	processInput()
}

func processFlags() {
	pflag.Parse()

	if *startPage == -1 {
		fmt.Fprintf(os.Stderr, "%s: please provide start page\n", programName)
		pflag.Usage()
		os.Exit(2)
	}

	if *endPage == -1 {
		fmt.Fprintf(os.Stderr, "%s: please provide end page\n", programName)
		pflag.Usage()
		os.Exit(4)
	}

	if *endPage < *startPage {
		fmt.Fprintf(os.Stderr, "%s: invalid end page %v\n", programName, *endPage)
		pflag.Usage()
		os.Exit(5)
	}

	inFileName := pflag.Arg(0)

	if inFileName == "-" || inFileName == "" {
		inFile = os.Stdin
	} else if inFileName != "" {
		var err error
		inFile, err = os.Open(inFileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: cannot open input file %s\n", programName, inFileName)
			os.Exit(10)
		}
	}
}

func processInput() {
	bufReader := bufio.NewReader(inFile)

	if *destination == "-" {
		outFile = os.Stdout
	} else {
		ld := exec.Command("lp", "-d", *destination)
		// ld := exec.Command("tee")
		ld.Stdin, outFile = io.Pipe()
		err := ld.Start()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: cannot open pipe to lp -d %s, %s\n", programName, *destination, err)
			os.Exit(13)
		}
	}

	var pageCurrent int

	if *isFormFeedMode {
		pageCurrent = 1
		for {
			c, err := bufReader.ReadByte()
			if err == io.EOF {
				break
			}
			if c == '\f' {
				pageCurrent++
			}
			if pageCurrent >= *startPage && pageCurrent <= *endPage {
				outFile.Write([]byte{c})
			}
		}
	} else {
		lineCurrent := 0
		pageCurrent = 1
		for {
			line, err := bufReader.ReadBytes('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Fprintf(os.Stderr, "%s: system error occurred: %e", programName, err)
			}
			lineCurrent++
			if lineCurrent > *linesPerPage {
				pageCurrent++
				lineCurrent = 1
			}
			if pageCurrent >= *startPage && pageCurrent <= *endPage {
				outFile.Write(line)
			}
		}
	}

	if pageCurrent < *startPage {
		fmt.Fprintf(os.Stderr, "%s: start_page (%d) greater than total pages (%d), no output written\n", programName, *startPage, pageCurrent)
	} else if pageCurrent < *endPage {
		fmt.Fprintf(os.Stderr, "%s: end_page (%d) greater than total pages (%d), less output than expected\n", programName, *endPage, pageCurrent)
	} else {
		fmt.Fprintf(os.Stderr, "%s: done\n", programName)
	}
}
