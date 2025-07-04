// Code generated by ent, DO NOT EDIT.

package generated

import (
	"github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/account"
	"github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/contract"
	"github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/deposit"
	"github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/goodbenefit"
	"github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/orderbenefit"
	"github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/payment"
	"github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/platform"
	"github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/transfer"
	"github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/user"
	"github.com/NpoolPlatform/kunman/middleware/account/db/ent/schema"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accountMixin := schema.Account{}.Mixin()
	accountMixinFields0 := accountMixin[0].Fields()
	_ = accountMixinFields0
	accountMixinFields1 := accountMixin[1].Fields()
	_ = accountMixinFields1
	accountFields := schema.Account{}.Fields()
	_ = accountFields
	// accountDescCreatedAt is the schema descriptor for created_at field.
	accountDescCreatedAt := accountMixinFields0[0].Descriptor()
	// account.DefaultCreatedAt holds the default value on creation for the created_at field.
	account.DefaultCreatedAt = accountDescCreatedAt.Default.(func() uint32)
	// accountDescUpdatedAt is the schema descriptor for updated_at field.
	accountDescUpdatedAt := accountMixinFields0[1].Descriptor()
	// account.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	account.DefaultUpdatedAt = accountDescUpdatedAt.Default.(func() uint32)
	// account.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	account.UpdateDefaultUpdatedAt = accountDescUpdatedAt.UpdateDefault.(func() uint32)
	// accountDescDeletedAt is the schema descriptor for deleted_at field.
	accountDescDeletedAt := accountMixinFields0[2].Descriptor()
	// account.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	account.DefaultDeletedAt = accountDescDeletedAt.Default.(func() uint32)
	// accountDescEntID is the schema descriptor for ent_id field.
	accountDescEntID := accountMixinFields1[1].Descriptor()
	// account.DefaultEntID holds the default value on creation for the ent_id field.
	account.DefaultEntID = accountDescEntID.Default.(func() uuid.UUID)
	// accountDescCoinTypeID is the schema descriptor for coin_type_id field.
	accountDescCoinTypeID := accountFields[0].Descriptor()
	// account.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	account.DefaultCoinTypeID = accountDescCoinTypeID.Default.(func() uuid.UUID)
	// accountDescAddress is the schema descriptor for address field.
	accountDescAddress := accountFields[1].Descriptor()
	// account.DefaultAddress holds the default value on creation for the address field.
	account.DefaultAddress = accountDescAddress.Default.(string)
	// accountDescUsedFor is the schema descriptor for used_for field.
	accountDescUsedFor := accountFields[2].Descriptor()
	// account.DefaultUsedFor holds the default value on creation for the used_for field.
	account.DefaultUsedFor = accountDescUsedFor.Default.(string)
	// accountDescPlatformHoldPrivateKey is the schema descriptor for platform_hold_private_key field.
	accountDescPlatformHoldPrivateKey := accountFields[3].Descriptor()
	// account.DefaultPlatformHoldPrivateKey holds the default value on creation for the platform_hold_private_key field.
	account.DefaultPlatformHoldPrivateKey = accountDescPlatformHoldPrivateKey.Default.(bool)
	// accountDescActive is the schema descriptor for active field.
	accountDescActive := accountFields[4].Descriptor()
	// account.DefaultActive holds the default value on creation for the active field.
	account.DefaultActive = accountDescActive.Default.(bool)
	// accountDescLocked is the schema descriptor for locked field.
	accountDescLocked := accountFields[5].Descriptor()
	// account.DefaultLocked holds the default value on creation for the locked field.
	account.DefaultLocked = accountDescLocked.Default.(bool)
	// accountDescLockedBy is the schema descriptor for locked_by field.
	accountDescLockedBy := accountFields[6].Descriptor()
	// account.DefaultLockedBy holds the default value on creation for the locked_by field.
	account.DefaultLockedBy = accountDescLockedBy.Default.(string)
	// accountDescBlocked is the schema descriptor for blocked field.
	accountDescBlocked := accountFields[7].Descriptor()
	// account.DefaultBlocked holds the default value on creation for the blocked field.
	account.DefaultBlocked = accountDescBlocked.Default.(bool)
	contractMixin := schema.Contract{}.Mixin()
	contractMixinFields0 := contractMixin[0].Fields()
	_ = contractMixinFields0
	contractMixinFields1 := contractMixin[1].Fields()
	_ = contractMixinFields1
	contractFields := schema.Contract{}.Fields()
	_ = contractFields
	// contractDescCreatedAt is the schema descriptor for created_at field.
	contractDescCreatedAt := contractMixinFields0[0].Descriptor()
	// contract.DefaultCreatedAt holds the default value on creation for the created_at field.
	contract.DefaultCreatedAt = contractDescCreatedAt.Default.(func() uint32)
	// contractDescUpdatedAt is the schema descriptor for updated_at field.
	contractDescUpdatedAt := contractMixinFields0[1].Descriptor()
	// contract.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	contract.DefaultUpdatedAt = contractDescUpdatedAt.Default.(func() uint32)
	// contract.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	contract.UpdateDefaultUpdatedAt = contractDescUpdatedAt.UpdateDefault.(func() uint32)
	// contractDescDeletedAt is the schema descriptor for deleted_at field.
	contractDescDeletedAt := contractMixinFields0[2].Descriptor()
	// contract.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	contract.DefaultDeletedAt = contractDescDeletedAt.Default.(func() uint32)
	// contractDescEntID is the schema descriptor for ent_id field.
	contractDescEntID := contractMixinFields1[1].Descriptor()
	// contract.DefaultEntID holds the default value on creation for the ent_id field.
	contract.DefaultEntID = contractDescEntID.Default.(func() uuid.UUID)
	// contractDescGoodID is the schema descriptor for good_id field.
	contractDescGoodID := contractFields[0].Descriptor()
	// contract.DefaultGoodID holds the default value on creation for the good_id field.
	contract.DefaultGoodID = contractDescGoodID.Default.(func() uuid.UUID)
	// contractDescDelegatedStakingID is the schema descriptor for delegated_staking_id field.
	contractDescDelegatedStakingID := contractFields[1].Descriptor()
	// contract.DefaultDelegatedStakingID holds the default value on creation for the delegated_staking_id field.
	contract.DefaultDelegatedStakingID = contractDescDelegatedStakingID.Default.(func() uuid.UUID)
	// contractDescAccountID is the schema descriptor for account_id field.
	contractDescAccountID := contractFields[2].Descriptor()
	// contract.DefaultAccountID holds the default value on creation for the account_id field.
	contract.DefaultAccountID = contractDescAccountID.Default.(func() uuid.UUID)
	// contractDescBackup is the schema descriptor for backup field.
	contractDescBackup := contractFields[3].Descriptor()
	// contract.DefaultBackup holds the default value on creation for the backup field.
	contract.DefaultBackup = contractDescBackup.Default.(bool)
	// contractDescContractOperatorType is the schema descriptor for contract_operator_type field.
	contractDescContractOperatorType := contractFields[4].Descriptor()
	// contract.DefaultContractOperatorType holds the default value on creation for the contract_operator_type field.
	contract.DefaultContractOperatorType = contractDescContractOperatorType.Default.(string)
	depositMixin := schema.Deposit{}.Mixin()
	depositMixinFields0 := depositMixin[0].Fields()
	_ = depositMixinFields0
	depositMixinFields1 := depositMixin[1].Fields()
	_ = depositMixinFields1
	depositFields := schema.Deposit{}.Fields()
	_ = depositFields
	// depositDescCreatedAt is the schema descriptor for created_at field.
	depositDescCreatedAt := depositMixinFields0[0].Descriptor()
	// deposit.DefaultCreatedAt holds the default value on creation for the created_at field.
	deposit.DefaultCreatedAt = depositDescCreatedAt.Default.(func() uint32)
	// depositDescUpdatedAt is the schema descriptor for updated_at field.
	depositDescUpdatedAt := depositMixinFields0[1].Descriptor()
	// deposit.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	deposit.DefaultUpdatedAt = depositDescUpdatedAt.Default.(func() uint32)
	// deposit.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	deposit.UpdateDefaultUpdatedAt = depositDescUpdatedAt.UpdateDefault.(func() uint32)
	// depositDescDeletedAt is the schema descriptor for deleted_at field.
	depositDescDeletedAt := depositMixinFields0[2].Descriptor()
	// deposit.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	deposit.DefaultDeletedAt = depositDescDeletedAt.Default.(func() uint32)
	// depositDescEntID is the schema descriptor for ent_id field.
	depositDescEntID := depositMixinFields1[1].Descriptor()
	// deposit.DefaultEntID holds the default value on creation for the ent_id field.
	deposit.DefaultEntID = depositDescEntID.Default.(func() uuid.UUID)
	// depositDescAppID is the schema descriptor for app_id field.
	depositDescAppID := depositFields[0].Descriptor()
	// deposit.DefaultAppID holds the default value on creation for the app_id field.
	deposit.DefaultAppID = depositDescAppID.Default.(func() uuid.UUID)
	// depositDescUserID is the schema descriptor for user_id field.
	depositDescUserID := depositFields[1].Descriptor()
	// deposit.DefaultUserID holds the default value on creation for the user_id field.
	deposit.DefaultUserID = depositDescUserID.Default.(func() uuid.UUID)
	// depositDescAccountID is the schema descriptor for account_id field.
	depositDescAccountID := depositFields[2].Descriptor()
	// deposit.DefaultAccountID holds the default value on creation for the account_id field.
	deposit.DefaultAccountID = depositDescAccountID.Default.(func() uuid.UUID)
	// depositDescIncoming is the schema descriptor for incoming field.
	depositDescIncoming := depositFields[3].Descriptor()
	// deposit.DefaultIncoming holds the default value on creation for the incoming field.
	deposit.DefaultIncoming = depositDescIncoming.Default.(decimal.Decimal)
	// depositDescOutcoming is the schema descriptor for outcoming field.
	depositDescOutcoming := depositFields[4].Descriptor()
	// deposit.DefaultOutcoming holds the default value on creation for the outcoming field.
	deposit.DefaultOutcoming = depositDescOutcoming.Default.(decimal.Decimal)
	// depositDescCollectingTid is the schema descriptor for collecting_tid field.
	depositDescCollectingTid := depositFields[5].Descriptor()
	// deposit.DefaultCollectingTid holds the default value on creation for the collecting_tid field.
	deposit.DefaultCollectingTid = depositDescCollectingTid.Default.(func() uuid.UUID)
	// depositDescScannableAt is the schema descriptor for scannable_at field.
	depositDescScannableAt := depositFields[6].Descriptor()
	// deposit.DefaultScannableAt holds the default value on creation for the scannable_at field.
	deposit.DefaultScannableAt = depositDescScannableAt.Default.(func() uint32)
	goodbenefitMixin := schema.GoodBenefit{}.Mixin()
	goodbenefitMixinFields0 := goodbenefitMixin[0].Fields()
	_ = goodbenefitMixinFields0
	goodbenefitMixinFields1 := goodbenefitMixin[1].Fields()
	_ = goodbenefitMixinFields1
	goodbenefitFields := schema.GoodBenefit{}.Fields()
	_ = goodbenefitFields
	// goodbenefitDescCreatedAt is the schema descriptor for created_at field.
	goodbenefitDescCreatedAt := goodbenefitMixinFields0[0].Descriptor()
	// goodbenefit.DefaultCreatedAt holds the default value on creation for the created_at field.
	goodbenefit.DefaultCreatedAt = goodbenefitDescCreatedAt.Default.(func() uint32)
	// goodbenefitDescUpdatedAt is the schema descriptor for updated_at field.
	goodbenefitDescUpdatedAt := goodbenefitMixinFields0[1].Descriptor()
	// goodbenefit.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	goodbenefit.DefaultUpdatedAt = goodbenefitDescUpdatedAt.Default.(func() uint32)
	// goodbenefit.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	goodbenefit.UpdateDefaultUpdatedAt = goodbenefitDescUpdatedAt.UpdateDefault.(func() uint32)
	// goodbenefitDescDeletedAt is the schema descriptor for deleted_at field.
	goodbenefitDescDeletedAt := goodbenefitMixinFields0[2].Descriptor()
	// goodbenefit.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	goodbenefit.DefaultDeletedAt = goodbenefitDescDeletedAt.Default.(func() uint32)
	// goodbenefitDescEntID is the schema descriptor for ent_id field.
	goodbenefitDescEntID := goodbenefitMixinFields1[1].Descriptor()
	// goodbenefit.DefaultEntID holds the default value on creation for the ent_id field.
	goodbenefit.DefaultEntID = goodbenefitDescEntID.Default.(func() uuid.UUID)
	// goodbenefitDescGoodID is the schema descriptor for good_id field.
	goodbenefitDescGoodID := goodbenefitFields[0].Descriptor()
	// goodbenefit.DefaultGoodID holds the default value on creation for the good_id field.
	goodbenefit.DefaultGoodID = goodbenefitDescGoodID.Default.(func() uuid.UUID)
	// goodbenefitDescAccountID is the schema descriptor for account_id field.
	goodbenefitDescAccountID := goodbenefitFields[1].Descriptor()
	// goodbenefit.DefaultAccountID holds the default value on creation for the account_id field.
	goodbenefit.DefaultAccountID = goodbenefitDescAccountID.Default.(func() uuid.UUID)
	// goodbenefitDescBackup is the schema descriptor for backup field.
	goodbenefitDescBackup := goodbenefitFields[2].Descriptor()
	// goodbenefit.DefaultBackup holds the default value on creation for the backup field.
	goodbenefit.DefaultBackup = goodbenefitDescBackup.Default.(bool)
	// goodbenefitDescTransactionID is the schema descriptor for transaction_id field.
	goodbenefitDescTransactionID := goodbenefitFields[3].Descriptor()
	// goodbenefit.DefaultTransactionID holds the default value on creation for the transaction_id field.
	goodbenefit.DefaultTransactionID = goodbenefitDescTransactionID.Default.(func() uuid.UUID)
	orderbenefitMixin := schema.OrderBenefit{}.Mixin()
	orderbenefitMixinFields0 := orderbenefitMixin[0].Fields()
	_ = orderbenefitMixinFields0
	orderbenefitMixinFields1 := orderbenefitMixin[1].Fields()
	_ = orderbenefitMixinFields1
	orderbenefitFields := schema.OrderBenefit{}.Fields()
	_ = orderbenefitFields
	// orderbenefitDescCreatedAt is the schema descriptor for created_at field.
	orderbenefitDescCreatedAt := orderbenefitMixinFields0[0].Descriptor()
	// orderbenefit.DefaultCreatedAt holds the default value on creation for the created_at field.
	orderbenefit.DefaultCreatedAt = orderbenefitDescCreatedAt.Default.(func() uint32)
	// orderbenefitDescUpdatedAt is the schema descriptor for updated_at field.
	orderbenefitDescUpdatedAt := orderbenefitMixinFields0[1].Descriptor()
	// orderbenefit.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	orderbenefit.DefaultUpdatedAt = orderbenefitDescUpdatedAt.Default.(func() uint32)
	// orderbenefit.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	orderbenefit.UpdateDefaultUpdatedAt = orderbenefitDescUpdatedAt.UpdateDefault.(func() uint32)
	// orderbenefitDescDeletedAt is the schema descriptor for deleted_at field.
	orderbenefitDescDeletedAt := orderbenefitMixinFields0[2].Descriptor()
	// orderbenefit.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	orderbenefit.DefaultDeletedAt = orderbenefitDescDeletedAt.Default.(func() uint32)
	// orderbenefitDescEntID is the schema descriptor for ent_id field.
	orderbenefitDescEntID := orderbenefitMixinFields1[1].Descriptor()
	// orderbenefit.DefaultEntID holds the default value on creation for the ent_id field.
	orderbenefit.DefaultEntID = orderbenefitDescEntID.Default.(func() uuid.UUID)
	// orderbenefitDescAppID is the schema descriptor for app_id field.
	orderbenefitDescAppID := orderbenefitFields[0].Descriptor()
	// orderbenefit.DefaultAppID holds the default value on creation for the app_id field.
	orderbenefit.DefaultAppID = orderbenefitDescAppID.Default.(func() uuid.UUID)
	// orderbenefitDescUserID is the schema descriptor for user_id field.
	orderbenefitDescUserID := orderbenefitFields[1].Descriptor()
	// orderbenefit.DefaultUserID holds the default value on creation for the user_id field.
	orderbenefit.DefaultUserID = orderbenefitDescUserID.Default.(func() uuid.UUID)
	// orderbenefitDescCoinTypeID is the schema descriptor for coin_type_id field.
	orderbenefitDescCoinTypeID := orderbenefitFields[2].Descriptor()
	// orderbenefit.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	orderbenefit.DefaultCoinTypeID = orderbenefitDescCoinTypeID.Default.(func() uuid.UUID)
	// orderbenefitDescAccountID is the schema descriptor for account_id field.
	orderbenefitDescAccountID := orderbenefitFields[3].Descriptor()
	// orderbenefit.DefaultAccountID holds the default value on creation for the account_id field.
	orderbenefit.DefaultAccountID = orderbenefitDescAccountID.Default.(func() uuid.UUID)
	// orderbenefitDescOrderID is the schema descriptor for order_id field.
	orderbenefitDescOrderID := orderbenefitFields[4].Descriptor()
	// orderbenefit.DefaultOrderID holds the default value on creation for the order_id field.
	orderbenefit.DefaultOrderID = orderbenefitDescOrderID.Default.(func() uuid.UUID)
	paymentMixin := schema.Payment{}.Mixin()
	paymentMixinFields0 := paymentMixin[0].Fields()
	_ = paymentMixinFields0
	paymentMixinFields1 := paymentMixin[1].Fields()
	_ = paymentMixinFields1
	paymentFields := schema.Payment{}.Fields()
	_ = paymentFields
	// paymentDescCreatedAt is the schema descriptor for created_at field.
	paymentDescCreatedAt := paymentMixinFields0[0].Descriptor()
	// payment.DefaultCreatedAt holds the default value on creation for the created_at field.
	payment.DefaultCreatedAt = paymentDescCreatedAt.Default.(func() uint32)
	// paymentDescUpdatedAt is the schema descriptor for updated_at field.
	paymentDescUpdatedAt := paymentMixinFields0[1].Descriptor()
	// payment.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	payment.DefaultUpdatedAt = paymentDescUpdatedAt.Default.(func() uint32)
	// payment.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	payment.UpdateDefaultUpdatedAt = paymentDescUpdatedAt.UpdateDefault.(func() uint32)
	// paymentDescDeletedAt is the schema descriptor for deleted_at field.
	paymentDescDeletedAt := paymentMixinFields0[2].Descriptor()
	// payment.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	payment.DefaultDeletedAt = paymentDescDeletedAt.Default.(func() uint32)
	// paymentDescEntID is the schema descriptor for ent_id field.
	paymentDescEntID := paymentMixinFields1[1].Descriptor()
	// payment.DefaultEntID holds the default value on creation for the ent_id field.
	payment.DefaultEntID = paymentDescEntID.Default.(func() uuid.UUID)
	// paymentDescAccountID is the schema descriptor for account_id field.
	paymentDescAccountID := paymentFields[0].Descriptor()
	// payment.DefaultAccountID holds the default value on creation for the account_id field.
	payment.DefaultAccountID = paymentDescAccountID.Default.(func() uuid.UUID)
	// paymentDescCollectingTid is the schema descriptor for collecting_tid field.
	paymentDescCollectingTid := paymentFields[1].Descriptor()
	// payment.DefaultCollectingTid holds the default value on creation for the collecting_tid field.
	payment.DefaultCollectingTid = paymentDescCollectingTid.Default.(func() uuid.UUID)
	// paymentDescAvailableAt is the schema descriptor for available_at field.
	paymentDescAvailableAt := paymentFields[2].Descriptor()
	// payment.DefaultAvailableAt holds the default value on creation for the available_at field.
	payment.DefaultAvailableAt = paymentDescAvailableAt.Default.(func() uint32)
	platformMixin := schema.Platform{}.Mixin()
	platformMixinFields0 := platformMixin[0].Fields()
	_ = platformMixinFields0
	platformMixinFields1 := platformMixin[1].Fields()
	_ = platformMixinFields1
	platformFields := schema.Platform{}.Fields()
	_ = platformFields
	// platformDescCreatedAt is the schema descriptor for created_at field.
	platformDescCreatedAt := platformMixinFields0[0].Descriptor()
	// platform.DefaultCreatedAt holds the default value on creation for the created_at field.
	platform.DefaultCreatedAt = platformDescCreatedAt.Default.(func() uint32)
	// platformDescUpdatedAt is the schema descriptor for updated_at field.
	platformDescUpdatedAt := platformMixinFields0[1].Descriptor()
	// platform.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	platform.DefaultUpdatedAt = platformDescUpdatedAt.Default.(func() uint32)
	// platform.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	platform.UpdateDefaultUpdatedAt = platformDescUpdatedAt.UpdateDefault.(func() uint32)
	// platformDescDeletedAt is the schema descriptor for deleted_at field.
	platformDescDeletedAt := platformMixinFields0[2].Descriptor()
	// platform.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	platform.DefaultDeletedAt = platformDescDeletedAt.Default.(func() uint32)
	// platformDescEntID is the schema descriptor for ent_id field.
	platformDescEntID := platformMixinFields1[1].Descriptor()
	// platform.DefaultEntID holds the default value on creation for the ent_id field.
	platform.DefaultEntID = platformDescEntID.Default.(func() uuid.UUID)
	// platformDescAccountID is the schema descriptor for account_id field.
	platformDescAccountID := platformFields[0].Descriptor()
	// platform.DefaultAccountID holds the default value on creation for the account_id field.
	platform.DefaultAccountID = platformDescAccountID.Default.(func() uuid.UUID)
	// platformDescUsedFor is the schema descriptor for used_for field.
	platformDescUsedFor := platformFields[1].Descriptor()
	// platform.DefaultUsedFor holds the default value on creation for the used_for field.
	platform.DefaultUsedFor = platformDescUsedFor.Default.(string)
	// platformDescBackup is the schema descriptor for backup field.
	platformDescBackup := platformFields[2].Descriptor()
	// platform.DefaultBackup holds the default value on creation for the backup field.
	platform.DefaultBackup = platformDescBackup.Default.(bool)
	transferMixin := schema.Transfer{}.Mixin()
	transferMixinFields0 := transferMixin[0].Fields()
	_ = transferMixinFields0
	transferMixinFields1 := transferMixin[1].Fields()
	_ = transferMixinFields1
	transferFields := schema.Transfer{}.Fields()
	_ = transferFields
	// transferDescCreatedAt is the schema descriptor for created_at field.
	transferDescCreatedAt := transferMixinFields0[0].Descriptor()
	// transfer.DefaultCreatedAt holds the default value on creation for the created_at field.
	transfer.DefaultCreatedAt = transferDescCreatedAt.Default.(func() uint32)
	// transferDescUpdatedAt is the schema descriptor for updated_at field.
	transferDescUpdatedAt := transferMixinFields0[1].Descriptor()
	// transfer.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	transfer.DefaultUpdatedAt = transferDescUpdatedAt.Default.(func() uint32)
	// transfer.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	transfer.UpdateDefaultUpdatedAt = transferDescUpdatedAt.UpdateDefault.(func() uint32)
	// transferDescDeletedAt is the schema descriptor for deleted_at field.
	transferDescDeletedAt := transferMixinFields0[2].Descriptor()
	// transfer.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	transfer.DefaultDeletedAt = transferDescDeletedAt.Default.(func() uint32)
	// transferDescEntID is the schema descriptor for ent_id field.
	transferDescEntID := transferMixinFields1[1].Descriptor()
	// transfer.DefaultEntID holds the default value on creation for the ent_id field.
	transfer.DefaultEntID = transferDescEntID.Default.(func() uuid.UUID)
	// transferDescAppID is the schema descriptor for app_id field.
	transferDescAppID := transferFields[0].Descriptor()
	// transfer.DefaultAppID holds the default value on creation for the app_id field.
	transfer.DefaultAppID = transferDescAppID.Default.(func() uuid.UUID)
	// transferDescUserID is the schema descriptor for user_id field.
	transferDescUserID := transferFields[1].Descriptor()
	// transfer.DefaultUserID holds the default value on creation for the user_id field.
	transfer.DefaultUserID = transferDescUserID.Default.(func() uuid.UUID)
	// transferDescTargetUserID is the schema descriptor for target_user_id field.
	transferDescTargetUserID := transferFields[2].Descriptor()
	// transfer.DefaultTargetUserID holds the default value on creation for the target_user_id field.
	transfer.DefaultTargetUserID = transferDescTargetUserID.Default.(func() uuid.UUID)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userMixinFields1 := userMixin[1].Fields()
	_ = userMixinFields1
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() uint32)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() uint32)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() uint32)
	// userDescDeletedAt is the schema descriptor for deleted_at field.
	userDescDeletedAt := userMixinFields0[2].Descriptor()
	// user.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	user.DefaultDeletedAt = userDescDeletedAt.Default.(func() uint32)
	// userDescEntID is the schema descriptor for ent_id field.
	userDescEntID := userMixinFields1[1].Descriptor()
	// user.DefaultEntID holds the default value on creation for the ent_id field.
	user.DefaultEntID = userDescEntID.Default.(func() uuid.UUID)
	// userDescAppID is the schema descriptor for app_id field.
	userDescAppID := userFields[0].Descriptor()
	// user.DefaultAppID holds the default value on creation for the app_id field.
	user.DefaultAppID = userDescAppID.Default.(func() uuid.UUID)
	// userDescUserID is the schema descriptor for user_id field.
	userDescUserID := userFields[1].Descriptor()
	// user.DefaultUserID holds the default value on creation for the user_id field.
	user.DefaultUserID = userDescUserID.Default.(func() uuid.UUID)
	// userDescCoinTypeID is the schema descriptor for coin_type_id field.
	userDescCoinTypeID := userFields[2].Descriptor()
	// user.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	user.DefaultCoinTypeID = userDescCoinTypeID.Default.(func() uuid.UUID)
	// userDescAccountID is the schema descriptor for account_id field.
	userDescAccountID := userFields[3].Descriptor()
	// user.DefaultAccountID holds the default value on creation for the account_id field.
	user.DefaultAccountID = userDescAccountID.Default.(func() uuid.UUID)
	// userDescUsedFor is the schema descriptor for used_for field.
	userDescUsedFor := userFields[4].Descriptor()
	// user.DefaultUsedFor holds the default value on creation for the used_for field.
	user.DefaultUsedFor = userDescUsedFor.Default.(string)
	// userDescLabels is the schema descriptor for labels field.
	userDescLabels := userFields[5].Descriptor()
	// user.DefaultLabels holds the default value on creation for the labels field.
	user.DefaultLabels = userDescLabels.Default.([]string)
	// userDescMemo is the schema descriptor for memo field.
	userDescMemo := userFields[6].Descriptor()
	// user.DefaultMemo holds the default value on creation for the memo field.
	user.DefaultMemo = userDescMemo.Default.(string)
}
