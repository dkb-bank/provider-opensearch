package elasticsearchindextemplate

import "github.com/upbound/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("elasticsearch_index_template", func(r *config.Resource) {
		r.ShortGroup = "index"
	})
}
