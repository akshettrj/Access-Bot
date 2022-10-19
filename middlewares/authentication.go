package middlewares

import (
	"Access-Bot/state"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"golang.org/x/exp/slices"
)

func CheckAuthorized(b *gotgbot.Bot, c *ext.Context) bool {
	cfg := state.State.Config

	var (
		ownerID   = cfg.Telegram.OwnerID
		sudoUsers = cfg.Telegram.SudoUsers
		authChats = cfg.Telegram.AuthorizedChats
	)

	sender := c.EffectiveSender.User
	chat := c.EffectiveChat

	var (
		isOwner    = sender != nil && sender.Id == ownerID
		isSudo     = sender != nil && slices.Contains(sudoUsers, sender.Id)
		isAuthUser = sender != nil && slices.Contains(authChats, sender.Id)
		isAuthChat = chat != nil && slices.Contains(authChats, chat.Id)
	)

	if isOwner || isSudo || isAuthUser || isAuthChat {
		return true
	}

	if c.CallbackQuery != nil {
		b.AnswerCallbackQuery(c.CallbackQuery.Id, &gotgbot.AnswerCallbackQueryOpts{
			Text:      "Not authorized",
			ShowAlert: false,
			CacheTime: 60,
		})
	}

	return false
}
