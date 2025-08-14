# Stoic Quote SVG

A backend written in GoLang to generate a stoic quote SVG. Inspired
from [Github Readme Quotes](https://github.com/PiyushSuthar/github-readme-quotes), Uses stoic quotes from [Stoic Quote API](https://stoic.tekloon.net/stoic-quote).

---

## Features

-   **Daily dose of Stoicism** – Get quotes from Marcus Aurelius, Seneca, Epictetus, and other Stoics.
-   **SVG output** – Perfect for embedding in GitHub READMEs, personal dashboards, or blogs.
-   **Customizable themes** – Choose from predefined color schemes.
-   **Random theme support** – Use `theme=random` to get a different look each time.
-   **Blazing fast** – Written in Go with minimal dependencies.

---

## Usage

```http

GET https://go-stoic-quotes-production.up.railway.app/stoic-quote-svg
```

with query parameter of theme,

```http

GET https://go-stoic-quotes-production.up.railway.app/stoic-quote-svg?theme=random
```

Embed in MarkDown for README,

```markdown
![Stoic Quotes](https://go-stoic-quotes-production.up.railway.app/stoic-quote-svg?theme=moonlight)
```

## Demo

Visit <a href="https://go-stoic-quotes-production.up.railway.app/">Demo</a> to view.

## Example

<p align="center">
    <img src="https://go-stoic-quotes-production.up.railway.app/stoic-quote-svg?theme=random&t=1221">
</p>

---