/*
 * Copyright 2022 Business Process Technologies. All rights reserved.
 */

package accounts

import (
	"reflect"
	"testing"

	accountstypes "git.ooo.ua/vipcoin/chain/x/accounts/types"
	"git.ooo.ua/vipcoin/lib/filter"
	"github.com/cosmos/cosmos-sdk/simapp"
	accountsdb "github.com/forbole/bdjuno/v2/database/vipcoin/chain/accounts"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func TestRepository_SaveAffiliateAddress(t *testing.T) {
	db, err := sqlx.Connect("pgx", "host=localhost port=5432 user=postgres dbname=juno password=postgres sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	// Create the codec
	codec := simapp.MakeTestEncodingConfig()

	type args struct {
		msg []*accountstypes.MsgSetAffiliateAddress
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				msg: []*accountstypes.MsgSetAffiliateAddress{
					{
						Creator:    "vcg1ljs7p2p9ae3en8knr3d3ke8srsfcj2zjvefv0g",
						Hash:       "a935ea2c467d7f666ea2a67870564f2efb902c05f0a2bb4b6202832aedd26cd1",
						OldAddress: "vcg1hwxejcutmafuedd8trjqumfdkst2498pggx45q",
						NewAddress: "vcg1hwxejcutmafuedd8trjqumfdkst2498pggx46q",
					},
					{
						Creator:    "vcg1ljs7p2p9ae3en8knr3d3ke8srsfcj2zjvefv1g",
						Hash:       "a935ea2c467d7f666ea2a67870564f2efb902c05f0a2bb4b6202832aedd26cd2",
						OldAddress: "vcg1hwxejcutmafuedd8trjqumfdkst2498pggx43q",
						NewAddress: "vcg1hwxejcutmafuedd8trjqumfdkst2498pggx44q",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := accountsdb.NewRepository(db, codec.Marshaler)

			if err := r.SaveAffiliateAddress(tt.args.msg...); (err != nil) != tt.wantErr {
				t.Errorf("Repository.SaveAffiliateAddress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_GetAffiliateAddress(t *testing.T) {
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
		want    []*accountstypes.MsgSetAffiliateAddress
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				accfilter: filter.NewFilter().SetArgument("hash", "a935ea2c467d7f666ea2a67870564f2efb902c05f0a2bb4b6202832aedd26cd1"),
			},
			want: []*accountstypes.MsgSetAffiliateAddress{
				{
					Creator:    "vcg1ljs7p2p9ae3en8knr3d3ke8srsfcj2zjvefv0g",
					Hash:       "a935ea2c467d7f666ea2a67870564f2efb902c05f0a2bb4b6202832aedd26cd1",
					OldAddress: "vcg1hwxejcutmafuedd8trjqumfdkst2498pggx45q",
					NewAddress: "vcg1hwxejcutmafuedd8trjqumfdkst2498pggx46q",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := accountsdb.NewRepository(db, codec.Marshaler)

			got, err := r.GetAffiliateAddress(tt.args.accfilter)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetAffiliateAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.GetAffiliateAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}