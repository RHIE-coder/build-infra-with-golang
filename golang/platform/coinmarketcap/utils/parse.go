package utils

import (
	"encoding/json"
	"fmt"
)

type PriceInfo struct {
	Name                  string `json:"name"`
	PriceChangePercentage string `json:"percent_change_24h"`
	Price                 string `json:"price"`
	Symbol                string `json:"symbol"`
	ErrorCode             int    `json:"error_code"`
}

type Response struct {
	Data   map[string]CoinInfo `json:"data"`
	Status struct {
		ErrorCode int `json:"error_code"`
	} `json:"status"`
}

type CoinInfo struct {
	ID     int              `json:"id"`
	Name   string           `json:"name"`
	Quote  map[string]Quote `json:"quote"`
	Slug   string           `json:"slug"`
	Symbol string           `json:"symbol"`
}

type Quote struct {
	Price            float64 `json:"price"`
	PercentChange24h float64 `json:"percent_change_24h"`
}

func ParseTest() {
	data := []byte(`{
  "data": {
    "1027": {
      "id": 1027,
      "name": "Ethereum",
      "quote": {
        "2798": {
          "last_updated": "2023-04-12T08:09:28.000Z",
          "percent_change_1h": 0.05714311,
          "percent_change_24h": -2.68148487,
          "price": 2481657.8934830828
        }
      },
      "slug": "ethereum",
      "symbol": "ETH"
    },
    "4647": {
      "id": 4647,
      "name": "PUBLISH",
      "quote": {
        "2798": {
          "last_updated": "2023-04-12T08:09:28.000Z",
          "percent_change_1h": 0.11955399,
          "percent_change_24h": 1.15709976,
          "price": 10.088188049159216
        }
      },
      "slug": "publish",
      "symbol": "NEWS"
    }
  },
  "status": {
    "error_code": 0,
    "timestamp": "2023-04-12T08:10:13.648Z"
  }
}`)

	var resp Response
	err := json.Unmarshal(data, &resp)
	if err != nil {
		panic(err)
	}

	var priceInfoList []PriceInfo

	for _, coin := range resp.Data {
		priceInfo := PriceInfo{
			Name:                  coin.Name,
			PriceChangePercentage: fmt.Sprintf("%f", coin.Quote["2798"].PercentChange24h),
			Price:                 fmt.Sprintf("%f", coin.Quote["2798"].Price),
			Symbol:                coin.Symbol,
			ErrorCode:             resp.Status.ErrorCode,
		}
		fmt.Printf("%+v\n", priceInfo)
		priceInfoList = append(priceInfoList, priceInfo)
	}
	list, _ := json.Marshal(priceInfoList)
	fmt.Println(string(list))
}

func ParseResponseData(data []byte) string {
	var resp Response
	err := json.Unmarshal(data, &resp)
	if err != nil {
		panic(err)
	}

	var priceInfoList []PriceInfo

	for _, coin := range resp.Data {
		priceInfo := PriceInfo{
			Name:                  coin.Name,
			PriceChangePercentage: fmt.Sprintf("%f", coin.Quote["2798"].PercentChange24h),
			Price:                 fmt.Sprintf("%f", coin.Quote["2798"].Price),
			Symbol:                coin.Symbol,
			ErrorCode:             resp.Status.ErrorCode,
		}
		fmt.Printf("%+v\n", priceInfo)
		priceInfoList = append(priceInfoList, priceInfo)
	}
	list, _ := json.Marshal(priceInfoList)
	return string(list)
}
