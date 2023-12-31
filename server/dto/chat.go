package dto

import (
	"github.com/sashabaranov/go-openai"
)

type ChatMessageDto struct {
	Content  string                         `json:"content" binding:"required"`
	Config   string                         `json:"config" binding:"required"`
	History  []openai.ChatCompletionMessage `json:"history" binding:"required"`
	DialogId uint64                         `json:"dialog_id" binding:"required"`
}

type RequestReChatMessageDto struct {
	ChatMessageDto
	MessageId uint64 `json:"message_id" binding:"required"`
}

type ResponseChatMessageDto struct {
	Content       string `json:"content" binding:"required"`
	UserMessageId uint64 `json:"user_message_id"`
}

type ChatConfig struct {
	Model             string  `json:"model"`
	Temperature       float32 `json:"temperature"`
	MaxTokens         int     `json:"max_tokens"`
	PresencePenalty   float32 `json:"presence_penalty"`
	HistoryCount      int     `json:"history_count"`
	CompressThreshold int     `json:"compress_threshold"`
	Memory            bool    `json:"memory"`
}
