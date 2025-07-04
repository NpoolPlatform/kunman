// Code generated by ent, DO NOT EDIT.

package generated

import (
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/apppool"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/coin"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/fractionwithdrawal"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/fractionwithdrawalrule"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/gooduser"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/orderuser"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/pool"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/rootuser"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/schema"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	apppoolMixin := schema.AppPool{}.Mixin()
	apppoolMixinFields0 := apppoolMixin[0].Fields()
	_ = apppoolMixinFields0
	apppoolMixinFields1 := apppoolMixin[1].Fields()
	_ = apppoolMixinFields1
	apppoolFields := schema.AppPool{}.Fields()
	_ = apppoolFields
	// apppoolDescCreatedAt is the schema descriptor for created_at field.
	apppoolDescCreatedAt := apppoolMixinFields0[0].Descriptor()
	// apppool.DefaultCreatedAt holds the default value on creation for the created_at field.
	apppool.DefaultCreatedAt = apppoolDescCreatedAt.Default.(func() uint32)
	// apppoolDescUpdatedAt is the schema descriptor for updated_at field.
	apppoolDescUpdatedAt := apppoolMixinFields0[1].Descriptor()
	// apppool.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	apppool.DefaultUpdatedAt = apppoolDescUpdatedAt.Default.(func() uint32)
	// apppool.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	apppool.UpdateDefaultUpdatedAt = apppoolDescUpdatedAt.UpdateDefault.(func() uint32)
	// apppoolDescDeletedAt is the schema descriptor for deleted_at field.
	apppoolDescDeletedAt := apppoolMixinFields0[2].Descriptor()
	// apppool.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	apppool.DefaultDeletedAt = apppoolDescDeletedAt.Default.(func() uint32)
	// apppoolDescEntID is the schema descriptor for ent_id field.
	apppoolDescEntID := apppoolMixinFields1[1].Descriptor()
	// apppool.DefaultEntID holds the default value on creation for the ent_id field.
	apppool.DefaultEntID = apppoolDescEntID.Default.(func() uuid.UUID)
	// apppoolDescAppID is the schema descriptor for app_id field.
	apppoolDescAppID := apppoolFields[0].Descriptor()
	// apppool.DefaultAppID holds the default value on creation for the app_id field.
	apppool.DefaultAppID = apppoolDescAppID.Default.(func() uuid.UUID)
	// apppoolDescPoolID is the schema descriptor for pool_id field.
	apppoolDescPoolID := apppoolFields[1].Descriptor()
	// apppool.DefaultPoolID holds the default value on creation for the pool_id field.
	apppool.DefaultPoolID = apppoolDescPoolID.Default.(func() uuid.UUID)
	coinMixin := schema.Coin{}.Mixin()
	coinMixinFields0 := coinMixin[0].Fields()
	_ = coinMixinFields0
	coinMixinFields1 := coinMixin[1].Fields()
	_ = coinMixinFields1
	coinFields := schema.Coin{}.Fields()
	_ = coinFields
	// coinDescCreatedAt is the schema descriptor for created_at field.
	coinDescCreatedAt := coinMixinFields0[0].Descriptor()
	// coin.DefaultCreatedAt holds the default value on creation for the created_at field.
	coin.DefaultCreatedAt = coinDescCreatedAt.Default.(func() uint32)
	// coinDescUpdatedAt is the schema descriptor for updated_at field.
	coinDescUpdatedAt := coinMixinFields0[1].Descriptor()
	// coin.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	coin.DefaultUpdatedAt = coinDescUpdatedAt.Default.(func() uint32)
	// coin.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	coin.UpdateDefaultUpdatedAt = coinDescUpdatedAt.UpdateDefault.(func() uint32)
	// coinDescDeletedAt is the schema descriptor for deleted_at field.
	coinDescDeletedAt := coinMixinFields0[2].Descriptor()
	// coin.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	coin.DefaultDeletedAt = coinDescDeletedAt.Default.(func() uint32)
	// coinDescEntID is the schema descriptor for ent_id field.
	coinDescEntID := coinMixinFields1[1].Descriptor()
	// coin.DefaultEntID holds the default value on creation for the ent_id field.
	coin.DefaultEntID = coinDescEntID.Default.(func() uuid.UUID)
	// coinDescPoolID is the schema descriptor for pool_id field.
	coinDescPoolID := coinFields[0].Descriptor()
	// coin.DefaultPoolID holds the default value on creation for the pool_id field.
	coin.DefaultPoolID = coinDescPoolID.Default.(func() uuid.UUID)
	// coinDescCoinTypeID is the schema descriptor for coin_type_id field.
	coinDescCoinTypeID := coinFields[1].Descriptor()
	// coin.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	coin.DefaultCoinTypeID = coinDescCoinTypeID.Default.(func() uuid.UUID)
	// coinDescCoinType is the schema descriptor for coin_type field.
	coinDescCoinType := coinFields[2].Descriptor()
	// coin.DefaultCoinType holds the default value on creation for the coin_type field.
	coin.DefaultCoinType = coinDescCoinType.Default.(string)
	// coinDescFeeRatio is the schema descriptor for fee_ratio field.
	coinDescFeeRatio := coinFields[3].Descriptor()
	// coin.DefaultFeeRatio holds the default value on creation for the fee_ratio field.
	coin.DefaultFeeRatio = coinDescFeeRatio.Default.(decimal.Decimal)
	// coinDescFixedRevenueAble is the schema descriptor for fixed_revenue_able field.
	coinDescFixedRevenueAble := coinFields[4].Descriptor()
	// coin.DefaultFixedRevenueAble holds the default value on creation for the fixed_revenue_able field.
	coin.DefaultFixedRevenueAble = coinDescFixedRevenueAble.Default.(bool)
	// coinDescLeastTransferAmount is the schema descriptor for least_transfer_amount field.
	coinDescLeastTransferAmount := coinFields[5].Descriptor()
	// coin.DefaultLeastTransferAmount holds the default value on creation for the least_transfer_amount field.
	coin.DefaultLeastTransferAmount = coinDescLeastTransferAmount.Default.(decimal.Decimal)
	// coinDescBenefitIntervalSeconds is the schema descriptor for benefit_interval_seconds field.
	coinDescBenefitIntervalSeconds := coinFields[6].Descriptor()
	// coin.DefaultBenefitIntervalSeconds holds the default value on creation for the benefit_interval_seconds field.
	coin.DefaultBenefitIntervalSeconds = coinDescBenefitIntervalSeconds.Default.(uint32)
	// coinDescRemark is the schema descriptor for remark field.
	coinDescRemark := coinFields[7].Descriptor()
	// coin.DefaultRemark holds the default value on creation for the remark field.
	coin.DefaultRemark = coinDescRemark.Default.(string)
	fractionwithdrawalMixin := schema.FractionWithdrawal{}.Mixin()
	fractionwithdrawalMixinFields0 := fractionwithdrawalMixin[0].Fields()
	_ = fractionwithdrawalMixinFields0
	fractionwithdrawalMixinFields1 := fractionwithdrawalMixin[1].Fields()
	_ = fractionwithdrawalMixinFields1
	fractionwithdrawalFields := schema.FractionWithdrawal{}.Fields()
	_ = fractionwithdrawalFields
	// fractionwithdrawalDescCreatedAt is the schema descriptor for created_at field.
	fractionwithdrawalDescCreatedAt := fractionwithdrawalMixinFields0[0].Descriptor()
	// fractionwithdrawal.DefaultCreatedAt holds the default value on creation for the created_at field.
	fractionwithdrawal.DefaultCreatedAt = fractionwithdrawalDescCreatedAt.Default.(func() uint32)
	// fractionwithdrawalDescUpdatedAt is the schema descriptor for updated_at field.
	fractionwithdrawalDescUpdatedAt := fractionwithdrawalMixinFields0[1].Descriptor()
	// fractionwithdrawal.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	fractionwithdrawal.DefaultUpdatedAt = fractionwithdrawalDescUpdatedAt.Default.(func() uint32)
	// fractionwithdrawal.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	fractionwithdrawal.UpdateDefaultUpdatedAt = fractionwithdrawalDescUpdatedAt.UpdateDefault.(func() uint32)
	// fractionwithdrawalDescDeletedAt is the schema descriptor for deleted_at field.
	fractionwithdrawalDescDeletedAt := fractionwithdrawalMixinFields0[2].Descriptor()
	// fractionwithdrawal.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	fractionwithdrawal.DefaultDeletedAt = fractionwithdrawalDescDeletedAt.Default.(func() uint32)
	// fractionwithdrawalDescEntID is the schema descriptor for ent_id field.
	fractionwithdrawalDescEntID := fractionwithdrawalMixinFields1[1].Descriptor()
	// fractionwithdrawal.DefaultEntID holds the default value on creation for the ent_id field.
	fractionwithdrawal.DefaultEntID = fractionwithdrawalDescEntID.Default.(func() uuid.UUID)
	// fractionwithdrawalDescAppID is the schema descriptor for app_id field.
	fractionwithdrawalDescAppID := fractionwithdrawalFields[0].Descriptor()
	// fractionwithdrawal.DefaultAppID holds the default value on creation for the app_id field.
	fractionwithdrawal.DefaultAppID = fractionwithdrawalDescAppID.Default.(func() uuid.UUID)
	// fractionwithdrawalDescUserID is the schema descriptor for user_id field.
	fractionwithdrawalDescUserID := fractionwithdrawalFields[1].Descriptor()
	// fractionwithdrawal.DefaultUserID holds the default value on creation for the user_id field.
	fractionwithdrawal.DefaultUserID = fractionwithdrawalDescUserID.Default.(func() uuid.UUID)
	// fractionwithdrawalDescOrderUserID is the schema descriptor for order_user_id field.
	fractionwithdrawalDescOrderUserID := fractionwithdrawalFields[2].Descriptor()
	// fractionwithdrawal.DefaultOrderUserID holds the default value on creation for the order_user_id field.
	fractionwithdrawal.DefaultOrderUserID = fractionwithdrawalDescOrderUserID.Default.(func() uuid.UUID)
	// fractionwithdrawalDescCoinTypeID is the schema descriptor for coin_type_id field.
	fractionwithdrawalDescCoinTypeID := fractionwithdrawalFields[3].Descriptor()
	// fractionwithdrawal.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	fractionwithdrawal.DefaultCoinTypeID = fractionwithdrawalDescCoinTypeID.Default.(func() uuid.UUID)
	// fractionwithdrawalDescFractionWithdrawState is the schema descriptor for fraction_withdraw_state field.
	fractionwithdrawalDescFractionWithdrawState := fractionwithdrawalFields[4].Descriptor()
	// fractionwithdrawal.DefaultFractionWithdrawState holds the default value on creation for the fraction_withdraw_state field.
	fractionwithdrawal.DefaultFractionWithdrawState = fractionwithdrawalDescFractionWithdrawState.Default.(string)
	// fractionwithdrawalDescWithdrawAt is the schema descriptor for withdraw_at field.
	fractionwithdrawalDescWithdrawAt := fractionwithdrawalFields[5].Descriptor()
	// fractionwithdrawal.DefaultWithdrawAt holds the default value on creation for the withdraw_at field.
	fractionwithdrawal.DefaultWithdrawAt = fractionwithdrawalDescWithdrawAt.Default.(uint32)
	// fractionwithdrawalDescPromisePayAt is the schema descriptor for promise_pay_at field.
	fractionwithdrawalDescPromisePayAt := fractionwithdrawalFields[6].Descriptor()
	// fractionwithdrawal.DefaultPromisePayAt holds the default value on creation for the promise_pay_at field.
	fractionwithdrawal.DefaultPromisePayAt = fractionwithdrawalDescPromisePayAt.Default.(uint32)
	// fractionwithdrawalDescMsg is the schema descriptor for msg field.
	fractionwithdrawalDescMsg := fractionwithdrawalFields[7].Descriptor()
	// fractionwithdrawal.DefaultMsg holds the default value on creation for the msg field.
	fractionwithdrawal.DefaultMsg = fractionwithdrawalDescMsg.Default.(string)
	fractionwithdrawalruleMixin := schema.FractionWithdrawalRule{}.Mixin()
	fractionwithdrawalruleMixinFields0 := fractionwithdrawalruleMixin[0].Fields()
	_ = fractionwithdrawalruleMixinFields0
	fractionwithdrawalruleMixinFields1 := fractionwithdrawalruleMixin[1].Fields()
	_ = fractionwithdrawalruleMixinFields1
	fractionwithdrawalruleFields := schema.FractionWithdrawalRule{}.Fields()
	_ = fractionwithdrawalruleFields
	// fractionwithdrawalruleDescCreatedAt is the schema descriptor for created_at field.
	fractionwithdrawalruleDescCreatedAt := fractionwithdrawalruleMixinFields0[0].Descriptor()
	// fractionwithdrawalrule.DefaultCreatedAt holds the default value on creation for the created_at field.
	fractionwithdrawalrule.DefaultCreatedAt = fractionwithdrawalruleDescCreatedAt.Default.(func() uint32)
	// fractionwithdrawalruleDescUpdatedAt is the schema descriptor for updated_at field.
	fractionwithdrawalruleDescUpdatedAt := fractionwithdrawalruleMixinFields0[1].Descriptor()
	// fractionwithdrawalrule.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	fractionwithdrawalrule.DefaultUpdatedAt = fractionwithdrawalruleDescUpdatedAt.Default.(func() uint32)
	// fractionwithdrawalrule.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	fractionwithdrawalrule.UpdateDefaultUpdatedAt = fractionwithdrawalruleDescUpdatedAt.UpdateDefault.(func() uint32)
	// fractionwithdrawalruleDescDeletedAt is the schema descriptor for deleted_at field.
	fractionwithdrawalruleDescDeletedAt := fractionwithdrawalruleMixinFields0[2].Descriptor()
	// fractionwithdrawalrule.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	fractionwithdrawalrule.DefaultDeletedAt = fractionwithdrawalruleDescDeletedAt.Default.(func() uint32)
	// fractionwithdrawalruleDescEntID is the schema descriptor for ent_id field.
	fractionwithdrawalruleDescEntID := fractionwithdrawalruleMixinFields1[1].Descriptor()
	// fractionwithdrawalrule.DefaultEntID holds the default value on creation for the ent_id field.
	fractionwithdrawalrule.DefaultEntID = fractionwithdrawalruleDescEntID.Default.(func() uuid.UUID)
	// fractionwithdrawalruleDescPoolCoinTypeID is the schema descriptor for pool_coin_type_id field.
	fractionwithdrawalruleDescPoolCoinTypeID := fractionwithdrawalruleFields[0].Descriptor()
	// fractionwithdrawalrule.DefaultPoolCoinTypeID holds the default value on creation for the pool_coin_type_id field.
	fractionwithdrawalrule.DefaultPoolCoinTypeID = fractionwithdrawalruleDescPoolCoinTypeID.Default.(func() uuid.UUID)
	// fractionwithdrawalruleDescWithdrawInterval is the schema descriptor for withdraw_interval field.
	fractionwithdrawalruleDescWithdrawInterval := fractionwithdrawalruleFields[1].Descriptor()
	// fractionwithdrawalrule.DefaultWithdrawInterval holds the default value on creation for the withdraw_interval field.
	fractionwithdrawalrule.DefaultWithdrawInterval = fractionwithdrawalruleDescWithdrawInterval.Default.(uint32)
	// fractionwithdrawalruleDescLeastWithdrawalAmount is the schema descriptor for least_withdrawal_amount field.
	fractionwithdrawalruleDescLeastWithdrawalAmount := fractionwithdrawalruleFields[2].Descriptor()
	// fractionwithdrawalrule.DefaultLeastWithdrawalAmount holds the default value on creation for the least_withdrawal_amount field.
	fractionwithdrawalrule.DefaultLeastWithdrawalAmount = fractionwithdrawalruleDescLeastWithdrawalAmount.Default.(decimal.Decimal)
	// fractionwithdrawalruleDescPayoutThreshold is the schema descriptor for payout_threshold field.
	fractionwithdrawalruleDescPayoutThreshold := fractionwithdrawalruleFields[3].Descriptor()
	// fractionwithdrawalrule.DefaultPayoutThreshold holds the default value on creation for the payout_threshold field.
	fractionwithdrawalrule.DefaultPayoutThreshold = fractionwithdrawalruleDescPayoutThreshold.Default.(decimal.Decimal)
	// fractionwithdrawalruleDescWithdrawFee is the schema descriptor for withdraw_fee field.
	fractionwithdrawalruleDescWithdrawFee := fractionwithdrawalruleFields[4].Descriptor()
	// fractionwithdrawalrule.DefaultWithdrawFee holds the default value on creation for the withdraw_fee field.
	fractionwithdrawalrule.DefaultWithdrawFee = fractionwithdrawalruleDescWithdrawFee.Default.(decimal.Decimal)
	gooduserMixin := schema.GoodUser{}.Mixin()
	gooduserMixinFields0 := gooduserMixin[0].Fields()
	_ = gooduserMixinFields0
	gooduserMixinFields1 := gooduserMixin[1].Fields()
	_ = gooduserMixinFields1
	gooduserFields := schema.GoodUser{}.Fields()
	_ = gooduserFields
	// gooduserDescCreatedAt is the schema descriptor for created_at field.
	gooduserDescCreatedAt := gooduserMixinFields0[0].Descriptor()
	// gooduser.DefaultCreatedAt holds the default value on creation for the created_at field.
	gooduser.DefaultCreatedAt = gooduserDescCreatedAt.Default.(func() uint32)
	// gooduserDescUpdatedAt is the schema descriptor for updated_at field.
	gooduserDescUpdatedAt := gooduserMixinFields0[1].Descriptor()
	// gooduser.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	gooduser.DefaultUpdatedAt = gooduserDescUpdatedAt.Default.(func() uint32)
	// gooduser.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	gooduser.UpdateDefaultUpdatedAt = gooduserDescUpdatedAt.UpdateDefault.(func() uint32)
	// gooduserDescDeletedAt is the schema descriptor for deleted_at field.
	gooduserDescDeletedAt := gooduserMixinFields0[2].Descriptor()
	// gooduser.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	gooduser.DefaultDeletedAt = gooduserDescDeletedAt.Default.(func() uint32)
	// gooduserDescEntID is the schema descriptor for ent_id field.
	gooduserDescEntID := gooduserMixinFields1[1].Descriptor()
	// gooduser.DefaultEntID holds the default value on creation for the ent_id field.
	gooduser.DefaultEntID = gooduserDescEntID.Default.(func() uuid.UUID)
	// gooduserDescRootUserID is the schema descriptor for root_user_id field.
	gooduserDescRootUserID := gooduserFields[0].Descriptor()
	// gooduser.DefaultRootUserID holds the default value on creation for the root_user_id field.
	gooduser.DefaultRootUserID = gooduserDescRootUserID.Default.(func() uuid.UUID)
	// gooduserDescName is the schema descriptor for name field.
	gooduserDescName := gooduserFields[1].Descriptor()
	// gooduser.DefaultName holds the default value on creation for the name field.
	gooduser.DefaultName = gooduserDescName.Default.(string)
	// gooduserDescReadPageLink is the schema descriptor for read_page_link field.
	gooduserDescReadPageLink := gooduserFields[2].Descriptor()
	// gooduser.DefaultReadPageLink holds the default value on creation for the read_page_link field.
	gooduser.DefaultReadPageLink = gooduserDescReadPageLink.Default.(string)
	orderuserMixin := schema.OrderUser{}.Mixin()
	orderuserMixinFields0 := orderuserMixin[0].Fields()
	_ = orderuserMixinFields0
	orderuserMixinFields1 := orderuserMixin[1].Fields()
	_ = orderuserMixinFields1
	orderuserFields := schema.OrderUser{}.Fields()
	_ = orderuserFields
	// orderuserDescCreatedAt is the schema descriptor for created_at field.
	orderuserDescCreatedAt := orderuserMixinFields0[0].Descriptor()
	// orderuser.DefaultCreatedAt holds the default value on creation for the created_at field.
	orderuser.DefaultCreatedAt = orderuserDescCreatedAt.Default.(func() uint32)
	// orderuserDescUpdatedAt is the schema descriptor for updated_at field.
	orderuserDescUpdatedAt := orderuserMixinFields0[1].Descriptor()
	// orderuser.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	orderuser.DefaultUpdatedAt = orderuserDescUpdatedAt.Default.(func() uint32)
	// orderuser.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	orderuser.UpdateDefaultUpdatedAt = orderuserDescUpdatedAt.UpdateDefault.(func() uint32)
	// orderuserDescDeletedAt is the schema descriptor for deleted_at field.
	orderuserDescDeletedAt := orderuserMixinFields0[2].Descriptor()
	// orderuser.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	orderuser.DefaultDeletedAt = orderuserDescDeletedAt.Default.(func() uint32)
	// orderuserDescEntID is the schema descriptor for ent_id field.
	orderuserDescEntID := orderuserMixinFields1[1].Descriptor()
	// orderuser.DefaultEntID holds the default value on creation for the ent_id field.
	orderuser.DefaultEntID = orderuserDescEntID.Default.(func() uuid.UUID)
	// orderuserDescGoodUserID is the schema descriptor for good_user_id field.
	orderuserDescGoodUserID := orderuserFields[0].Descriptor()
	// orderuser.DefaultGoodUserID holds the default value on creation for the good_user_id field.
	orderuser.DefaultGoodUserID = orderuserDescGoodUserID.Default.(func() uuid.UUID)
	// orderuserDescUserID is the schema descriptor for user_id field.
	orderuserDescUserID := orderuserFields[1].Descriptor()
	// orderuser.DefaultUserID holds the default value on creation for the user_id field.
	orderuser.DefaultUserID = orderuserDescUserID.Default.(func() uuid.UUID)
	// orderuserDescAppID is the schema descriptor for app_id field.
	orderuserDescAppID := orderuserFields[2].Descriptor()
	// orderuser.DefaultAppID holds the default value on creation for the app_id field.
	orderuser.DefaultAppID = orderuserDescAppID.Default.(func() uuid.UUID)
	// orderuserDescName is the schema descriptor for name field.
	orderuserDescName := orderuserFields[3].Descriptor()
	// orderuser.DefaultName holds the default value on creation for the name field.
	orderuser.DefaultName = orderuserDescName.Default.(string)
	// orderuserDescReadPageLink is the schema descriptor for read_page_link field.
	orderuserDescReadPageLink := orderuserFields[4].Descriptor()
	// orderuser.DefaultReadPageLink holds the default value on creation for the read_page_link field.
	orderuser.DefaultReadPageLink = orderuserDescReadPageLink.Default.(string)
	poolMixin := schema.Pool{}.Mixin()
	poolMixinFields0 := poolMixin[0].Fields()
	_ = poolMixinFields0
	poolMixinFields1 := poolMixin[1].Fields()
	_ = poolMixinFields1
	poolFields := schema.Pool{}.Fields()
	_ = poolFields
	// poolDescCreatedAt is the schema descriptor for created_at field.
	poolDescCreatedAt := poolMixinFields0[0].Descriptor()
	// pool.DefaultCreatedAt holds the default value on creation for the created_at field.
	pool.DefaultCreatedAt = poolDescCreatedAt.Default.(func() uint32)
	// poolDescUpdatedAt is the schema descriptor for updated_at field.
	poolDescUpdatedAt := poolMixinFields0[1].Descriptor()
	// pool.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	pool.DefaultUpdatedAt = poolDescUpdatedAt.Default.(func() uint32)
	// pool.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	pool.UpdateDefaultUpdatedAt = poolDescUpdatedAt.UpdateDefault.(func() uint32)
	// poolDescDeletedAt is the schema descriptor for deleted_at field.
	poolDescDeletedAt := poolMixinFields0[2].Descriptor()
	// pool.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	pool.DefaultDeletedAt = poolDescDeletedAt.Default.(func() uint32)
	// poolDescEntID is the schema descriptor for ent_id field.
	poolDescEntID := poolMixinFields1[1].Descriptor()
	// pool.DefaultEntID holds the default value on creation for the ent_id field.
	pool.DefaultEntID = poolDescEntID.Default.(func() uuid.UUID)
	// poolDescMiningPoolType is the schema descriptor for mining_pool_type field.
	poolDescMiningPoolType := poolFields[0].Descriptor()
	// pool.DefaultMiningPoolType holds the default value on creation for the mining_pool_type field.
	pool.DefaultMiningPoolType = poolDescMiningPoolType.Default.(string)
	// poolDescName is the schema descriptor for name field.
	poolDescName := poolFields[1].Descriptor()
	// pool.DefaultName holds the default value on creation for the name field.
	pool.DefaultName = poolDescName.Default.(string)
	// poolDescSite is the schema descriptor for site field.
	poolDescSite := poolFields[2].Descriptor()
	// pool.DefaultSite holds the default value on creation for the site field.
	pool.DefaultSite = poolDescSite.Default.(string)
	// poolDescLogo is the schema descriptor for logo field.
	poolDescLogo := poolFields[3].Descriptor()
	// pool.DefaultLogo holds the default value on creation for the logo field.
	pool.DefaultLogo = poolDescLogo.Default.(string)
	// poolDescDescription is the schema descriptor for description field.
	poolDescDescription := poolFields[4].Descriptor()
	// pool.DefaultDescription holds the default value on creation for the description field.
	pool.DefaultDescription = poolDescDescription.Default.(string)
	rootuserMixin := schema.RootUser{}.Mixin()
	rootuserMixinFields0 := rootuserMixin[0].Fields()
	_ = rootuserMixinFields0
	rootuserMixinFields1 := rootuserMixin[1].Fields()
	_ = rootuserMixinFields1
	rootuserFields := schema.RootUser{}.Fields()
	_ = rootuserFields
	// rootuserDescCreatedAt is the schema descriptor for created_at field.
	rootuserDescCreatedAt := rootuserMixinFields0[0].Descriptor()
	// rootuser.DefaultCreatedAt holds the default value on creation for the created_at field.
	rootuser.DefaultCreatedAt = rootuserDescCreatedAt.Default.(func() uint32)
	// rootuserDescUpdatedAt is the schema descriptor for updated_at field.
	rootuserDescUpdatedAt := rootuserMixinFields0[1].Descriptor()
	// rootuser.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	rootuser.DefaultUpdatedAt = rootuserDescUpdatedAt.Default.(func() uint32)
	// rootuser.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	rootuser.UpdateDefaultUpdatedAt = rootuserDescUpdatedAt.UpdateDefault.(func() uint32)
	// rootuserDescDeletedAt is the schema descriptor for deleted_at field.
	rootuserDescDeletedAt := rootuserMixinFields0[2].Descriptor()
	// rootuser.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	rootuser.DefaultDeletedAt = rootuserDescDeletedAt.Default.(func() uint32)
	// rootuserDescEntID is the schema descriptor for ent_id field.
	rootuserDescEntID := rootuserMixinFields1[1].Descriptor()
	// rootuser.DefaultEntID holds the default value on creation for the ent_id field.
	rootuser.DefaultEntID = rootuserDescEntID.Default.(func() uuid.UUID)
	// rootuserDescName is the schema descriptor for name field.
	rootuserDescName := rootuserFields[0].Descriptor()
	// rootuser.DefaultName holds the default value on creation for the name field.
	rootuser.DefaultName = rootuserDescName.Default.(string)
	// rootuserDescPoolID is the schema descriptor for pool_id field.
	rootuserDescPoolID := rootuserFields[1].Descriptor()
	// rootuser.DefaultPoolID holds the default value on creation for the pool_id field.
	rootuser.DefaultPoolID = rootuserDescPoolID.Default.(func() uuid.UUID)
	// rootuserDescEmail is the schema descriptor for email field.
	rootuserDescEmail := rootuserFields[2].Descriptor()
	// rootuser.DefaultEmail holds the default value on creation for the email field.
	rootuser.DefaultEmail = rootuserDescEmail.Default.(string)
	// rootuserDescAuthToken is the schema descriptor for auth_token field.
	rootuserDescAuthToken := rootuserFields[3].Descriptor()
	// rootuser.DefaultAuthToken holds the default value on creation for the auth_token field.
	rootuser.DefaultAuthToken = rootuserDescAuthToken.Default.(string)
	// rootuserDescAuthTokenSalt is the schema descriptor for auth_token_salt field.
	rootuserDescAuthTokenSalt := rootuserFields[4].Descriptor()
	// rootuser.DefaultAuthTokenSalt holds the default value on creation for the auth_token_salt field.
	rootuser.DefaultAuthTokenSalt = rootuserDescAuthTokenSalt.Default.(string)
	// rootuserDescAuthed is the schema descriptor for authed field.
	rootuserDescAuthed := rootuserFields[5].Descriptor()
	// rootuser.DefaultAuthed holds the default value on creation for the authed field.
	rootuser.DefaultAuthed = rootuserDescAuthed.Default.(bool)
	// rootuserDescRemark is the schema descriptor for remark field.
	rootuserDescRemark := rootuserFields[6].Descriptor()
	// rootuser.DefaultRemark holds the default value on creation for the remark field.
	rootuser.DefaultRemark = rootuserDescRemark.Default.(string)
}
