package elasticsearchcomponenttemplate

import (
	"github.com/upbound/upjet/pkg/config"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("elasticsearch_component_template", func(r *config.Resource) {
		r.ShortGroup = "component"
	})
}
