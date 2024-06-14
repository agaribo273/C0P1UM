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

		if args[1] == "help" {
			commands := []string{
				"**help**: Displays all the current integrated commands",
				"**activity**: Gives a random activity to do in Old School Runescape",
				"**boss**: Gives a random boss to do in Old School Runescape",
				"**test**: Replies with initiated (only meant as a test)",
			}

			commandDisplay := strings.Join(commands, "\n")

			s.ChannelMessageSend(m.ChannelID, commandDisplay)
		}

		if args[1] == "test" {
			s.ChannelMessageSend(m.ChannelID, "initiated!")
		}

		if args[1] == "activity" {
			activities := []string{
				"Get a smithing level.",
				"Get any unique from any wilderness boss.",
				"Get 1 of any skilling level of your choice.",
				"Run LMS until you achieve at least 2 kills.",
				"Attempt a Combat achievement of your choice, you have 10 attempts!",
				"Get a construction level through Mahogany Homes.",
				"Run 3 games of Guardians of the Rift.",
				"Get a sub 1.15 in Zulrah within 10 attempts.",
				"Do an hour of crafting using any method.",
				"Create a fashionscape outfit and attempt to get a complement on it.",
				"Earn 15 golden nuggets at Motherlode Mine",
				"Get 10 unidentified minerals at the mining guild.",
				"Do 5 different bosses with 1 inventory without banking.",
				"Follow someone until they ask you something.",
			}

			selection := rand.Intn(len(activities))

			s.ChannelMessageSend(m.ChannelID, activities[selection])
		}

		if args[1] == "boss" {
			bosses := []string{
				"Run Tombs of Amascut",
				"Run Chambers of Xeric",
				"Run Theatre of Blood",
				"Run Colosseum",
				"Run Inferno",
				"Do Vorkath",
				"Do Leviathan",
				"Do Barrows",
				"Do Giant Mole",
				"Do Dagannoth Kings",
				"Do Bandos",
				"Do Zulrah",
				"Do Muspah",
				"Do Duke Sucellus",
				"Do Vardorvis",
				"Do Corrupted Gauntlet",
				"Do slayer instead loser",
			}

			selection := rand.Intn(len(bosses))

			// ### This section of code is if we want to make an embedded message ###
			// author := discordgo.MessageEmbedAuthor{
			// 	Name: "That dude Jeff",
			// 	URL:  "https://github.com/ItsJeff510",
			// }

			// embed := discordgo.MessageEmbed{
			// 	Title:  proverbs[selection],
			// 	Author: &author,
			// }

			s.ChannelMessageSend(m.ChannelID, bosses[selection])
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
