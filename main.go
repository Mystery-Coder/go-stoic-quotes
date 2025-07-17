package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

const QuotesURL = "https://stoic.tekloon.net/stoic-quote"
const JSONFileName = "quoteOfTheDay.json"

// Match the API response shape: { "data": { "author": "...", "quote": "..." } }
type QuoteData struct {
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

type QuoteResponse struct {
	Data QuoteData `json:"data"`
}

type QuoteDataFile struct {
	Data QuoteData `json:"data"`
	Date string    `json:"date"`
}

func main() {

	r := gin.Default()

	r.GET("/stoic-quote-svg", func(c *gin.Context) { // Takes a Query Parameter of theme, light by default

		jsonFile, err := os.Open(JSONFileName)

		if err != nil {
			fmt.Print("Could not open file", err)
		}
		defer jsonFile.Close()

		bytes, err := io.ReadAll(jsonFile)
		if err != nil {
			fmt.Print("Could not read file", err)
		}

		var quoteFile QuoteDataFile
		if err := json.Unmarshal(bytes, &quoteFile); err != nil {
			fmt.Print("Could not unmarshal", err)
		}

		date_today := time.Now().Format("2006-01-02")

		themeName := c.DefaultQuery("theme", "default")

		if date_today != quoteFile.Date { //If quote is not today's, Get a new one and write to file
			res, err := http.Get(QuotesURL)

			if err != nil {
				fmt.Print("Quote Fetch failed", err)
			}
			defer res.Body.Close()

			var quoteRes QuoteResponse
			if err := json.NewDecoder(res.Body).Decode(&quoteRes); err != nil {
				fmt.Print("Decode failed", err)
			}

			jsonFileData := QuoteDataFile{
				Data: quoteRes.Data,
				Date: date_today,
			}

			file, err := os.Create(JSONFileName) //Overwrites the file
			if err != nil {
				fmt.Print("Could not create file", err)
			}
			defer file.Close()

			encoder := json.NewEncoder(file)
			encoder.SetIndent("", "  ")
			if err := encoder.Encode(jsonFileData); err != nil {
				fmt.Print("Could not encode", err)
			}

			svg := renderQuoteSVG(quoteRes.Data.Author, quoteRes.Data.Quote, themeName)

			c.Data(200, "image/svg+xml", []byte(svg))
		} else {
			svg := renderQuoteSVG(quoteFile.Data.Author, quoteFile.Data.Quote, themeName)

			c.Data(200, "image/svg+xml", []byte(svg))
		}

	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local dev
	}
	r.Run(":" + port)
}
