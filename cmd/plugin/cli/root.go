package cli

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/logan-bobo/abx/pkg/logger"
	"github.com/logan-bobo/abx/pkg/plugin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var (
	KubernetesConfigFlags *genericclioptions.ConfigFlags
)

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "abx",
		Short:         "Executes kubectl commands based on attributes",
		Long:          `Executes kubectl commands over multiple clusters based on custom attributes assigned to contexts`,
		SilenceErrors: true,
		SilenceUsage:  true,
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlags(cmd.Flags())
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			log := logger.NewLogger()

			if err := plugin.RunPlugin(KubernetesConfigFlags); err != nil {
				return errors.Unwrap(err)
			}

			log.Info("")

			return nil
		},
	}

	cobra.OnInitialize(initConfig)

	cmd.Flags().StringSlice("sa", []string{}, "a group of select atributes to filter over")
	cmd.Flags().StringSlice("na", []string{}, "a group of negate atributes to filter over")
	KubernetesConfigFlags = genericclioptions.NewConfigFlags(false)
	KubernetesConfigFlags.AddFlags(cmd.Flags())
	fmt.Println(cmd.Flags())

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	return cmd
}

func InitAndExecute() {
	if err := RootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	viper.AutomaticEnv()
}
