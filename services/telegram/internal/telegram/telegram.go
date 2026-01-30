package telegram

import (
	"fmt"
	"telegram-service/internal/dto"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Telegram struct {
	bot           *tgbotapi.BotAPI
	updates       tgbotapi.UpdatesChannel
	providerToken string
}

func New(telegramToken, providerToken string) (*Telegram, error) {
	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		return nil, err
	}
	telegram := &Telegram{
		bot:           bot,
		providerToken: providerToken,
	}

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

// маппинг кнопок главного меню
func keyboardMainMenu() tgbotapi.InlineKeyboardMarkup {
	options := []string{"получить конфиг", "помощь", "протестировать", "стоимость", "оплатить"}

	var rows [][]tgbotapi.InlineKeyboardButton
	for _, opt := range options {
		btn := tgbotapi.NewInlineKeyboardButtonData(opt, dto.EncodeCallbackData(opt))
		row := tgbotapi.NewInlineKeyboardRow(btn)
		rows = append(rows, row)
	}
	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}

// маппинг кнопок для выхода из опции
func keyboardBackMenu() tgbotapi.InlineKeyboardMarkup {
	opt := "<- назад"
	btn := tgbotapi.NewInlineKeyboardButtonData(opt, dto.EncodeCallbackData(opt))
	row := tgbotapi.NewInlineKeyboardRow(btn)
	return tgbotapi.NewInlineKeyboardMarkup(row)
}

// создает новое сообщение и отправляет меню
func (t *Telegram) Menu(chatID int64) error {

	msg := tgbotapi.NewMessage(chatID, "📱 Главное меню")
	msg.ReplyMarkup = keyboardMainMenu()

	_, err := t.bot.Send(msg)
	return err
}

// меняет текущее сообщение и отправляет меню
func (t *Telegram) UpdateMainMenu(update tgbotapi.Update) error {

	msg := tgbotapi.NewEditMessageTextAndMarkup(
		update.CallbackQuery.Message.Chat.ID,
		update.CallbackQuery.Message.MessageID,
		"📱 главное меню",
		keyboardMainMenu(),
	)

	_, err := t.bot.Send(msg)
	return err
}

// меняет текущее сообщение и отправляет заданный текст с маппингом выхода из опции
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

func (t *Telegram) SendFile(chatID int64, bufer []byte) error {
	// create document struct
	unix := time.Now().Unix()
	file := tgbotapi.FileBytes{
		Name:  fmt.Sprintf("awg%d.conf", unix),
		Bytes: bufer,
	}
	msg := tgbotapi.NewDocument(chatID, file)
	_, err := t.bot.Send(msg)
	return err
}

func (t *Telegram) SendText(chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := t.bot.Send(msg)
	return err
}
