import logging
from config import config
from telegram.ext import Updater
from database.connect import db_init

from commands.ban_chat import ban_chat_command
from commands.purge_message import purge_messages_command
from commands.start import start_bot_command
from commands.unban_chat import unban_chat_command
from commands.pin_message import pin_message_command
from commands.find_member import find_member_command
from commands.find_req_on_gg import find_request_command

from forward_handlers.private_chat import private_message_forwarder
from forward_handlers.control_group_chat import control_group_message_forwarder

from callbacks.approve_user import approval_callback_handler
from callbacks.go_back import approval_go_back_callback_handler
from callbacks.warning_message import warn_message_callback_handler
from callbacks.no_request_found import no_request_callback_handler
from callbacks.go_back import no_request_go_back_callback_handler
from callbacks.undo_no_req import undo_no_request_callback_handler
from callbacks.banned_message import banned_message_callback_handler

from wserver import keep_alive

def main():
    updater = Updater(config.bot_token, use_context=True)
    dispatcher = updater.dispatcher

    for command_handler in command_handlers:
        dispatcher.add_handler(command_handler.handler)

    for callback_handler in callback_handlers:
        dispatcher.add_handler(callback_handler.handler)

    updater.start_polling()


if __name__ == "__main__":

    logging.basicConfig(
        format='%(asctime)s - %(name)s - %(levelname)s - %(message)s',
        level=logging.INFO
    )

    db_init(config)

    command_handlers = [
        ban_chat_command,
        purge_messages_command,
        start_bot_command,
        private_message_forwarder,
        control_group_message_forwarder,
        unban_chat_command,
        pin_message_command,
        find_member_command,
        find_request_command
    ]

    callback_handlers = [
        approval_callback_handler,
        approval_go_back_callback_handler,
        no_request_go_back_callback_handler,
        warn_message_callback_handler,
        no_request_callback_handler,
        undo_no_request_callback_handler,
        banned_message_callback_handler
    ]

    keep_alive()
    main()
