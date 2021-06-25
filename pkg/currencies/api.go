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

	request, err := http.NewRequest(http.MethodGet, urlRequest, nil)
	if err != nil {
		return 0, err
	}

	resp, err := client.Do(request)
	if err != nil {
		return 0, err
	}

	var result Currency

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return 0, err
	}

	return result.Rates[SecondCurrency].(float64), nil
}
