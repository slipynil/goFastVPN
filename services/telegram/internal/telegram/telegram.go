package telegram

import (
	"fmt"
	"telegram-service/internal/dto"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Telegram struct {
	bot     *tgbotapi.BotAPI
	updates tgbotapi.UpdatesChannel
}

func New(token string) (*Telegram, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	telegram := &Telegram{bot: bot}

	u := tgbotapi.NewUpdate(0)
	telegram.updates = bot.GetUpdatesChan(u)

	// меню слева внизу с командами
	commands := []tgbotapi.BotCommand{
		{Command: "menu", Description: "Меню"},
	}
	_, err = telegram.bot.Request(tgbotapi.NewSetMyCommands(commands...))
	if err != nil {
		return nil, err
	}
	return telegram, nil
}

func (t *Telegram) Chan() tgbotapi.UpdatesChannel {
	return t.updates
}

func keyboardMainMenu() tgbotapi.InlineKeyboardMarkup {
	options := []string{"получить конфиг", "помощь", "стоимость", "оплатить"}

	var rows [][]tgbotapi.InlineKeyboardButton
	for _, opt := range options {
		btn := tgbotapi.NewInlineKeyboardButtonData(opt, dto.EncodeCallbackData(opt))
		row := tgbotapi.NewInlineKeyboardRow(btn)
		rows = append(rows, row)
	}
	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}

func keyboardBackMenu() tgbotapi.InlineKeyboardMarkup {
	opt := "назад"
	btn := tgbotapi.NewInlineKeyboardButtonData(opt, dto.EncodeCallbackData(opt))
	row := tgbotapi.NewInlineKeyboardRow(btn)
	return tgbotapi.NewInlineKeyboardMarkup(row)
}

// создает новое сообщение и отправляет меню
func (t *Telegram) Menu(chatID int64) error {

	msg := tgbotapi.NewMessage(chatID, "главное меню")
	msg.ReplyMarkup = keyboardMainMenu()

	_, err := t.bot.Send(msg)
	return err
}

// меняет текущее сообщение и отправляет меню
func (t *Telegram) UpdateMainMenu(update tgbotapi.Update) error {

	msg := tgbotapi.NewEditMessageTextAndMarkup(
		update.CallbackQuery.Message.Chat.ID,
		update.CallbackQuery.Message.MessageID,
		"главное меню",
		keyboardMainMenu(),
	)

	_, err := t.bot.Send(msg)
	return err
}

func (t *Telegram) UpdateSendText(update tgbotapi.Update, text string) error {
	msg := tgbotapi.NewEditMessageTextAndMarkup(
		update.CallbackQuery.Message.Chat.ID,
		update.CallbackQuery.Message.MessageID,
		text,
		keyboardBackMenu(),
	)

	_, err := t.bot.Send(msg)

	return err
}

func (t *Telegram) SendFile(chat *tgbotapi.Chat, bufer []byte) error {
	// create document struct
	file := tgbotapi.FileBytes{
		Name:  fmt.Sprintf("%s.conf", chat.UserName),
		Bytes: bufer,
	}
	msg := tgbotapi.NewDocument(chat.ID, file)
	_, err := t.bot.Send(msg)
	return err
}
