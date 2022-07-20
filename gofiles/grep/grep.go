package grep

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type options struct {
	printLineNumbers     bool
	printFileNames       bool
	caseInsensitiveMatch bool
	invertPatternMatch   bool
	matchEntireLine      bool
	prefixFilenames      bool
}

func Search(pattern string, flags, files []string) []string {
	result := []string{}
	opts := parseOptions(flags, len(files))

	// done := make(chan bool)
	// searchCh := make(chan []string)

	for _, file := range files {
		// go func(f string) {
		// 	searchResult := searchFile(pattern, f, opts)
		// 	select {
		// 	case searchCh <- searchResult:
		// 	case <-done:
		// 	}
		// }(file)

		searchResult := searchFile(pattern, file, opts)
		result = append(result, searchResult...)
	}

	// result = append(result, <-searchCh...)
	// close(done)

	return result
}

// searchFile searches for a pattern in the given file using the provided options
// returns a list of results from the search
func searchFile(pattern, file string, opts *options) []string {
	searchResult := []string{}

	f, err := os.Open(file)
	if err != nil {
		return nil
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("Failed to close file: %v Err: %s", f.Name(), err)
		}
	}(f)

	var filePrefix string
	if opts.prefixFilenames {
		filePrefix = file + ":"
	}

	scanner := bufio.NewScanner(f)

	for lineNum := 1; scanner.Scan(); lineNum++ {
		var match bool
		line := scanner.Text()

		if opts.matchEntireLine {
			if opts.caseInsensitiveMatch {
				match = strings.EqualFold(line, pattern)
			} else {
				match = line == pattern
			}
		} else {
			if opts.caseInsensitiveMatch {
				match = strings.Contains(strings.ToLower(line), strings.ToLower(pattern))
			} else {
				match = strings.Contains(line, pattern)
			}
		}

		if match && !opts.invertPatternMatch {
			switch {
			case opts.printFileNames:
				return append(searchResult, file)
			case opts.printLineNumbers:
				lineResult := fmt.Sprintf("%s%d:%s", filePrefix, lineNum, line)
				searchResult = append(searchResult, lineResult)
			default:
				searchResult = append(searchResult, filePrefix+line)
			}
		} else if !match && opts.invertPatternMatch {
			switch {
			case opts.printFileNames:
				return append(searchResult, file)
			case opts.printLineNumbers:
				lineResult := fmt.Sprintf("%s%d:%s", filePrefix, lineNum, line)
				searchResult = append(searchResult, lineResult)
			default:
				searchResult = append(searchResult, filePrefix+line)
			}
		}

	}

	return searchResult
}

// parseOptions parses the flags passed in and number of files returning options
func parseOptions(flags []string, nfiles int) *options {
	opt := &options{}

	for _, flag := range flags {
		switch flag {
		case "-n":
			opt.printLineNumbers = true
		case "-l":
			opt.printFileNames = true
		case "-i":
			opt.caseInsensitiveMatch = true
		case "-v":
			opt.invertPatternMatch = true
		case "-x":
			opt.matchEntireLine = true
		}
	}

	if nfiles > 1 {
		opt.prefixFilenames = true
	}

	return opt
}
