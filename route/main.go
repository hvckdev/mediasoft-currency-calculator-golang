package route

import (
	"awesomeProject3/connection"
	"awesomeProject3/currencies"
	"awesomeProject3/table"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Create(c *gin.Context) {
	db := connection.Database()

	currency1 := c.Query("currency1")
	currency2 := c.Query("currency2")

	dt := time.Now().Format(time.RFC822Z)

	var values table.Main

	err := db.Get(&values, "SELECT * FROM main WHERE currency1=? AND currency2=?", currency1, currency2)
	if err != nil {
		_, err = db.NamedExec("INSERT INTO `main` (`id`, `currency1`, `currency2`, `rate`, `updated_at`) VALUES (NULL,:currency1,:currency2, :result, :date)",
			map[string]interface{}{
				"currency1": currency1,
				"currency2": currency2,
				"result":    currencies.GetRate(currency1, currency2),
				"date":      dt,
			})
		if err != nil {
			c.JSON(200, err)
		}

		c.JSON(http.StatusOK, gin.H{"message": "successfully"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "database already had this values"})
	}
}

func Convert(c *gin.Context) {
	currency1 := c.Query("currency1")
	currency2 := c.Query("currency2")
	value, _ := strconv.ParseFloat(c.Query("value"), 64)

	c.JSON(200, gin.H{"result": value * currencies.GetRate(currency1, currency2)})
}
