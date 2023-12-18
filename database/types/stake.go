package types

type (
	// StakeMsgSell - db model for 'overgold_stake_sell'
	StakeMsgSell struct {
		ID      uint64 `db:"id"`
		TxHash  string `db:"tx_hash"`
		Creator string `db:"creator"`
		Amount  uint64 `db:"amount"`
	}

	// StakeMsgSellCancel - db model for 'overgold_stake_sell_cancel'
	StakeMsgSellCancel struct {
		ID      uint64 `db:"id"`
		TxHash  string `db:"tx_hash"`
		Creator string `db:"creator"`
		Amount  uint64 `db:"amount"`
	}

	// StakeMsgBuy - db model for 'overgold_stake_buy'
	StakeMsgBuy struct {
		ID      uint64 `db:"id"`
		TxHash  string `db:"tx_hash"`
		Creator string `db:"creator"`
		Amount  uint64 `db:"amount"`
	}

	// StakeMsgDistribute - db model for 'overgold_stake_distribute_rewards'
	StakeMsgDistribute struct {
		ID      uint64 `db:"id"`
		TxHash  string `db:"tx_hash"`
		Creator string `db:"creator"`
	}

	// StakeMsgClaim - db model for 'overgold_stake_claim_reward'
	StakeMsgClaim struct {
		ID      uint64 `db:"id"`
		TxHash  string `db:"tx_hash"`
		Creator string `db:"creator"`
		Amount  uint64 `db:"amount"`
	}
)