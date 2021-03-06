package main

import (
	"io"

	"github.com/apprenda/kuberang/pkg/kuberang"
	"github.com/spf13/cobra"
)

// NewKismaticCommand creates the kismatic command
func NewKuberangCommand(version string, in io.Reader, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kuberang",
		Short: "kuberang tests your kubernetes cluster using kubectl",
		RunE: func(cmd *cobra.Command, args []string) error {
			return doCheckKubernetes(args)
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	return cmd
}

func doCheckKubernetes(args []string) error {
	return kuberang.CheckKubernetes()
}
