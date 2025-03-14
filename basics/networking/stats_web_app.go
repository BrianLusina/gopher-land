package main

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type statistics struct {
	numbers []float64
	mean    float64
	median  float64
}

const (
	statisticsForm = `<html><body><form action="/" method="POST">
<h1>Statistics</h1>
<h5>Compute base statistics for a given list of numbers</h5>
<label for="numbers">Numbers (comma or space-separated):</label><br>
<input type="text" name="numbers" size="30"><br />
<input type="submit" value="Calculate">
</form></html></body>`
	errorTemplate = `<p class="error">%s</p>`
)

var (
	pageTop    = ""
	pageBottom = ""
)

func run_stats_web() {
	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		fmt.Println(err.Error())
	}
}

// Write an HTML header, parse the form, write form to writer and make request for numbers
func homePage(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")

	err := request.ParseForm()
	_, err = fmt.Fprint(writer, pageTop, statisticsForm)
	if err != nil {
		return
	}

	if err != nil {
		_, err := fmt.Fprint(writer, pageTop, errorTemplate, err.Error(), pageBottom)
		if err != nil {
			return
		}
		return
	}

	switch request.Method {
	case "GET":
		_, _ = writer.Write([]byte(pageTop))
		_, _ = writer.Write([]byte(statisticsForm))
		_, _ = writer.Write([]byte(pageBottom))
	case "POST":
		numbers, message, ok := processRequest(request)
		if ok {
			stats := getStats(numbers)
			_, err := fmt.Fprintf(writer, formatStats(stats))
			if err != nil {
				return
			}
			_, _ = writer.Write([]byte(message))
		} else {
			_, err := fmt.Fprintf(writer, errorTemplate, message)
			if err != nil {
				return
			}
		}
	default:
		_, err := fmt.Fprint(writer, pageBottom)
		if err != nil {
			return
		}
	}
}

func formatStats(stats statistics) string {
	return fmt.Sprintf(`<table border="1">
<tr><th colspan="2">Results</th></tr>
<tr><td>Numbers</td><td>%v</td></tr>
<tr><td>Count</td><td>%d</td></tr>
<tr><td>Mean</td><td>%f</td></tr>
<tr><td>Median</td><td>%f</td></tr>
</table>`, stats.numbers, len(stats.numbers), stats.mean, stats.median)
}

// Capture the numbers from the request, and format the data and check for errors
func processRequest(request *http.Request) ([]float64, string, bool) {
	var numbers []float64
	if slice, found := request.Form["numbers"]; found && len(slice) > 0 {
		text := strings.Replace(slice[0], ",", " ", -1)
		for _, field := range strings.Fields(text) {
			if x, err := strconv.ParseFloat(field, 64); err != nil {
				return numbers, "'" + field + "' is invalid", false
			} else {
				numbers = append(numbers, x)
			}
		}
	}
	if len(numbers) == 0 {
		return numbers, "", false // no data first time form is shown
	}
	return numbers, "", true
}

// sort the values to get mean and median
func getStats(numbers []float64) (stats statistics) {
	stats.numbers = numbers
	sort.Float64s(stats.numbers)
	stats.mean = sum(numbers) / float64(len(numbers))
	stats.median = median(numbers)
	return
}

// seperate function to calculate the sum for mean
func sum(numbers []float64) (total float64) {
	for _, x := range numbers {
		total += x
	}
	return
}

// seperate function to calculate the median
func median(numbers []float64) float64 {
	middle := len(numbers) / 2
	result := numbers[middle]
	if len(numbers)%2 == 0 {
		result = (result + numbers[middle-1]) / 2
	}
	return result
}
