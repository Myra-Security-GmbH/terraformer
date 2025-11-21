package myrasec

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	mgo "github.com/Myra-Security-GmbH/myrasec-go/v2"
)

// SettingGenerator
type SettingsGenerator struct {
	MyrasecService
}

// createSettingResources
func (g *SettingsGenerator) createSettingResources(api *mgo.API, domainId int, vhost mgo.VHost, wg *sync.WaitGroup) error {
	defer wg.Done()

	params := map[string]string{}

	s, err := api.ListSettingsFull(domainId, vhost.Label, params)
	if err != nil {
		return err
	}

	r := terraformutils.NewResource(
		strconv.Itoa(vhost.ID),
		fmt.Sprintf("%s_%d", vhost.Label, vhost.ID),
		"myrasec_settings",
		"myrasec",
		map[string]string{
			"subdomain_name": vhost.Label,
		},
		[]string{},
		map[string]any{},
	)

	appendIgnoreKeys(s, r)

	r.IgnoreKeys = append(r.IgnoreKeys, "cdn")
	g.Resources = append(g.Resources, r)
	return nil
}

// in terraform we only want the attributes that are configured on the current level, all other attributes should be ignored
func appendIgnoreKeys(response any, r terraformutils.Resource) {
	data := response.(*map[string]any)
	domain := (*data)["domain"]
	parent := (*data)["parent"]

	domainSettings, _ := domain.(map[string]any)
	parentSettings := parent.(map[string]any)
	for i := range parentSettings {
		if _, ok := domainSettings[i]; !ok {
			r.IgnoreKeys = append(r.IgnoreKeys, i)
		}
	}
}

// InitResources
func (g *SettingsGenerator) InitResources() error {
	wg := sync.WaitGroup{}

	api, err := g.initializeAPI()
	if err != nil {
		return nil
	}

	funcs := []func(*mgo.API, int, mgo.VHost, *sync.WaitGroup) error{
		g.createSettingResources,
	}

	err = createResourcesPerSubDomain(api, funcs, &wg, true)
	if err != nil {
		return nil
	}

	wg.Wait()

	return nil
}
