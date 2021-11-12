package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://api.binance.com/sapi/v1/capital/deposit/address?coin=BTC&recvWindow=60000&timestamp=1636623540552&signature=214061f8332201809ce3039cb91506a81c7031aba19eb23f31ff63a4bedf5a92"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("X-MBX-APIKEY", "459GGkaLLZuP7leFGfqF9FrmZPXix7d6Z30wgWLoPuQqUVv8Q1UyIVBSwTS98T2F")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
