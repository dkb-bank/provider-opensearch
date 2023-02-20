package opensearchrolemapping

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("elasticsearch_opensearch_roles_mapping", func(r *config.Resource) {
		r.ShortGroup = ""
		r.References = config.References{
			"role_name": {
				Type: "github.com/dkb-bank/provider-opensearch/apis/opensearch/v1alpha1.Role",
			},
		}

	})
}
