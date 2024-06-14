package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

const prefix string = "!cope"

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	secret := os.Getenv("DISCORD_BOT_ID")

	sess, err := discordgo.New(string(secret))
	if err != nil {
		log.Fatal(err)
	}

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		args := strings.Split(m.Content, " ")

		if args[0] != prefix {
			return
		}

		if args[1] == "test" {
			s.ChannelMessageSend(m.ChannelID, "initiated!")
		}

		if args[1] == "boss" {
			proverbs := []string{
				"Do Vorkath",
				"Do Leviathan",
				"Do Barrows",
				"Run Tombs of Amascut",
				"Run Chambers of Xeric",
				"Run Theatre of Blood",
				"Do Giant Mole",
				"Do Dagannoth Kings",
				"Do Bandos",
				"Do Zulrah",
				"Do Muspah",
				"Do Duke Sucellus",
				"Do Vardorvis",
				"Do Corrupted Gauntlet",
				"Run Colosseum",
				"Run Inferno",
				"Do slayer instead loser",
			}

			selection := rand.Intn(len(proverbs))

			// ### This section of code is if we want to make an embedded message ###
			// author := discordgo.MessageEmbedAuthor{
			// 	Name: "That dude Jeff",
			// 	URL:  "https://github.com/ItsJeff510",
			// }

			// embed := discordgo.MessageEmbed{
			// 	Title:  proverbs[selection],
			// 	Author: &author,
			// }

			s.ChannelMessageSend(m.ChannelID, proverbs[selection])
		}
	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = sess.Open()

	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	fmt.Println("The bot is online!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
