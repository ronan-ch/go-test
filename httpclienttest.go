package main

import (
	"context"
	"net/http"
	"time"
)

func testHttpClient() {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	url := "https://gn-shop.csan.goodnotesapp.cn/nest/api/marketplace/products/related?locale=en&supportId=E7Z34N4HGV&id=64a6986868111ad6c54b869a&appVersion=6.6000.1&draft=false&countryCode=CHN)"
	req, err := http.NewRequestWithContext(context.Background(),
		http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "users API-Key 9434d256-6c9d-4c4c-89cd-b7c922064c9a")
	req.Header.Add("x-gn-account-id", "aaeab55c-f1b5-4ec3-899c-69294a7c65bb")
	res, err := client.Do(req)
	print(res.Header)

}
