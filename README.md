Wordle For Discord
==================

Play Wordle with a Discord bot.

## Gameplay Demo
TBD

## Functionality
TBD

## Developer setup

### Required setup
1. Golang installed - developed on 1.17.6
1. See `go.mod` file for the libraries for development

### Testing
Run unit tests:
```
go test -v ./...
```

Run benchmark tests:
```
go test -v -bench . ./...
```

### Running the bot yourself

#### Required
1. Register a discord bot application - see [intro docs](https://discord.com/developers/docs/intro)
1. Get the bot token, application ID, and a test server's guild ID
1. Clone the repo and navigate to it
1. Build the executable: `go build`
1. Run the executable with the credentials passed in, for example on Windows: `.\wordlego.exe --guild 12345 --token abc123 --app 98765`
1. You should see `Bot is up!` when the bot is stood up successfully, and properly registers the slash commands
1. Go to the server you used as an input to the bot executable, and test out the slash commands yourself
1. To stop the bot, just CTRL+C or SIGINTERRUPT the process.
