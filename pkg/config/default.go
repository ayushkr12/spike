package config

import (
	"os"
	"path/filepath"

	log "github.com/ayushkr12/logger"
)

var HomeDir = os.Getenv("HOME")
var AppName = "spike"

func DefaultConfig() *Config {
	log.Debugf("Using $HOME directory as: %s", HomeDir)
	return &Config{
		AppConfig: AppConfig{
			AppDir:            filepath.Join(HomeDir, AppName),
			DefaultDBPath:     filepath.Join(HomeDir, AppName, "spike.db"),
			DefaultConfigPath: filepath.Join(HomeDir, AppName, "config.yaml"),
		},
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
			Nuclei: NucleiConfig{
				Enabled: true,
				Threads: 25,
				TemplatePaths: NucleiTemplatePaths{
					Default: filepath.Join(HomeDir, "nuclei-templates"),
					Custom:  filepath.Join(HomeDir, "custom-nuclei-templates"),
				},
				TemplateSettings: NucleiTemplateSettings{
					DefaultEnabled:  true,
					CustomEnabled:   true,
					DastEnabled:     true,
					HeadlessEnabled: false,
				},
				CustomScanOptions: NucleiCustomScanOptions{
					DomXSS: true,
				},
				CustomScanTemplates: NucleiCustomScanTemplates{
					DomXSS: filepath.Join(HomeDir, "nuclei-templates/dast/vulnerabilities/xss/dom-xss.yaml"),
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
