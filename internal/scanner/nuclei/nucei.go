package nuclei

import (
	"path/filepath"
	"strconv"

	"github.com/ayushkr12/spike/internal/pkg/utils"
	"github.com/ayushkr12/spike/pkg/config"
)

func GenericScan(cfg *config.NucleiConfig, urls []string) ([]string, error) {
	cfgCopy := *cfg
	cfgCopy.TemplateSettings.DastEnabled = false
	return scan(urls, &cfgCopy)
}

func DastScan(cfg *config.NucleiConfig, urls []string) ([]string, error) {
	cfgCopy := *cfg
	cfgCopy.TemplateSettings.DastEnabled = true
	return scan(urls, &cfgCopy)
}

func scan(urls []string, cfg *config.NucleiConfig) ([]string, error) {
	args, err := buildNucleiArgs(cfg)
	if err != nil {
		return nil, err
	}
	return utils.RunCommandWithStdinInput("nuclei", args, urls)
}

func buildNucleiArgs(cfg *config.NucleiConfig) ([]string, error) {
	if err := ValidateNucleiConfig(cfg); err != nil {
		return nil, err
	}
	var args []string
	if cfg.TemplateSettings.DastEnabled {
		args = append(args, "-t", "dast")
		if cfg.TemplateSettings.DefaultEnabled {
			args = append(args, "-t", filepath.Join(cfg.TemplatePaths.Default, "dast"))
		}
		if cfg.TemplateSettings.CustomEnabled {
			args = append(args, "-t", filepath.Join(cfg.TemplatePaths.Custom, "dast"))
		}
	} else {
		if cfg.TemplateSettings.DefaultEnabled {
			args = append(args, "-t", cfg.TemplatePaths.Default)
		}
		if cfg.TemplateSettings.CustomEnabled {
			args = append(args, "-t", cfg.TemplatePaths.Custom)
		}
	}
	if cfg.TemplateSettings.HeadlessEnabled {
		args = append(args, "-headless")
	}
	args = append(args, "-c", strconv.Itoa(cfg.Threads))
	args = append(args, "-silent")

	return args, nil
}
