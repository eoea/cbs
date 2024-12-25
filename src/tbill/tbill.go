package tbill

// This package is for the Treasury Bill (tbill).
// Created by Emile O.E. Antat (eoea) <eoea754@gmail.com>

import (
	"errors"
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

	browser "gitlab.com/eoea/cbs/src/browser"
)

// purchase_value: returns the amount of money you need to spend for the given
// period, interest rate, and face value (fv).
//
// r is the applicable market interest rate for the respective term to maturity.
// t is the number of days remaining until maturity of the T-bill.
func purchase_value(fv, r, t float64) float64 {
	return math.Round(fv / (1 + (r * t / 365)))
}

// tFmt: Takes an HTML content string and returns the day and corresponding
// rates. Error is returned if the format of the HTML content is not correct,
// often meaning that the HTML content has changed or is incorrect.
func tFmt(content string) (map[string]string, error) {
	// Regex patterns to match the headers and values
	headerPattern := `<th[^>]*>([^<]+)</th>`
	valuePattern := `<td[^>]*>([^<]+)</td>`
	// Extract headers
	headerRegex := regexp.MustCompile(headerPattern)
	headers := headerRegex.FindAllStringSubmatch(content, -1)
	// Extract values
	valueRegex := regexp.MustCompile(valuePattern)
	values := valueRegex.FindAllStringSubmatch(content, -1)

	var tbl = make(map[string]string)
	if len(headers) != len(values) && (len(headers) == 3 && len(values) == 3) {
		return nil, errors.New("Count of headers and rates does not match.")
	}
	for i := 0; i < 3; i++ {
		h := strings.TrimSpace(headers[i][1])
		v := strings.TrimSpace(values[i][1])
		tbl[h] = v
	}
	return tbl, nil
}

// CbsTbill:
// This is the main function that acts as a helper to pull all the Treasury Bill
// rates. Error is returned if the user passes an invalid input or if the format
// of the HTML changes on SCB's site.
func CbsTbill(fv float64) error {
	content, err := browser.FetchHTMLPage("https://www.cbs.sc/marketinfo/TBILL.html")
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	tbl, err := tFmt(content)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for k, v := range tbl {
		r, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return errors.New("Error converting string to float:")
		}
		r /= 100 // Converts the rates to decimal instead of percentage
		switch k {
		case "91-day":
			fmt.Println("At 91-day, rate:", r)
			fmt.Println("Investment amount:", purchase_value(fv, r, 91))
			fmt.Println("Profit:", fv-purchase_value(fv, r, 91))
			fmt.Println("------------------------------")
			//fallthrough
		case "182-day":
			fmt.Println("At 182-day, rate:", r)
			fmt.Println("Investment amount:", purchase_value(fv, r, 182))
			fmt.Println("Profit:", fv-purchase_value(fv, r, 182))
			fmt.Println("------------------------------")
			//fallthrough
		case "365-day":
			fmt.Println("At 365-day rate:", r)
			fmt.Println("Investment amount:", purchase_value(fv, r, 365))
			fmt.Println("Profit:", fv-purchase_value(fv, r, 365))
			fmt.Println("------------------------------")
			//fallthrough
		default:
			return errors.New("Invalid date content from HTML.")
		}
	}
	return nil
}
