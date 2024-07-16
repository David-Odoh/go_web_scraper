# Concurrent Web Scraper

This project is a concurrent web scraper written in Go. It fetches data from multiple URLs concurrently, processes the data, and aggregates the results. The solution utilizes goroutines and channels to achieve concurrency and synchronization.

## Features

- Concurrently fetches data from multiple URLs.
- Processes the fetched data.
- Aggregates and displays the results.
- Handles errors gracefully.

## How It Works

1. **Fetcher**: Fetches data from URLs and sends results to a channel.
2. **Processor**: Processes the fetched data and sends results to a channel.
3. **Aggregator**: Aggregates the processed results and prints them.

## Running the Project

To run the project, execute the following commands:

```sh
go run main.go
