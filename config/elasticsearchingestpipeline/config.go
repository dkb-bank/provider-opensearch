package elasticsearchingestpipeline

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("elasticsearch_ingest_pipeline", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "IngestPipeline"
	})
}
