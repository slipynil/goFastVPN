package service

import (
	"errors"
	"telegram-service/internal/dto"
	"telegram-service/logger"
)

func (s *service) Update(logger *logger.MyLogger) {

	for u := range s.telegram.Chan() {

		// если структура message не пустая и является командой
		if u.Message != nil && u.Message.Command() == "menu" {
			s.telegram.Menu(u.Message.Chat.ID)
			continue
			// если это inline кнопка
		} else if u.CallbackQuery != nil {

			// add peer command send conf file for user
			callBackData, err := dto.DecodeCallbackData(u.CallbackQuery.Data)
			logger.IsErr("fail decoding callback data", err)
			switch callBackData.Action {
			case "<- назад":
				err := s.telegram.UpdateMainMenu(u)
				logger.IsErr("fail to redirect main menu", err)
			case "получить конфиг":
				err := s.add(u.CallbackQuery.Message.Chat)
				if errors.Is(err, dto.ErrUserExist) {
					err := s.telegram.UpdateSendText(u, "у вас уже есть конфиг")
					logger.IsErr("", err)
				} else {
					logger.IsErr("fail adding peer", err)
				}
			case "помощь":
				err := s.telegram.UpdateSendText(u, `
🛠 Техническая поддержка GopherSecure
Возникли вопросы по настройке? Мы поможем!
📖 Наш канал: @GopherSecure — как настроить соединение на iPhone, Android, Windows и Mac.
👤 Живая поддержка: Напишите нашему инженеру @w3berr, если возникли проблемы с подключением.`)
				logger.IsErr("", err)
			case "стоимость":
				err := s.telegram.UpdateSendText(u, `
💳 Тарифы GopherSecure

🔘 Тестовый (24ч): 0 ₽ — Попробуйте качество связи.
🔘 Стандарт (30 дней): 200 ₽ — Оптимально для работы и серфинга.

Все тарифы включают безлимитный трафик на скорости до 1 Гбит/с.`,
				)
				logger.IsErr("", err)
			case "оплатить":
				err := s.telegram.UpdateSendText(u, "ссылка на оплату")
				logger.IsErr("", err)
			}
		}
	}
}
