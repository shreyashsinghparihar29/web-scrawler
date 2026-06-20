# Web-Scrawler

> Web-Scrawler is a lightweight HTML scraping and extraction tool written in Go. It allows users to fetch web pages, parse HTML documents, and extract structured data using familiar CSS selectors and jQuery-style syntax. The project is designed to be simple, extensible, and suitable for integration into automation pipelines, ETL workflows, monitoring systems, and data-processing applications.

---

# Key Features

### Multiple Interfaces

Supports three different ways of interaction:

* CLI for scripting and automation
* HTTP API for service integration
* Interactive Shell (REPL) for debugging and experimentation

### JavaScript-Based Extraction Engine

Instead of hardcoding selectors in Go, extraction logic is executed dynamically through an embedded JavaScript runtime.

Example:

```javascript
$("title").text()
$(".price").text()
$("meta[name=description]").attr("content")
```

This allows extraction rules to be modified without recompiling the application.

### jQuery-Style DOM Traversal

HTML documents are parsed into a DOM tree and exposed through a familiar jQuery-like API, making data extraction intuitive for frontend and scraping developers.

### Structured JSON Output

Extraction results are returned in a machine-readable JSON format, making integration with downstream services straightforward.

### Interactive Debugging Shell

A built-in REPL allows developers to inspect requests, responses, and extraction expressions interactively.

### HTTP Service Mode

The scraper can run as a standalone HTTP service, allowing other applications to perform extraction through REST APIs.

### Lightweight Deployment

The application is containerized using Docker and supports multi-stage builds for minimal image size and efficient deployment.

---

# Architecture

The project follows a layered architecture that separates presentation concerns from business logic.

```text
                 +------------------+
                 |      User        |
                 +--------+---------+
                          |
          ---------------------------------
          |               |               |
          |               |               |
          v               v               v

      CLI Interface   HTTP API      Interactive Shell

                \         |         /
                 \        |        /
                  \       |       /
                   v      v      v

                    Fetch Engine

                           |
                           v

                     HTTP Request

                           |
                           v

                     Target Website

                           |
                           v

                     HTML Response

                           |
                           v

                     DOM Parser
                       (GoQuery)

                           |
                           v

                 JavaScript Runtime
                        (Goja)

                           |
                           v

                  Extraction Execution

                           |
                           v

                     JSON Response
```

---

# Design Decisions

## Single Fetch + Local DOM Processing

One of the core architectural decisions in Web-Scrawler is downloading the webpage only once and performing all extraction operations against a locally parsed DOM tree.

### Why?

Instead of:

```text
Query 1 -> Fetch Webpage
Query 2 -> Fetch Webpage Again
Query 3 -> Fetch Webpage Again
```

Web-Scrawler performs:

```text
Fetch Once
    |
    v
HTML Snapshot
    |
    v
DOM Tree
    |
    +--> Query 1
    +--> Query 2
    +--> Query 3
```

### Benefits

#### Reduced Network Overhead

Network requests are significantly slower than in-memory DOM traversal.

```text
DOM Lookup     -> Microseconds
HTTP Request   -> Milliseconds
```

Fetching once dramatically improves performance.

#### Consistent Data Snapshot

All extraction expressions operate on the exact same version of the webpage.

Without a shared DOM snapshot:

```text
Title -> Version A
Price -> Version B
Author -> Version C
```

With a DOM snapshot:

```text
Title -> Same Page Version
Price -> Same Page Version
Author -> Same Page Version
```

This guarantees consistency across extracted fields.

#### Lower Server Load

Only a single request is sent to the target website regardless of the number of extraction expressions executed.

#### Better Reliability

Extraction remains independent of temporary network fluctuations after the page has been downloaded.

#### Efficient JavaScript Execution

The embedded JavaScript engine operates directly on an in-memory DOM representation, enabling fast and flexible extraction logic.

---

# Core Technologies

## Go

Primary implementation language used for performance, simplicity, and efficient concurrency.

## GoQuery

GoQuery provides jQuery-like HTML traversal and CSS selector support.

Example:

```go
doc.Find("title").Text()
```

## Goja

Goja is an embedded JavaScript engine written in Go.

It enables execution of user-defined extraction expressions at runtime.

Example:

```javascript
$("title").text()
```

## Fiber

Used for building the HTTP API layer.

Provides:

* Routing
* Middleware support
* Error handling
* Response compression

## Docker

Used for packaging and deployment.

Multi-stage builds are used to:

* Reduce image size
* Minimize attack surface
* Improve deployment efficiency

---

# Overview

You can use `web-scrawler` within your stack via `cli` or `http`.

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

---

# HTTP Mode

Run the HTTP server:

```bash
$ web-scrawler serve
```

Then:

```bash
$ curl http://localhost:8010/extract \
    -H "Content-Type: application/json" \
    -s \
    -d '{"url": "https://github.com/shreyashsinghparihar29/web-scrawler", "extractors": {"title": "$(\"title\").text()"}, "return_body": false, "user_agent": "CustomUserAgent"}'
```

---

# Interactive Shell

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

---

# Future Improvements

* Concurrent extraction execution
* Caching layer for repeated requests
* Proxy rotation support
* Rate limiting support
* Headless browser integration for JavaScript-heavy websites
* Authentication and session support
* Distributed scraping workers

---

# About

Made by **Shreyash Singh Parihar**
