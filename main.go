package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"Access-Bot/database"
	mw "Access-Bot/middlewares"
	"Access-Bot/state"
	"Access-Bot/utils"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func main() {
	// Load configuration
	cfg := state.State.Config
	if len(os.Args) > 1 {
		cfg.Path = os.Args[1]
	} else {
		cfg.Path = "config.yaml"
	}
	cfg.LoadConfig()

	// Connect to database
	db, err := database.Connect()
	if err != nil {
		log.Fatalln("Could not connect to the database : ", err)
	}
	state.State.Database = db

	// Create the bot
	// TODO: Add HTML Parse Mode in all and Authorization checking
	bot, err := gotgbot.NewBot(cfg.Telegram.BotToken, &gotgbot.BotOpts{
		Client: http.Client{},
		DefaultRequestOpts: &gotgbot.RequestOpts{
			Timeout: gotgbot.DefaultTimeout,
			APIURL:  cfg.Telegram.ApiURL,
		},
	})
	if err != nil {
		log.Fatalln("Could not initialize bot : ", err)
	}
	state.State.Bot = bot
	state.State.StartTime = time.Now()

	// Clean downloads
	os.RemoveAll("tempdata")

	bot.UseMiddleware(mw.SendWithoutReply)
	bot.UseMiddleware(mw.ParseAsHTML)

	updater := ext.NewUpdater(&ext.UpdaterOpts{
		ErrorLog: nil,
		DispatcherOpts: ext.DispatcherOpts{
			// If an error is returned by a handler, log it and continue going.
			Error: func(_ *gotgbot.Bot, _ *ext.Context, err error) ext.DispatcherAction {
				fmt.Println("an error occurred while handling update:", err.Error())
				return ext.DispatcherActionNoop
			},
			MaxRoutines: ext.DefaultMaxRoutines,
		},
	})
	dispatcher := updater.Dispatcher
	state.State.Dispatcher = dispatcher

	dispatcher.AddHandler(handlers.NewCommand(
		"start",
		func(b *gotgbot.Bot, c *ext.Context) error {
			if !mw.CheckAuthorized(b, c) {
				return nil
			}
			_, err := b.SendMessage(
				c.EffectiveChat.Id,
				"Hi, I am alive!",
				&gotgbot.SendMessageOpts{
					ReplyToMessageId: c.EffectiveMessage.MessageId,
				})
			return err
		}))

	utils.RegisterBotCommand(bot, state.State.Commands...)

	// Start the poller
	err = updater.StartPolling(bot, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: gotgbot.GetUpdatesOpts{
			Timeout: 9,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 10,
			},
		},
	})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}
	log.Printf("Telegram bot logged in as %s [ @%s ]\n", bot.FirstName, bot.Username)

	// Idle, to keep updates coming in, and avoid bot stopping.
	updater.Idle()
}
