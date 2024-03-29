// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package seaport

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AdditionalRecipient is an auto generated low-level Go binding around an user-defined struct.
type AdditionalRecipient struct {
	Amount    *big.Int
	Recipient common.Address
}

// AdvancedOrder is an auto generated low-level Go binding around an user-defined struct.
type AdvancedOrder struct {
	Parameters  OrderParameters
	Numerator   *big.Int
	Denominator *big.Int
	Signature   []byte
	ExtraData   []byte
}

// BasicOrderParameters is an auto generated low-level Go binding around an user-defined struct.
type BasicOrderParameters struct {
	ConsiderationToken                common.Address
	ConsiderationIdentifier           *big.Int
	ConsiderationAmount               *big.Int
	Offerer                           common.Address
	Zone                              common.Address
	OfferToken                        common.Address
	OfferIdentifier                   *big.Int
	OfferAmount                       *big.Int
	BasicOrderType                    uint8
	StartTime                         *big.Int
	EndTime                           *big.Int
	ZoneHash                          [32]byte
	Salt                              *big.Int
	OffererConduitKey                 [32]byte
	FulfillerConduitKey               [32]byte
	TotalOriginalAdditionalRecipients *big.Int
	AdditionalRecipients              []AdditionalRecipient
	Signature                         []byte
}

// ConsiderationItem is an auto generated low-level Go binding around an user-defined struct.
type ConsiderationItem struct {
	ItemType             uint8
	Token                common.Address
	IdentifierOrCriteria *big.Int
	StartAmount          *big.Int
	EndAmount            *big.Int
	Recipient            common.Address
}

// CriteriaResolver is an auto generated low-level Go binding around an user-defined struct.
type CriteriaResolver struct {
	OrderIndex    *big.Int
	Side          uint8
	Index         *big.Int
	Identifier    *big.Int
	CriteriaProof [][32]byte
}

// Execution is an auto generated low-level Go binding around an user-defined struct.
type Execution struct {
	Item       ReceivedItem
	Offerer    common.Address
	ConduitKey [32]byte
}

// Fulfillment is an auto generated low-level Go binding around an user-defined struct.
type Fulfillment struct {
	OfferComponents         []FulfillmentComponent
	ConsiderationComponents []FulfillmentComponent
}

// FulfillmentComponent is an auto generated low-level Go binding around an user-defined struct.
type FulfillmentComponent struct {
	OrderIndex *big.Int
	ItemIndex  *big.Int
}

// OfferItem is an auto generated low-level Go binding around an user-defined struct.
type OfferItem struct {
	ItemType             uint8
	Token                common.Address
	IdentifierOrCriteria *big.Int
	StartAmount          *big.Int
	EndAmount            *big.Int
}

// Order is an auto generated low-level Go binding around an user-defined struct.
type Order struct {
	Parameters OrderParameters
	Signature  []byte
}

// OrderComponents is an auto generated low-level Go binding around an user-defined struct.
type OrderComponents struct {
	Offerer       common.Address
	Zone          common.Address
	Offer         []OfferItem
	Consideration []ConsiderationItem
	OrderType     uint8
	StartTime     *big.Int
	EndTime       *big.Int
	ZoneHash      [32]byte
	Salt          *big.Int
	ConduitKey    [32]byte
	Counter       *big.Int
}

// OrderParameters is an auto generated low-level Go binding around an user-defined struct.
type OrderParameters struct {
	Offerer                         common.Address
	Zone                            common.Address
	Offer                           []OfferItem
	Consideration                   []ConsiderationItem
	OrderType                       uint8
	StartTime                       *big.Int
	EndTime                         *big.Int
	ZoneHash                        [32]byte
	Salt                            *big.Int
	ConduitKey                      [32]byte
	TotalOriginalConsiderationItems *big.Int
}

// ReceivedItem is an auto generated low-level Go binding around an user-defined struct.
type ReceivedItem struct {
	ItemType   uint8
	Token      common.Address
	Identifier *big.Int
	Amount     *big.Int
	Recipient  common.Address
}

// SpentItem is an auto generated low-level Go binding around an user-defined struct.
type SpentItem struct {
	ItemType   uint8
	Token      common.Address
	Identifier *big.Int
	Amount     *big.Int
}

// SeaportMetaData contains all meta data concerning the Seaport contract.
var SeaportMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduitController\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BadContractSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadFraction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BadReturnValueFromERC20OnTransfer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"name\":\"BadSignatureV\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotCancelOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ConsiderationCriteriaResolverOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ConsiderationLengthNotEqualToTotalOriginal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"considerationIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortfallAmount\",\"type\":\"uint256\"}],\"name\":\"ConsiderationNotMet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CriteriaNotEnabledForItem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"identifiers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"ERC1155BatchTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InexactFraction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientNativeTokensSupplied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Invalid1155BatchTransferEncoding\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBasicOrderParameterEncoding\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"InvalidCallToConduit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"InvalidConduit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"InvalidContractOrder\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"InvalidERC721TransferAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFulfillmentComponentData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNativeOfferItem\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"InvalidRestrictedOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"InvalidTime\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fulfillmentIndex\",\"type\":\"uint256\"}],\"name\":\"MismatchedFulfillmentOfferAndConsiderationComponents\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"}],\"name\":\"MissingFulfillmentComponentOnAggregation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingItemAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingOriginalConsiderationItems\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NativeTokenTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"NoContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoReentrantCalls\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSpecifiedOrdersAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OfferAndConsiderationRequiredOnFulfillment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OfferCriteriaResolverOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderAlreadyFilled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"}],\"name\":\"OrderCriteriaResolverOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderIsCancelled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderPartiallyFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PartialFillsNotEnabledForOrder\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"considerationIndex\",\"type\":\"uint256\"}],\"name\":\"UnresolvedConsiderationCriteria\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offerIndex\",\"type\":\"uint256\"}],\"name\":\"UnresolvedOfferCriteria\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnusedItemParameters\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCounter\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"}],\"name\":\"CounterIncremented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSpentItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structReceivedItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"}],\"name\":\"OrderFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structOrderParameters\",\"name\":\"orderParameters\",\"type\":\"tuple\"}],\"name\":\"OrderValidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"orderHashes\",\"type\":\"bytes32[]\"}],\"name\":\"OrdersMatched\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"counter\",\"type\":\"uint256\"}],\"internalType\":\"structOrderComponents[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"cancel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"cancelled\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"uint120\",\"name\":\"numerator\",\"type\":\"uint120\"},{\"internalType\":\"uint120\",\"name\":\"denominator\",\"type\":\"uint120\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structAdvancedOrder\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"criteriaProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structCriteriaResolver[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"fulfillAdvancedOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"uint120\",\"name\":\"numerator\",\"type\":\"uint120\"},{\"internalType\":\"uint120\",\"name\":\"denominator\",\"type\":\"uint120\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structAdvancedOrder[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"criteriaProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structCriteriaResolver[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[][]\",\"name\":\"\",\"type\":\"tuple[][]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[][]\",\"name\":\"\",\"type\":\"tuple[][]\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maximumFulfilled\",\"type\":\"uint256\"}],\"name\":\"fulfillAvailableAdvancedOrders\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"},{\"components\":[{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structReceivedItem\",\"name\":\"item\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"internalType\":\"structExecution[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[][]\",\"name\":\"\",\"type\":\"tuple[][]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[][]\",\"name\":\"\",\"type\":\"tuple[][]\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maximumFulfilled\",\"type\":\"uint256\"}],\"name\":\"fulfillAvailableOrders\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"},{\"components\":[{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structReceivedItem\",\"name\":\"item\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"internalType\":\"structExecution[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"considerationToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"considerationIdentifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"considerationAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"offerToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offerIdentifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offerAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumBasicOrderType\",\"name\":\"basicOrderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"offererConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalAdditionalRecipients\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structAdditionalRecipient[]\",\"name\":\"additionalRecipients\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structBasicOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"}],\"name\":\"fulfillBasicOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"considerationToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"considerationIdentifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"considerationAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"offerToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offerIdentifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offerAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumBasicOrderType\",\"name\":\"basicOrderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"offererConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalAdditionalRecipients\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structAdditionalRecipient[]\",\"name\":\"additionalRecipients\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structBasicOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"}],\"name\":\"fulfillBasicOrder_efficient_6GL6yc\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"}],\"name\":\"fulfillOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractOfferer\",\"type\":\"address\"}],\"name\":\"getContractOffererNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"}],\"name\":\"getCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"counter\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"counter\",\"type\":\"uint256\"}],\"internalType\":\"structOrderComponents\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"getOrderHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"getOrderStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidated\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isCancelled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"totalFilled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSize\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incrementCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newCounter\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"information\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"conduitController\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"uint120\",\"name\":\"numerator\",\"type\":\"uint120\"},{\"internalType\":\"uint120\",\"name\":\"denominator\",\"type\":\"uint120\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structAdvancedOrder[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"criteriaProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structCriteriaResolver[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[]\",\"name\":\"offerComponents\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[]\",\"name\":\"considerationComponents\",\"type\":\"tuple[]\"}],\"internalType\":\"structFulfillment[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"matchAdvancedOrders\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structReceivedItem\",\"name\":\"item\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"internalType\":\"structExecution[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[]\",\"name\":\"offerComponents\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[]\",\"name\":\"considerationComponents\",\"type\":\"tuple[]\"}],\"internalType\":\"structFulfillment[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"name\":\"matchOrders\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structReceivedItem\",\"name\":\"item\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"internalType\":\"structExecution[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"name\":\"validate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6101c06040523480156200001257600080fd5b5060405162005e1938038062005e19833981016040819052620000359162000245565b808080808080808080806200004962000119565b610120526101005260e05260c081905260a0828152466101408190526040805160009485526020879052948152606091825230608090815292842085825293909152939052610160526001600160a01b038316610180819052630a96ad3960e01b825282519092630a96ad3992600480820193918290030181865afa158015620000d7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000fd919062000277565b506101a0525050600160005550620002fb975050505050505050565b600080808080806200014560408051808201909152600781526614d9585c1bdc9d60ca1b602082015290565b805190602001209550604051806040016040528060038152602001620c4b8d60ea1b81525080519060200120945060006040518060a00160405280606a815260200162005daf606a9139905060006040518060c001604052806084815260200162005c05608491399050600060405180610100016040528060d4815260200162005cdb60d49139905060405180608001604052806052815260200162005c8960529139805190602001209650828051906020012095508180519060200120945060008183856040516020016200021e93929190620002ce565b60405160208183030381529060405290508080519060200120945050505050909192939495565b6000602082840312156200025857600080fd5b81516001600160a01b03811681146200027057600080fd5b9392505050565b600080604083850312156200028b57600080fd5b505080516020909101519092909150565b6000815160005b81811015620002bf5760208185018101518683015201620002a3565b50600093019283525090919050565b6000620002f2620002eb620002e484886200029c565b866200029c565b846200029c565b95945050505050565b60805160a05160c05160e05161010051610120516101405161016051610180516101a051615872620003936000396000612bc10152600081816111c60152612b9001526000612a87015260006129ce015260008181610b450152611577015260008181610ad301526113a5015260008181610a6c01526114f0015260006129fe01526000612a4801526000612a2401526158726000f3fe6080604052600436106100fd5760003560e01c8063a900866b11610095578063f07ec37311610064578063f07ec373146102e6578063f2d12b1214610306578063f47b774014610319578063fb0f3ee114610111578063fd9f1e101461033d57600080fd5b8063a900866b14610277578063b3a34c4c146102ad578063e7acab24146102c0578063ed98a574146102d357600080fd5b806379df72bd116100d157806379df72bd146101f657806387201b41146102165780638814773214610237578063a81744041461025757600080fd5b801561011157806306fdde031461013957806346423aa71461015b5780635b34b966146101d357600080fd5b3661010c5761010a61035d565b005b600080fd5b61012461011f366004614dfd565b610372565b60405190151581526020015b60405180910390f35b34801561014557600080fd5b5061014e610383565b6040516101309190614e7e565b34801561016757600080fd5b506101b1610176366004614e91565b60009081526002602052604090205460ff808216926101008304909116916001600160781b03620100008204811692600160881b9092041690565b6040805194151585529215156020850152918301526060820152608001610130565b3480156101df57600080fd5b506101e8610392565b604051908152602001610130565b34801561020257600080fd5b506101e8610211366004614eaa565b61039c565b610229610224366004614f50565b6103d4565b604051610130929190615102565b34801561024357600080fd5b50610124610252366004615151565b61044d565b61026a610265366004615192565b61046c565b60405161013091906151fd565b34801561028357600080fd5b506101e8610292366004615210565b6001600160a01b031660009081526003602052604090205490565b6101246102bb36600461522d565b61050b565b6101246102ce366004615276565b610589565b6102296102e1366004615305565b6105c8565b3480156102f257600080fd5b506101e8610301366004615210565b61066f565b61026a6103143660046153ad565b61068d565b34801561032557600080fd5b5061032e6106d6565b6040516101309392919061545a565b34801561034957600080fd5b50610124610358366004615151565b6106ee565b60036000541461037057610370346106fa565b565b600061037d8261070c565b92915050565b606061038d61093b565b905090565b600061038d610953565b6000806103a960046109c4565b90506103cd6103be826109d55b63ffffffff16565b610140830135610a54565b3590565b9392505050565b6060806104396103ef6103e760046109c4565b610b9b6103b6565b6104066103fe60046020610c05565b610c246103b6565b61041d61041560046040610c05565b610c7e6103b6565b61042c61041560046060610c05565b89338a15028a0189610cd8565b915091509b509b9950505050505050505050565b60006103cd61046761045f60046109c4565b610d116103b6565b610d6b565b606061050061048661047e60046109c4565b610e916103b6565b60408051600080825260208201909252906104e2565b6104cf6040805160a081019091526000808252602082019081526020016000815260200160008152602001606081525090565b81526020019060019003908161049c5790505b506104fa6104f260046020610c05565b610eeb6103b6565b33610f45565b90505b949350505050565b60006103cd61052561051d60046109c4565b610f736103b6565b6040805160008082526020820190925290610581565b61056e6040805160a081019091526000808252602082019081526020016000815260200160008152602001606081525090565b81526020019060019003908161053b5790505b508433610fea565b60006105be6105a361059b60046109c4565b6111546103b6565b6105b26103fe60046020610c05565b85338615028601610fea565b9695505050505050565b60608061065e6105db61047e60046109c4565b6040805160008082526020820190925290610637565b6106246040805160a081019091526000808252602082019081526020016000815260200160008152602001606081525090565b8152602001906001900390816105f15790505b5061064761041560046020610c05565b61065661041560046040610c05565b883389610cd8565b915091509850989650505050505050565b6001600160a01b03811660009081526001602052604081205461037d565b60606106c861069f6103e760046109c4565b6106ae6103fe60046020610c05565b6106bd6104f260046040610c05565b338615028601610f45565b90505b979650505050505050565b60606000806106e36111b1565b925092509250909192565b60006103cd8383611202565b63a61be9f0600052806020526024601cfd5b600061012435600281901c906003166001821183341582148061073257610732346106fa565b506003841160a0810260240135906502030203010160d01b861a90630101020360d01b871a6107658a8883898888611344565b9450506101c4600582901b0135600086600581111561078657610786615038565b036107bc5760443560243517156107a557636ab37ce76000526004601cfd5b6107af8382611642565b6107b76116fb565b610916565b60408051602080825281830190925260009160208201818036833701905050905060028960058111156107f1576107f1615038565b036108305761082b61080960c08d0160a08e01615210565b61081960808e0160608f01615210565b338e60c001358f60e00135878761177c565b610901565b600389600581111561084457610844615038565b0361087e5761082b61085c60c08d0160a08e01615210565b61086c60808e0160608f01615210565b338e60c001358f60e0013587876117c7565b600489600581111561089257610892615038565b036108cc5761082b6108a760208d018d615210565b338d60600160208101906108bb9190615210565b8e602001358f60400135878761177c565b6109016108dc60208d018d615210565b338d60600160208101906108f09190615210565b8e602001358f6040013587876117c7565b61090b83826117fd565b610914816118bb565b505b61092185888c6118e4565b61092b6001600055565b5060019998505050505050505050565b6060602080526707536561706f727460475260606020f35b600061095d61194b565b600143034060801c33600052600160205260406000208054820192508281555050336001600160a01b03167f721c20121297512b72821b97f5326877ea8ecf4bb9948fea5bfcb6453074d37f826040516109b991815260200190565b60405180910390a290565b6000813563ffffffff16820161037d565b60006109e961016060408051918201905290565b90506109f8828261014061195d565b610a15610a0e610a09846040610c05565b611966565b6040830152565b6000610a2a610a25846060610c05565b6119ba565b9050610a37816060840152565b610a4a610a42825190565b610140840152565b50919050565b0190565b610140820151604080519084015180516000939284927f000000000000000000000000000000000000000000000000000000000000000092602090910190845b81811015610ac1578251601f1901805186825260c082208652905260209384019390920191600101610a94565b508060051b60405120945050505060007f00000000000000000000000000000000000000000000000000000000000000009150604051602060608901510160005b86811015610b2f578151601f1901805186825260e082208552905260209283019290910190600101610b02565b505060408051600587901b9020601f198a0180517f00000000000000000000000000000000000000000000000000000000000000008252928b01805197815260608c018051938152610140909c019a8b5261018082209390915295909552939097525050925250919050565b60008063ffffffff8335169050600581901b610bc06020820160408051918201905290565b8281529250602083810190850160005b83811015610bfb57610bf3610bed610be88484610c05565b611154565b82850152565b602001610bd0565b5050505050919050565b60006103cd63ffffffff610c1e6103c98686610a508516565b16840190565b60008063ffffffff8335169050600581901b610c496020820160408051918201905290565b8281529250602083810190850160005b83811015610bfb57610c76610bed610c718484610c05565b6119fe565b602001610c59565b60008063ffffffff8335169050600581901b610ca36020820160408051918201905290565b8281529250602083810190850160005b83811015610bfb57610cd0610bed610ccb8484610c05565b611a41565b602001610cb3565b606080600080610cec8b8b6000888a611a85565b91509150610cff8b8a8a8a8a8787611d85565b93509350505097509795505050505050565b60008063ffffffff8335169050600581901b610d366020820160408051918201905290565b8281529250602083810190850160005b83811015610bfb57610d63610bed610d5e8484610c05565b611f37565b602001610d46565b6000610d7561194b565b6000806000808551905060005b81811015610e84576000878281518110610d9e57610d9e6154a3565b60209081029190910101518051909150600481608001516004811115610dc657610dc6615038565b03610dd2575050610e7c565b80519450610ddf81611f7b565b60008181526002602052604081209850909650610e0190879089906001611fb6565b50865460ff16610e795780610140015181606001515114610e2457610e24612047565b610e3385878460200151612055565b865460ff191660011787556040517ff280791efe782edcf06ce15c8f4dff17601db3b88eb3805a0db7d77faf757f0490610e7090889084906155aa565b60405180910390a15b50505b600101610d82565b5060019695505050505050565b60008063ffffffff8335169050600581901b610eb66020820160408051918201905290565b8281529250602083810190850160005b83811015610bfb57610ee3610bed610ede8484610c05565b610f73565b602001610ec6565b60008063ffffffff8335169050600581901b610f106020820160408051918201905290565b8281529250602083810190850160005b83811015610bfb57610f3d610bed610f388484610c05565b6120f4565b602001610f20565b6060600080610f59878760018a5188611a85565b91509150610f6682612128565b6106cb8786848785612167565b6000610f8761020060408051918201905290565b60a08101808252909150610fa3610f9d846109c4565b8261227a565b610faf60016020840152565b610fbb60016040840152565b610fd8610fd1610fcc856020610c05565b6122b0565b6060840152565b610a4a610fe36122d8565b6080840152565b600061100e6004865160800151600481111561100857611008615038565b146122f4565b600080600061101e886001612304565b604080516001808252818301909252939650919450925060009190816020015b611046614d09565b81526020019060019003908161103e579050509050888160008151811061106f5761106f6154a3565b60200260200101819052506110848189612532565b600081600081518110611099576110996154a3565b60200260200101516000015190506110b48185858b8b612714565b6040805160018082528183019092526000916020808301908036833701905050905085816000815181106110ea576110ea6154a3565b60200260200101818152505061111b8360008151811061110c5761110c6154a3565b602002602001015182886128b8565b61113986836000015184602001518b86604001518760600151612965565b6111436001600055565b5060019a9950505050505050505050565b600061116861020060408051918201905290565b905061117c6020838101908301604061195d565b60a0810180825261118f610f9d846109c4565b6111a0610fd1610fcc856060610c05565b610a4a610fe3610fcc856080610c05565b606060008060006111c06129ca565b905060007f00000000000000000000000000000000000000000000000000000000000000009050606060005281602052806040526303312e3460635260a06000f35b600061120c61194b565b60008083815b81811015611329573687878381811061122d5761122d6154a3565b905060200281019061123f919061568d565b905060006112506020830183615210565b905060006112646040840160208501615210565b9050600061127860a08501608086016156ae565b9050813314833314171560048214178717965060006112ad6112a261129a8790565b6109d56103b6565b866101400135610a54565b60008181526002602052604090819020805461ffff19166101001781559051909a509091506001600160a01b0380851691908616907f6bacc01dbe442496068f7d234edd811f1a5f833243e0aec824f86ab861f3c90d906113119085815260200190565b60405180910390a38560010195505050505050611212565b5050801561133957611339612aa9565b506001949350505050565b600061135060006122f4565b611358612ab7565b426101643511154261014435111715611387576321ccfeb760005261014435602052610164356040526044601cfd5b610204356102643510156113a35763466aa6166000526004601cfd5b7f0000000000000000000000000000000000000000000000000000000000000000608081905260a08690526060602460c037604060646101203760e060802061016052610264358060051b6102a0016001820181526020810190508781526080602460208301376101608760a0528660c052600060e052600061020435935060005b84811015611479578060400261028401602081610100376040816101203760208101358317925060208401935060e0608020845260a0850194508a8552896020860152604081606087013750600101611425565b6001850160051b610160206060526102643594505b848110156114c9578060400261028401925060a08401935089845288602085015260408360608601376020830135919091179060010161148e565b506001600160a01b038111156114e7576339f3e3fd6000526004601cfd5b505050505060007f00000000000000000000000000000000000000000000000000000000000000009050806080528260a052606060c460c03760206101046101203760c0608020600052602060002060e0526102643560051b6102000160018152836020820152606060c4604083013750506084356001600160a01b0381166000908152600160205260408120547f000000000000000000000000000000000000000000000000000000000000000060808190529091506040608460a03760605161010052886101205260a061014461014037816101e05261018060802093505050506102643560051b6101800181815233602082015260806040820152610120606082015260a061026435026101e00160a4356084357f9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f318385a36000606052016040526105be8161163d6102208a018a6156cf565b612b05565b60c43560843560e4356101043584156116ab5760006040519050632671a55160e11b815260206004820152600160248201528660448201528460648201528360848201523360a48201528260c48201528160e48201526116a58682610104612b8a565b506116f3565b60028660058111156116bf576116bf615038565b036116e657806001146116d5576116d581612c47565b6116e184843385612c59565b6116f3565b6116f38484338585612d16565b505050505050565b346064356084356102643560061b600080805b8381101561174d5761028481013592506102a481013591508683111561173657611736612df6565b82870396506117458284612e04565b60400161170e565b508585111561175e5761175e612df6565b6117688486612e04565b848611156116f3576116f333868803612e04565b6117868183612e3e565b816117ad578260011461179c5761179c83612c47565b6117a887878787612c59565b6117be565b6117be828260028a8a8a8a8a612e5d565b50505050505050565b6117d083612edd565b6117da8183612e3e565b816117ec576117a88787878787612d16565b6117be828260038a8a8a8a8a612e5d565b6000806000806000861561182557505060843592503391505060c4356101043560e43561183a565b50339350506084359150506024356064356044355b801561184857611848612ef0565b50600586901b6101e403356102643560061b600080805b838110156118a05761028481013592506102a481013591508a1561188a57611887838761572b565b95505b611898878a8486898f612efe565b60400161185f565b506118af86898988888e612efe565b50505050505050505050565b60408151146118c75750565b60006118d4826020015190565b90506118e08183612f34565b5050565b61190c826118f860a0840160808501615210565b331415600182116004909210919091161690565b156119465760008061191e8584612f58565b909250905061194361193660a0850160808601615210565b86848463fb5014fc613038565b50505b505050565b6001600054146103705761037061308d565b80838337505050565b600063ffffffff8235166040519150808252602082018160051b81018060a084026020870183378293505b818410156119aa5780845260209093019260a001611991565b60405250919392505050565b9052565b600063ffffffff8235166040519150808252602082018160051b81018060c084026020870183378293505b818410156119aa5780845260209093019260c0016119e5565b6000611a1160a060408051918201905290565b9050611a1f8282608061195d565b611a3c611a35611a30846080610c05565b61309b565b6080830152565b919050565b600063ffffffff8235166040519150808252602082018160051b8101808360061b6020870183378293505b818410156119aa57808452602090930192604001611a6c565b60606000611a9360016122f4565b600160e61b60003516611aa4614d09565b8851600090806001600160401b03811115611ac157611ac161548d565b604051908082528060200260200182016040528015611aea578160200160208202803683370190505b50955060010160051b905060608060205b83811015611d0357808d0151945089600003611b1d5760006020860152611cfb565b6000806000611b2c888f612304565b92509250925081600003611b495750506000602087015250611cfb565b82848c01528c600190039c506000886000015160a0015190506000896000015160c001519050896000015160400151975060008851905060008b60000151608001519050600481108060005250600181118e179d505060005b81811015611c435760008a8281518110611bbe57611bbe6154a3565b602002602001015190506000518151108e179d506000611be3888884608001516130d3565b90508160800151826060015103611c005760608201819052611c15565b611c0f888884606001516130d3565b60608301525b6000611c2983606001518389896000613112565b606084018190526080909301929092525050600101611ba2565b508a5160600151805190985060005b81811015611cf25760008a8281518110611c6e57611c6e6154a3565b602002602001015190506000611c89898984608001516130d3565b90508160800151826060015103611ca65760608201819052611cbb565b611cb5898984606001516130d3565b60608301525b6000611ccf8360600151838a8a6001613112565b6060840181905260a0840180516080909501949094529092525050600101611c52565b50505050505050505b602001611afb565b5050506001600160e61b018303611d1c57611d1c613167565b611d268a8a612532565b600060205b82811015611d76578681015191508115611d6e57808c01519350600084600001519050611d6c83826000015183602001518c85604001518660600151612965565b505b602001611d2b565b50505050509550959350505050565b855185516060918291611d98818361573e565b6001600160401b03811115611daf57611daf61548d565b604051908082528060200260200182016040528015611de857816020015b611dd5614d3d565b815260200190600190039081611dcd5790505b5092506000805b83811015611e71576000611e208e60008f8581518110611e1157611e116154a3565b60200260200101518e8e613175565b805180516020830151608090920151929350151591141615611e4757826001019250611e68565b808684840381518110611e5c57611e5c6154a3565b60200260200101819052505b50600101611def565b5060005b82811015611efa576000611ea78e60018e8581518110611e9757611e976154a3565b60200260200101518e6000613175565b805180516020830151608090920151929350151591141615611ece57826001019250611ef1565b8086848785010381518110611ee557611ee56154a3565b60200260200101819052505b50600101611e75565b508015611f08578084510384525b508251600003611f1a57611f1a61320b565b611f278b84888a89613219565b9350505097509795505050505050565b6000611f496040808051918201905290565b9050611f63611f5f611f5a846109c4565b61354d565b8252565b611a3c611f74610fcc846020610c05565b6020830152565b6000611f9182606001515183610140015161356d565b81516001600160a01b031660009081526001602052604090205461037d908390610a54565b8254600090610100900460ff1615611fdf578115611fd757611fd78561357d565b506000610503565b83546201000090046001600160781b0316801561203b57831561200a576120058661358f565b61203b565b8454600160881b90046001600160781b0316811061203b57821561203157612031866135a1565b6000915050610503565b50600195945050505050565b632165628a6000526004601cfd5b33831480156120645750505050565b600061206e6129ca565b61190160f01b60009081526002828152602287815260428320908390528651939450929190601f601d840116106102e260621984011016156120da576120b486886135b3565b61190160f01b6000908152600286905260228281526042822091905290975090506120dd565b50815b6120ea888285858a61364f565b5050505050505050565b60006121066040808051918201905290565b9050612117611f5f610ccb846109c4565b611a3c611f74610ccb846020610c05565b80518060051b6040019050602082038051602082527f4b9f2d36e1b4c93de62cc077b00b1a91d84b6c31b4a14e012718dcca230689e78383a190525050565b8351606090806001600160401b038111156121845761218461548d565b6040519080825280602002602001820160405280156121bd57816020015b6121aa614d3d565b8152602001906001900390816121a25790505b5091506000805b828110156122525760008882815181106121e0576121e06154a3565b6020026020010151905060006122008b83600001518460200151866137a7565b80518051602083015160809092015192935015159114161561222757836001019350612248565b80868585038151811061223c5761223c6154a3565b60200260200101819052505b50506001016121c4565b508015612260578083510383525b5061226e8783878787613219565b50505b95945050505050565b612287828261016061195d565b612298610a0e610a09846040610c05565b6118e06122a9610a25846060610c05565b6060830152565b6040518135601f0163ffffffe01660200180838337913563ffffffff16815290810160405290565b60006122eb602060408051918201905290565b60008152905090565b6122fc61194b565b600201600055565b600080600080856000015190506123248160a001518260c00151876139b8565b61233857506000925082915081905061252b565b602086015160408701516001600160781b039182169450169150600060048260800151600481111561236c5761236c615038565b036123a35760018385021890508015612387576123876139db565b612396828860800151886139e9565b945094509450505061252b565b5081831183151780156123b8576123b86139db565b60808201516001161584841116156123d2576123d2613c06565b6123db82611f7b565b60008181526002602052604081209196506123fa90879083908a611fb6565b61240e57506000935083925061252b915050565b805460ff1661242a5761242a8360000151878a60600151612055565b8080548060881c8061243e57879150612516565b6001600160781b038260101c1691506001870361246357908103965094508480612516565b80870361248157908701868103878211029788900397900390612516565b80870297810291909602810187810388821102918290039796919003906001600160781b038083119088111715612516576124cb565b60005b8215610a4a579082900691906124ba565b6124de6124d888846124b7565b896124b7565b80150197889004979687900496909104906001600160781b03808311908811171561251657634e487b7160005260116020526024601cfd5b508560881b8160101b17600117825550505050505b9250925092565b8051825160005b8281101561262a576000848281518110612555576125556154a3565b6020026020010151905060008160000151905083811061257c5761257c8260200151613c14565b6000878281518110612590576125906154a3565b6020026020010151905080602001516001600160781b03166000036125b757505050612622565b80516040808201519085015163bfb3f8ce6000876020015160018111156125e0576125e0615038565b146125fd5760006125f2856060613c26565b9350636088d7de9150505b8251821061260f57806000526004601cfd5b61261a838389613c32565b505050505050505b600101612539565b5060005b8181101561194357600085828151811061264a5761264a6154a3565b6020026020010151905080602001516001600160781b031660000361266f575061270c565b805160608101515160005b818110156126c4576126ad8360600151828151811061269b5761269b6154a3565b60200260200101516000015160031090565b156126bc576126bc8582613ce8565b60010161267a565b505060408101515160005b81811015612707576126f08360400151828151811061269b5761269b6154a3565b156126ff576126ff8582613cfe565b6001016126cf565b505050505b60010161262e565b60a085015160c08601516040805160208082528183019092526000916020820181803683375050506040890151519091506000805b828110156127ca5760008b604001518281518110612769576127696154a3565b6020026020010151905060008160000151905080158417935050600061279d826060015183608001518e8e8c8c6000613d14565b606083015250608081018890528b516101208d01516127c191839188613d506103b6565b50600101612749565b5060808a015160048110821680156127e4576127e4613167565b50505060608901515160009150815b818110156128905760008b606001518281518110612813576128136154a3565b602002602001015190506000612837826060015183608001518e8e8c8c6001613d14565b6060830181905260a08301516080840152905060008251600581111561285f5761285f615038565b03612877574794508481111561287757612877612df6565b61288682338c89613d506103b6565b50506001016127f3565b505061289b826118bb565b504780156128ad576128ad3382612e04565b505050505050505050565b82516080810151602082015160009283928392839291600481106001909111163390911415161561290b576128f386828a608001518a613e46565b60208301519194509250945063fb5014fc9350612958565b60048160800151600481111561292357612923615038565b036120ea578051945060008560601b905061294587838b608001518b85613f6d565b6393979285965090945092506129589050565b6120ea8587858588613038565b60608290506060829050856001600160a01b0316876001600160a01b03167f9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f318a8886866040516129b8949392919061578b565b60405180910390a35050505050505050565b60007f00000000000000000000000000000000000000000000000000000000000000004614612a84575060408051608080517f000000000000000000000000000000000000000000000000000000000000000060009081527f00000000000000000000000000000000000000000000000000000000000000006020527f0000000000000000000000000000000000000000000000000000000000000000855246606090815230845260a08220949095529093529190915290565b507f000000000000000000000000000000000000000000000000000000000000000090565b63fed398fc6000526004601cfd5b600435602014610224356102401416610244356102606102643560061b01141660186101243510600160a01b60843560a4351760c435602435171710161680612b0257612b0261403f565b50565b600083815260026020526040902060843590612b248582600180611fb6565b50805460ff16612b6f57612b6f828686868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061205592505050565b71010000000000000000000000000000010001905550505050565b604080517f000000000000000000000000000000000000000000000000000000000000000060ff60a01b17600090815260208690527f000000000000000000000000000000000000000000000000000000000000000083526055600b20919092526001600160a01b03169050600080600080526020600085876000875af19150600051905081612c2557612c1c61404d565b612c2583614094565b6001600160e01b03198116632671a55160e11b146116f3576116f386846140a6565b6369f95827600052806020526024601cfd5b833b612c7157635f15d672600052836020526024601cfd5b6040516323b872dd60e01b6000528360045282602452816044526000806064600080895af180612d07573d15612ce457601f3d0160051c8260051c8160030281831115612ccb578183036003028280028480020360091c01015b5a602082011015612ce0573d6000803e3d6000fd5b5050505b63f486bc8760005285602052846040528360605282608052600160a05260a4601cfd5b50604052505060006060525050565b843b612d2e57635f15d672600052846020526024601cfd5b60405160805160a05160c051637921219560e11b6000528760045286602452856044528460645260a0608452600060a45260008060c46000808d5af180612dda573d15612db857601f3d0160051c8560051c8160030281831115612d9f578183036003028280028480020360091c01015b5a602082011015612db4573d6000803e3d6000fd5b5050505b63f486bc87600052896020528860405287606052866080528560a05260a4601cfd5b5060809290925260a05260c05260405250506000606052505050565b638ffff9806000526004601cfd5b612e0d81612edd565b600080600080600085875af190508061194657612e2861404d565b63bc806b9660005282602052816040526044601cfd5b6000612e4b836020015190565b905081811461194657611946836118bb565b60006020885103612e985750604080885260208089018a9052632671a55160e11b918901919091526044880152600160648801819052612ea7565b50606487018051600101908190525b603c60c082028901038781528660208201528560408201528460608201528360808201528260a082015250505050505050505050565b80612b02576391b3e5146000526004601cfd5b636ab37ce76000526004601cfd5b612f0783612edd565b612f118183612e3e565b81612f22576116e1868686866140bc565b6116f38282600189898960008a612e5d565b6064810151604082019060c002604401612f4f848383612b8a565b50506020905250565b600080612f636141bb565b6317b1f942815260208082015260408101858152336060830152601c9091019250612f97610a0e6080860160608701615210565b612fa861014460e08301606061195d565b610140612fb6816060840152565b612fc460a082016080840152565b61026435602081026102000160a0820261016001612fe582858701836141c6565b929092019150612ff890508160a0840152565b613003600082840152565b6020016130118160c0840152565b61301c600182840152565b6130298682840160200152565b80606401925050509250929050565b6000806001600160e01b03198551166000805260206000868860008c5af16000519093501490508161307a5761306c61404d565b82600052856020526024601cfd5b806117be5782600052856020526024601cfd5b637fa8a9876000526004601cfd5b60008063ffffffff83351690506001810160051b6130bf8160408051918201905290565b92506130cc84848361195d565b5050919050565b60008284036130e35750806103cd565b82848309156130fa5763c63cf0896000526004601cfd5b60006131068584615825565b93909304949350505050565b600084861461315d57838303428590038082036000613131838a615825565b61313b838c615825565b613145919061573e565b90508584878303040181151502945050505050612271565b5092949350505050565b6312d3f5a36000526004601cfd5b61317d614d3d565b835160000361318f5761318f856141dd565b805160008660018111156131a5576131a5615038565b036131c9576001600160a01b03831660808201526131c48786846141ef565b6131e2565b6131d487868461435a565b336020830152604082018490525b8060600151600003613201576000602083018190526080820152600181525b5095945050505050565b63d5da9a1b6000526004601cfd5b84516060906000816001600160401b038111156132385761323861548d565b604051908082528060200260200182016040528015613261578160200160208202803683370190505b5060408051602080825281830190925291925060009190602082018180368337019050508851909150600090815b8181101561330a5760008b82815181106132ab576132ab6154a3565b602090810291909101015180519091506000815160058111156132d0576132d0615038565b036132ec5747945084816060015111156132ec576132ec612df6565b613300818360200151846040015189613d50565b505060010161328f565b50505060005b838110156134a35760008a828151811061332c5761332c6154a3565b6020026020010151905080602001516001600160781b031660000361337557600084838151811061335f5761335f6154a3565b911515602092830291909101909101525061349b565b6001848381518110613389576133896154a3565b9115156020928302919091019091015280516040810151805160005b818110156134325760008382815181106133c1576133c16154a3565b60200260200101519050806060015160001461341f5761341f613411828f6040805160a0810182526000808252602082018190529181018290526060810182905260809081019190915282015290565b86516101208801518b613d50565b60808101516060909101526001016133a5565b5050506060810151805160005b8181101561349557600083828151811061345b5761345b6154a3565b60200260200101519050600081606001519050806000146134815761348188848361447b565b5060a081015160609091015260010161343f565b50505050505b600101613310565b506134ad816118bb565b4780156134be576134be3382612e04565b85156135355760005b84811015613533578381815181106134e1576134e16154a3565b60200260200101511561352b5761352b8b8281518110613503576135036154a3565b60200260200101518a8b848151811061351e5761351e6154a3565b60200260200101516128b8565b6001016134c7565b505b61353f6001600055565b509098975050505050505050565b600061356161016060408051918201905290565b9050611a3c828261227a565b808210156118e0576118e0614495565b631a515574600052806020526024601cfd5b63ee9e0e63600052806020526024601cfd5b6310fda3e1600052806020526024601cfd5b600080600084516001811660410380820360051c9250808752806020018701915050805160e81c6003820191506001811660051b868152825160208218525060015b8381101561362257604060002082821c60051b6020908116918252938401805191909418526001016135f5565b505050604060002091506000613637826144a3565b60009081526020939093525050604090209392505050565b6000806000528151602083038051826041036000600182116136b8576040870151606088015160001a831561369757601b8260ff1c0190506001600160ff1b03821660408a01525b88528a85526020600060808760015afa508385528588526040880152506000515b8a148a151516945084905061378957858552604082526044850380516040870351630b135d3f60e11b83528960408903526020600060648b01858f5afa9650861561377d57630b135d3f60e11b6000511461377d578b3b1561372257634f7fb80d6000526004601cfd5b600186604103111561373c57638baa579f6000526004601cfd5b640101000000606089015160001a1a1560418714161561376f57631f003d0a600052606088015160001a6020526024601cfd5b63815e1d646000526004601cfd5b8385529152603f198601525b505050806116f35761379961404d565b634f7fb80d6000526004601cfd5b6137af614d3d565b835115806137bc57508251155b156137c9576137c96148b4565b6137d1614d3d565b6137dc86858361435a565b8051606081015160000361380757506000602082018190528151608001528051600190529050610503565b6138128787856141ef565b8060400151836000015160400151188160200151846000015160200151188260000151600581111561384657613846615038565b855151600581111561385a5761385a615038565b1860ff16176001600160a01b03161760001461387957613879846148c2565b8260000151606001518160600151111561390c576000856000815181106138a2576138a26154a3565b60200260200101519050836000015160600151826060015103888260000151815181106138d1576138d16154a3565b602002602001015160000151606001518260200151815181106138f6576138f66154a3565b6020026020010151606001818152505050613999565b600086600081518110613921576139216154a3565b6020026020010151905081606001518460000151606001510388826000015181518110613950576139506154a3565b60200260200101516000015160400151826020015181518110613975576139756154a3565b60200260200101516060018181525050816060015184600001516060018181525050505b60809081015183516001600160a01b0390911691015250949350505050565b4280841115908311168180156139cc575080155b156103cd576103cd84846148d4565b635a052b326000526004601cfd5b600080600085610140015186606001515114613a0757613a07612047565b855160008080613a178a8a6148ea565b9150915060008082846000885af16001600160a01b0385166000908152600360205260409020805460018101909155606086901b189750925082613a6e57613a5f888861499a565b96509650965050505050613bfd565b505050506000806000613a836103b66149bc90565b92509250925082600014613a9a57613a9a86614b00565b604089015151825180821115613ab357613ab388614b00565b60005b82811015613b34576000613ae48d604001518381518110613ad957613ad96154a3565b602002602001015190565b90506000613afd878481518110613ad957613ad96154a3565b9050613b098282614b12565b613b24613b17606084015190565b60608501511190565b1190565b1797909717965050600101613ab6565b50505060408901829052606089015181518151811115613b5757613b5788614b00565b60005b81811015613bd9576000613b79848381518110613ad957613ad96154a3565b90506000613b92868481518110613ad957613ad96154a3565b9050613bb0613ba260a084015190565b60a083015181159114171590565b613bba8383614b12565b613bc8613b17606086015190565b171797909717965050600101613b5a565b505050606089018190528215613bf257613bf286614b00565b506001935083925050505b93509350939050565b63a11b63ff6000526004601cfd5b63133c37c6600052806020526024601cfd5b60006103cd8284015190565b6000838381518110613c4657613c466154a3565b60200260200101519050600081600001519050613c638160031090565b613c6f57613c6f614b6f565b60408201518015613c9257613c8d8460600151828660800151614b7d565b613ca5565b60808401515115613ca557613ca5614bcb565b6004821460030383816005811115613cbf57613cbf615038565b90816005811115613cd257613cd2615038565b9052505050606090920151604090910152505050565b63a8930e9a60005281602052806040526044601cfd5b63d692933260005281602052806040526044601cfd5b6000868803613d2f57613d288686896130d3565b90506106cb565b6106c8613d3d87878b6130d3565b613d4888888b6130d3565b868686613112565b600084516005811115613d6557613d65615038565b03613da257604084015160208501516001600160a01b03161715613d8b57613d8b612ef0565b613d9d84608001518560600151612e04565b613e40565b600184516005811115613db757613db7615038565b03613de857604084015115613dce57613dce612ef0565b613d9d846020015184866080015187606001518686612efe565b600284516005811115613dfd57613dfd615038565b03613e2157613d9d846020015184866080015187604001518860600151878761177c565b613e4084602001518486608001518760400151886060015187876117c7565b50505050565b600080613e516141bb565b6317b1f942815260208082015260408101878152336060830152601c90910192508551604082015285613e8f613e8860a083015190565b60e0840152565b613ea5613e9d60c083015190565b610100840152565b613ebb613eb360e083015190565b610120840152565b610140613ec9816060850152565b6000613ed6604084015190565b90506000613ee682848701614bd9565b928301929050613ef7836080870152565b6000613f04606086015190565b90506000613f1482868901614c40565b948501949050613f258560a0890152565b6000613f338c878a01614cae565b958601959050613f448660c08a0152565b6000613f528c888b01614cce565b9a9f96909a016024019d50949b505050505050505050505050565b600080613f786141bb565b63f4dd92ce815287841860a0820152601c8101925060200160a0808252876000613fa3604083015190565b90506000613fb382858701614bd9565b938401939050613fc4846020870152565b6000613fd1606085015190565b90506000613fe182878901614c40565b958601959050613ff2866040890152565b60006140008d888a01614cae565b9687019690506140118760608a0152565b600061401f8d898b01614cce565b905080880197508760040199505050505050505050509550959350505050565b6339f3e3fd6000526004601cfd5b3d1561037057601f3d0160051c60405160051c816003028183111561407f578183036003028280028480020360091c01015b5a602082011015611946573d6000803e3d6000fd5b63d13d53d4600052806020526024601cfd5b631cf99b2660005281602052806040526044601cfd5b6040516323b872dd60e01b600052836004528260245281604452602060006064600080895af1803d15601f3d116001600051141617163d151581166141ab5780873b1515166141ab5780614199578161417b573d1561415857601f3d0160051c8360051c816003028183111561413f578183036003028280028480020360091c01015b5a602082011015614154573d6000803e3d6000fd5b5050505b63f486bc8760005286602052856040528460605260006080528360a05260a4601cfd5b6398891923600052866020528560405284606052836080526084601cfd5b635f15d672600052866020526024601cfd5b5050604052505060006060525050565b600061038d60405190565b8082828560045afa80153d151715613e4057600080fd5b63375c24c1600052806020526024601cfd5b600080600084855160051b86015b8082101561430b576020820191508151518851811061421e5761421e614339565b8060051b60208a0101519050805160208451015160006040830151602085015115815184101517156142545750505050506141fd565b8260051b60208201015191505060608101935083518901915083511589831060011b178817975081985060008452895193508615600181146142be5760608220881860408c01516101208601511860208d01518651181717156142b9576142b9614339565b614301565b815185526020820151602086015260408201516040860152835160208c015261012084015160408c015260608520975060208c0192508683181561430157865183525b50505050506141fd565b5050508160608451015280156119435760018103614331576391b3e5146000526004601cfd5b611943614347565b637fda72796000526004601cfd5b634e487b7160005260116020526024601cfd5b600080600084855160051b86015b80821015614454576020820191508151518851811061438957614389614339565b8060051b60208a01015190506020835101516000606083510151602084015115815184101517156143bd5750505050614368565b8260051b60208201015191505060608101925082518801915082511588831060011b178717965081975060008352885192508515600181146144105760a08220871461440b5761440b614339565b61444b565b8151845260208201516020850152604082015160408501526080820151608085015260a08220965060208b0192508583181561444b57855183525b50505050614368565b505083516060018390525080156119435760018103614331576391b3e5146000526004601cfd5b63a5f542086000528260205281604052806060526064601cfd5b63466aa6166000526004601cfd5b60006148ab565b60006009821015614604576005821015614561576003821015614514577f832c58a5b611aadcfa6a082ac9d04bace53d8278387f10040347b7e98eb5b30260018314027fbf8e29b89f29ed9b529c154a63038ffca562f8d7cd1e2545dda53a1b582dde301861037d565b7ff3e8417a785f980bdaf134fa0274a6bf891eeb8195cd94b09d2aa651046e28bc60038314027fa02eb7ff164c884e5e2c336dc85f81c6a93329d8e9adf214b32729b894de2af11861037d565b60078210156145b7577f25d02425402d882d211a7ab774c0ed6eca048c4d03d9af40132475744753b2a360058314027f1c19f71958cdd8f081b4c31f7caf5c010b29d12950be2fa1c95070dc47e30b551861037d565b7fb58d772fb09b426b9dece637f61ca9065f2b994f1464b51e9207f55f7c8f594860078314027f7ff98d9d4e55d876c5cfac10b43c04039522f3ddfb0ea9bfe70c68cfb5c7cc141861037d565b601182101561475c57600d8210156146b957600b82101561466c577f6f0ec38c21f6f583ab7f3c5413c773ffd5344c34fde1d390958e438bf667448f60098314027fd1d97d1ef5eaa37a4ee5fbf234e6f6d64eb511eb562221cd7edfbdde0848da051861037d565b7f32f4e7485d6485f9f6c255929b9905c62ba919758bbe231f231eaeecf33d810c600b8314027fbb98d87cc12922b83759626c5f07d72266da9702d19ffad6a514c73a89002f5f1861037d565b600f82101561470f577f8df51df98847160517f5b1186b4bc3f418d98b8a7f17f1292f392d79d600d79e600d8314027f6b5b04cbae4fcb1a9d78e7b2dfc51a36933d023cf6e347e03d517b472a8525901861037d565b7fcc4886e37eedd9aacd6c1c2c9247197a621a71282e87a7cbc673f3736d9aa141600f8314027f1da3eed3ecef6ebaa6e5023c057ec2c75150693fd0dac5c90f4a142f9879fde81861037d565b60158210156148085760138210156147bb577f2d7a3ed6dab270fdb8e054b2ad525f0ce2a8b89cc76c17f0965434740f673a5560118314027fc3939feff011e53ab8c35ca3370aad54c5df1fc2938cd62543174fa6e7d858771861037d565b7f54b3212a178782f104e0d514b41a9a5c4ca9c980bf6597c3cecbf280917e202a60138314027f5a4f867d3d458dabecad65f6201ceeaba0096df2d0c491cc32e6ea4e643500171861037d565b601782101561485e577fbb40bf8cea3a5a716e2b6eb08bbdac8ec159f82f380783db3c56904f15a43d0460158314027f3bd8cff538aba49a9c374c806d277181e9651624b3e31111bc0624574f8bca1d1861037d565b7f403be09941a31d05cfc2f896505811353d45d38743288b016630cce39435476a60178314027f1d51df90cba8de7637ca3e8fe1e3511d1dc2f23487d05dbdecb781860c21ac1c1861037d565b61037d826144aa565b6398e9db6e6000526004601cfd5b63bced929d600052806020526024601cfd5b6321ccfeb760005281602052806040526044601cfd5b600080836148f66141bb565b639891976581523360208201908152608060408301819052601c9092019450906000614923604085015190565b9050600061493382848601614bd9565b928301929050614944836040860152565b6000614951606087015190565b9050600061496182868801614bd9565b948501949050614972856060880152565b89600061498182898901614cae565b9a9d96909a016004019b50949950505050505050505050565b600080600084156149ae576149ae84614b00565b509193600093508392509050565b60803d106000808080808086614a1e5760406000803e6000516020513d81113d8311179850909450925086614a1e5760208460003e600051915060208360203e506020516000805261ffff828217113d608060a08402600786901b0101111796505b86614a4257614a308260208601614a4b565b9550614a3f8160208501614aaa565b94505b50505050909192565b6000604051905060c08302602001810160405282815260208360010160051b8083015b81831015614aa1578083850152608085823e6060810151608082015260808501945060a081019050602083019250614a6e565b50505092915050565b6000604051905060e08302602001810160405282815260208360010160051b8083015b81831015614aa1578083850152608085823e604060608601608083013e60a094909401936020929092019160c001614acd565b6393979285600052806020526024601cfd5b6000825160408401518015600383111615614b3857600482146003039150604084015190505b6060850151606085015150608086015181146040860151831416855184146020870151602089015114161615935050505092915050565b6394eb6af66000526004601cfd5b600083600052602060002060208301835160051b81015b80821015614bbc57815180841160051b93845260209384185260406000209290910190614b94565b50508314905080613e4057613e405b6309bde3396000526004601cfd5b6000825180835260208401602084018260051b82015b80831015614c2d5782518051835260208101516020840152604081015160408401526060810151606084015250602083019250608082019150614bef565b5050508060071b60200191505092915050565b600080614c4b845190565b8084529050602084810190600583901b860181019085015b614c6d8383613b20565b15614c9b576000614c7d84614cff565b9050614c8b818360a06141c6565b506020929092019160a001614c63565b60a0840260200194505050505092915050565b600063ffffffe0603f614cbf855190565b0116905061037d8383836141c6565b600080614cd9845190565b8084529050600581901b614cf46020868101908601836141c6565b602001949350505050565b600061037d825190565b6040518060a00160405280614d1c614d80565b81526000602082018190526040820152606080820181905260809091015290565b60408051610100810182526000606082018181526080830182905260a0830182905260c0830182905260e083018290528252602082018190529181019190915290565b60405180610160016040528060006001600160a01b0316815260200160006001600160a01b03168152602001606081526020016060815260200160006004811115614dcd57614dcd615038565b815260006020820181905260408201819052606082018190526080820181905260a0820181905260c09091015290565b600060208284031215614e0f57600080fd5b81356001600160401b03811115614e2557600080fd5b820161024081850312156103cd57600080fd5b6000815180845260005b81811015614e5e57602081850181015186830182015201614e42565b506000602082860101526020601f19601f83011685010191505092915050565b6020815260006103cd6020830184614e38565b600060208284031215614ea357600080fd5b5035919050565b600060208284031215614ebc57600080fd5b81356001600160401b03811115614ed257600080fd5b820161016081850312156103cd57600080fd5b60008083601f840112614ef757600080fd5b5081356001600160401b03811115614f0e57600080fd5b6020830191508360208260051b8501011115614f2957600080fd5b9250929050565b6001600160a01b0381168114612b0257600080fd5b8035611a3c81614f30565b600080600080600080600080600080600060e08c8e031215614f7157600080fd5b6001600160401b03808d351115614f8757600080fd5b614f948e8e358f01614ee5565b909c509a5060208d0135811015614faa57600080fd5b614fba8e60208f01358f01614ee5565b909a50985060408d0135811015614fd057600080fd5b614fe08e60408f01358f01614ee5565b909850965060608d0135811015614ff657600080fd5b506150078d60608e01358e01614ee5565b909550935060808c0135925061501f60a08d01614f45565b915060c08c013590509295989b509295989b9093969950565b634e487b7160e01b600052602160045260246000fd5b600681106119b6576119b6615038565b61506982825161504e565b6020818101516001600160a01b0390811691840191909152604080830151908401526060808301519084015260809182015116910152565b600081518084526020808501945080840160005b838110156150f75781516150ca88825161505e565b808401516001600160a01b031660a08901526040015160c088015260e090960195908201906001016150b5565b509495945050505050565b604080825283519082018190526000906020906060840190828701845b8281101561513d57815115158452928401929084019060010161511f565b505050838103828501526105be81866150a1565b6000806020838503121561516457600080fd5b82356001600160401b0381111561517a57600080fd5b61518685828601614ee5565b90969095509350505050565b600080600080604085870312156151a857600080fd5b84356001600160401b03808211156151bf57600080fd5b6151cb88838901614ee5565b909650945060208701359150808211156151e457600080fd5b506151f187828801614ee5565b95989497509550505050565b6020815260006103cd60208301846150a1565b60006020828403121561522257600080fd5b81356103cd81614f30565b6000806040838503121561524057600080fd5b82356001600160401b0381111561525657600080fd5b83016040818603121561526857600080fd5b946020939093013593505050565b60008060008060006080868803121561528e57600080fd5b85356001600160401b03808211156152a557600080fd5b9087019060a0828a0312156152b957600080fd5b909550602087013590808211156152cf57600080fd5b506152dc88828901614ee5565b9095509350506040860135915060608601356152f781614f30565b809150509295509295909350565b60008060008060008060008060a0898b03121561532157600080fd5b88356001600160401b038082111561533857600080fd5b6153448c838d01614ee5565b909a50985060208b013591508082111561535d57600080fd5b6153698c838d01614ee5565b909850965060408b013591508082111561538257600080fd5b5061538f8b828c01614ee5565b999c989b509699959896976060870135966080013595509350505050565b60008060008060008060006080888a0312156153c857600080fd5b87356001600160401b03808211156153df57600080fd5b6153eb8b838c01614ee5565b909950975060208a013591508082111561540457600080fd5b6154108b838c01614ee5565b909750955060408a013591508082111561542957600080fd5b506154368a828b01614ee5565b909450925050606088013561544a81614f30565b8091505092959891949750929550565b60608152600061546d6060830186614e38565b6020830194909452506001600160a01b0391909116604090910152919050565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b600081518084526020808501945080840160005b838110156150f75781516154e288825161504e565b838101516001600160a01b03168885015260408082015190890152606080820151908901526080908101519088015260a090960195908201906001016154cd565b600081518084526020808501945080840160005b838110156150f757815161554c88825161504e565b808401516001600160a01b0390811689860152604080830151908a0152606080830151908a0152608080830151908a015260a091820151169088015260c09096019590820190600101615537565b600581106119b6576119b6615038565b828152604060208201526155ca6040820183516001600160a01b03169052565b600060208301516155e660608401826001600160a01b03169052565b5060408301516101608060808501526156036101a08501836154b9565b91506060850151603f198584030160a08601526156208382615523565b925050608085015161563560c086018261559a565b5060a085015160e085015260c0850151610100818187015260e0870151915061012082818801528188015192506101409150828288015280880151848801525080870151610180870152505050809150509392505050565b6000823561015e198336030181126156a457600080fd5b9190910192915050565b6000602082840312156156c057600080fd5b8135600581106103cd57600080fd5b6000808335601e198436030181126156e657600080fd5b8301803591506001600160401b0382111561570057600080fd5b602001915036819003821315614f2957600080fd5b634e487b7160e01b600052601160045260246000fd5b8181038181111561037d5761037d615715565b8082018082111561037d5761037d615715565b600081518084526020808501945080840160005b838110156150f75761577887835161505e565b60a0969096019590820190600101615765565b60006080808301878452602060018060a01b03808916828701526040848188015283895180865260a089019150848b01955060005b818110156158015786516157d584825161504e565b8087015186168488015284810151858501526060908101519084015295850195918701916001016157c0565b50508781036060890152615815818a615751565b9c9b505050505050505050505050565b808202811582820484141761037d5761037d61571556fea26469706673582212201b8f503dc30f3f6d05c20f2e1b9a982bb074fecd1cc7cbf6dc0b7adecc58853564736f6c63430008110033436f6e73696465726174696f6e4974656d2875696e7438206974656d547970652c6164647265737320746f6b656e2c75696e74323536206964656e7469666965724f7243726974657269612c75696e74323536207374617274416d6f756e742c75696e7432353620656e64416d6f756e742c6164647265737320726563697069656e7429454950373132446f6d61696e28737472696e67206e616d652c737472696e672076657273696f6e2c75696e7432353620636861696e49642c6164647265737320766572696679696e67436f6e7472616374294f72646572436f6d706f6e656e74732861646472657373206f6666657265722c61646472657373207a6f6e652c4f666665724974656d5b5d206f666665722c436f6e73696465726174696f6e4974656d5b5d20636f6e73696465726174696f6e2c75696e7438206f72646572547970652c75696e7432353620737461727454696d652c75696e7432353620656e6454696d652c62797465733332207a6f6e65486173682c75696e743235362073616c742c6279746573333220636f6e647569744b65792c75696e7432353620636f756e746572294f666665724974656d2875696e7438206974656d547970652c6164647265737320746f6b656e2c75696e74323536206964656e7469666965724f7243726974657269612c75696e74323536207374617274416d6f756e742c75696e7432353620656e64416d6f756e7429",
}

// SeaportABI is the input ABI used to generate the binding from.
// Deprecated: Use SeaportMetaData.ABI instead.
var SeaportABI = SeaportMetaData.ABI

// SeaportBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SeaportMetaData.Bin instead.
var SeaportBin = SeaportMetaData.Bin

// DeploySeaport deploys a new Ethereum contract, binding an instance of Seaport to it.
func DeploySeaport(auth *bind.TransactOpts, backend bind.ContractBackend, conduitController common.Address) (common.Address, *types.Transaction, *Seaport, error) {
	parsed, err := SeaportMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SeaportBin), backend, conduitController)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Seaport{SeaportCaller: SeaportCaller{contract: contract}, SeaportTransactor: SeaportTransactor{contract: contract}, SeaportFilterer: SeaportFilterer{contract: contract}}, nil
}

// Seaport is an auto generated Go binding around an Ethereum contract.
type Seaport struct {
	SeaportCaller     // Read-only binding to the contract
	SeaportTransactor // Write-only binding to the contract
	SeaportFilterer   // Log filterer for contract events
}

// SeaportCaller is an auto generated read-only Go binding around an Ethereum contract.
type SeaportCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeaportTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SeaportTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeaportFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SeaportFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeaportSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SeaportSession struct {
	Contract     *Seaport          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SeaportCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SeaportCallerSession struct {
	Contract *SeaportCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SeaportTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SeaportTransactorSession struct {
	Contract     *SeaportTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SeaportRaw is an auto generated low-level Go binding around an Ethereum contract.
type SeaportRaw struct {
	Contract *Seaport // Generic contract binding to access the raw methods on
}

// SeaportCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SeaportCallerRaw struct {
	Contract *SeaportCaller // Generic read-only contract binding to access the raw methods on
}

// SeaportTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SeaportTransactorRaw struct {
	Contract *SeaportTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSeaport creates a new instance of Seaport, bound to a specific deployed contract.
func NewSeaport(address common.Address, backend bind.ContractBackend) (*Seaport, error) {
	contract, err := bindSeaport(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Seaport{SeaportCaller: SeaportCaller{contract: contract}, SeaportTransactor: SeaportTransactor{contract: contract}, SeaportFilterer: SeaportFilterer{contract: contract}}, nil
}

// NewSeaportCaller creates a new read-only instance of Seaport, bound to a specific deployed contract.
func NewSeaportCaller(address common.Address, caller bind.ContractCaller) (*SeaportCaller, error) {
	contract, err := bindSeaport(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SeaportCaller{contract: contract}, nil
}

// NewSeaportTransactor creates a new write-only instance of Seaport, bound to a specific deployed contract.
func NewSeaportTransactor(address common.Address, transactor bind.ContractTransactor) (*SeaportTransactor, error) {
	contract, err := bindSeaport(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SeaportTransactor{contract: contract}, nil
}

// NewSeaportFilterer creates a new log filterer instance of Seaport, bound to a specific deployed contract.
func NewSeaportFilterer(address common.Address, filterer bind.ContractFilterer) (*SeaportFilterer, error) {
	contract, err := bindSeaport(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SeaportFilterer{contract: contract}, nil
}

// bindSeaport binds a generic wrapper to an already deployed contract.
func bindSeaport(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SeaportABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Seaport *SeaportRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Seaport.Contract.SeaportCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Seaport *SeaportRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Seaport.Contract.SeaportTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Seaport *SeaportRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Seaport.Contract.SeaportTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Seaport *SeaportCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Seaport.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Seaport *SeaportTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Seaport.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Seaport *SeaportTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Seaport.Contract.contract.Transact(opts, method, params...)
}

// GetContractOffererNonce is a free data retrieval call binding the contract method 0xa900866b.
//
// Solidity: function getContractOffererNonce(address contractOfferer) view returns(uint256 nonce)
func (_Seaport *SeaportCaller) GetContractOffererNonce(opts *bind.CallOpts, contractOfferer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Seaport.contract.Call(opts, &out, "getContractOffererNonce", contractOfferer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContractOffererNonce is a free data retrieval call binding the contract method 0xa900866b.
//
// Solidity: function getContractOffererNonce(address contractOfferer) view returns(uint256 nonce)
func (_Seaport *SeaportSession) GetContractOffererNonce(contractOfferer common.Address) (*big.Int, error) {
	return _Seaport.Contract.GetContractOffererNonce(&_Seaport.CallOpts, contractOfferer)
}

// GetContractOffererNonce is a free data retrieval call binding the contract method 0xa900866b.
//
// Solidity: function getContractOffererNonce(address contractOfferer) view returns(uint256 nonce)
func (_Seaport *SeaportCallerSession) GetContractOffererNonce(contractOfferer common.Address) (*big.Int, error) {
	return _Seaport.Contract.GetContractOffererNonce(&_Seaport.CallOpts, contractOfferer)
}

// GetCounter is a free data retrieval call binding the contract method 0xf07ec373.
//
// Solidity: function getCounter(address offerer) view returns(uint256 counter)
func (_Seaport *SeaportCaller) GetCounter(opts *bind.CallOpts, offerer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Seaport.contract.Call(opts, &out, "getCounter", offerer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCounter is a free data retrieval call binding the contract method 0xf07ec373.
//
// Solidity: function getCounter(address offerer) view returns(uint256 counter)
func (_Seaport *SeaportSession) GetCounter(offerer common.Address) (*big.Int, error) {
	return _Seaport.Contract.GetCounter(&_Seaport.CallOpts, offerer)
}

// GetCounter is a free data retrieval call binding the contract method 0xf07ec373.
//
// Solidity: function getCounter(address offerer) view returns(uint256 counter)
func (_Seaport *SeaportCallerSession) GetCounter(offerer common.Address) (*big.Int, error) {
	return _Seaport.Contract.GetCounter(&_Seaport.CallOpts, offerer)
}

// GetOrderHash is a free data retrieval call binding the contract method 0x79df72bd.
//
// Solidity: function getOrderHash((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256) ) view returns(bytes32 orderHash)
func (_Seaport *SeaportCaller) GetOrderHash(opts *bind.CallOpts, arg0 OrderComponents) ([32]byte, error) {
	var out []interface{}
	err := _Seaport.contract.Call(opts, &out, "getOrderHash", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOrderHash is a free data retrieval call binding the contract method 0x79df72bd.
//
// Solidity: function getOrderHash((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256) ) view returns(bytes32 orderHash)
func (_Seaport *SeaportSession) GetOrderHash(arg0 OrderComponents) ([32]byte, error) {
	return _Seaport.Contract.GetOrderHash(&_Seaport.CallOpts, arg0)
}

// GetOrderHash is a free data retrieval call binding the contract method 0x79df72bd.
//
// Solidity: function getOrderHash((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256) ) view returns(bytes32 orderHash)
func (_Seaport *SeaportCallerSession) GetOrderHash(arg0 OrderComponents) ([32]byte, error) {
	return _Seaport.Contract.GetOrderHash(&_Seaport.CallOpts, arg0)
}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns(bool isValidated, bool isCancelled, uint256 totalFilled, uint256 totalSize)
func (_Seaport *SeaportCaller) GetOrderStatus(opts *bind.CallOpts, orderHash [32]byte) (struct {
	IsValidated bool
	IsCancelled bool
	TotalFilled *big.Int
	TotalSize   *big.Int
}, error) {
	var out []interface{}
	err := _Seaport.contract.Call(opts, &out, "getOrderStatus", orderHash)

	outstruct := new(struct {
		IsValidated bool
		IsCancelled bool
		TotalFilled *big.Int
		TotalSize   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsValidated = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.IsCancelled = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.TotalFilled = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalSize = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns(bool isValidated, bool isCancelled, uint256 totalFilled, uint256 totalSize)
func (_Seaport *SeaportSession) GetOrderStatus(orderHash [32]byte) (struct {
	IsValidated bool
	IsCancelled bool
	TotalFilled *big.Int
	TotalSize   *big.Int
}, error) {
	return _Seaport.Contract.GetOrderStatus(&_Seaport.CallOpts, orderHash)
}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns(bool isValidated, bool isCancelled, uint256 totalFilled, uint256 totalSize)
func (_Seaport *SeaportCallerSession) GetOrderStatus(orderHash [32]byte) (struct {
	IsValidated bool
	IsCancelled bool
	TotalFilled *big.Int
	TotalSize   *big.Int
}, error) {
	return _Seaport.Contract.GetOrderStatus(&_Seaport.CallOpts, orderHash)
}

// Information is a free data retrieval call binding the contract method 0xf47b7740.
//
// Solidity: function information() view returns(string version, bytes32 domainSeparator, address conduitController)
func (_Seaport *SeaportCaller) Information(opts *bind.CallOpts) (struct {
	Version           string
	DomainSeparator   [32]byte
	ConduitController common.Address
}, error) {
	var out []interface{}
	err := _Seaport.contract.Call(opts, &out, "information")

	outstruct := new(struct {
		Version           string
		DomainSeparator   [32]byte
		ConduitController common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Version = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.DomainSeparator = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.ConduitController = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Information is a free data retrieval call binding the contract method 0xf47b7740.
//
// Solidity: function information() view returns(string version, bytes32 domainSeparator, address conduitController)
func (_Seaport *SeaportSession) Information() (struct {
	Version           string
	DomainSeparator   [32]byte
	ConduitController common.Address
}, error) {
	return _Seaport.Contract.Information(&_Seaport.CallOpts)
}

// Information is a free data retrieval call binding the contract method 0xf47b7740.
//
// Solidity: function information() view returns(string version, bytes32 domainSeparator, address conduitController)
func (_Seaport *SeaportCallerSession) Information() (struct {
	Version           string
	DomainSeparator   [32]byte
	ConduitController common.Address
}, error) {
	return _Seaport.Contract.Information(&_Seaport.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Seaport *SeaportCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Seaport.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Seaport *SeaportSession) Name() (string, error) {
	return _Seaport.Contract.Name(&_Seaport.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Seaport *SeaportCallerSession) Name() (string, error) {
	return _Seaport.Contract.Name(&_Seaport.CallOpts)
}

// Cancel is a paid mutator transaction binding the contract method 0xfd9f1e10.
//
// Solidity: function cancel((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256)[] orders) returns(bool cancelled)
func (_Seaport *SeaportTransactor) Cancel(opts *bind.TransactOpts, orders []OrderComponents) (*types.Transaction, error) {
	return _Seaport.contract.Transact(opts, "cancel", orders)
}

// Cancel is a paid mutator transaction binding the contract method 0xfd9f1e10.
//
// Solidity: function cancel((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256)[] orders) returns(bool cancelled)
func (_Seaport *SeaportSession) Cancel(orders []OrderComponents) (*types.Transaction, error) {
	return _Seaport.Contract.Cancel(&_Seaport.TransactOpts, orders)
}

// Cancel is a paid mutator transaction binding the contract method 0xfd9f1e10.
//
// Solidity: function cancel((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256)[] orders) returns(bool cancelled)
func (_Seaport *SeaportTransactorSession) Cancel(orders []OrderComponents) (*types.Transaction, error) {
	return _Seaport.Contract.Cancel(&_Seaport.TransactOpts, orders)
}

// FulfillAdvancedOrder is a paid mutator transaction binding the contract method 0xe7acab24.
//
// Solidity: function fulfillAdvancedOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes) , (uint256,uint8,uint256,uint256,bytes32[])[] , bytes32 fulfillerConduitKey, address recipient) payable returns(bool fulfilled)
func (_Seaport *SeaportTransactor) FulfillAdvancedOrder(opts *bind.TransactOpts, arg0 AdvancedOrder, arg1 []CriteriaResolver, fulfillerConduitKey [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _Seaport.contract.Transact(opts, "fulfillAdvancedOrder", arg0, arg1, fulfillerConduitKey, recipient)
}

// FulfillAdvancedOrder is a paid mutator transaction binding the contract method 0xe7acab24.
//
// Solidity: function fulfillAdvancedOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes) , (uint256,uint8,uint256,uint256,bytes32[])[] , bytes32 fulfillerConduitKey, address recipient) payable returns(bool fulfilled)
func (_Seaport *SeaportSession) FulfillAdvancedOrder(arg0 AdvancedOrder, arg1 []CriteriaResolver, fulfillerConduitKey [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _Seaport.Contract.FulfillAdvancedOrder(&_Seaport.TransactOpts, arg0, arg1, fulfillerConduitKey, recipient)
}

// FulfillAdvancedOrder is a paid mutator transaction binding the contract method 0xe7acab24.
//
// Solidity: function fulfillAdvancedOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes) , (uint256,uint8,uint256,uint256,bytes32[])[] , bytes32 fulfillerConduitKey, address recipient) payable returns(bool fulfilled)
func (_Seaport *SeaportTransactorSession) FulfillAdvancedOrder(arg0 AdvancedOrder, arg1 []CriteriaResolver, fulfillerConduitKey [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _Seaport.Contract.FulfillAdvancedOrder(&_Seaport.TransactOpts, arg0, arg1, fulfillerConduitKey, recipient)
}

// FulfillAvailableAdvancedOrders is a paid mutator transaction binding the contract method 0x87201b41.
//
// Solidity: function fulfillAvailableAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] , (uint256,uint8,uint256,uint256,bytes32[])[] , (uint256,uint256)[][] , (uint256,uint256)[][] , bytes32 fulfillerConduitKey, address recipient, uint256 maximumFulfilled) payable returns(bool[], ((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_Seaport *SeaportTransactor) FulfillAvailableAdvancedOrders(opts *bind.TransactOpts, arg0 []AdvancedOrder, arg1 []CriteriaResolver, arg2 [][]FulfillmentComponent, arg3 [][]FulfillmentComponent, fulfillerConduitKey [32]byte, recipient common.Address, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _Seaport.contract.Transact(opts, "fulfillAvailableAdvancedOrders", arg0, arg1, arg2, arg3, fulfillerConduitKey, recipient, maximumFulfilled)
}

// FulfillAvailableAdvancedOrders is a paid mutator transaction binding the contract method 0x87201b41.
//
// Solidity: function fulfillAvailableAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] , (uint256,uint8,uint256,uint256,bytes32[])[] , (uint256,uint256)[][] , (uint256,uint256)[][] , bytes32 fulfillerConduitKey, address recipient, uint256 maximumFulfilled) payable returns(bool[], ((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_Seaport *SeaportSession) FulfillAvailableAdvancedOrders(arg0 []AdvancedOrder, arg1 []CriteriaResolver, arg2 [][]FulfillmentComponent, arg3 [][]FulfillmentComponent, fulfillerConduitKey [32]byte, recipient common.Address, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _Seaport.Contract.FulfillAvailableAdvancedOrders(&_Seaport.TransactOpts, arg0, arg1, arg2, arg3, fulfillerConduitKey, recipient, maximumFulfilled)
}

// FulfillAvailableAdvancedOrders is a paid mutator transaction binding the contract method 0x87201b41.
//
// Solidity: function fulfillAvailableAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] , (uint256,uint8,uint256,uint256,bytes32[])[] , (uint256,uint256)[][] , (uint256,uint256)[][] , bytes32 fulfillerConduitKey, address recipient, uint256 maximumFulfilled) payable returns(bool[], ((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_Seaport *SeaportTransactorSession) FulfillAvailableAdvancedOrders(arg0 []AdvancedOrder, arg1 []CriteriaResolver, arg2 [][]FulfillmentComponent, arg3 [][]FulfillmentComponent, fulfillerConduitKey [32]byte, recipient common.Address, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _Seaport.Contract.FulfillAvailableAdvancedOrders(&_Seaport.TransactOpts, arg0, arg1, arg2, arg3, fulfillerConduitKey, recipient, maximumFulfilled)
}

// FulfillAvailableOrders is a paid mutator transaction binding the contract method 0xed98a574.
//
// Solidity: function fulfillAvailableOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] , (uint256,uint256)[][] , (uint256,uint256)[][] , bytes32 fulfillerConduitKey, uint256 maximumFulfilled) payable returns(bool[], ((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_Seaport *SeaportTransactor) FulfillAvailableOrders(opts *bind.TransactOpts, arg0 []Order, arg1 [][]FulfillmentComponent, arg2 [][]FulfillmentComponent, fulfillerConduitKey [32]byte, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _Seaport.contract.Transact(opts, "fulfillAvailableOrders", arg0, arg1, arg2, fulfillerConduitKey, maximumFulfilled)
}

// FulfillAvailableOrders is a paid mutator transaction binding the contract method 0xed98a574.
//
// Solidity: function fulfillAvailableOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] , (uint256,uint256)[][] , (uint256,uint256)[][] , bytes32 fulfillerConduitKey, uint256 maximumFulfilled) payable returns(bool[], ((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_Seaport *SeaportSession) FulfillAvailableOrders(arg0 []Order, arg1 [][]FulfillmentComponent, arg2 [][]FulfillmentComponent, fulfillerConduitKey [32]byte, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _Seaport.Contract.FulfillAvailableOrders(&_Seaport.TransactOpts, arg0, arg1, arg2, fulfillerConduitKey, maximumFulfilled)
}

// FulfillAvailableOrders is a paid mutator transaction binding the contract method 0xed98a574.
//
// Solidity: function fulfillAvailableOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] , (uint256,uint256)[][] , (uint256,uint256)[][] , bytes32 fulfillerConduitKey, uint256 maximumFulfilled) payable returns(bool[], ((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_Seaport *SeaportTransactorSession) FulfillAvailableOrders(arg0 []Order, arg1 [][]FulfillmentComponent, arg2 [][]FulfillmentComponent, fulfillerConduitKey [32]byte, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _Seaport.Contract.FulfillAvailableOrders(&_Seaport.TransactOpts, arg0, arg1, arg2, fulfillerConduitKey, maximumFulfilled)
}

// FulfillBasicOrder is a paid mutator transaction binding the contract method 0xfb0f3ee1.
//
// Solidity: function fulfillBasicOrder((address,uint256,uint256,address,address,address,uint256,uint256,uint8,uint256,uint256,bytes32,uint256,bytes32,bytes32,uint256,(uint256,address)[],bytes) parameters) payable returns(bool fulfilled)
func (_Seaport *SeaportTransactor) FulfillBasicOrder(opts *bind.TransactOpts, parameters BasicOrderParameters) (*types.Transaction, error) {
	return _Seaport.contract.Transact(opts, "fulfillBasicOrder", parameters)
}

// FulfillBasicOrder is a paid mutator transaction binding the contract method 0xfb0f3ee1.
//
// Solidity: function fulfillBasicOrder((address,uint256,uint256,address,address,address,uint256,uint256,uint8,uint256,uint256,bytes32,uint256,bytes32,bytes32,uint256,(uint256,address)[],bytes) parameters) payable returns(bool fulfilled)
func (_Seaport *SeaportSession) FulfillBasicOrder(parameters BasicOrderParameters) (*types.Transaction, error) {
	return _Seaport.Contract.FulfillBasicOrder(&_Seaport.TransactOpts, parameters)
}

// FulfillBasicOrder is a paid mutator transaction binding the contract method 0xfb0f3ee1.
//
// Solidity: function fulfillBasicOrder((address,uint256,uint256,address,address,address,uint256,uint256,uint8,uint256,uint256,bytes32,uint256,bytes32,bytes32,uint256,(uint256,address)[],bytes) parameters) payable returns(bool fulfilled)
func (_Seaport *SeaportTransactorSession) FulfillBasicOrder(parameters BasicOrderParameters) (*types.Transaction, error) {
	return _Seaport.Contract.FulfillBasicOrder(&_Seaport.TransactOpts, parameters)
}

// FulfillBasicOrderEfficient6GL6yc is a paid mutator transaction binding the contract method 0x00000000.
//
// Solidity: function fulfillBasicOrder_efficient_6GL6yc((address,uint256,uint256,address,address,address,uint256,uint256,uint8,uint256,uint256,bytes32,uint256,bytes32,bytes32,uint256,(uint256,address)[],bytes) parameters) payable returns(bool fulfilled)
func (_Seaport *SeaportTransactor) FulfillBasicOrderEfficient6GL6yc(opts *bind.TransactOpts, parameters BasicOrderParameters) (*types.Transaction, error) {
	return _Seaport.contract.Transact(opts, "fulfillBasicOrder_efficient_6GL6yc", parameters)
}

// FulfillBasicOrderEfficient6GL6yc is a paid mutator transaction binding the contract method 0x00000000.
//
// Solidity: function fulfillBasicOrder_efficient_6GL6yc((address,uint256,uint256,address,address,address,uint256,uint256,uint8,uint256,uint256,bytes32,uint256,bytes32,bytes32,uint256,(uint256,address)[],bytes) parameters) payable returns(bool fulfilled)
func (_Seaport *SeaportSession) FulfillBasicOrderEfficient6GL6yc(parameters BasicOrderParameters) (*types.Transaction, error) {
	return _Seaport.Contract.FulfillBasicOrderEfficient6GL6yc(&_Seaport.TransactOpts, parameters)
}

// FulfillBasicOrderEfficient6GL6yc is a paid mutator transaction binding the contract method 0x00000000.
//
// Solidity: function fulfillBasicOrder_efficient_6GL6yc((address,uint256,uint256,address,address,address,uint256,uint256,uint8,uint256,uint256,bytes32,uint256,bytes32,bytes32,uint256,(uint256,address)[],bytes) parameters) payable returns(bool fulfilled)
func (_Seaport *SeaportTransactorSession) FulfillBasicOrderEfficient6GL6yc(parameters BasicOrderParameters) (*types.Transaction, error) {
	return _Seaport.Contract.FulfillBasicOrderEfficient6GL6yc(&_Seaport.TransactOpts, parameters)
}

// FulfillOrder is a paid mutator transaction binding the contract method 0xb3a34c4c.
//
// Solidity: function fulfillOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes) , bytes32 fulfillerConduitKey) payable returns(bool fulfilled)
func (_Seaport *SeaportTransactor) FulfillOrder(opts *bind.TransactOpts, arg0 Order, fulfillerConduitKey [32]byte) (*types.Transaction, error) {
	return _Seaport.contract.Transact(opts, "fulfillOrder", arg0, fulfillerConduitKey)
}

// FulfillOrder is a paid mutator transaction binding the contract method 0xb3a34c4c.
//
// Solidity: function fulfillOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes) , bytes32 fulfillerConduitKey) payable returns(bool fulfilled)
func (_Seaport *SeaportSession) FulfillOrder(arg0 Order, fulfillerConduitKey [32]byte) (*types.Transaction, error) {
	return _Seaport.Contract.FulfillOrder(&_Seaport.TransactOpts, arg0, fulfillerConduitKey)
}

// FulfillOrder is a paid mutator transaction binding the contract method 0xb3a34c4c.
//
// Solidity: function fulfillOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes) , bytes32 fulfillerConduitKey) payable returns(bool fulfilled)
func (_Seaport *SeaportTransactorSession) FulfillOrder(arg0 Order, fulfillerConduitKey [32]byte) (*types.Transaction, error) {
	return _Seaport.Contract.FulfillOrder(&_Seaport.TransactOpts, arg0, fulfillerConduitKey)
}

// IncrementCounter is a paid mutator transaction binding the contract method 0x5b34b966.
//
// Solidity: function incrementCounter() returns(uint256 newCounter)
func (_Seaport *SeaportTransactor) IncrementCounter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Seaport.contract.Transact(opts, "incrementCounter")
}

// IncrementCounter is a paid mutator transaction binding the contract method 0x5b34b966.
//
// Solidity: function incrementCounter() returns(uint256 newCounter)
func (_Seaport *SeaportSession) IncrementCounter() (*types.Transaction, error) {
	return _Seaport.Contract.IncrementCounter(&_Seaport.TransactOpts)
}

// IncrementCounter is a paid mutator transaction binding the contract method 0x5b34b966.
//
// Solidity: function incrementCounter() returns(uint256 newCounter)
func (_Seaport *SeaportTransactorSession) IncrementCounter() (*types.Transaction, error) {
	return _Seaport.Contract.IncrementCounter(&_Seaport.TransactOpts)
}

// MatchAdvancedOrders is a paid mutator transaction binding the contract method 0xf2d12b12.
//
// Solidity: function matchAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] , (uint256,uint8,uint256,uint256,bytes32[])[] , ((uint256,uint256)[],(uint256,uint256)[])[] , address recipient) payable returns(((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_Seaport *SeaportTransactor) MatchAdvancedOrders(opts *bind.TransactOpts, arg0 []AdvancedOrder, arg1 []CriteriaResolver, arg2 []Fulfillment, recipient common.Address) (*types.Transaction, error) {
	return _Seaport.contract.Transact(opts, "matchAdvancedOrders", arg0, arg1, arg2, recipient)
}

// MatchAdvancedOrders is a paid mutator transaction binding the contract method 0xf2d12b12.
//
// Solidity: function matchAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] , (uint256,uint8,uint256,uint256,bytes32[])[] , ((uint256,uint256)[],(uint256,uint256)[])[] , address recipient) payable returns(((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_Seaport *SeaportSession) MatchAdvancedOrders(arg0 []AdvancedOrder, arg1 []CriteriaResolver, arg2 []Fulfillment, recipient common.Address) (*types.Transaction, error) {
	return _Seaport.Contract.MatchAdvancedOrders(&_Seaport.TransactOpts, arg0, arg1, arg2, recipient)
}

// MatchAdvancedOrders is a paid mutator transaction binding the contract method 0xf2d12b12.
//
// Solidity: function matchAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] , (uint256,uint8,uint256,uint256,bytes32[])[] , ((uint256,uint256)[],(uint256,uint256)[])[] , address recipient) payable returns(((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_Seaport *SeaportTransactorSession) MatchAdvancedOrders(arg0 []AdvancedOrder, arg1 []CriteriaResolver, arg2 []Fulfillment, recipient common.Address) (*types.Transaction, error) {
	return _Seaport.Contract.MatchAdvancedOrders(&_Seaport.TransactOpts, arg0, arg1, arg2, recipient)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xa8174404.
//
// Solidity: function matchOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] , ((uint256,uint256)[],(uint256,uint256)[])[] ) payable returns(((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_Seaport *SeaportTransactor) MatchOrders(opts *bind.TransactOpts, arg0 []Order, arg1 []Fulfillment) (*types.Transaction, error) {
	return _Seaport.contract.Transact(opts, "matchOrders", arg0, arg1)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xa8174404.
//
// Solidity: function matchOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] , ((uint256,uint256)[],(uint256,uint256)[])[] ) payable returns(((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_Seaport *SeaportSession) MatchOrders(arg0 []Order, arg1 []Fulfillment) (*types.Transaction, error) {
	return _Seaport.Contract.MatchOrders(&_Seaport.TransactOpts, arg0, arg1)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xa8174404.
//
// Solidity: function matchOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] , ((uint256,uint256)[],(uint256,uint256)[])[] ) payable returns(((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_Seaport *SeaportTransactorSession) MatchOrders(arg0 []Order, arg1 []Fulfillment) (*types.Transaction, error) {
	return _Seaport.Contract.MatchOrders(&_Seaport.TransactOpts, arg0, arg1)
}

// Validate is a paid mutator transaction binding the contract method 0x88147732.
//
// Solidity: function validate(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] ) returns(bool)
func (_Seaport *SeaportTransactor) Validate(opts *bind.TransactOpts, arg0 []Order) (*types.Transaction, error) {
	return _Seaport.contract.Transact(opts, "validate", arg0)
}

// Validate is a paid mutator transaction binding the contract method 0x88147732.
//
// Solidity: function validate(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] ) returns(bool)
func (_Seaport *SeaportSession) Validate(arg0 []Order) (*types.Transaction, error) {
	return _Seaport.Contract.Validate(&_Seaport.TransactOpts, arg0)
}

// Validate is a paid mutator transaction binding the contract method 0x88147732.
//
// Solidity: function validate(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] ) returns(bool)
func (_Seaport *SeaportTransactorSession) Validate(arg0 []Order) (*types.Transaction, error) {
	return _Seaport.Contract.Validate(&_Seaport.TransactOpts, arg0)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Seaport *SeaportTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Seaport.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Seaport *SeaportSession) Receive() (*types.Transaction, error) {
	return _Seaport.Contract.Receive(&_Seaport.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Seaport *SeaportTransactorSession) Receive() (*types.Transaction, error) {
	return _Seaport.Contract.Receive(&_Seaport.TransactOpts)
}

// SeaportCounterIncrementedIterator is returned from FilterCounterIncremented and is used to iterate over the raw logs and unpacked data for CounterIncremented events raised by the Seaport contract.
type SeaportCounterIncrementedIterator struct {
	Event *SeaportCounterIncremented // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SeaportCounterIncrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeaportCounterIncremented)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SeaportCounterIncremented)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SeaportCounterIncrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeaportCounterIncrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeaportCounterIncremented represents a CounterIncremented event raised by the Seaport contract.
type SeaportCounterIncremented struct {
	NewCounter *big.Int
	Offerer    common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCounterIncremented is a free log retrieval operation binding the contract event 0x721c20121297512b72821b97f5326877ea8ecf4bb9948fea5bfcb6453074d37f.
//
// Solidity: event CounterIncremented(uint256 newCounter, address indexed offerer)
func (_Seaport *SeaportFilterer) FilterCounterIncremented(opts *bind.FilterOpts, offerer []common.Address) (*SeaportCounterIncrementedIterator, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}

	logs, sub, err := _Seaport.contract.FilterLogs(opts, "CounterIncremented", offererRule)
	if err != nil {
		return nil, err
	}
	return &SeaportCounterIncrementedIterator{contract: _Seaport.contract, event: "CounterIncremented", logs: logs, sub: sub}, nil
}

// WatchCounterIncremented is a free log subscription operation binding the contract event 0x721c20121297512b72821b97f5326877ea8ecf4bb9948fea5bfcb6453074d37f.
//
// Solidity: event CounterIncremented(uint256 newCounter, address indexed offerer)
func (_Seaport *SeaportFilterer) WatchCounterIncremented(opts *bind.WatchOpts, sink chan<- *SeaportCounterIncremented, offerer []common.Address) (event.Subscription, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}

	logs, sub, err := _Seaport.contract.WatchLogs(opts, "CounterIncremented", offererRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeaportCounterIncremented)
				if err := _Seaport.contract.UnpackLog(event, "CounterIncremented", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCounterIncremented is a log parse operation binding the contract event 0x721c20121297512b72821b97f5326877ea8ecf4bb9948fea5bfcb6453074d37f.
//
// Solidity: event CounterIncremented(uint256 newCounter, address indexed offerer)
func (_Seaport *SeaportFilterer) ParseCounterIncremented(log types.Log) (*SeaportCounterIncremented, error) {
	event := new(SeaportCounterIncremented)
	if err := _Seaport.contract.UnpackLog(event, "CounterIncremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SeaportOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the Seaport contract.
type SeaportOrderCancelledIterator struct {
	Event *SeaportOrderCancelled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SeaportOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeaportOrderCancelled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SeaportOrderCancelled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SeaportOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeaportOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeaportOrderCancelled represents a OrderCancelled event raised by the Seaport contract.
type SeaportOrderCancelled struct {
	OrderHash [32]byte
	Offerer   common.Address
	Zone      common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0x6bacc01dbe442496068f7d234edd811f1a5f833243e0aec824f86ab861f3c90d.
//
// Solidity: event OrderCancelled(bytes32 orderHash, address indexed offerer, address indexed zone)
func (_Seaport *SeaportFilterer) FilterOrderCancelled(opts *bind.FilterOpts, offerer []common.Address, zone []common.Address) (*SeaportOrderCancelledIterator, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _Seaport.contract.FilterLogs(opts, "OrderCancelled", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return &SeaportOrderCancelledIterator{contract: _Seaport.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0x6bacc01dbe442496068f7d234edd811f1a5f833243e0aec824f86ab861f3c90d.
//
// Solidity: event OrderCancelled(bytes32 orderHash, address indexed offerer, address indexed zone)
func (_Seaport *SeaportFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *SeaportOrderCancelled, offerer []common.Address, zone []common.Address) (event.Subscription, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _Seaport.contract.WatchLogs(opts, "OrderCancelled", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeaportOrderCancelled)
				if err := _Seaport.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOrderCancelled is a log parse operation binding the contract event 0x6bacc01dbe442496068f7d234edd811f1a5f833243e0aec824f86ab861f3c90d.
//
// Solidity: event OrderCancelled(bytes32 orderHash, address indexed offerer, address indexed zone)
func (_Seaport *SeaportFilterer) ParseOrderCancelled(log types.Log) (*SeaportOrderCancelled, error) {
	event := new(SeaportOrderCancelled)
	if err := _Seaport.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SeaportOrderFulfilledIterator is returned from FilterOrderFulfilled and is used to iterate over the raw logs and unpacked data for OrderFulfilled events raised by the Seaport contract.
type SeaportOrderFulfilledIterator struct {
	Event *SeaportOrderFulfilled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SeaportOrderFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeaportOrderFulfilled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SeaportOrderFulfilled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SeaportOrderFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeaportOrderFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeaportOrderFulfilled represents a OrderFulfilled event raised by the Seaport contract.
type SeaportOrderFulfilled struct {
	OrderHash     [32]byte
	Offerer       common.Address
	Zone          common.Address
	Recipient     common.Address
	Offer         []SpentItem
	Consideration []ReceivedItem
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOrderFulfilled is a free log retrieval operation binding the contract event 0x9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f31.
//
// Solidity: event OrderFulfilled(bytes32 orderHash, address indexed offerer, address indexed zone, address recipient, (uint8,address,uint256,uint256)[] offer, (uint8,address,uint256,uint256,address)[] consideration)
func (_Seaport *SeaportFilterer) FilterOrderFulfilled(opts *bind.FilterOpts, offerer []common.Address, zone []common.Address) (*SeaportOrderFulfilledIterator, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _Seaport.contract.FilterLogs(opts, "OrderFulfilled", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return &SeaportOrderFulfilledIterator{contract: _Seaport.contract, event: "OrderFulfilled", logs: logs, sub: sub}, nil
}

// WatchOrderFulfilled is a free log subscription operation binding the contract event 0x9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f31.
//
// Solidity: event OrderFulfilled(bytes32 orderHash, address indexed offerer, address indexed zone, address recipient, (uint8,address,uint256,uint256)[] offer, (uint8,address,uint256,uint256,address)[] consideration)
func (_Seaport *SeaportFilterer) WatchOrderFulfilled(opts *bind.WatchOpts, sink chan<- *SeaportOrderFulfilled, offerer []common.Address, zone []common.Address) (event.Subscription, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _Seaport.contract.WatchLogs(opts, "OrderFulfilled", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeaportOrderFulfilled)
				if err := _Seaport.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOrderFulfilled is a log parse operation binding the contract event 0x9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f31.
//
// Solidity: event OrderFulfilled(bytes32 orderHash, address indexed offerer, address indexed zone, address recipient, (uint8,address,uint256,uint256)[] offer, (uint8,address,uint256,uint256,address)[] consideration)
func (_Seaport *SeaportFilterer) ParseOrderFulfilled(log types.Log) (*SeaportOrderFulfilled, error) {
	event := new(SeaportOrderFulfilled)
	if err := _Seaport.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SeaportOrderValidatedIterator is returned from FilterOrderValidated and is used to iterate over the raw logs and unpacked data for OrderValidated events raised by the Seaport contract.
type SeaportOrderValidatedIterator struct {
	Event *SeaportOrderValidated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SeaportOrderValidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeaportOrderValidated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SeaportOrderValidated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SeaportOrderValidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeaportOrderValidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeaportOrderValidated represents a OrderValidated event raised by the Seaport contract.
type SeaportOrderValidated struct {
	OrderHash       [32]byte
	OrderParameters OrderParameters
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOrderValidated is a free log retrieval operation binding the contract event 0xf280791efe782edcf06ce15c8f4dff17601db3b88eb3805a0db7d77faf757f04.
//
// Solidity: event OrderValidated(bytes32 orderHash, (address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256) orderParameters)
func (_Seaport *SeaportFilterer) FilterOrderValidated(opts *bind.FilterOpts) (*SeaportOrderValidatedIterator, error) {

	logs, sub, err := _Seaport.contract.FilterLogs(opts, "OrderValidated")
	if err != nil {
		return nil, err
	}
	return &SeaportOrderValidatedIterator{contract: _Seaport.contract, event: "OrderValidated", logs: logs, sub: sub}, nil
}

// WatchOrderValidated is a free log subscription operation binding the contract event 0xf280791efe782edcf06ce15c8f4dff17601db3b88eb3805a0db7d77faf757f04.
//
// Solidity: event OrderValidated(bytes32 orderHash, (address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256) orderParameters)
func (_Seaport *SeaportFilterer) WatchOrderValidated(opts *bind.WatchOpts, sink chan<- *SeaportOrderValidated) (event.Subscription, error) {

	logs, sub, err := _Seaport.contract.WatchLogs(opts, "OrderValidated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeaportOrderValidated)
				if err := _Seaport.contract.UnpackLog(event, "OrderValidated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOrderValidated is a log parse operation binding the contract event 0xf280791efe782edcf06ce15c8f4dff17601db3b88eb3805a0db7d77faf757f04.
//
// Solidity: event OrderValidated(bytes32 orderHash, (address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256) orderParameters)
func (_Seaport *SeaportFilterer) ParseOrderValidated(log types.Log) (*SeaportOrderValidated, error) {
	event := new(SeaportOrderValidated)
	if err := _Seaport.contract.UnpackLog(event, "OrderValidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SeaportOrdersMatchedIterator is returned from FilterOrdersMatched and is used to iterate over the raw logs and unpacked data for OrdersMatched events raised by the Seaport contract.
type SeaportOrdersMatchedIterator struct {
	Event *SeaportOrdersMatched // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SeaportOrdersMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeaportOrdersMatched)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SeaportOrdersMatched)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SeaportOrdersMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeaportOrdersMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeaportOrdersMatched represents a OrdersMatched event raised by the Seaport contract.
type SeaportOrdersMatched struct {
	OrderHashes [][32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOrdersMatched is a free log retrieval operation binding the contract event 0x4b9f2d36e1b4c93de62cc077b00b1a91d84b6c31b4a14e012718dcca230689e7.
//
// Solidity: event OrdersMatched(bytes32[] orderHashes)
func (_Seaport *SeaportFilterer) FilterOrdersMatched(opts *bind.FilterOpts) (*SeaportOrdersMatchedIterator, error) {

	logs, sub, err := _Seaport.contract.FilterLogs(opts, "OrdersMatched")
	if err != nil {
		return nil, err
	}
	return &SeaportOrdersMatchedIterator{contract: _Seaport.contract, event: "OrdersMatched", logs: logs, sub: sub}, nil
}

// WatchOrdersMatched is a free log subscription operation binding the contract event 0x4b9f2d36e1b4c93de62cc077b00b1a91d84b6c31b4a14e012718dcca230689e7.
//
// Solidity: event OrdersMatched(bytes32[] orderHashes)
func (_Seaport *SeaportFilterer) WatchOrdersMatched(opts *bind.WatchOpts, sink chan<- *SeaportOrdersMatched) (event.Subscription, error) {

	logs, sub, err := _Seaport.contract.WatchLogs(opts, "OrdersMatched")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeaportOrdersMatched)
				if err := _Seaport.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOrdersMatched is a log parse operation binding the contract event 0x4b9f2d36e1b4c93de62cc077b00b1a91d84b6c31b4a14e012718dcca230689e7.
//
// Solidity: event OrdersMatched(bytes32[] orderHashes)
func (_Seaport *SeaportFilterer) ParseOrdersMatched(log types.Log) (*SeaportOrdersMatched, error) {
	event := new(SeaportOrdersMatched)
	if err := _Seaport.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
