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
There is one main slash command, `/wordle`. This is the main hook into the game. 

From here, functionality is divided into sub "actions". The `action` parameter is required for the
slash command, and is a choice from a static set of allowed string inputs.

| Action | Description                               |
| ------ | ----------------------------------------- |
| start  | Initiates a new game for the user         |
| stop   | Cancels an ongoing game for the user      |
| guess  | Execute a single guess for an active game |
| help   | Prints help info for the command          |

#### Optional arguments
Because all of the actions are in the umbrella of the `wordle` command, all of the sub parameters
are lumped into this command, and therefore are not technically required by the slash command on discord,
but are required in the bot logic.

* `word`: The word to guess
    * Required for: `guess`
* `puzzle-num`: Specific puzzle to attempt. If not provided, defaults to the current day's word
    * Optional for: `start`
* `max-guesses`: Configuration for the maximum number of guesses for the puzzle when starting a new game
    * Optional for: `start`

## Developer setup

### Required setup
1. Golang installed - developed on 1.17.6
1. See `go.mod` file for the libraries for development

### dotenv secrets
This uses the [godotenv port](https://github.com/joho/godotenv) in order to load secrets as environment
variables. Look at the `./env.template` file as a reference. The main code loads in the environment variables
so that it can default to those values, rather than relying on the `flag` inputs, i.e.:
```bash
./wordlego --guild 123 --app 456 --token abc123
```

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
    * Note that you can optionally configure the `.env` file to store the secrets so you don't have to pass them in on the command line.
1. You should see `Bot is up!` when the bot is stood up successfully, and properly registers the slash commands
1. Go to the server you used as an input to the bot executable, and test out the slash commands yourself
1. To stop the bot, just CTRL+C or SIGINTERRUPT the process.
