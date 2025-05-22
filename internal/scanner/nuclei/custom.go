package nuclei

import (
	"github.com/ayushkr12/spike/internal/pkg/utils"
	"github.com/ayushkr12/spike/pkg/config"
)

func ScanDomXSS(cfg *config.NucleiConfig, urls []string) ([]string, error) {
	if err := ValidateNucleiConfig(cfg); err != nil {
		return nil, err
	}
	var nucleiArgs []string
	if err := validateTemplatePath(cfg.CustomScanTemplates.DomXSS); err != nil {
		return nil, err
	}
	nucleiArgs = append(nucleiArgs, "-t", cfg.CustomScanTemplates.DomXSS)
	nucleiArgs = append(nucleiArgs, "-dast", "-headless", "-silent")
	return utils.RunCommandWithStdinInput("nuclei", nucleiArgs, urls)
}
