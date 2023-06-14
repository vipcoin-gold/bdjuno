package banking

import (
	"git.ooo.ua/vipcoin/chain/x/banking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v2/types"
)

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	if len(tx.Logs) == 0 {
		return nil
	}

	switch bankingMsg := msg.(type) {
	case *types.MsgSetRewardManagerAddress:
		return m.handleMsgSetRewardManagerAddress(tx, index, bankingMsg)
	case *types.MsgSystemTransfer:
		return m.handleMsgSystemTransfer(tx, index, bankingMsg)
	case *types.MsgIssue:
		return m.handleMsgIssue(tx, index, bankingMsg)
	case *types.MsgPayment:
		return m.handleMsgPayments(tx, index, bankingMsg)
	case *types.MsgSetTransferExtra:
		return m.handleMsgSetTransferExtra(tx, index, bankingMsg)
	case *types.MsgSystemRewardTransfer:
		return m.handleMsgSystemRewardTransfer(tx, index, bankingMsg)
	case *types.MsgWithdraw:
		return m.handleMsgWithdraw(tx, index, bankingMsg)
	default:
		return nil
	}
}