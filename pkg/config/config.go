package config

type Config struct {
	ToolsConfig ToolsConfig    `yaml:"tools"`
	Reporter    ReporterConfig `yaml:"reporter"`
}

type ToolsConfig struct {
	HTTPX     HTTPXConfig     `yaml:"httpx"`
	Subfinder SubfinderConfig `yaml:"subfinder"`
	Crawler   CrawlerConfig   `yaml:"crawler"`
	Nuclei    NucleiConfig    `yaml:"nuclei"`
}

type HTTPXConfig struct {
	Threads int `yaml:"threads"`
}

type SubfinderConfig struct {
	Threads int `yaml:"threads"`
}

type CrawlerConfig struct {
	Katana KatanaConfig `yaml:"katana"`
	Gau    GauConfig    `yaml:"gau"`
}

type NucleiConfig struct {
	Enabled           bool                    `yaml:"enabled"`
	Threads           int                     `yaml:"threads"`
	TemplatePaths     NucleiTemplatePaths     `yaml:"template_paths"`
	TemplateSettings  NucleiTemplateSettings  `yaml:"template_settings"`
	CustomScanOptions NucleiCustomScanOptions `yaml:"custom_scan_options"`
}

type NucleiTemplatePaths struct {
	Default string `yaml:"default"`
	Custom  string `yaml:"custom"`
}

type NucleiTemplateSettings struct {
	DefaultEnabled  bool `yaml:"default_enabled"`
	CustomEnabled   bool `yaml:"custom_enabled"`
	DastEnabled     bool `yaml:"dast_enabled"`
	HeadlessEnabled bool `yaml:"headless_enabled"`
}

type NucleiCustomScanOptions struct {
	DomXSS bool `yaml:"dom_xss"`
}

type GauConfig struct {
	Enabled bool `yaml:"enabled"`
	Threads int  `yaml:"threads"`
}

type KatanaConfig struct {
	Enabled            bool   `yaml:"enabled"`
	Threads            int    `yaml:"threads"`
	CrawlDepth         int    `yaml:"crawl_depth"`
	MaxCrawlTime       string `yaml:"max_crawl_time"`
	ParallelismThreads int    `yaml:"parallelism_threads"`
	Headless           bool   `yaml:"headless"`
	NoSandbox          bool   `yaml:"no_sandbox"`
}

type ReporterConfig struct {
	Telegram TelegramConfig `yaml:"telegram"`
}

type TelegramConfig struct {
	Enabled  bool   `yaml:"enabled"`
	BotToken string `yaml:"bot_token"`
	ChatID   string `yaml:"chat_id"`
}
