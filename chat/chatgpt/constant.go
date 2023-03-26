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
	GPT3Dot5Turbo      = "gpt-3.5-turbo"
	GPT3Dot5Turbo0301  = "gpt-3.5-turbo-0301"
	GPT3TextDavinci003 = "text-davinci-003"
	GPT3TextDavinci002 = "text-davinci-002"
)

// Chat message role defined by the OpenAI API.
const (
	ChatMessageRoleSystem    = "system"
	ChatMessageRoleUser      = "user"
	ChatMessageRoleAssistant = "assistant"
)

func (c *chatGPT) GetModels() []string {
	return []string{GPT3Dot5Turbo, GPT3Dot5Turbo0301, GPT3TextDavinci003, GPT3TextDavinci002}
}
