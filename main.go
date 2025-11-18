package main

import (
	"encoding/json"
	"fmt"
	"io"

	"math/rand/v2"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

const StoicQuotesURL = "https://stoic.tekloon.net/stoic-quote" //Stoic Quotes
const DuneQuotesURL = "https://api.duniverse.space/v1/random"
const JSONFileName = "quoteOfTheDay.json"

// Match the API response shape: { "data": { "author": "...", "quote": "..." } }
type StoicQuoteData struct {
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

type StoicQuoteResponse struct {
	Data StoicQuoteData `json:"data"`
}

type DuneQuoteAuthor struct {
	Name string `json:"name"`
}

type DuneQuoteData struct {
	Title  string          `json:"title"`
	Author DuneQuoteAuthor `json:"author"`
}

type DuneQuoteResponse struct {
	Id   string `json:"id"`
	Text string `json:"text"`
	Book DuneQuoteData
}

type QuoteData struct {
	Quote  string
	Author string
}

// QuoteDataFile will be same regardless of stoic or dune
type QuoteDataFile struct {
	Data QuoteData `json:"data"`
	Date string    `json:"date"`
}

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.File("./demo/index.html")
	})
	r.Static("/static", "./demo")

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
		c.Header("Cache-Control", "public, max-age=86400")

		if date_today != quoteFile.Date { //If quote is not today's, Get a new one and write to file
			fmt.Println("Herer")
			toss := rand.IntN(2)
			var res *http.Response
			var err error
			var qData QuoteData
			var jsonFileData QuoteDataFile

			if toss == 0 {
				res, err = http.Get(StoicQuotesURL)
				if err != nil {
					fmt.Print("Quote Fetch failed", err)
				}
				defer res.Body.Close()

				var stoicQuoteRes StoicQuoteResponse
				if err := json.NewDecoder(res.Body).Decode(&stoicQuoteRes); err != nil {
					fmt.Print("Decode failed", err)
				}

				qData.Author = stoicQuoteRes.Data.Author
				qData.Quote = stoicQuoteRes.Data.Quote
				jsonFileData = QuoteDataFile{
					Data: qData,
					Date: date_today,
				}
			} else {
				res, err = http.Get(DuneQuotesURL)
				if err != nil {
					fmt.Print("Quote Fetch failed", err)
				}
				defer res.Body.Close()

				var duneQuoteRes DuneQuoteResponse
				if err := json.NewDecoder(res.Body).Decode(&duneQuoteRes); err != nil {
					fmt.Print("Decode failed", err)
				}

				qData.Author = duneQuoteRes.Book.Author.Name + "(" + duneQuoteRes.Book.Title + ")"
				qData.Quote = duneQuoteRes.Text
				jsonFileData = QuoteDataFile{
					Data: qData,
					Date: date_today,
				}
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

			svg := renderQuoteSVG(jsonFileData.Data.Author, jsonFileData.Data.Quote, themeName)

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
