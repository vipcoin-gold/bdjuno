package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/cosmos/cosmos-sdk/types"
)

type (
	// DBWallets represents a single row inside the "vipcoin_chain_wallets_wallets" table
	DBWallets struct {
		Address        string    `db:"address"`
		AccountAddress string    `db:"account_address"`
		Kind           int32     `db:"kind"`
		State          int32     `db:"state"`
		Balance        BalanceDB `db:"balance"`
		Extras         ExtraDB   `db:"extras"`
		DefaultStatus  bool      `db:"default_status"`
	}

	// DBSetWalletKind represents a single row inside the "vipcoin_chain_wallets_set_wallet_kind" table
	DBSetWalletKind struct {
		Creator string `db:"creator"`
		Address string `db:"address"`
		Kind    int32  `db:"kind"`
	}

	// DBSetWalletState represents a single row inside the "vipcoin_chain_wallets_set_wallet_state" table
	DBSetWalletState struct {
		Creator string `db:"creator"`
		Address string `db:"address"`
		State   int32  `db:"state"`
	}

	// DBCreateWallet represents a single row inside the "vipcoin_chain_wallets_create_wallet" table
	DBCreateWallet struct {
		Creator        string  `db:"creator"`
		Address        string  `db:"address"`
		AccountAddress string  `db:"account_address"`
		Kind           int32   `db:"kind"`
		State          int32   `db:"state"`
		Extras         ExtraDB `db:"extras"`
	}

	// DBCreateWalletWithBalance represents a single row inside the "vipcoin_chain_wallets_create_wallet_with_balance" table
	DBCreateWalletWithBalance struct {
		Creator        string    `db:"creator"`
		Address        string    `db:"address"`
		AccountAddress string    `db:"account_address"`
		Kind           int32     `db:"kind"`
		State          int32     `db:"state"`
		Extras         ExtraDB   `db:"extras"`
		DefaultStatus  bool      `db:"default_status"`
		Balance        BalanceDB `db:"balance"`
	}

	// DBSetDefaultWallet represents a single row inside the "vipcoin_chain_wallets_set_default_wallet" table
	DBSetDefaultWallet struct {
		Creator string `db:"creator"`
		Address string `db:"address"`
	}

	// DBSetExtra represents a single row inside the "vipcoin_chain_wallets_set_extra" table
	DBSetExtra struct {
		Creator string  `db:"creator"`
		Address string  `db:"address"`
		Extras  ExtraDB `db:"extras"`
	}

	// BalanceDB helpers type
	BalanceDB struct {
		Balance types.Coins
	}
)

// Make the BalanceDB struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (b BalanceDB) Value() (driver.Value, error) {
	return json.Marshal(b.Balance)
}

// Make the BalanceDB struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (b *BalanceDB) Scan(value interface{}) error {
	v, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(v, &b.Balance)
}
