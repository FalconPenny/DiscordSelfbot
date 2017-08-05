package main

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

var (
	Config Configuration
)

func main() {
	_, err := toml.DecodeFile("config.toml", &Config)
	if os.IsNotExist(err) {
		log.Infoln("You need to change the config for the bot to work.")
		log.Infoln("It has been created and saved.")
		Config = Configuration{
			LogMode: uint32(log.GetLevel()),
			Prefix:  "--",
			Token:   "SET THE TOKEN HERE",
		}
		WriteConfig(&Config)
		return
	}

	bot, err := discordgo.New(Config.Token)
	if err != nil {
		log.Errorln("Couldn't create discordgo bot!", err)
		return
	}
	defer func() {
		err = bot.Close()
		if err != nil {
			log.Debugln("An error occurred during close.", err)
		}
	}()
	_, err = bot.User("@me") // To make sure the token isn't invalid.
	if err != nil {
		log.Errorln("Token is invalid! Update it in the config.")
		log.Debugln(err)
		return
	}


}

func WriteConfig(config *Configuration) {
	file, err := os.OpenFile("config.toml", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0660)
	if err != nil {
		log.Errorln("Error while writing to configuration!", err)
		return
	}
	defer file.Close()
	err = toml.NewEncoder(file).Encode(config)
	if err != nil {
		log.Errorln("Error while writing to configuration!", err)
		return
	}
}
