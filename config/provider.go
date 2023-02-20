/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	elasticsearchcomponenttemplate "github.com/dkb-bank/provider-opensearch/config/elasticsearchcomponenttemplate"
	elasticsearchdatastream "github.com/dkb-bank/provider-opensearch/config/elasticsearchdatastream"
	elasticsearchindextemplate "github.com/dkb-bank/provider-opensearch/config/elasticsearchindextemplate"
	opensearchismpolicy "github.com/dkb-bank/provider-opensearch/config/opensearchismpolicy"
	opensearchismpolicymapping "github.com/dkb-bank/provider-opensearch/config/opensearchismpolicymapping"
	opensearchrole "github.com/dkb-bank/provider-opensearch/config/opensearchrole"
	opensearchrolemapping "github.com/dkb-bank/provider-opensearch/config/opensearchrolemapping"
)

const (
	resourcePrefix = "opensearch"
	modulePath     = "github.com/dkb-bank/provider-opensearch"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithShortName("opensearch"),
		ujconfig.WithRootGroup("opensearch.crossplane.io"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		opensearchrole.Configure,
		elasticsearchindextemplate.Configure,
		elasticsearchdatastream.Configure,
		opensearchismpolicy.Configure,
		opensearchismpolicymapping.Configure,
		opensearchrolemapping.Configure,
		elasticsearchcomponenttemplate.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
