package accounts

import (
	"strings"

	"github.com/Dorol-Chain/Prysm-drl/v5/cmd"
	"github.com/Dorol-Chain/Prysm-drl/v5/cmd/validator/flags"
	"github.com/Dorol-Chain/Prysm-drl/v5/validator/accounts"
	"github.com/Dorol-Chain/Prysm-drl/v5/validator/accounts/userprompt"
	"github.com/Dorol-Chain/Prysm-drl/v5/validator/client"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func accountsDelete(c *cli.Context) error {
	w, km, err := walletWithKeymanager(c)
	if err != nil {
		return err
	}
	dialOpts := client.ConstructDialOptions(
		c.Int(cmd.GrpcMaxCallRecvMsgSizeFlag.Name),
		c.String(flags.CertFlag.Name),
		c.Uint(flags.GRPCRetriesFlag.Name),
		c.Duration(flags.GRPCRetryDelayFlag.Name),
	)
	grpcHeaders := strings.Split(c.String(flags.GRPCHeadersFlag.Name), ",")

	opts := []accounts.Option{
		accounts.WithWallet(w),
		accounts.WithKeymanager(km),
		accounts.WithGRPCDialOpts(dialOpts),
		accounts.WithBeaconRPCProvider(c.String(flags.BeaconRPCProviderFlag.Name)),
		accounts.WithBeaconRESTApiProvider(c.String(flags.BeaconRESTApiProviderFlag.Name)),
		accounts.WithGRPCHeaders(grpcHeaders),
	}

	// Get full set of public keys from the keymanager.
	validatingPublicKeys, err := km.FetchValidatingPublicKeys(c.Context)
	if err != nil {
		return err
	}
	if len(validatingPublicKeys) == 0 {
		return errors.New("wallet is empty, no accounts to delete")
	}
	// Filter keys either from CLI flag or from interactive session.
	filteredPubKeys, err := accounts.FilterPublicKeysFromUserInput(
		c,
		flags.DeletePublicKeysFlag,
		validatingPublicKeys,
		userprompt.SelectAccountsDeletePromptText,
	)
	if err != nil {
		return errors.Wrap(err, "could not filter public keys for deletion")
	}
	opts = append(opts, accounts.WithFilteredPubKeys(filteredPubKeys))
	opts = append(opts, accounts.WithWalletKeyCount(len(validatingPublicKeys)))
	opts = append(opts, accounts.WithDeletePublicKeys(c.IsSet(flags.DeletePublicKeysFlag.Name)))

	acc, err := accounts.NewCLIManager(opts...)
	if err != nil {
		return err
	}
	return acc.Delete(c.Context)
}
