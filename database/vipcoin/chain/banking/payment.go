package banking

import (
	"context"
	"database/sql"

	bankingtypes "git.ooo.ua/vipcoin/chain/x/banking/types"
	"git.ooo.ua/vipcoin/lib/filter"

	"github.com/forbole/bdjuno/v2/database/types"
)

// SavePayments - method that create payments to the "vipcoin_chain_banking_payment" table
func (r Repository) SavePayments(payments ...*bankingtypes.Payment) error {
	if len(payments) == 0 {
		return nil
	}

	tx, err := r.db.BeginTxx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	defer tx.Rollback()

	queryBaseTransfer := `INSERT INTO vipcoin_chain_banking_base_transfers 
       ("id", "asset", "amount", "kind", "extras", "timestamp", "tx_hash") 
     VALUES 
       (:id, :asset, :amount, :kind, :extras, :timestamp, :tx_hash)`

	queryPayment := `INSERT INTO vipcoin_chain_banking_payment
			("id", "wallet_from", "wallet_to", "fee")
			VALUES
			(:id,:wallet_from,:wallet_to,:fee)`

	for _, p := range payments {
		paymentDB := toPaymentDatabase(p)

		if _, err := tx.NamedExec(queryBaseTransfer, paymentDB); err != nil {
			return err
		}

		if _, err := tx.NamedExec(queryPayment, paymentDB); err != nil {
			return err
		}

	}

	return tx.Commit()
}

// GetPayments - method that get payments from the "vipcoin_chain_banking_payment" table
func (r Repository) GetPayments(filter filter.Filter) ([]*bankingtypes.Payment, error) {
	query, args := filter.ToJoiner().
		PrepareTable("vipcoin_chain_banking_base_transfers", "id", "asset", "amount", "kind", "extras", "timestamp", "tx_hash").
		PrepareTable("vipcoin_chain_banking_payment", "id", "wallet_from", "wallet_to", "fee").
		PrepareJoinStatement("INNER JOIN vipcoin_chain_banking_base_transfers on vipcoin_chain_banking_base_transfers.id = vipcoin_chain_banking_payment.id").
		Build("vipcoin_chain_banking_payment")

	var result []types.DBPayment
	if err := r.db.Select(&result, query, args...); err != nil {
		return []*bankingtypes.Payment{}, err
	}

	payments := make([]*bankingtypes.Payment, 0, len(result))
	for _, payment := range result {
		payments = append(payments, toPaymentDomain(payment))
	}

	return payments, nil
}

// UpdatePayments - method that update the payment in the "vipcoin_chain_banking_payment" table
func (r Repository) UpdatePayments(payments ...*bankingtypes.Payment) error {
	if len(payments) == 0 {
		return nil
	}

	tx, err := r.db.BeginTxx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	defer tx.Rollback()

	queryBaseTransfer := `UPDATE vipcoin_chain_banking_base_transfers SET
	id =:id, asset =:asset, amount =:amount, kind =:kind, extras =:extras, timestamp =:timestamp, tx_hash =:tx_hash
	WHERE id =:id;
	`

	queryPayment := `UPDATE vipcoin_chain_banking_payment SET
	id =:id, wallet_from =:wallet_from, wallet_to =:wallet_to, fee =:fee
	WHERE id =:id;
	`

	for _, payment := range payments {
		paymentDB := toPaymentDatabase(payment)

		if _, err := tx.NamedExec(queryBaseTransfer, paymentDB); err != nil {
			return err
		}

		if _, err := tx.NamedExec(queryPayment, paymentDB); err != nil {
			return err
		}
	}

	return tx.Commit()
}
