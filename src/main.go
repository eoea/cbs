package main

import (
	"flag"
	"log"
	"strconv"

	rates "github.com/eoea/cbs/src/rates"
	tbill "github.com/eoea/cbs/src/tbill"
)

func main() {
	cRates := flag.Bool("rates", false, "Gets the Seychelles Central Bank rates for EUR, USD, and GBP.")
	cTbill := flag.Bool("tbill", false, "Gets the Seychelles Central Bank Treasury Bill rates.")

	flag.BoolVar(cRates, "r", *cRates, "Gets the Seychelles Central Bank rates for EUR, USD, and GBP. [short version]")
	flag.BoolVar(cTbill, "t", *cTbill, "Gets the Seychelles Central Bank Treasury Bill rates. [short version]")

	flag.Parse()

	if *cRates {
		rates.CbsRates()
	} else if *cTbill {
		arg := flag.Args()
		if len(arg) < 1 || len(arg) > 1 {
			log.Fatalf("Error: Invalid Face Value, you passed %v arguments.", len(arg))
		}
		fv, err := strconv.ParseFloat(arg[0], 64)
		if err != nil {
			log.Fatalf("Error: Invalid Face Value, expected a number but got %v.", arg)
		}
		if err := tbill.CbsTbill(fv); err != nil {
			log.Fatalf("Error: %v", err)
		}
	}
}
