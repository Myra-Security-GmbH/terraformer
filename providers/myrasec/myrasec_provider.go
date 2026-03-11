package myrasec

import (
	"errors"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

// Provider
type Provider struct {
	terraformutils.Provider
}

// Init
func (p *Provider) Init(_ []string) error {
	return nil
}

// GetName
func (p *Provider) GetName() string {
	return "myrasec"
}

// GetProviderData
func (p *Provider) GetProviderData(_ ...string) map[string]any {
	return map[string]any{}
}

// GetResourceConnections
func (Provider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

// GetSupportedService
func (p *Provider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"cache_setting":     &CacheSettingGenerator{},
		"dns_record":        &DNSGenerator{},
		"domain":            &DomainGenerator{},
		"error_page":        &ErrorPageGenerator{},
		"ip_filter":         &IPFilterGenerator{},
		"maintenance":       &MaintenanceGenerator{},
		"redirect":          &RedirectGenerator{},
		"settings":          &SettingsGenerator{},
		"tag":               &TagGenerator{},
		"tag_cache_setting": &TagCacheSettingGenerator{},
		"tag_setting":       &TagSettingGenerator{},
		"tag_waf_rule":      &TagWafRuleGenerator{},
		"waf_rule":          &WafRuleGenerator{},
	}
}

// InitService
func (p *Provider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("myrasec: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())

	return nil
}
