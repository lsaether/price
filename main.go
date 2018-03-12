package main

import (
	"os";
	"flag";
	"fmt";
	"strings"
)

func main() {
	priceSource := flag.String("source", "coinmarketcap", "source of price feed")
	flag.Parse()

	validSources := map[string]bool {
		"coinmarketcap": true,
	}

	if !validSources[*priceSource] {
		panic("Not a valid price source!")
	}

	// Collect the args minus the program running.
	args := os.Args[1:]

	// We define the currency to check as the first one after the flags.
	var curr string 
	for _, elem := range args {
		if strings.HasPrefix(elem, "-") {
			continue
		} else {
			curr = elem
			break
		}
	}
	
	fmt.Println(curr)
}