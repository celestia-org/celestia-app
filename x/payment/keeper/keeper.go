package keeper

import (
	"context"
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/celestiaorg/celestia-app/x/payment/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TokenDenomination = "tia"
)

// Keeper handles all the state changes for the celestia-app module.
type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey sdk.StoreKey
	memKey   sdk.StoreKey
	bank     BankKeeper
}

func NewKeeper(cdc codec.BinaryCodec, bank BankKeeper, storeKey, memKey sdk.StoreKey) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
		bank:     bank,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

//  MsgPayForMessage moves a user's coins to the module address and burns them.
func (k Keeper) PayForMessage(goCtx context.Context, msg *types.MsgPayForMessage) (*types.MsgPayForMessageResponse, error) {
	return &types.MsgPayForMessageResponse{}, nil
}

// BankKeeper restricts the funtionality of the bank keeper used in the payment keeper
type BankKeeper interface {
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
}
