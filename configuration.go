package main

type Configuration struct {
	Token string `toml:"bot.token"`
	Prefix string `toml:"bot.prefix"`

	LogMode uint32 `toml:"bot.logging"`
}