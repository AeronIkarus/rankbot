# rankbot

Very simple rank bot for [mee6](https://mee6.xyz) on discord.

## Why?
With mee6, each time you post a message you get a random number of XP between 5 and 10.

To avoid flood, mee6 only allows you to gain xp once per minute.

This bot sends out a message every minute allowing you to rank up very fast.

## Install
```zsh
go get github.com/AeronIkarus/rankbot
```

## Usage
```zsh
[$] rankbot --help
Usage of rankbot:
  -chan string
        guild channel
  -email string
        account email
  -guild string
        account guild
  -msg string
        message to be sent (default "_")
  -pass string
        account password
```

### Example usage
```zsh
rankbot -email="me@domain.com" -pass="secret" -guild="server" -chan="channel"
```
