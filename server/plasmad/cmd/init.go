package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	app "github.com/FourthState/plasma-mvp-sidechain"
	pConfig "github.com/FourthState/plasma-mvp-sidechain/server/plasmad/config"
	gaiaInit "github.com/cosmos/cosmos-sdk/cmd/gaia/init"

	"github.com/cosmos/cosmos-sdk/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	tmConfig "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/libs/cli"
	tmCommon "github.com/tendermint/tendermint/libs/common"
)

const (
	flagOverwrite = "overwrite"
	flagMoniker   = "moniker"
	flagChainID   = "chainId"
)

type chainInfo struct {
	Moniker    string          `json:"moniker"`
	ChainID    string          `json:"chain_id"`
	NodeID     string          `json:"node_id"`
	AppMessage json.RawMessage `json:"app_message"`
}

// get cmd to initialize all files for tendermint and application
// nolint
func InitCmd(ctx *server.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize private validator, p2p, genesis, and application configuration files",
		Long:  `Initialize validators's and node's configuration files.`,
		Args:  cobra.NoArgs,
		RunE: func(_ *cobra.Command, _ []string) error {
			config := ctx.Config
			config.SetRoot(viper.GetString(cli.HomeFlag))

			chainID := viper.GetString(client.FlagChainID)
			if chainID == "" {
				chainID = fmt.Sprintf("test-chain-%v", common.RandStr(6))
			}
			if viper.GetString(flagMoniker) != "" {
				config.Moniker = viper.GetString(flagMoniker)
			}

			_, _, err := gaiaInit.InitializeNodeValidatorFiles(config)
			if err != nil {
				return err
			}
			var appState json.RawMessage
			genFile := config.GenesisFile()
			if !viper.GetBool(flagOverwrite) && tmCommon.FileExists(genFile) {
				return fmt.Errorf("genesis.json file already exists: %v", genFile)
			}

			appState, err = json.MarshalIndent(app.GenesisState{}, "\t")
			if err != nil {
				return err
			}

			pubKey := gaiaInit.ReadOrCreatePrivValidator(config.PrivValidatorFile())

			if err = gaiaInit.ExportGenesisFile(genFile, chainID, nil, appState); err != nil {
				return err
			}

			// write tendermint and plasma config files to disk
			tmConfig.WriteConfigFile(filepath.Join(config.RootDir, "config", "config.toml"), config)
			plasmaConfig := pConfig.DefaultPlasmaConfig()
			pConfig.WritePlasmaConfigFile(filepath.Join(config.RootDir, "config", "plasma.toml"), plasmaConfig)

			// display chain info
			info, err := json.MarshalIndent(chainInfo{
				ChainID:    chainID,
				Moniker:    config.Moniker,
				NodeID:     nodeID,
				AppMessage: appState,
			}, "\t")
			if err != nil {
				return err
			}
			fmt.Printf("%s\n", string(info))

			return nil
		},
	}

	cmd.Flags().String(cli.HomeFlag, app.DefaultNodeHome, "node's home directory")
	cmd.Flags().BoolP(flagOverwrite, "o", false, "overwrite the genesis.json file")
	cmd.Flags().String(flagChainID, "", "genesis file chain-id, if left blank will be randomly created")
	cmd.Flags().String(flagMoniker, "m", "set the validator's moniker")
	return cmd
}