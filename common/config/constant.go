package config

const (
	DeveloperOpenAI = "OpenAI"
)

const (
	GPT3Dot5Turbo      = "gpt-3.5-turbo"
	GPT3Dot5Turbo0301  = "gpt-3.5-turbo-0301"
	GPT3TextDavinci003 = "text-davinci-003"
	GPT3TextDavinci002 = "text-davinci-002"

	GPT3TextCurie001        = "text-curie-001"
	GPT3TextBabbage001      = "text-babbage-001"
	GPT3TextAda001          = "text-ada-001"
	GPT3TextDavinci001      = "text-davinci-001"
	GPT3DavinciInstructBeta = "davinci-instruct-beta"
	GPT3Davinci             = "davinci"
	GPT3CurieInstructBeta   = "curie-instruct-beta"
	GPT3Curie               = "curie"
	GPT3Ada                 = "ada"
	GPT3Babbage             = "babbage"
)

const (
	CodexCodeDavinci002 = "code-davinci-002"
	CodexCodeCushman001 = "code-cushman-001"
	CodexCodeDavinci001 = "code-davinci-001"
)

const (
	ColorRed    string = "\033[31m"
	ColorGreen  string = "\033[32m"
	ColorYellow string = "\033[33m"
	ColorBlue   string = "\033[34m"
	ColorPurple string = "\033[35m"
	ColorCyan   string = "\033[36m"
)

var Colors = map[string]string{
	"Red":    ColorRed,
	"Green":  ColorGreen,
	"Yellow": ColorYellow,
	"Blue":   ColorBlue,
	"Purple": ColorPurple,
	"Cyan":   ColorCyan,
}
