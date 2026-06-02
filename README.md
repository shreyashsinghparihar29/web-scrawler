Web-Scrawler
============
> Web-Scrawler is a simple HTML scraping tool. If you're familiar with CSS selectors and jQuery-style syntax, you can start extracting data immediately. The goal is to remain lightweight, easy to use, and suitable for integration into larger automation and data-processing systems.

Overview
========
> You can use `web-scrawler` within your stack via `cli` or `http`.

```bash
# CLI usage examples

# Extracting the title and description from the GitHub repository page
$ web-scrawler extract \
    -u "https://github.com/shreyashsinghparihar29/web-scrawler" \
    -x title='$("title").text()' \
    -x description='$("meta[name=description]").attr("content")'

# Same thing but with a custom user agent
$ web-scrawler extract \
    -u "https://github.com/shreyashsinghparihar29/web-scrawler" \
    -ua "OptionalCustomUserAgent" \
    -x title='$("title").text()' \
    -x description='$("meta[name=description]").attr("content")'

# Same thing but returning the response body for debugging purposes
$ web-scrawler extract \
    --return-body \
    -u "https://github.com/shreyashsinghparihar29/web-scrawler" \
    -x title='$("title").text()' \
    -x description='$("meta[name=description]").attr("content")'
```

> For `http` usage, run the HTTP server and use any HTTP client to interact with it.

```bash
# Running the HTTP server
# By default it listens on address ":8010" which equals "0.0.0.0:8010"
# For more information execute `$ web-scrawler help`
$ web-scrawler serve

# Then in another shell execute:
$ curl http://localhost:8010/extract \
    -H "Content-Type: application/json" \
    -s \
    -d '{"url": "https://github.com/shreyashsinghparihar29/web-scrawler", "extractors": {"title": "$(\"title\").text()"}, "return_body": false, "user_agent": "CustomUserAgent"}'
```

> For debugging, there is an interactive `shell`.

```bash
$ web-scrawler shell -u https://github.com/shreyashsinghparihar29/web-scrawler

➜ (web-scrawler) > $("title").text()
GitHub - shreyashsinghparihar29/web-scrawler: A lightweight HTML scraping and extraction tool for collecting structured web data

➜ (web-scrawler) > request.url
https://github.com/shreyashsinghparihar29/web-scrawler

➜ (web-scrawler) > response.status_code
200

➜ (web-scrawler) > response.url
https://github.com/shreyashsinghparihar29/web-scrawler

➜ (web-scrawler) > response.body
<html>.....
```




About
======
Made by Shreyash Singh Parihar
