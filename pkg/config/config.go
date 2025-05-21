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
	Threads           int    `yaml:"threads"`
	CrawlDepth        int    `yaml:"crawl_depth"`
	MaxCrawlTime      string `yaml:"max_crawl_time"`
	ParllelismThreads int    `yaml:"parallelism_threads"`
	Headless          bool   `yaml:"headless"`
	NoSandbox         bool   `yaml:"no_sandbox"`
}

type ReporterConfig struct {
	Telegram TelegramConfig `yaml:"telegram"`
}

type TelegramConfig struct {
	Enabled  bool   `yaml:"enabled"`
	BotToken string `yaml:"bot_token"`
	ChatID   string `yaml:"chat_id"`
}
