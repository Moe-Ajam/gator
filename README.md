# rss-blog-aggregator

A blog aggregator built in Go

## Instructions

You will need Postgres and Go installed on your machine to run this app

You can install it by running `go install https://github.com/Moe-Ajam/gator`

You will need to create a file in the root of your directory to hold the configuration for the app, on MacOs thats in `~/.gatorconfig.json`

The configuration file should contain the below:

```json
{
  "db_url": "postgres://mahmoudajam:@localhost:5432/gator?sslmode=disable",
  "current_user_name": "kahya"
}
```

Where `db_url` is the url to your local postgres database and the `current_user_name` is the currently logged in user (this will be filled automatically by the application once you login using the `login` command)
