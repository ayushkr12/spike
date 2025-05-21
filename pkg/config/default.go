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
			Crawler: CrawlerConfig{
				Gau: GauConfig{
					Enabled: true,
					Threads: 10,
				},
				Katana: KatanaConfig{
					Enabled:            true,
					Threads:            10,
					CrawlDepth:         3,
					MaxCrawlTime:       "10m",
					ParallelismThreads: 10,
					Headless:           false,
					NoSandbox:          false,
				},
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
