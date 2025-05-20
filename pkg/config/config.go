package config

type Config struct {
	ToolsConfig ToolsConfig    `yaml:"tools"`
	Reporter    ReporterConfig `yaml:"reporter"`
}

type ToolsConfig struct {
	HTTPX     HTTPXConfig     `yaml:"httpx"`
	Subfinder SubfinderConfig `yaml:"subfinder"`
	Katana    KatanaConfig    `yaml:"katana"`
}

type HTTPXConfig struct {
	Threads int `yaml:"threads"`
}

type SubfinderConfig struct {
	Threads int `yaml:"threads"`
}

type KatanaConfig struct {
	Threads           int `yaml:"threads"`
	ParllelismThreads int `yaml:"parallelism_threads"`
}

type ReporterConfig struct {
	Telegram TelegramConfig `yaml:"telegram"`
}

type TelegramConfig struct {
	Enabled  bool   `yaml:"enabled"`
	BotToken string `yaml:"bot_token"`
	ChatID   string `yaml:"chat_id"`
}
