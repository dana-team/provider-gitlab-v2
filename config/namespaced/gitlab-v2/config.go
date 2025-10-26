package gitlab_v2

import "github.com/crossplane/upjet/v2/pkg/config"

const (
	apiVersion = "v1alpha1"
	shortGroup = "gitlab-v2"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("gitlab_group", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Group"
		r.Version = apiVersion
	})

	p.AddResourceConfigurator("gitlab_project", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Project"
		r.Version = apiVersion
	})

	p.AddResourceConfigurator("gitlab_project", func(r *config.Resource) {
		r.References["namespace_id"] = config.Reference{
			TerraformName: "gitlab_group",
			Extractor:     "github.com/crossplane/upjet/v2/pkg/resource.ExtractResourceID()",
		}
	})

	p.AddResourceConfigurator("gitlab_repository_file", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "File"
		r.Version = apiVersion
	})

	p.AddResourceConfigurator("gitlab_repository_file", func(r *config.Resource) {
		r.References["project"] = config.Reference{
			TerraformName: "gitlab_project",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("gitlab_project_membership", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ProjectMembership"
		r.Version = apiVersion
	})

	p.AddResourceConfigurator("gitlab_project_membership", func(r *config.Resource) {
		r.References["project"] = config.Reference{
			TerraformName: "gitlab_project",
			Extractor:     "github.com/crossplane/upjet/v2/pkg/resource.ExtractResourceID()",
		}

		r.References["user_id"] = config.Reference{
			TerraformName: "gitlab_user",
			Extractor:     "github.com/crossplane/upjet/v2/pkg/resource.ExtractResourceID()",
		}
	})

	p.AddResourceConfigurator("gitlab_group_membership", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "GroupMembership"
		r.Version = apiVersion
	})

	p.AddResourceConfigurator("gitlab_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "User"
		r.Version = apiVersion
		r.UseAsync = true
	})

	p.AddResourceConfigurator("gitlab_project_share_group", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ProjectShareGroup"
		r.Version = apiVersion
	})

	p.AddResourceConfigurator("gitlab_project_share_group", func(r *config.Resource) {
		r.References["project"] = config.Reference{
			TerraformName: "gitlab_project",
			Extractor:     "github.com/crossplane/upjet/v2/pkg/resource.ExtractResourceID()",
		}

		r.References["group_id"] = config.Reference{
			TerraformName: "gitlab_group",
			Extractor:     "github.com/crossplane/upjet/v2/pkg/resource.ExtractResourceID()",
		}
	})
}
