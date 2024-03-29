from config import config
from database.Chat import DBChat
from filters.control_group_filter import ControlGroupFilter
from filters.private_chat_filter import PrivateChatFilter
from helpers.status_control import get_invites_status
from telegram import (InlineKeyboardButton, InlineKeyboardMarkup, ParseMode,
                      Update)
from telegram.ext import CallbackContext, CommandHandler

class StartBot:

    def __init__(self):

        self.filters = ControlGroupFilter() | PrivateChatFilter()
        self.command_trigger: str = 'start'
        self.handler = CommandHandler(
            self.command_trigger,
            StartBot.callback_function,
            filters=self.filters,
            pass_args=True,
            run_async=True
        )

    @staticmethod
    def callback_function(update: Update, _: CallbackContext):

        bot_status = get_invites_status()

        if update.effective_chat.id == config.control_group_id:
            update.effective_message.reply_markdown("*Okay! I am alive.*")

        elif bot_status == 'off':
            update.effective_message.reply_markdown(
                text=config.requests_closed_message,
                reply_markup=InlineKeyboardMarkup([
                    [
                        InlineKeyboardButton(
                            text="Doubts/Problems",
                            url=f"{config.help_group_link}"
                        )
                    ]
                ])
            )

        else:

            chat_object = DBChat.objects(
                chat_id=update.effective_chat.id
            ).first()

            if not chat_object:
                chat_object = DBChat()
                chat_object.chat_id = update.effective_chat.id
                chat_object.chat_type = 'private'
                chat_object.save()

            update.effective_message.reply_photo(
                photo=config.private_chat_start_image,
                caption=f'''
*{config.group_name} Access Bot*

1. Only for members of {config.invite_channel_link}
2. Click on the Button "Join Google Group"
3. Click *"Ask to join the group"* Button and complete the steps
4. *Send the email used to this bot*


- *⚠️⚠️ This is voluntary work done by admins. So have patience*
- *❌❌ Spam requests will get you banned from the group.*
*P.S*: 📪📪 _Only_ `@gmail.com` _Emails Are Allowed._
''',
                parse_mode=ParseMode.MARKDOWN,
                reply_markup=InlineKeyboardMarkup([
                    [
                        InlineKeyboardButton(
                            text="Join Google Group",
                            url=config.google_group_link
                        )
                    ], [
                        InlineKeyboardButton(
                            text="Doubts/Problems",
                            url=f"{config.help_group_link}"
                        )
                    ]
                ])
            )


start_bot_command = StartBot()
