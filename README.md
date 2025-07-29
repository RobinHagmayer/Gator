# Gator

Gator 🐊 is a guided project from [boot.dev](https://boot.dev): a lightweight Go-powered CLI that tames your RSS chaos.
Subscribe to blogs, podcasts, or news sites, auto-fetch new posts into PostgreSQL, and browse bite-sized, link-enabled summaries—all from your terminal.

## Features

- Add & remove RSS feeds
- Store fetched posts in PostgreSQL
- Follow/unfollow feeds added by others
- View post summaries with links to full content
- Continuous background fetching

## Prerequisites

- Go 1.23+
- PostgreSQL

## Installation

How to install the Gator CLI on Linux:

1. Clone the repository `git clone https://github.com/RobinHagmayer/Gator.git`
2. Change into the directory and build the executable `go build -o gator`

## Config

To use this tool we need a `.gatorconfig.json` file inside your home directory.
It stores the database connection and the current user.
The application just needs the database connection to work.
Change the variables to your configuration.

```json
{
  "db_url": "postgres://username:password@localhost:5432/database?sslmode=disable"
}
```

## Usage

TODO

## Learning Goals

- Integrate a Go app with PostgreSQL
- Write & run SQL migrations with goose
- Generate typesafe SQL using sqlc
- Build a long-running service that polls RSS feeds
