package currency

import (
	"awesomeProject3/internal/models"
	"awesomeProject3/pkg/currencies"
	"awesomeProject3/pkg/pg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Create(c *gin.Context) {
	firstCurrency := c.Query("currency1")
	secondCurrency := c.Query("currency2")

	rate, err := currencies.GetRate(firstCurrency, secondCurrency)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	var model models.Currency

	checkQuery := `select * from currency where currency1=$1 and currency2=$2`
	insertQuery := `
		insert into currency (currency1, currency2, rate) VALUES ($1, $2, $3)
	`

	err = pg.DB.Get(&model, checkQuery, firstCurrency, secondCurrency)
	if err != nil {
		_, err := pg.DB.Queryx(insertQuery, firstCurrency, secondCurrency, rate)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		c.JSON(http.StatusOK, gin.H{"result": "successfully created"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "this pair already has in db"})
	}
}

func Convert(c *gin.Context) {
	firstCurrency := c.Query("currencyFrom")
	secondCurrency := c.Query("currencyTo")
	rate, err := currencies.GetRate(firstCurrency, secondCurrency)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		value, _ := strconv.ParseFloat(c.Query("value"), 64)
		c.JSON(http.StatusOK, gin.H{"result": value * rate})
	}
}
