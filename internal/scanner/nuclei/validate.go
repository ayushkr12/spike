package nuclei

import (
	"errors"
	"fmt"

	"github.com/ayushkr12/spike/internal/pkg/utils"
	"github.com/ayushkr12/spike/pkg/config"
)

var ErrConflictHeadlessAndDomXSS = errors.New("internal: headless and DOM XSS conflict")

func ValidateNucleiConfig(cfg *config.NucleiConfig) error {
	if cfg.TemplateSettings.DefaultEnabled {
		if err := validateTemplatePath(cfg.TemplatePaths.Default); err != nil {
			return fmt.Errorf("invalid default template path: %v", err)
		}
	}
	if cfg.TemplateSettings.CustomEnabled {
		if err := validateTemplatePath(cfg.TemplatePaths.Custom); err != nil {
			return fmt.Errorf("invalid custom template path: %v", err)
		}
	}
	if cfg.TemplateSettings.HeadlessEnabled && cfg.CustomScanOptions.DomXSS {
		return ErrConflictHeadlessAndDomXSS
	}
	return nil
}

func validateTemplatePath(path string) error {
	if path == "" {
		return fmt.Errorf("path cannot be empty")
	}
	if !utils.IsDirectory(path) {
		return fmt.Errorf("%s is not a directory", path)
	}
	return nil
}
