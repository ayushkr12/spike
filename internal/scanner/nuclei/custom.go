package nuclei

import (
	"path/filepath"

	"github.com/ayushkr12/spike/internal/pkg/utils"
	"github.com/ayushkr12/spike/pkg/config"
)

var DomXSSTemplatePath = "dast/vulnerabilities/xss/dom-xss.yaml"

func ScanDomXSS(cfg *config.NucleiConfig, urls []string) ([]string, error) {
	if err := ValidateNucleiConfig(cfg); err != nil {
		return nil, err
	}
	var nucleiArgs []string
	if cfg.TemplateSettings.DefaultEnabled {
		nucleiArgs = append(nucleiArgs, "-t", filepath.Join(cfg.TemplatePaths.Default, DomXSSTemplatePath))
	}
	if cfg.TemplateSettings.CustomEnabled {
		nucleiArgs = append(nucleiArgs, "-t", filepath.Join(cfg.TemplatePaths.Custom, DomXSSTemplatePath))
	}
	nucleiArgs = append(nucleiArgs, "-headless")
	nucleiArgs = append(nucleiArgs, "-silent")
	return utils.RunCommandWithStdinInput("nuclei", nucleiArgs, urls)
}
