package main

import (
	"fmt"
	"strings"
)

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
  <svg width="300" height="300" fill="none" xmlns="http://www.w3.org/2000/svg">
    <foreignObject width="100vw" height="100vh">
      <div xmlns="http://www.w3.org/1999/xhtml">
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
            font-size: 50px;
            display: block;
            margin-bottom: -20px;
          }
          .container h3 {
            margin-bottom: 15px;
          }
          .container p {
            font-style: italic;
          }
          ${themeStyles}
        </style>
        <div class="container">
          <h3>%s</h3>
          <p>- %s</p>
        </div>
      </div>
    </foreignObject>
  </svg>
  `, escapeForSVG(quote), escapeForSVG(author))
}
