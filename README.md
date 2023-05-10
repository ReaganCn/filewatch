<p align="center"><img src="etc/assets/gopher.png" width="250"></p>
<p align="center">
  <a href="https://goreportcard.com/report/go.mongodb.org/mongo-driver"><img src="https://goreportcard.com/badge/go.mongodb.org/mongo-driver"></a>
  <!-- <a href="https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo"><img src="etc/assets/godev-mongo-blue.svg" alt="docs"></a>
  <a href="https://pkg.go.dev/go.mongodb.org/mongo-driver/bson"><img src="etc/assets/godev-bson-blue.svg" alt="docs"></a>
  <a href="https://www.mongodb.com/docs/drivers/go/current/"><img src="etc/assets/docs-mongodb-green.svg"></a> -->
</p>

# FileWatch

A file monitoring golang command line tool that can watch for modifications, creations and deletions of files and folders and send an event 

-------------------------
## Requirements

- Go 1.19 or higher. We aim to support the latest versions of Go.

- MongoDB 3.6 and higher.

-------------------------

## Usage

```
filewatch is a Go command-line tool that watches files and directories for changes.

Commands:
		set --path <path>		Sets the path to the file to watch
		run				Runs the file watcher
```

## Installation

Make sure you have docker, and docker compose installed.

Start the mongo db server with `docker run -d -p 27017:27017 --name mong-db mongo:latest`

Clone the project and add the below environments to `.env` in the project root folder:
```
OPEN_AI_KEY="your-open-ai-key"
TELEGRAM_BOT_TOKEN="your-telegram-bot-token"
MONGO_URI="your-mongo-db-URI" or simply "mongodb://127.0.0.1:27017/"
```

Install packages with `go mod download`

Execute `go run .`

## Contribution

For help with the bot, for now, you are free to create an issue. Contribution details are still work in progress.

-------------------------

## License

This Golang filewatch project is licensed under the [GNU](LICENSE).
