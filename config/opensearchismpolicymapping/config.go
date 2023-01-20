package opensearchismpolicymapping

import "github.com/upbound/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("elasticsearch_opensearch_ism_policy_mapping", func(r *config.Resource) {
		r.ShortGroup = "opensearch"
	})
}
