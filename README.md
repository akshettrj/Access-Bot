# Access Bot

Join [@AccessBotHelp](https://t.me/AccessBotHelp) for help.

## Deploy

1. Copy `sample_config.py` to `config.py`
2. Fill out the variables in `config.py`
3. Make the bot admin in log channel, control group and main group
4. Send `/start` in the control group to start the bot
5. You know the rest

## Deploy and RUN of REPL.IT for free 😂

[<img height=15 src="https://repl.it/badge/github/akshettrj/Access-Bot">](https://repl.it/github/akshettrj/Access-Bot)

1. Copy `sample_config.py` to `config.py`
2. Fill out the variables in `config.py`
3. Make the bot admin in log channel, control group and main group
4. Tap On the Green **`▶️ RUN`** logo and wait for the bot setup to finish and run the bot with backend ...
5. Copy the Repl.it Web--App url from above of the repl.it terminal... it should be something like **`https://Access-Bot.example.repl.co`**
6. Sign-up / log-in and go to https://console.cron-job.org and paste your repl.it web app url then save the job also don't forget to set the interval as 5 minutes...
7. Send `/start` in the control group to start the bot
8. You know the rest

### Same method can be done on heroku...


## Features:

* Detects if the mail is gmail or edu
* Gives you search links and members link based on the email
* Easy keyboard interface, very less commands
* Reply to forwarded messages to reply to original sender
* Ban the spammers using `/ban` command on the forwarded messages
* Keeps a log of the members approved
* Checks if the person is a member of the main group or not

## Commands

All these commands work in the **Control Group** (specified in config.py)

* `ban/unban:` Pass in user id (not username) or reply to a forwarded message
* `findmember:` Pass in email id or reply to a message containing email id to get google grp membership search results
* `findreq:` Pass in email id or reply to a message containing email id to get google grp request search results
* `pin:` pin messages
* `purge:` purge messages

## Warning

* Any member of the control group can use the bot so make it private and add only trusted members.

TODO:

* [ ] Add help command
