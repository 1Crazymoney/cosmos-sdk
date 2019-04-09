package auth

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// name of this module
const ModuleName = "auth"

// app module object
type AppModule struct {
	accountKeeper AccountKeeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(accountKeeper AccountKeeper) AppModule {
	return AppModule{
		accountKeeper: accountKeeper,
	}
}

var _ sdk.AppModule = AppModule{}

// module name
func (AppModule) Name() string {
	return ModuleName
}

// register codec
func (AppModule) RegisterCodec(cdc *codec.Codec) {
	RegisterCodec(cdc)
}

// register invariants
func (AppModule) RegisterInvariants(_ sdk.InvariantRouter) {}

// module message route name
func (AppModule) Route() string { return "" }

// module handler
func (AppModule) NewHandler() sdk.Handler { return nil }

// module querier route name
func (AppModule) QuerierRoute() string {
	return QuerierRoute
}

// module querier
func (a AppModule) NewQuerierHandler() sdk.Querier {
	return NewQuerier(a.accountKeeper)
}

// module begin-block
func (AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) (sdk.Tags, error) {
	return sdk.EmptyTags(), nil
}

// module end-block
func (AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) ([]abci.ValidatorUpdate, sdk.Tags, error) {
	return []abci.ValidatorUpdate{}, sdk.EmptyTags(), nil
}
