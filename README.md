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
- PostgreSQL 16+

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

The general way to use Gator is `gator <command> [args...]`.

The usual way to use the CLI works like this:

1. Register a new user: `gator register <username>`
2. Add a new RSS feed: `gator addfeed <feed_name> <feed_url>`
3. Start the aggregator: `gator agg <time_between_requests>` (`gator agg 1m`)
4. View the posts: `gator browse <posts_limit>`

All available commands:

| Command                                | Description                         |
| -------------------------------------- | ----------------------------------- |
| `gator register <user_name>`           | Register a new user                 |
| `gator login <user_name>`              | Login an existing user              |
| `gator users`                          | List all users                      |
| `gator addfeed <feed_name> <feed_url>` | Add a new feed and follow it        |
| `gator feeds`                          | List all feeds                      |
| `gator follow <feed_url>`              | Follow an existing feed             |
| `gator unfollow <feed_url>`            | Unfollow a feed                     |
| `gator following`                      | List all followed feeds             |
| `gator agg <time_between_requests>`    | Start the aggregation of blog posts |
| `gator browse <limit>`                 | Browse the latest posts             |

## Learning Goals

- Integrate a Go app with PostgreSQL
- Write & run SQL migrations with goose
- Generate typesafe SQL using sqlc
- Build a long-running service that polls RSS feeds
