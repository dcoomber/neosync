package neosync_cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	accounts_cmd "github.com/nucleuscloud/neosync/cli/internal/cmds/neosync/accounts"
	connections_cmd "github.com/nucleuscloud/neosync/cli/internal/cmds/neosync/connections"
	jobs_cmd "github.com/nucleuscloud/neosync/cli/internal/cmds/neosync/jobs"
	login_cmd "github.com/nucleuscloud/neosync/cli/internal/cmds/neosync/login"
	sync_cmd "github.com/nucleuscloud/neosync/cli/internal/cmds/neosync/sync"
	version_cmd "github.com/nucleuscloud/neosync/cli/internal/cmds/neosync/version"
	whoami_cmd "github.com/nucleuscloud/neosync/cli/internal/cmds/neosync/whoami"
	"github.com/nucleuscloud/neosync/cli/internal/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc/metadata"
)

const (
	neosyncDirName           = ".neosync"
	cliSettingsFileNameNoExt = ".neosync-cli"
	cliSettingsFileExt       = "yaml"

	apiKeyEnvVarName = "NEOSYNC_API_KEY" //nolint:gosec
	apiKeyFlag       = "api-key"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "neosync",
		Short: "Terminal UI that interfaces with the Neosync system.",
		Long:  "",
		PersistentPreRun: func(cmd *cobra.Command, _ []string) {
			cmd.SilenceErrors = true
			cmd.SetContext(metadata.NewOutgoingContext(cmd.Context(), version.Get().GrpcMetadata()))
		},
	}

	var cfgFile string
	cobra.OnInitialize(
		func() { initConfig(cfgFile) },
		func() {
			apiKey, err := rootCmd.Flags().GetString(apiKeyFlag)
			if err != nil {
				panic(err)
			}
			envApiKey := viper.GetString(apiKeyEnvVarName)
			if apiKey == "" && envApiKey != "" {
				err = rootCmd.Flags().Set(apiKeyFlag, envApiKey)
				if err != nil {
					panic(err)
				}
			}
		},
	)

	rootCmd.Version = version.Get().GitVersion
	rootCmd.SetVersionTemplate(`{{printf "%s\n" .Version}}`)

	rootCmd.PersistentFlags().StringVar(
		&cfgFile, "config", "", fmt.Sprintf("config file (default is $HOME/%s.%s)", cliSettingsFileNameNoExt, cliSettingsFileExt),
	)

	rootCmd.PersistentFlags().String(apiKeyFlag, "", fmt.Sprintf("Neosync API Key. Takes precedence over $%s", apiKeyEnvVarName))
	rootCmd.AddCommand(jobs_cmd.NewCmd())
	rootCmd.AddCommand(version_cmd.NewCmd())
	rootCmd.AddCommand(whoami_cmd.NewCmd())
	rootCmd.AddCommand(login_cmd.NewCmd())
	rootCmd.AddCommand(sync_cmd.NewCmd())
	rootCmd.AddCommand(accounts_cmd.NewCmd())
	rootCmd.AddCommand(connections_cmd.NewCmd())

	cobra.CheckErr(rootCmd.Execute())
}

// initConfig reads in config file and ENV variables if set.
func initConfig(cfgFilePath string) {
	if cfgFilePath != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFilePath)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		fullNeosyncSettingsDir := filepath.Join(home, neosyncDirName)
		neosyncConfigDir := os.Getenv("NEOSYNC_CONFIG_DIR") // helpful for tools such as direnv and people who want it somewhere interesting
		xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")       // linux users expect this to be respected

		viper.AddConfigPath(".")
		viper.AddConfigPath(fullNeosyncSettingsDir)
		viper.AddConfigPath(home)
		viper.AddConfigPath(neosyncConfigDir)
		viper.AddConfigPath(xdgConfigHome)

		viper.SetConfigType(cliSettingsFileExt)
		viper.SetConfigName(cliSettingsFileNameNoExt)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		if !errors.As(err, &viper.ConfigFileNotFoundError{}) {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
			return
		}
	}
}
