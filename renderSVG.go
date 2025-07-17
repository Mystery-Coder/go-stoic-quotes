package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

type Theme struct {
	Quote      string `json:"quote"`
	Author     string `json:"author"`
	Background string `json:"background"`
	Symbol     string `json:"symbol"`
}

var Themes = map[string]Theme{
	"default": {
		Quote:      "333",
		Author:     "2f80ed",
		Background: "fffefe",
		Symbol:     "4c71f2",
	},
	"defaultDarkModeSupport": {
		Quote:      "9f9f9f",
		Author:     "fff",
		Background: "151515",
		Symbol:     "79ff97",
	},
	"chartreuse-dark": {
		Quote:      "fff",
		Author:     "7fff00",
		Background: "000",
		Symbol:     "00AEFF",
	},
	"radical": {
		Quote:      "a9fef7",
		Author:     "fe428e",
		Background: "141321",
		Symbol:     "f8d847",
	},
	"merko": {
		Quote:      "68b587",
		Author:     "abd200",
		Background: "0a0f0b",
		Symbol:     "b7d364",
	},
	"gruvbox": {
		Quote:      "8ec07c",
		Author:     "fabd2f",
		Background: "282828",
		Symbol:     "fe8019",
	},
	"tokyonight": {
		Quote:      "38bdae",
		Author:     "70a5fd",
		Background: "1a1b27",
		Symbol:     "bf91f3",
	},
	"catppuccin": {
		Quote:      "96CDFB",
		Author:     "D9E0EE",
		Background: "161320",
		Symbol:     "DDB6F2",
	},
	"catppuccin_latte": {
		Quote:      "179299",
		Author:     "4c4f69",
		Background: "eff1f5",
		Symbol:     "8839ef",
	},
	"catppuccin_frappe": {
		Quote:      "81c8be",
		Author:     "c6d0f5",
		Background: "303446",
		Symbol:     "ca9ee6",
	},
	"catppuccin_macchiato": {
		Quote:      "8bd5ca",
		Author:     "cad3f5",
		Background: "24273a",
		Symbol:     "c6a0f6",
	},
	"catppuccin_mocha": {
		Quote:      "94e2d5",
		Author:     "cdd6f4",
		Background: "1e1e2e",
		Symbol:     "cba6f7",
	},
	"algolia": {
		Quote:      "00ADFE",
		Author:     "FEFEFE",
		Background: "050F2C",
		Symbol:     "26BB85",
	},
	"monokai": {
		Quote:      "EA1F6A",
		Author:     "CFCFC9",
		Background: "272822",
		Symbol:     "E18905",
	},
	"dracula": {
		Quote:      "F8F8F2",
		Author:     "6272A4",
		Background: "282A36",
		Symbol:     "FF79c6",
	},
	"nord": {
		Quote:      "D8DEE9",
		Author:     "4C566A",
		Background: "2E3440",
		Symbol:     "88C0D0",
	},
	"github": {
		Quote:      "FFFFFF",
		Author:     "4C566A",
		Background: "0D1117",
		Symbol:     "43C293",
	},
	"github_dark": {
		Quote:      "C3D1D9",
		Author:     "58A6FF",
		Background: "0D1117",
		Symbol:     "1F6FEB",
	},
	"github_blue": {
		Quote:      "C7D5E0",
		Author:     "56A1F7",
		Background: "0D1117",
		Symbol:     "F9826C",
	},
	"graywhite": {
		Quote:      "24292E",
		Author:     "24292E",
		Background: "FFFFFF",
		Symbol:     "24292E",
	},
	"moonlight": {
		Quote:      "F8F8F8",
		Author:     "FF757F",
		Background: "222436",
		Symbol:     "599DFF",
	},
	"hackerman": {
		Quote:      "00B3D6",
		Author:     "00B3D6",
		Background: "000000",
		Symbol:     "00B3D6",
	},
	"shadow_red": {
		Quote:      "9A0000",
		Author:     "9A0000",
		Background: "151515",
		Symbol:     "4F0000",
	},
	"shadow_green": {
		Quote:      "007A00",
		Author:     "007A00",
		Background: "151515",
		Symbol:     "003D00",
	},
	"shadow_blue": {
		Quote:      "00779A",
		Author:     "00779A",
		Background: "151515",
		Symbol:     "004490",
	},
}

func getRandomTheme() string {

	var keys []string

	for key := range Themes {
		keys = append(keys, key)
	}

	randIdx := rand.IntN(len(keys))

	return keys[randIdx]

}

func makeThemeStyle(themeName string) string {

	if themeName == "random" {
		themeName = getRandomTheme()
	}

	return fmt.Sprintf(`
      .container {
        background-color: #%s;
      }
      .container h3 {
        color: #%s;
      }
      .container h3::before, .container h3::after {
        color: #%s;
      }
      .container p {
        color: #%s;
      }
    `, Themes[themeName].Background, Themes[themeName].Quote, Themes[themeName].Symbol, Themes[themeName].Author)
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

func renderQuoteSVG(author, quote, themeName string) string {
	return fmt.Sprintf(`
  
  <svg width="300" height="300" fill="none" xmlns="http://www.w3.org/2000/svg">
    <foreignObject width="100vw" height="100vh">
      <div xmlns="http://www.w3.org/1999/xhtml">
        <defs>
          <style type="text/css">
            <![CDATA[
                @font-face {
                    font-family: 'Poppins';
                    font-style: normal;
                    font-weight: 400;
                    src: url(https://fonts.gstatic.com/s/poppins/v15/pxiEyp8kv8JHgFVrJJfecg.woff2) format('woff2');
                    unicode-range: U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD;
                }
            ]]>
          </style>
        </defs>
        <style>
          * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
          }
          .container {
            width: 300px;
            height: 300px;
            font-family: Poppins, Arial, Helvetica, sans-serif;
            padding: 15px;
            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: center;
            text-align: center;
            border-radius: 10px;
          }
          .container h3::before {
            content: open-quote;
            font-size: 50px;
            display: block;
            margin-bottom: -20px;
          }
          .container h3::after {
            content: close-quote;
            font-size: 45px;
            display: block;
            margin-bottom: -20px;
          }
          .container h3 {
            margin-bottom: 12px;
          }
          .container p {
            font-style: italic;
          }
          %s
        </style>
        <div class="container">
          <h3>%s</h3>
          <p>- %s</p>
        </div>
      </div>
    </foreignObject>
  </svg>
  `, makeThemeStyle(themeName), escapeForSVG(quote), escapeForSVG(author))
}
