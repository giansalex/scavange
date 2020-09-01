// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/giansalex/scavenge/x/wasm/internal/types
// ALIASGEN: github.com/giansalex/scavenge/x/wasm/internal/keeper
package wasm

import (
	"github.com/giansalex/scavenge/x/wasm/internal/keeper"
	"github.com/giansalex/scavenge/x/wasm/internal/types"
)

const (
	DefaultParamspace               = types.DefaultParamspace
	ModuleName                      = types.ModuleName
	StoreKey                        = types.StoreKey
	TStoreKey                       = types.TStoreKey
	QuerierRoute                    = types.QuerierRoute
	RouterKey                       = types.RouterKey
	MaxWasmSize                     = types.MaxWasmSize
	MaxLabelSize                    = types.MaxLabelSize
	BuildTagRegexp                  = types.BuildTagRegexp
	MaxBuildTagSize                 = types.MaxBuildTagSize
	CustomEventType                 = types.CustomEventType
	AttributeKeyContractAddr        = types.AttributeKeyContractAddr
	ProposalTypeStoreCode           = types.ProposalTypeStoreCode
	ProposalTypeInstantiateContract = types.ProposalTypeInstantiateContract
	ProposalTypeMigrateContract     = types.ProposalTypeMigrateContract
	ProposalTypeUpdateAdmin         = types.ProposalTypeUpdateAdmin
	ProposalTypeClearAdmin          = types.ProposalTypeClearAdmin
	GasMultiplier                   = keeper.GasMultiplier
	MaxGas                          = keeper.MaxGas
	QueryListContractByCode         = keeper.QueryListContractByCode
	QueryGetContract                = keeper.QueryGetContract
	QueryGetContractState           = keeper.QueryGetContractState
	QueryGetCode                    = keeper.QueryGetCode
	QueryListCode                   = keeper.QueryListCode
	QueryMethodContractStateSmart   = keeper.QueryMethodContractStateSmart
	QueryMethodContractStateAll     = keeper.QueryMethodContractStateAll
	QueryMethodContractStateRaw     = keeper.QueryMethodContractStateRaw
)

var (
	// functions aliases
	RegisterCodec             = types.RegisterCodec
	ValidateGenesis           = types.ValidateGenesis
	ConvertToProposals        = types.ConvertToProposals
	GetCodeKey                = types.GetCodeKey
	GetContractAddressKey     = types.GetContractAddressKey
	GetContractStorePrefixKey = types.GetContractStorePrefixKey
	NewCodeInfo               = types.NewCodeInfo
	NewAbsoluteTxPosition     = types.NewAbsoluteTxPosition
	NewContractInfo           = types.NewContractInfo
	NewEnv                    = types.NewEnv
	NewWasmCoins              = types.NewWasmCoins
	ParseEvents               = types.ParseEvents
	DefaultWasmConfig         = types.DefaultWasmConfig
	DefaultParams             = types.DefaultParams
	InitGenesis               = keeper.InitGenesis
	ExportGenesis             = keeper.ExportGenesis
	NewMessageHandler         = keeper.NewMessageHandler
	DefaultEncoders           = keeper.DefaultEncoders
	EncodeBankMsg             = keeper.EncodeBankMsg
	NoCustomMsg               = keeper.NoCustomMsg
	EncodeStakingMsg          = keeper.EncodeStakingMsg
	EncodeWasmMsg             = keeper.EncodeWasmMsg
	NewKeeper                 = keeper.NewKeeper
	NewQuerier                = keeper.NewQuerier
	DefaultQueryPlugins       = keeper.DefaultQueryPlugins
	BankQuerier               = keeper.BankQuerier
	NoCustomQuerier           = keeper.NoCustomQuerier
	StakingQuerier            = keeper.StakingQuerier
	WasmQuerier               = keeper.WasmQuerier
	NewWasmProposalHandler    = keeper.NewWasmProposalHandler

	// variable aliases
	ModuleCdc            = types.ModuleCdc
	DefaultCodespace     = types.DefaultCodespace
	ErrCreateFailed      = types.ErrCreateFailed
	ErrAccountExists     = types.ErrAccountExists
	ErrInstantiateFailed = types.ErrInstantiateFailed
	ErrExecuteFailed     = types.ErrExecuteFailed
	ErrGasLimit          = types.ErrGasLimit
	ErrInvalidGenesis    = types.ErrInvalidGenesis
	ErrNotFound          = types.ErrNotFound
	ErrQueryFailed       = types.ErrQueryFailed
	ErrInvalidMsg        = types.ErrInvalidMsg
	KeyLastCodeID        = types.KeyLastCodeID
	KeyLastInstanceID    = types.KeyLastInstanceID
	CodeKeyPrefix        = types.CodeKeyPrefix
	ContractKeyPrefix    = types.ContractKeyPrefix
	ContractStorePrefix  = types.ContractStorePrefix
	EnableAllProposals   = types.EnableAllProposals
	DisableAllProposals  = types.DisableAllProposals
)

type (
	ProposalType            = types.ProposalType
	GenesisState            = types.GenesisState
	Code                    = types.Code
	Contract                = types.Contract
	MsgStoreCode            = types.MsgStoreCode
	MsgInstantiateContract  = types.MsgInstantiateContract
	MsgExecuteContract      = types.MsgExecuteContract
	MsgMigrateContract      = types.MsgMigrateContract
	MsgUpdateAdmin          = types.MsgUpdateAdmin
	MsgClearAdmin           = types.MsgClearAdmin
	Model                   = types.Model
	CodeInfo                = types.CodeInfo
	ContractInfo            = types.ContractInfo
	CreatedAt               = types.AbsoluteTxPosition
	WasmConfig              = types.WasmConfig
	MessageHandler          = keeper.MessageHandler
	BankEncoder             = keeper.BankEncoder
	CustomEncoder           = keeper.CustomEncoder
	StakingEncoder          = keeper.StakingEncoder
	WasmEncoder             = keeper.WasmEncoder
	MessageEncoders         = keeper.MessageEncoders
	Keeper                  = keeper.Keeper
	ContractInfoWithAddress = keeper.ContractInfoWithAddress
	GetCodeResponse         = keeper.GetCodeResponse
	ListCodeResponse        = keeper.ListCodeResponse
	QueryHandler            = keeper.QueryHandler
	CustomQuerier           = keeper.CustomQuerier
	QueryPlugins            = keeper.QueryPlugins
)
