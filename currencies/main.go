package currencies

import (
	"encoding/json"
	"net/http"
)

func GetRate(currency1 string, currency2 string) float64 {
	urlRequest := "https://api.exchangerate.host/latest?base=" + currency1

	client := http.Client{}

	request, _ := http.NewRequest("GET", urlRequest, nil)
	resp, _ := client.Do(request)
	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	return result["rates"].(map[string]interface{})[currency2].(float64)
}
