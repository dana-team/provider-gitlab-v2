// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	file "github.com/dana-team/provider-gitlab-v2/internal/controller/namespaced/gitlab-v2/file"
	group "github.com/dana-team/provider-gitlab-v2/internal/controller/namespaced/gitlab-v2/group"
	groupmembership "github.com/dana-team/provider-gitlab-v2/internal/controller/namespaced/gitlab-v2/groupmembership"
	project "github.com/dana-team/provider-gitlab-v2/internal/controller/namespaced/gitlab-v2/project"
	projectmembership "github.com/dana-team/provider-gitlab-v2/internal/controller/namespaced/gitlab-v2/projectmembership"
	projectsharegroup "github.com/dana-team/provider-gitlab-v2/internal/controller/namespaced/gitlab-v2/projectsharegroup"
	user "github.com/dana-team/provider-gitlab-v2/internal/controller/namespaced/gitlab-v2/user"
	providerconfig "github.com/dana-team/provider-gitlab-v2/internal/controller/namespaced/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		file.Setup,
		group.Setup,
		groupmembership.Setup,
		project.Setup,
		projectmembership.Setup,
		projectsharegroup.Setup,
		user.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		file.SetupGated,
		group.SetupGated,
		groupmembership.SetupGated,
		project.SetupGated,
		projectmembership.SetupGated,
		projectsharegroup.SetupGated,
		user.SetupGated,
		providerconfig.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
