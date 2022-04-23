package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Order struct {
	OrderType  int     `json:"orderType"`
	LimitPrice float64 `json:"limitPrice"`
	Amount     float64 `json:"amount"`
}

func main() {
	fmt.Println("💲 On chain limit buying 💲")
	fmt.Println("Version 1.0")
	fmt.Println("Commands:")
	fmt.Println("buy [amount in USDC] at [limit price](ex. To buy 10USDC worth of JEWEL at $2.4, type 'buy 10 at 2.4')")
	fmt.Println("sell [amount in JEWEL] at [limit price](ex. To sell 1 JEWEL worth at $2.4, type 'sell 1 at 2.4')")
	fmt.Println("orders (Gets all current limit orders placed)")
	for {
		var line string
		input := bufio.NewScanner(os.Stdin)
		fmt.Print(">")
		if input.Scan() {
			line = input.Text()
		}
		split := strings.Fields(line)
		switch strings.TrimSpace(split[0]) {
		case "orders":
			orders()
		case "buy":
			buy(split[3], split[1])
		case "sell":
			sell(split[3], split[1])
		default:
			fmt.Println("Invalid command")
		}
	}

}
func orders() {
	response, err := http.Get("http://" + os.Args[1] + ":5000/orders")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}
func buy(limitPrice string, amount string) {
	p, err := strconv.ParseFloat(limitPrice, 64)
	if err != nil {
		log.Fatal(err)
	}
	a, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		log.Fatal(err)
	}

	order := Order{
		OrderType:  1,
		LimitPrice: p,
		Amount:     a,
	}
	body, err := json.Marshal(order)
	if err != nil {
		log.Fatal(err)
	}
	response, err := http.Post("http://"+os.Args[1]+":5000/orders", "application/json", bytes.NewBuffer(body))

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	fmt.Println("Response Status:", response.Status)
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}
func sell(limitPrice string, amount string) {
	p, err := strconv.ParseFloat(limitPrice, 64)
	if err != nil {
		log.Fatal(err)
	}
	a, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		log.Fatal(err)
	}

	order := Order{
		OrderType:  0,
		LimitPrice: p,
		Amount:     a,
	}
	body, err := json.Marshal(order)
	if err != nil {
		log.Fatal(err)
	}
	response, err := http.Post("http://"+os.Args[1]+":5000/orders", "application/json", bytes.NewBuffer(body))

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	fmt.Println("Response Status:", response.Status)

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}
