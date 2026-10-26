# gator cli

A simple Go-based CLI for managing and browsing blog feeds. It lets you register users, add RSS feeds, follow feeds, aggregate updates from them, and browse the latest posts in a personalized feed.

The project is built as a command-line tool in Go and uses PostgreSQL as its main database. SQL access is generated with `sqlc`, while feed parsing is handled with Go's standard XML support and HTTP client.

## What this project uses

- Go 1.27.1
- PostgreSQL
- `lib/pq` for PostgreSQL driver support
- `google/uuid` for UUID generation
- `sqlc` for type-safe SQL code generation
- Go standard library packages:
  - `net/http`
  - `encoding/xml`
  - `database/sql`
  - `encoding/json`
  - `os`
  - `time`
  - `context`
- RSS feed parsing using XML-based feed structures
- Local config persistence in `~/.gatorconfig.json`

## Project structure

```text
.
├── go.mod
├── go.sum
├── main.go
├── sqlc.yaml
├── sql/
│   ├── queries/
│   │   ├── feed_follows.sql
│   │   ├── feeds.sql
│   │   ├── posts.sql
│   │   └── users.sql
│   └── schema/
│       ├── 001_users.sql
│       ├── 002_feeds.sql
│       ├── 003_feed_follows.sql
│       ├── 004_feeds_last_fetched.sql
│       └── 005_posts.sql
├── internal/
│   ├── command/
│   │   ├── addFeedHandler.go
│   │   ├── aggHandler.go
│   │   ├── browseHandler.go
│   │   ├── command.go
│   │   ├── feedsHandler.go
│   │   ├── followHandler.go
│   │   ├── followingHandler.go
│   │   ├── loginHandler.go
│   │   ├── middlewareLoggedIn.go
│   │   ├── registerHandler.go
│   │   ├── resetHandler.go
│   │   ├── scrapeFeeds.go
│   │   ├── state.go
│   │   ├── unfollowHandler.go
│   │   └── usersHandler.go
│   ├── config/
│   │   ├── config.go
│   │   └── read_write.go
│   ├── database/
│   │   ├── db.go
│   │   ├── feed_follows.sql.go
│   │   ├── feeds.sql.go
│   │   ├── models.go
│   │   ├── posts.sql.go
│   │   └── users.sql.go
│   └── rss/
│       └── rss.go
└── .gitignore
```

## Features

- User registration and login
- Store a current user in a config file
- Add RSS feeds to the app
- Follow feeds as a user
- Display all feeds and followed feeds
- Browse posts for the current user
- Aggregate feeds on a timer using `agg`
- Scrape RSS content and save posts in PostgreSQL

## Database design

The app uses PostgreSQL tables for:

- `users`
- `feeds`
- `feed_follows`
- `posts`

The schema is managed via SQL migration files in `sql/schema/`, and the Go database layer is generated with `sqlc` from the SQL in `sql/queries`.

## Configuration

The CLI stores configuration in `~/.gatorconfig.json` and stores values like:

```json
{
  "db_url": "postgresql://username:password@localhost:5432/gator",
  "current_user_name": "your_username"
}
```

`db_url` is used to connect to PostgreSQL, and `current_user_name` tracks the active logged-in user.

## Prerequisites

Before running the project, make sure you have:

- Go installed
- PostgreSQL running locally or on a server
- A database created for the app
- `sqlc` installed if you want to regenerate the database models

## Setup

1. Clone the repo:

```bash
git clone https://github.com/abdullahshahad-w/blog-aggregator-go.git
cd blog-aggregator-go
```

2. Create a PostgreSQL database. For example:

```bash
createdb gator
```

3. Set the database URL in `~/.gatorconfig.json`:

```bash
cat > ~/.gatorconfig.json <<'EOF'
{
  "db_url": "postgresql://postgres:postgres@localhost:5432/gator",
  "current_user_name": ""
}
EOF
```

4. Run the app:

```bash
go run .
```

## How to use gator cli

### Register a user

```bash
go run . register alice
```

This creates a user named `alice` and makes them the active user.

### Log in as a user

```bash
go run . login alice
```

### Add a feed

```bash
go run . addfeed "My Blog" "https://example.com/feed.xml"
```

This creates a feed and automatically adds a follow record for the current user.

### View all feeds

```bash
go run . feeds
```

### Follow an existing feed by URL

```bash
go run . follow "https://example.com/feed.xml"
```

### View followed feeds

```bash
go run . following
```

### View all users

```bash
go run . users
```

### Browse posts for the current user

```bash
go run . browse
```

Optional limit:

```bash
go run . browse 10
```

### Start feed aggregation

```bash
go run . agg 1m
```

This starts a loop that fetches feeds every minute. The `agg` command uses the RSS parser and stores new posts in the database.

### Reset the app state

```bash
go run . reset
```

## Command reference

| Command | Description |
| --- | --- |
| `register <username>` | Create a user and set them as the current user |
| `login <username>` | Log in to an existing user |
| `users` | List all users |
| `addfeed <name> <url>` | Create a new feed and follow it |
| `feeds` | List all registered feeds |
| `follow <feed-url>` | Follow an existing feed |
| `following` | Show followed feeds for the active user |
| `unfollow <feed-url>` | Remove a follow relationship |
| `browse [limit]` | Show posts for the current user |
| `agg <duration>` | Poll RSS feeds on a timer |
| `reset` | Reset or clean up app state |

## Notes on the implementation

This CLI is intentionally compact and practical:

- The app entrypoint is `main.go` and wires the command system together.
- Each command sits under `internal/command`.
- PostgreSQL access is generated into `internal/database` using `sqlc`.
- Feed parsing is isolated under `internal/rss`.
- App config is saved locally so the user does not have to re-login each time.

## Example workflow

```bash
go run . register alice
go run . login alice
go run . addfeed "Go Blog" "https://go.dev/blog/feed.atom"
go run . browse 5
go run . agg 5m
```

## License

This project does not currently include a license file. Add one if you want to publish or distribute it publicly.

## Summary

`gator cli` is a lightweight RSS blog aggregator built in Go. It uses PostgreSQL for persistence, `sqlc` for data access, and a simple custom command framework to provide a terminal-based feed experience.
