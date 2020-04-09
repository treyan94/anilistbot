# anilistbot
An unofficial [inline](https://core.telegram.org/bots/inline) Telegram bot for [anilist.co](https://anilist.co/) that searches for anime and anime characters. I wrote this when I was learning Go and GraphQL.

# Installation
The bot requires a Telegram api key, which you can get from the BotFather. All the required info can be found at https://core.telegram.org/bots. Once you've got the key, you either create an env variable name `ANI_BOT_KEY`, or you can pass it as the first argument. The second option takes precedence.

# Usage
By default, the bot is searching for anime. This is what a query might look like: `@anilist_unofficial_bot evangelion`. You can add ` /a` to your query to make search for adult titles.

To search for a character, you need to add ` /char` to your query (`@anilist_unofficial_bot ikari shinji /char`).
