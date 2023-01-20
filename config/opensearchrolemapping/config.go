package opensearchrolemapping

import "github.com/upbound/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("elasticsearch_opensearch_roles_mapping", func(r *config.Resource) {
		r.ShortGroup = "opensearch"
		r.References = config.References{
			"role_name": {
				Type: "github.com/dkb-bank/provider-opensearch/apis/opensearch/v1alpha1.Role",
			},
		}

	})
}
