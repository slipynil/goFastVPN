package service

import (
	"telegram-service/internal/dto"
	"telegram-service/logger"
)

func (s *service) Update(logger *logger.MyLogger) {

	for u := range s.telegram.Chan() {
		// если есть сигнал об оплате
		if u.PreCheckoutQuery != nil {
			err := s.telegram.PreCheckoutQuery(u)
			logger.IsErr("failed to answer pre_checkout_query", err)
			// если структура message не пустая и является командой
		} else if u.Message != nil {
			if u.Message.SuccessfulPayment != nil {
				if dto, err := s.telegram.HandleSuccessfulPayment(u); err != nil {
					logger.IsErr("payment canceled", err)
				} else {
					err := s.postgres.SuccessfulPaymentStatus(dto.InvoicePayload)
					logger.IsErr("failed to change true status on payment", err)
					err = s.postgres.StatusTrue(u.Message.Chat.ID)
					logger.IsErr("failed to change true status on client", err)
					err = s.add(u.Message.Chat, dto.TotalAmount)
					logger.IsErr("failed to add peer", err)
				}
			}
			// если команда
			switch u.Message.Command() {
			case "menu":
				s.telegram.Menu(u.Message.Chat.ID)
			case "start":
				err := s.postgres.AddClient(u.Message.Chat.UserName, u.Message.Chat.ID)
				logger.IsErr("failed to add client to postgres", err)
			}
			continue
			// если это inline кнопка
		} else if u.CallbackQuery != nil {

			// add peer command send conf file for user
			callBackData, err := dto.DecodeCallbackData(u.CallbackQuery.Data)
			logger.IsErr("failed decoding callback data", err)

			switch callBackData.Action {

			case "<- назад":
				err := s.telegram.UpdateMainMenu(u)
				logger.IsErr("failed to redirect main menu", err)

			case "получить конфиг":
				err := s.getConfFile(u)
				logger.IsErr("", err)

			case "помощь":
				err := s.telegram.UpdateSendText(u, HelpText)
				logger.IsErr("", err)

			case "стоимость":
				err := s.telegram.UpdateSendText(u, PricingText)
				logger.IsErr("", err)

			case "оплатить":
				if s.postgres.CheckStatus(u.CallbackQuery.Message.Chat.ID) {
					err := s.telegram.UpdateSendText(u, "Вы уже оплатили")
					logger.IsErr("", err)
				} else {
					err := s.Invoice(u)
					logger.IsErr("failed to create invoice", err)
				}

			case "протестировать":
				if s.postgres.IsTested(u.CallbackQuery.Message.Chat.ID) {
					err := s.telegram.UpdateSendText(u, "У вас уже был тестовый доступ")
					logger.IsErr("", err)
				} else {
					err := s.add(u.CallbackQuery.Message.Chat, 0)
					logger.IsErr("", err)
				}
			}
		}
	}
}
