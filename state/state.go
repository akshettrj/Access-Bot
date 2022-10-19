package state

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"gorm.io/gorm"
)

type state struct {
	Bot        *gotgbot.Bot
	Dispatcher *ext.Dispatcher
	Config     *Config
	Database   *gorm.DB
	Commands   []gotgbot.BotCommand

	StartTime time.Time
}

var State state

func init() {
	State.Config = &Config{}
}
