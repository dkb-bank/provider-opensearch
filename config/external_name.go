/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"elasticsearch_opensearch_role":               config.IdentifierFromProvider,
	"elasticsearch_opensearch_roles_mapping":      config.IdentifierFromProvider,
	"elasticsearch_index_template":                config.IdentifierFromProvider,
	"elasticsearch_data_stream":                   config.IdentifierFromProvider,
	"elasticsearch_opensearch_ism_policy":         config.IdentifierFromProvider,
	"elasticsearch_opensearch_ism_policy_mapping": config.IdentifierFromProvider,
	"elasticsearch_component_template":            config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
