package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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


func escapeForSVG(text string) string {
	replacer := strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
		"\"", "&quot;",
		"'", "&apos;",
	)
	return replacer.Replace(text)
}


func renderQuoteSVG(author, quote string) string {
	return fmt.Sprintf(`
<svg width="600" height="400" viewBox="0 0 600 400" xmlns="http://www.w3.org/2000/svg">
  <!-- Background -->
  <rect width="600" height="400" fill="#f8f9fa" stroke="#e9ecef" stroke-width="2" rx="12"/>
  
  <!-- Decorative elements -->
  <circle cx="50" cy="50" r="3" fill="#6c757d" opacity="0.3"/>
  <circle cx="550" cy="350" r="3" fill="#6c757d" opacity="0.3"/>
  <circle cx="550" cy="50" r="2" fill="#6c757d" opacity="0.2"/>
  <circle cx="50" cy="350" r="2" fill="#6c757d" opacity="0.2"/>
  
  <!-- Quote marks -->
  <text x="80" y="120" font-family="Georgia, serif" font-size="60" fill="#dee2e6" font-weight="bold">"</text>
  <text x="480" y="280" font-family="Georgia, serif" font-size="60" fill="#dee2e6" font-weight="bold">"</text>
  
  <!-- Quote text -->
  <text x="300" y="160" font-family="Georgia, serif" font-size="24" fill="#212529" text-anchor="middle" font-style="italic">
    <tspan x="300" dy="0">%s</tspan>
  </text>
  
  <!-- Divider line -->
  <line x1="200" y1="240" x2="400" y2="240" stroke="#6c757d" stroke-width="1" opacity="0.4"/>
  
  <!-- Author -->
  <text x="300" y="280" font-family="Arial, sans-serif" font-size="18" fill="#6c757d" text-anchor="middle" font-weight="500">
    — %s
  </text>
  
  <!-- Subtle border accent -->
  <rect x="8" y="8" width="584" height="384" fill="none" stroke="#dee2e6" stroke-width="1" rx="8" opacity="0.5"/>
</svg>`, escapeForSVG(quote), escapeForSVG(author))
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
