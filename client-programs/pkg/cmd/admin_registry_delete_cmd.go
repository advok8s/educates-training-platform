// Copyright 2022-2023 The Educates Authors.

package cmd

import (
	"github.com/spf13/cobra"

	"github.com/vmware-tanzu-labs/educates-training-platform/client-programs/pkg/registry"
)

func (p *ProjectInfo) NewAdminRegistryDeleteCmd() *cobra.Command {
	var c = &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "delete",
		Short: "Deletes the local image registry",
		RunE:  func(_ *cobra.Command, _ []string) error { return registry.DeleteRegistry() },
	}

	return c
}
