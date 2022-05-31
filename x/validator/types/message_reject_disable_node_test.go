package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
)

func TestMsgRejectDisableValidator_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRejectDisableValidator
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgRejectDisableValidator{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgRejectDisableValidator{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)

				return
			}
			require.NoError(t, err)
		})
	}
}