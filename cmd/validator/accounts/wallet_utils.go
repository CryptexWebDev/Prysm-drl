package accounts

import (
	"strings"

	"github.com/Dorol-Chain/Prysm-drl/v5/validator/accounts"
	"github.com/Dorol-Chain/Prysm-drl/v5/validator/accounts/iface"
	"github.com/Dorol-Chain/Prysm-drl/v5/validator/accounts/wallet"
	"github.com/Dorol-Chain/Prysm-drl/v5/validator/keymanager"
	remote_web3signer "github.com/Dorol-Chain/Prysm-drl/v5/validator/keymanager/remote-web3signer"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func walletWithKeymanager(c *cli.Context) (*wallet.Wallet, keymanager.IKeymanager, error) {
	w, err := wallet.OpenWalletOrElseCli(c, func(cliCtx *cli.Context) (*wallet.Wallet, error) {
		return nil, wallet.ErrNoWalletFound
	})
	if err != nil {
		return nil, nil, errors.Wrap(err, "could not open wallet")
	}
	km, err := w.InitializeKeymanager(c.Context, iface.InitKeymanagerConfig{ListenForChanges: false})
	if err != nil && strings.Contains(err.Error(), keymanager.IncorrectPasswordErrMsg) {
		return nil, nil, errors.New("wrong wallet password entered")
	}
	if err != nil {
		return nil, nil, errors.Wrap(err, accounts.ErrCouldNotInitializeKeymanager)
	}
	return w, km, nil
}

func walletWithWeb3SignerKeymanager(c *cli.Context, config *remote_web3signer.SetupConfig) (*wallet.Wallet, keymanager.IKeymanager, error) {
	w := wallet.NewWalletForWeb3Signer(c)
	km, err := w.InitializeKeymanager(c.Context, iface.InitKeymanagerConfig{ListenForChanges: false, Web3SignerConfig: config})
	if err != nil {
		return nil, nil, err
	}
	return w, km, nil
}
