package bank

import (
	"context"
	"database/sql"
	"errors"

	"git.ooo.ua/vipcoin/lib/errs"
	"git.ooo.ua/vipcoin/lib/filter"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/lib/pq"
)

// GetAllMsgMultiSend - method that get data from a db (msg_multi_send).
func (r Repository) GetAllMsgMultiSend(filter filter.Filter) ([]bank.MsgMultiSend, error) {
	query, args := filter.Build(tableMsgMultiSend)

	var result []msgMultiSend
	if err := r.db.Select(&result, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NotFound{What: tableMsgMultiSend}
		}

		return nil, errs.Internal{Cause: err.Error()}
	}
	if len(result) == 0 {
		return nil, errs.NotFound{What: tableMsgMultiSend}
	}

	return toGetMsgMultiSendDomainList(result)
}

// InsertMsgMultiSend - insert a new MsgCreateAddresses in a database (msg_multi_send).
func (r Repository) InsertMsgMultiSend(hash string, msgs ...bank.MsgMultiSend) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	tx, err := r.db.BeginTxx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return errs.Internal{Cause: err.Error()}
	}

	defer func() {
		_ = tx.Rollback()
	}()

	q := `
		INSERT INTO msg_multi_send (
			tx_hash, inputs, outputs
		) VALUES (
			$1, $2, $3
		) RETURNING
			id, tx_hash, inputs, outputs
	`

	// NOTE: use tx.Exec for custom type pq.Array(DbSendDataList)
	for _, msg := range msgs {
		m := toMsgMultiSendDatabase(hash, msg)
		if _, err = tx.Exec(q, m.TxHash, pq.Array(m.Inputs), pq.Array(m.Ouputs)); err != nil {
			return errs.Internal{Cause: err.Error()}
		}
	}

	return tx.Commit()
}