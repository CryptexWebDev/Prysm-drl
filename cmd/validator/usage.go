// This code was adapted from https://github.com/ethereum/go-ethereum/blob/master/cmd/geth/usage.go
package main

import (
	"io"
	"sort"

	"github.com/Dorol-Chain/Prysm-drl/v5/cmd"
	"github.com/Dorol-Chain/Prysm-drl/v5/cmd/validator/flags"
	"github.com/Dorol-Chain/Prysm-drl/v5/config/features"
	"github.com/Dorol-Chain/Prysm-drl/v5/runtime/debug"
	"github.com/urfave/cli/v2"
)

var appHelpTemplate = `NAME:
   {{.App.Name}} - {{.App.Usage}}

USAGE:
   {{.App.HelpName}} [options]{{if .App.Commands}} command [command options]{{end}} {{if .App.ArgsUsage}}{{.App.ArgsUsage}}{{else}}[arguments...]{{end}}
{{if .App.Version}}
VERSION:
	{{.App.Version}}
{{end -}}
{{if len .App.Authors}}
AUTHORS:
   {{range .App.Authors}}{{ . }}
   {{end -}}
{{end -}}
{{if .App.Commands}}
global OPTIONS:
   {{range .App.Commands}}{{join .Names ", "}}{{ "\t" }}{{.Usage}}
   {{end -}}
{{end -}}
{{if .FlagGroups}}
{{range .FlagGroups}}{{.Name}} OPTIONS:
  {{range .Flags}}{{.}}
  {{end}}
{{end -}}
{{end -}}
{{if .App.Copyright }}
COPYRIGHT:
   {{.App.Copyright}}
{{end -}}
`

type flagGroup struct {
	Name  string
	Flags []cli.Flag
}

var appHelpFlagGroups = []flagGroup{
	{
		Name: "cmd",
		Flags: []cli.Flag{
			cmd.MinimalConfigFlag,
			cmd.E2EConfigFlag,
			cmd.VerbosityFlag,
			cmd.DataDirFlag,
			flags.WalletDirFlag,
			flags.WalletPasswordFileFlag,
			cmd.ClearDB,
			cmd.ForceClearDB,
			cmd.EnableBackupWebhookFlag,
			cmd.BackupWebhookOutputDir,
			cmd.EnableTracingFlag,
			cmd.TracingProcessNameFlag,
			cmd.TracingEndpointFlag,
			cmd.TraceSampleFractionFlag,
			cmd.MonitoringHostFlag,
			flags.MonitoringPortFlag,
			cmd.DisableMonitoringFlag,
			cmd.LogFormat,
			cmd.LogFileName,
			cmd.ConfigFileFlag,
			cmd.ChainConfigFileFlag,
			cmd.GrpcMaxCallRecvMsgSizeFlag,
			cmd.AcceptTosFlag,
			cmd.ApiTimeoutFlag,
		},
	},
	{
		Name: "debug",
		Flags: []cli.Flag{
			debug.PProfFlag,
			debug.PProfAddrFlag,
			debug.PProfPortFlag,
			debug.MemProfileRateFlag,
			debug.CPUProfileFlag,
			debug.TraceFlag,
			debug.BlockProfileRateFlag,
			debug.MutexProfileFractionFlag,
		},
	},
	{
		Name: "rpc",
		Flags: []cli.Flag{
			flags.CertFlag,
			flags.BeaconRPCProviderFlag,
			flags.EnableRPCFlag,
			flags.RPCHost,
			flags.RPCPort,
			flags.HTTPServerPort,
			flags.HTTPServerHost,
			flags.GRPCRetriesFlag,
			flags.GRPCRetryDelayFlag,
			flags.HTTPServerCorsDomain,
			flags.GRPCHeadersFlag,
			flags.BeaconRESTApiProviderFlag,
		},
	},
	{
		Name: "proposer",
		Flags: []cli.Flag{
			flags.ProposerSettingsFlag,
			flags.ProposerSettingsURLFlag,
			flags.SuggestedFeeRecipientFlag,
			flags.EnableBuilderFlag,
			flags.BuilderGasLimitFlag,
			flags.ValidatorsRegistrationBatchSizeFlag,
			flags.GraffitiFlag,
			flags.GraffitiFileFlag,
		},
	},
	{
		Name: "remote signer",
		Flags: []cli.Flag{
			flags.Web3SignerURLFlag,
			flags.Web3SignerPublicValidatorKeysFlag,
			flags.Web3SignerKeyFileFlag,
		},
	},
	{
		Name: "slasher",
		Flags: []cli.Flag{
			flags.SlasherRPCProviderFlag,
			flags.SlasherCertFlag,
		},
	},
	{
		Name: "misc",
		Flags: []cli.Flag{
			flags.EnableWebFlag,
			flags.DisablePenaltyRewardLogFlag,
			flags.DisableAccountMetricsFlag,
			flags.EnableDistributed,
			flags.AuthTokenPathFlag,
		},
	},
	{
		Name:  "features",
		Flags: features.ActiveFlags(features.ValidatorFlags),
	},
	{
		Name: "interop",
		Flags: []cli.Flag{
			flags.InteropNumValidators,
			flags.InteropStartIndex,
		},
	},
}

func init() {
	cli.AppHelpTemplate = appHelpTemplate

	type helpData struct {
		App        interface{}
		FlagGroups []flagGroup
	}

	originalHelpPrinter := cli.HelpPrinter
	cli.HelpPrinter = func(w io.Writer, tmpl string, data interface{}) {
		if tmpl == appHelpTemplate {
			for _, group := range appHelpFlagGroups {
				sort.Sort(cli.FlagsByName(group.Flags))
			}
			originalHelpPrinter(w, tmpl, helpData{data, appHelpFlagGroups})
		} else {
			originalHelpPrinter(w, tmpl, data)
		}
	}
}