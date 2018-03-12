package main

import (
	// "encoding/json";
	"flag";
	"fmt";
	"io";
	"net/http";
	"os";
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

	resp, err := http.Get("https://api.coinmarketcap.com/v1/ticker/?limit=250")
	if err != nil {
		panic(err)
	} else {
		defer resp.Body.Close()
		_, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
				panic(err)
		}
	}
	
	// fmt.Println(resp.Body)
	fmt.Println(curr)
}