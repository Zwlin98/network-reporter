package main

import (
	"fmt"
	"os/exec"

	viper "github.com/spf13/viper"
	tele "gopkg.in/telebot.v3"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.config/network-reporter")
	viper.AddConfigPath("/etc/network-reporter")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func main() {
	bot_token := viper.GetString("BOT_TOKEN")
	chat_id := viper.GetInt64("CHAT_ID")
	device := viper.GetString("DEVICE")
	header := viper.GetString("HEADER")

	perf := tele.Settings{
		Token: bot_token,
	}

	bot, err := tele.NewBot(perf)

	if err != nil {
		panic(err)
	}

	cmd := exec.Command("vnstati", "--headertext", header, "-d", "-hs", "-i", device, "-o", "/tmp/vnstat.png")

	err = cmd.Run()

	if err != nil {
		panic(err)
	}

	photo := &tele.Photo{File: tele.FromDisk("/tmp/vnstat.png")}

	bot.Send(tele.ChatID(chat_id), photo)
}
