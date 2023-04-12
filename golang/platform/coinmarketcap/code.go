package main

import (
	"golang/platform/coinmarketcap/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"
)

func READ_ENV_FILE() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	/*
		You can supply your API Key in REST API calls in one of two ways:
			Preferred method: Via a custom header named X-CMC_PRO_API_KEY
			Convenience method: Via a query string parameter named CMC_PRO_API_KEY
	*/
	READ_ENV_FILE()
	API_KEY := os.Getenv("API_KEY")
	// API_KEY := "b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c" // public sandbox key

	url := "https://pro-api.coinmarketcap.com"
	// url := "https://sandbox-api.coinmarketcap.com"
	client := utils.NewClient(url)

	client.SetBeforeRequest(func(req *fasthttp.Request) {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-CMC_PRO_API_KEY", API_KEY)
	})

	// flat_map := &utils.Requester{
	// 	Method: fasthttp.MethodGet,
	// 	URL:    "/v1/fiat/map",
	// }

	// cryptocurrency_categories := &utils.Requester{
	// 	Method: fasthttp.MethodGet,
	// 	URL:    "/v1/cryptocurrency/categories",
	// }

	// symbol := "ETH"
	// symbol := "NEWS"
	// requester := &utils.Requester{
	// 	Method: fasthttp.MethodGet,
	// 	URL:    "/v1/cryptocurrency/map" + "?symbol=" + symbol,
	// }

	// param := "publish"
	// requester := &utils.Requester{
	// 	Method: fasthttp.MethodGet,
	// 	URL:    "/v2/cryptocurrency/info" + "?slug=" + param,
	// }

	// param := "publish"

	// requester := &utils.Requester{
	// 	Method: fasthttp.MethodGet,
	// 	URL:    "/v2/cryptocurrency/market-pairs/latest" + "?slug=" + param,
	// }

	// requester := &utils.Requester{
	// 	Method: fasthttp.MethodGet,
	// 	URL:    "/v1/exchange/market-pairs/latest?id=4647&convert=KRW",
	// }

	// requester := &utils.Requester{
	// 	Method: fasthttp.MethodGet,
	// 	URL:    "/v1/exchange/map",
	// }

	// requester := &utils.Requester{
	// 	Method: fasthttp.MethodGet,
	// 	URL:    "/v1/exchange/assets?id=270",
	// }

	// param := "publish"
	param := "4647,1027"
	requester := &utils.Requester{
		Method: fasthttp.MethodGet,
		URL:    "/v2/cryptocurrency/quotes/latest?id=" + param + "&convert_id=2798",
	}

	client.Dial(requester)

}
