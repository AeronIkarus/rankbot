package main

import (
	"flag"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/AeronIkarus/color/log"
)

var (
	email    = flag.String("email", "", "account email")
	pass     = flag.String("pass", "", "account password")
	guild    = flag.String("guild", "", "guild (server) to join")
	channel  = flag.String("chan", "", "channel to join")
	message  = flag.String("msg", "_", "message to be sent")
	interval = flag.Int64("int", 60, "interval between messages in seconds")
)

func main() {
	flag.Parse()
	if *email == "" || *pass == "" {
		log.Fatal("please provide an email and password")
	}
	s, err := discordgo.New(*email, *pass)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("logged in")

	g := findGuild(s)
	if g == nil {
		log.Fatal("could not find guild")
	}
	id := findChannel(s, g)
	if id == "" {
		log.Fatal("could not find channel")
	}

	for t := time.Tick(time.Duration(*interval) * time.Second); ; <-t {
		if _, err := s.ChannelMessageSend(id, *message); err != nil {
			log.Print(err)
		} else {
			log.Print("sent message")
		}
	}
}

func findGuild(s *discordgo.Session) *discordgo.UserGuild {
	gs, err := s.UserGuilds()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("got guilds")
	for _, g := range gs {
		if g.Name == *guild {
			log.Print("found guild")
			return g
		}
	}
	return nil
}

func findChannel(s *discordgo.Session, g *discordgo.UserGuild) string {
	chs, err := s.GuildChannels(g.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("got channels")
	for _, ch := range chs {
		if ch.Name == *channel {
			log.Print("found channel")
			return ch.ID
		}
	}
	return ""
}
