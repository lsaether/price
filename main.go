package main

// TODO: allow for input of just the Ticker symbol.

import (
	"bytes";
	"encoding/json";
	"flag";
	"fmt";
	// "io";
	"net/http";
	"os";
	"strings"
)

// TODO: create struct for JSON parsing

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

	lowered := strings.ToLower(curr)
	req := fmt.Sprintf("https://api.coinmarketcap.com/v1/ticker/%s", lowered)


	resp, err := http.Get(req)
	if err != nil {
		panic(err)
	} else {
		defer resp.Body.Close()
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		var f interface{}
		err := json.Unmarshal(buf.Bytes(), &f)
		m := f.([]interface{})
		// fmt.Println(m)
		r := m[0]
		c := r.(map[string]interface{})
		fmt.Println(c["price_btc"])
		if err != nil {
			panic(err)
		}
	}

	// This gets all the currency symbols, but is not really needed.
	//
	// resp, err := http.Get("https://api.coinmarketcap.com/v1/ticker/?limit=250")
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	defer resp.Body.Close()
	// 	buf := new(bytes.Buffer)
	// 	buf.ReadFrom(resp.Body)
	// 	var f interface{}
	// 	err := json.Unmarshal(buf.Bytes(), &f)
	// 	m := f.([]interface{})
	// 	var symbols []string
	// 	for _, elem := range m {
	// 		d := elem.(map[string]interface{})
	// 		fmt.Println(d["symbol"]))
	// 	}
	// 	if err != nil {
	// 			panic(err)
	// 	}
	// }
	
	// fmt.Println(resp.Body)
	// fmt.Println(curr)
}