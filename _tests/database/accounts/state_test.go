package accounts

import (
	"reflect"
	"testing"

	accountstypes "git.ooo.ua/vipcoin/chain/x/accounts/types"
	"git.ooo.ua/vipcoin/lib/filter"
	"github.com/cosmos/cosmos-sdk/simapp"
	accountsdb "github.com/forbole/bdjuno/v2/database/vipcoin/chain/accounts"
	"github.com/forbole/bdjuno/v2/modules/vipcoin/chain/accounts"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func TestRepository_SaveState(t *testing.T) {
	db, err := sqlx.Connect("pgx", "host=localhost port=5432 user=postgres dbname=juno password=postgres sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	// Create the codec
	codec := simapp.MakeTestEncodingConfig()

	type args struct {
		msg []*accountstypes.MsgSetState
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				msg: []*accountstypes.MsgSetState{
					{
						Creator: "vcg1ljs7p2p9ae3en8knr3d3ke8srsfcj2zjvefv0g",
						Hash:    "a935ea2c467d7f666ea2a67870564f2efb902c05f0a2bb4b6202832aedd26cd1",
						State:   accountstypes.ACCOUNT_STATE_ACTIVE,
						Reason:  "valid account",
					},
					{
						Creator: "vcg1ljs7p2p9ae3en8knr3d3ke8srsfcj2zjvefv1g",
						Hash:    "a935ea2c467d7f666ea2a67870564f2efb902c05f0a2bb4b6202832aedd26cd2",
						State:   accountstypes.ACCOUNT_STATE_BLOCKED,
						Reason:  "fraud",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := accountsdb.NewRepository(db, codec.Marshaler)

			if err := r.SaveState(tt.args.msg...); (err != nil) != tt.wantErr {
				t.Errorf("Repository.SaveState() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_GetState(t *testing.T) {
	db, err := sqlx.Connect("pgx", "host=localhost port=5432 user=postgres dbname=juno password=postgres sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	// Create the codec
	codec := simapp.MakeTestEncodingConfig()

	type args struct {
		accfilter filter.Filter
	}
	tests := []struct {
		name    string
		args    args
		want    []*accountstypes.MsgSetState
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				filter.NewFilter().SetArgument(accounts.FieldHash, "a935ea2c467d7f666ea2a67870564f2efb902c05f0a2bb4b6202832aedd26cd2"),
			},
			want: []*accountstypes.MsgSetState{
				{
					Creator: "vcg1ljs7p2p9ae3en8knr3d3ke8srsfcj2zjvefv1g",
					Hash:    "a935ea2c467d7f666ea2a67870564f2efb902c05f0a2bb4b6202832aedd26cd2",
					State:   accountstypes.ACCOUNT_STATE_BLOCKED,
					Reason:  "fraud",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := accountsdb.NewRepository(db, codec.Marshaler)

			got, err := r.GetState(tt.args.accfilter)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.GetState() = %v, want %v", got, tt.want)
			}
		})
	}
}