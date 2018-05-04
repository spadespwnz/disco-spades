package main

import (
	"fmt"
	"github.com/DavidBelicza/TextRank"

	"github.com/JesusIslam/tldr"
	"github.com/bwmarrin/discordgo"
	"github.com/subosito/gotenv"
	"os"
	"os/signal"
	"syscall"
)

var codyCount int
var sentanceCount = 1
var tldrCount = 5
var message string

var tylerRant string
var testText string = "he should run as a independent???? interesting idea man but i think he'd need party support to win. republican's prolly his best best TBH especially if democrats win in 2020. imagine a republicn party in the wake of donald trump. what do they do? there's a chance if they're smart they become more progressive and shift politics to the left. how to do that? new leadership from Kanye West. a bold new direction for the republic party. trump is literally the republican party. he exposes them thats how republican he is. thats whyt hey hate him. he's pushing the exact agenda they want but he makes it too obvious. but after this there will be no republican party left if we're lucky. trump is destroying them from the inside by being so cancer and setting such a terrible example. and so they could rebuild with Kanye at the head. they could be the flashy celebrity good times lit party. while the democrats remain boring and toothless best bet* trump is so republican that it might embarass the republican party into becoming liberal/socialist with kanye as their new leadership and the face of their new direction with the principles of bernie sanders hell ya my guy thats a man to vote for i would register as republican for that"

func init() {
	gotenv.Load()
}
func main() {
	dg, err := discordgo.New("Bot " + os.Getenv("DG_TOKEN"))
	if err != nil {
		fmt.Println("error creaing discord session,", err)
		return
	}

	dg.AddHandler(MessageHandler)

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()

}

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
	if m.Author.Username == "spades" {
		if m.Content == "!chanId" {
			s.ChannelMessageSend(m.ChannelID, m.ChannelID)
		}

	}
	if m.Author.Username == "spades" && m.ChannelID == "438599733309865984" {
		if m.Content == "!sum" {
			bag := tldr.New()
			result, err := bag.Summarize(testText, 1)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, err.Error())
				return
			}
			s.ChannelMessageSend(m.ChannelID, result)
		}
	}
	if m.Author.Username == "fmjester" && m.ChannelID == "390339861816803356" {
		codyCount++
		message = message + m.Content + "."
		fmt.Println("Count: %d", codyCount)
		if codyCount >= tldrCount {

			tr := textrank.NewTextRank()
			rule := textrank.NewDefaultRule()
			lang := textrank.NewDefaultLanguage()
			algo := textrank.NewDefaultAlgorithm()

			tr.Populate(message, lang, rule)
			tr.Ranking(algo)
			phrases := textrank.FindPhrases(tr)
			words := make([]string, 0)
			for i := 0; i < len(phrases); i++ {
				if len(words) >= 10 {
					break
				}
				if !stringInSlice(phrases[i].Left, words) {
					words = append(words, phrases[i].Left)
				}
				if !stringInSlice(phrases[i].Right, words) {
					words = append(words, phrases[i].Right)
				}
			}
			s.ChannelMessageSend(m.ChannelID, stringFromSlice(words))

			codyCount = 0
			message = ""
		}
	} else {
		codyCount = 0
		message = ""
	}
}
func stringFromSlice(list []string) string {
	var msg string
	for _, b := range list {
		msg += b + " "
	}
	return msg
}
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
