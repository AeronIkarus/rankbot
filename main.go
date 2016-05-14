package main

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/nhooyr/color/log"
)

const (
	email   = "anmol@aubble.com"
	pass    = "566348aA"
	guild   = "ezekiel"
	channel = "izi"
	message = "_"
)

func main() {
	s, err := discordgo.New(email, pass)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("logged in")
	g := findGuild(s)
	if g == nil {
		log.Fatalln("could not find guild")
	}
	id := findChannel(s, g)
	if id == "" {
		log.Fatalln("could not find channel")
	}
	sendLoop(s, id)
}

func findGuild(s *discordgo.Session) *discordgo.Guild {
	gs, err := s.UserGuilds()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("got guilds")
	for _, g := range gs {
		if g.Name == guild {
			log.Println("found guild")
			return g
		}
	}
	return nil
}

func findChannel(s *discordgo.Session, g *discordgo.Guild) string {
	chs, err := s.GuildChannels(g.ID)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("got channels")
	for _, ch := range chs {
		if ch.Name == channel {
			log.Println("found channel")
			return ch.ID
		}
	}
	return ""
}

func sendLoop(s *discordgo.Session, id string) {
	for t := time.Tick(time.Minute); ; <-t {
		_, err := s.ChannelMessageSend(id, message)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("sent message")
	}
}
