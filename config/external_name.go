package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"gitlab_group":              config.IdentifierFromProvider,
	"gitlab_project":            config.IdentifierFromProvider,
	"gitlab_repository_file":    config.IdentifierFromProvider,
	"gitlab_project_membership": config.IdentifierFromProvider,
	"gitlab_group_membership":   config.IdentifierFromProvider,
	"gitlab_user": {
		SetIdentifierArgumentFn: func(base map[string]any, name string) {
			base["username"] = name
		},
		GetExternalNameFn: config.IDAsExternalName,
		GetIDFn:           config.ExternalNameAsID,
		OmittedFields:     []string{"username"},
	},
	"gitlab_project_share_group": config.IdentifierFromProvider,
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
