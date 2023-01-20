package opensearchrole

import "github.com/upbound/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("elasticsearch_opensearch_role", func(r *config.Resource) {
		r.ShortGroup = "opensearch"
	})
}
