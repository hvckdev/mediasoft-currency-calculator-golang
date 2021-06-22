package currencies

import (
	"encoding/json"
	"net/http"
)

type Currency struct {
	Base  string                 `json:"base"`
	Rates map[string]interface{} `json:"rates"`
}

func GetRate(FirstCurrency string, SecondCurrency string) (float64, error) {
	urlRequest := "https://api.exchangerate.host/latest?base=" + FirstCurrency

	client := http.Client{}

	request, _ := http.NewRequest(http.MethodGet, urlRequest, nil)
	resp, _ := client.Do(request)

	var result Currency

	err := json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return 0, err
	}

	return result.Rates[SecondCurrency].(float64), nil
}
