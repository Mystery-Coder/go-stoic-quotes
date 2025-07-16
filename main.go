package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)
const QuotesURL = "https://stoic.tekloon.net/stoic-quote";

// Match the API response shape: { "data": { "author": "...", "quote": "..." } }
type QuoteData struct {
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

type QuoteResponse struct {
	Data QuoteData `json:"data"`
}



func main() {

    r := gin.Default()

    r.GET("/stoic-quote-svg", func(c *gin.Context) {
		res, err := http.Get(QuotesURL)

		if err != nil {
			fmt.Print("Quote Fetch failed", err)
		}
		defer res.Body.Close()

		var quoteRes QuoteResponse
		if err := json.NewDecoder(res.Body).Decode(&quoteRes); err != nil {
			fmt.Print("Decode failed", err)
		}

		// fmt.Print(quote.Data.Author, quote.Data.Quote)

		svg := renderQuoteSVG(quoteRes.Data.Author, quoteRes.Data.Quote)

		c.Data(200, "image/svg+xml", []byte(svg))

    })

    r.Run(":8080")
}
