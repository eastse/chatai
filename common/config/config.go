package config

type SettingConfig struct {
	MultilineInput    bool   `yaml:"multilineInput"`
	UsingChatID       int64  `yaml:"usingChatId"`
	QuickChatID       int64  `yaml:"quickChatId"`
	InputPrompt       string `yaml:"inputPrompt"`
	SinglePromptColor string `yaml:"singlePromptColor"`
	MultiPromptColor  string `yaml:"multiPromptColor"`
	SingleTextColor   string `yaml:"singleTextColor"`
	MultiTextColor    string `yaml:"multiTextColor"`
}

type ChatConfig struct {
	Title       string  `yaml:"title"`       // chat title
	ID          int64   `yaml:"id"`          // ID
	Type        int64   `yaml:"type"`        // chat type
	Prompt      string  `yaml:"prompt"`      // The prompt(s) to generate completions
	Model       string  `yaml:"model"`       // ID of the model to use. See the model endpoint compatibility table for details on which models work with the Chat API.
	Temperature float32 `yaml:"temperature"` // Defaults 1 What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic.We generally recommend altering this or top_p but not both.
	Review      int64   `yaml:"review"`      // Reviewing the number of information in context.
	Sort        int64   `yaml:"sort"`        // The heavier the sorting weight, the more priority it receives.
	Developer   string  `yaml:"developer"`   // api developer
	HistoryFile string  `yaml:"historyFile"` // chat history
}
