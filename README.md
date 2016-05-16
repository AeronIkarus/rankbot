# rankbot

Very simple rank bot for [mee6](https://mee6.xyz) on discord.

## Why?
With mee6, each time you post a message you get a random number of XP between 5 and 10.

To avoid flood, mee6 only allows you to gain xp once per minute.

This bot sends out a message every minute allowing you to rank up very fast.

## Install
```
go get github.com/nhooyr/rankbot
```

## Usage
```
rankbot --help
```

### Example usage
```
rankbot -email="me@domain.com" -pass="secret" -guild="server" -chan="channel" -int=1
```
