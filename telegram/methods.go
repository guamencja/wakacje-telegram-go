package telegram

type User struct {
	Id           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
	IsPremium    bool   `json:"is_premium"`
}

// getMe
func (b *Bot) GetMe() (User, error) {
	var u User
	if err := b.request("getMe", nil, &u); err != nil {
		return User{}, err
	}
	return u, nil
}

// editMessageText
func (b *Bot) EditMessageText(chatId, messageId, content string) error {
	data := map[string]interface{}{
		"chat_id":    chatId,
		"message_id": messageId,
		"text":       content,
	}
	if err := b.request("editMessageText", data, nil); err != nil {
		return err
	}
	return nil
}
