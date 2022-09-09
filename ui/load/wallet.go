package load

import (
	"errors"
	"fmt"
	"sort"

	"github.com/decred/dcrd/dcrutil/v4"
	"gitlab.com/raedah/cryptopower/libwallet"
	"gitlab.com/raedah/cryptopower/wallet"
)

type WalletItem struct {
	Wallet       *libwallet.Wallet
	TotalBalance string
}

type WalletLoad struct {
	MultiWallet *libwallet.MultiWallet
	TxAuthor    libwallet.TxAuthor

	UnspentOutputs *wallet.UnspentOutputs
	Wallet         *wallet.Wallet

	SelectedWallet  *WalletItem
	SelectedAccount *int
}

func (wl *WalletLoad) SortedWalletList() []*libwallet.Wallet {
	wallets := wl.MultiWallet.AllWallets()

	sort.Slice(wallets, func(i, j int) bool {
		return wallets[i].ID < wallets[j].ID
	})

	return wallets
}

func (wl *WalletLoad) TotalWalletsBalance() (dcrutil.Amount, error) {
	totalBalance := int64(0)
	for _, w := range wl.MultiWallet.AllWallets() {
		accountsResult, err := w.GetAccountsRaw()
		if err != nil {
			return -1, err
		}

		for _, account := range accountsResult.Acc {
			totalBalance += account.TotalBalance
		}
	}

	return dcrutil.Amount(totalBalance), nil
}

func (wl *WalletLoad) TotalWalletBalance(walletID int) (dcrutil.Amount, error) {
	totalBalance := int64(0)
	wallet := wl.MultiWallet.WalletWithID(walletID)
	if wallet == nil {
		return -1, errors.New(libwallet.ErrNotExist)
	}

	accountsResult, err := wallet.GetAccountsRaw()
	if err != nil {
		return -1, err
	}

	for _, account := range accountsResult.Acc {
		totalBalance += account.TotalBalance
	}

	return dcrutil.Amount(totalBalance), nil
}

func (wl *WalletLoad) SpendableWalletBalance(walletID int) (dcrutil.Amount, error) {
	spendableBal := int64(0)
	wallet := wl.MultiWallet.WalletWithID(walletID)
	if wallet == nil {
		return -1, errors.New(libwallet.ErrNotExist)
	}

	accountsResult, err := wallet.GetAccountsRaw()
	if err != nil {
		return -1, err
	}

	for _, account := range accountsResult.Acc {
		spendableBal += account.Balance.Spendable
	}

	return dcrutil.Amount(spendableBal), nil
}

func (wl *WalletLoad) HDPrefix() string {
	switch wl.Wallet.Net {
	case libwallet.Testnet3:
		return libwallet.TestnetHDPath
	case "mainnet":
		return libwallet.MainnetHDPath
	default:
		return ""
	}
}

func (wl *WalletLoad) WalletDirectory() string {
	return fmt.Sprintf("%s/%s", wl.Wallet.Root, wl.Wallet.Net)
}

func (wl *WalletLoad) DataSize() string {
	v, err := wl.MultiWallet.RootDirFileSizeInBytes()
	if err != nil {
		return "Unknown"
	}
	return fmt.Sprintf("%f GB", float64(v)*1e-9)
}
