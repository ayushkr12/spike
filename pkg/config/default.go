package config

func DefaultConfig() *Config {
	return &Config{
		ToolsConfig: ToolsConfig{
			HTTPX: HTTPXConfig{
				Threads: 50,
			},
			Subfinder: SubfinderConfig{
				Threads: 10,
			},
			Katana: KatanaConfig{
				Threads:           10,
				ParllelismThreads: 10,
			},
		},
		Reporter: ReporterConfig{
			Telegram: TelegramConfig{
				Enabled:  false,
				BotToken: "",
				ChatID:   "",
			},
		},
	}
}
