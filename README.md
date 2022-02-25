Wordle For Discord
==================

[![codecov](https://codecov.io/gh/SaxyPandaBear/wordle-in-discord/branch/main/graph/badge.svg?token=20EO8DCJPJ)](https://codecov.io/gh/SaxyPandaBear/wordle-in-discord)

Play Wordle with a Discord bot.

## Gameplay Demo
TBD

## Functionality

### Implementation
I'm following with a minimal understanding of Wordle, taking note
of [double letter edge cases](https://www.reddit.com/r/wordle/comments/ry49ne/illustration_of_what_happens_when_your_guess_has/). 

I am using the lists of words found [here](https://github.com/CrispyConductor/wordle-solver/tree/71b9f7c4c7f9e7fe57b7df85bb624265b0b8e17d). 
It seems that Wordle keeps track of a static, deterministic list of words for solutions, and also
maintains a separate list of allowed guesses.

### Commands
TBD

## Developer setup

### Required setup
1. Golang installed - developed on 1.17.6
1. See `go.mod` file for the libraries for development

### dotenv secrets (optional)
This is not a required part of setup, and it's not actively used in the development process,
but I like to use `.env` for keeping track of my sensitive credentials while developing. As
such, `.env` is ignored in Git, and `.env.template` is provided as a template for contributors
to keep track of the required secrets.

I am not using any dotenv style library to lookup/use the secrets currently for development. 
The bot gets run by passing in the secrets as argument flags to the executable currently.

TODO: Change to default to environment variables when this is ready to deploy.

### Testing
Run unit tests:
```bash
go test -v ./...
```

Run benchmark tests:
```bash
go test -v -bench . ./...
```

### Running the bot yourself

#### Required
1. Register a discord bot application - see [intro docs](https://discord.com/developers/docs/intro)
1. Get the bot token, application ID, and a test server's guild ID
1. Clone the repo and navigate to it
1. Build the executable: `go build`
1. Run the executable with the credentials passed in, for example on Windows: `.\wordlego.exe --guild 12345 --token abc123 --app 98765` (order doesn't matter for the flags)
1. You should see `Bot is up!` when the bot is stood up successfully, and properly registers the slash commands
1. Go to the server you used as an input to the bot executable, and test out the slash commands yourself
1. To stop the bot, just CTRL+C or SIGINTERRUPT the process.
