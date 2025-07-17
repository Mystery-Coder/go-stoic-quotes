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
