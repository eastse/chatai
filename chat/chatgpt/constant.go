package chatgpt

const (
	BASE_URL = "https://api.openai.com/v1"
	CHAT_URL = "/chat/completions"
)

// GPT3 Defines the models provided by OpenAI to use when generating
// completions from OpenAI.
// GPT3 Models are designed for text-based tasks. For code-specific
// tasks, please refer to the Codex series of models.
const (
	GPT3Dot5Turbo     = "gpt-3.5-turbo"
	GPT3Dot5Turbo0301 = "gpt-3.5-turbo-0301"
)

// Codex Defines the models provided by OpenAI.
// These models are designed for code-specific tasks, and use
// a different tokenizer which optimizes for whitespace.
const (
	CodexCodeDavinci002 = "code-davinci-002"
)

// Chat message role defined by the OpenAI API.
const (
	ChatMessageRoleSystem    = "system"
	ChatMessageRoleUser      = "user"
	ChatMessageRoleAssistant = "assistant"
)

func (c *chatGPT) GetModels() []string {
	return []string{GPT3Dot5Turbo, GPT3Dot5Turbo0301}
}
