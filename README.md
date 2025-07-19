# gator

Guided Python Project from Boot.dev

## Prerequisites

- Go
- PostgreSQL
- goose (database migration tool)

## Clone the repository

```bash

git clone https://github.com/Waterbootdev/gator.git

```

## Create a `.gatorconfig.json` file in your home directotry

```json
{
  "db_url": "postgres://<user_name>:<password>@localhost:5432/gator?sslmode=disable",
}
```

## Navigate to the project directory

```bash

cd gator

```

## Do a up migration

```bash

goose postgres "postgres://<user_name>:<password>@localhost:5432/gator" up

```

## Install gator

```bash
go install .

```

## Usage

Once installed and configured, you can start using Gator with the following commands:

- `gator register <user_name>`: register as new database user
- `gator login <user_name>`: login as a registered user
- `gator reset`: reset the database to a clean state
- `gator users`: list all users in the database
- `gator feeds`: list all feeds in the database
- `gator addfeed <feed> <url>`: add and follow a new feed to the database for current user
- `gator follow <url>`: follow a feed for current user
- `gator unfollow <url>`: unfollow a feed for current user
- `gator following`: list followed feeds for current user
- `gator browse <limit>`: quick look at posts of the feeds followed by current user
- `gator agg <duration>`: scrape feeds(posts) followed by current user
