package requestProject

import (
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func SendTelegramNewDepot() error {
	var chatId int = 5465604654
	var text string = "🤑 Пришел новый деп!"
	var telegramApi string = "https://api.telegram.org/bot" + os.Getenv("TELEGRAM_BOT_SEND") + "/sendMessage"
	response, err := http.PostForm(
		telegramApi,
		url.Values{
			"chat_id": {strconv.Itoa(chatId)},
			"text":    {text},
		})
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return nil
}
