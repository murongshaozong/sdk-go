package types

import (
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = &Params{}

// Exchange params default values
const (
	// DefaultFundingIntervalSeconds is 3600. This represents the number of seconds in one hour which is the frequency at which
	// funding is applied by default on derivative markets.
	DefaultFundingIntervalSeconds int64 = 3600

	// DefaultFundingMultipleSeconds is 3600. This represents the number of seconds in one hour which is multiple of the
	// unix time seconds timestamp that each perpetual market's funding timestamp should be. This ensures that
	// funding is consistently applied on the hour for all perpetual markets.
	DefaultFundingMultipleSeconds int64 = 3600

	// SpotMarketInstantListingFee is 20 INJ
	SpotMarketInstantListingFee int64 = 20

	// DerivativeMarketInstantListingFee is 20 INJ
	DerivativeMarketInstantListingFee int64 = 20

	// BinaryOptionsMarketInstantListingFee is 100 INJ
	BinaryOptionsMarketInstantListingFee int64 = 100

	// MaxDerivativeOrderSideCount is 20
	MaxDerivativeOrderSideCount uint32 = 20

	MaxOracleScaleFactor uint32 = 18

	MaxDecimals uint32 = 18

	MaxTickerLength int = 40

	// MaxHistoricalTradeRecordAge is the maximum age of trade records to track.
	MaxHistoricalTradeRecordAge = 60 * 5

	// MaxSubaccountNonceLength restricts the size of a subaccount number from 0 to 999
	MaxSubaccountNonceLength = 3

	// MaxGranterDelegations is the maximum number of delegations that are checked for stake granter
	MaxGranterDelegations = 25
)

var DefaultInjAuctionMaxCap = math.NewIntWithDecimal(10_000, 18)

var MaxBinaryOptionsOrderPrice = math.LegacyOneDec()

// would be $0.000001 for USDT
var MinDerivativeOrderPrice = math.LegacyOneDec()

// MaxOrderPrice equals 10^32
var MaxOrderPrice = math.LegacyMustNewDecFromStr("100000000000000000000000000000000")

// MaxOrderMargin equals 10^32
var MaxOrderMargin = math.LegacyMustNewDecFromStr("100000000000000000000000000000000")

// MaxTokenInt equals 100,000,000 * 10^18
var MaxTokenInt, _ = math.NewIntFromString("100000000000000000000000000")

var MaxOrderQuantity = math.LegacyMustNewDecFromStr("100000000000000000000000000000000")
var MaxFeeMultiplier = math.LegacyMustNewDecFromStr("100")

var minMarginRatio = math.LegacyNewDecWithPrec(5, 3)

// Parameter keys
var (
	KeySpotMarketInstantListingFee                 = []byte("SpotMarketInstantListingFee")
	KeyDerivativeMarketInstantListingFee           = []byte("DerivativeMarketInstantListingFee")
	KeyDefaultSpotMakerFeeRate                     = []byte("DefaultSpotMakerFeeRate")
	KeyDefaultSpotTakerFeeRate                     = []byte("DefaultSpotTakerFeeRate")
	KeyDefaultDerivativeMakerFeeRate               = []byte("DefaultDerivativeMakerFeeRate")
	KeyDefaultDerivativeTakerFeeRate               = []byte("DefaultDerivativeTakerFeeRate")
	KeyDefaultInitialMarginRatio                   = []byte("DefaultInitialMarginRatio")
	KeyDefaultMaintenanceMarginRatio               = []byte("DefaultMaintenanceMarginRatio")
	KeyDefaultFundingInterval                      = []byte("DefaultFundingInterval")
	KeyFundingMultiple                             = []byte("FundingMultiple")
	KeyRelayerFeeShareRate                         = []byte("RelayerFeeShareRate")
	KeyDefaultHourlyFundingRateCap                 = []byte("DefaultHourlyFundingRateCap")
	KeyDefaultHourlyInterestRate                   = []byte("DefaultHourlyInterestRate")
	KeyMaxDerivativeOrderSideCount                 = []byte("MaxDerivativeOrderSideCount")
	KeyInjRewardStakedRequirementThreshold         = []byte("KeyInjRewardStakedRequirementThreshold")
	KeyTradingRewardsVestingDuration               = []byte("TradingRewardsVestingDuration")
	KeyLiquidatorRewardShareRate                   = []byte("LiquidatorRewardShareRate")
	KeyBinaryOptionsMarketInstantListingFee        = []byte("BinaryOptionsMarketInstantListingFee")
	KeyAtomicMarketOrderAccessLevel                = []byte("AtomicMarketOrderAccessLevel")
	KeySpotAtomicMarketOrderFeeMultiplier          = []byte("SpotAtomicMarketOrderFeeMultiplier")
	KeyDerivativeAtomicMarketOrderFeeMultiplier    = []byte("DerivativeAtomicMarketOrderFeeMultiplier")
	KeyBinaryOptionsAtomicMarketOrderFeeMultiplier = []byte("BinaryOptionsAtomicMarketOrderFeeMultiplier")
	KeyMinimalProtocolFeeRate                      = []byte("MinimalProtocolFeeRate")
	KeyIsInstantDerivativeMarketLaunchEnabled      = []byte("IsInstantDerivativeMarketLaunchEnabled")
	KeyPostOnlyModeHeightThreshold                 = []byte("PostOnlyModeHeightThreshold")
)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	spotMarketInstantListingFee sdk.Coin,
	derivativeMarketInstantListingFee sdk.Coin,
	defaultSpotMakerFee math.LegacyDec,
	defaultSpotTakerFee math.LegacyDec,
	defaultDerivativeMakerFee math.LegacyDec,
	defaultDerivativeTakerFee math.LegacyDec,
	defaultInitialMarginRatio math.LegacyDec,
	defaultMaintenanceMarginRatio math.LegacyDec,
	defaultFundingInterval int64,
	fundingMultiple int64,
	relayerFeeShare math.LegacyDec,
	defaultHourlyFundingRateCap math.LegacyDec,
	defaultHourlyInterestRate math.LegacyDec,
	maxDerivativeSideOrderCount uint32,
	injRewardStakedRequirementThreshold math.Int,
	tradingRewardsVestingDuration int64,
	liquidatorRewardShareRate math.LegacyDec,
	binaryOptionsMarketInstantListingFee sdk.Coin,
	atomicMarketOrderAccessLevel AtomicMarketOrderAccessLevel,
	spotAtomicMarketOrderFeeMultiplier math.LegacyDec,
	derivativeAtomicMarketOrderFeeMultiplier math.LegacyDec,
	binaryOptionsAtomicMarketOrderFeeMultiplier math.LegacyDec,
	minimalProtocolFeeRate math.LegacyDec,
	postOnlyModeHeightThreshold int64,
) Params {
	return Params{
		SpotMarketInstantListingFee:                 spotMarketInstantListingFee,
		DerivativeMarketInstantListingFee:           derivativeMarketInstantListingFee,
		DefaultSpotMakerFeeRate:                     defaultSpotMakerFee,
		DefaultSpotTakerFeeRate:                     defaultSpotTakerFee,
		DefaultDerivativeMakerFeeRate:               defaultDerivativeMakerFee,
		DefaultDerivativeTakerFeeRate:               defaultDerivativeTakerFee,
		DefaultInitialMarginRatio:                   defaultInitialMarginRatio,
		DefaultMaintenanceMarginRatio:               defaultMaintenanceMarginRatio,
		DefaultFundingInterval:                      defaultFundingInterval,
		FundingMultiple:                             fundingMultiple,
		RelayerFeeShareRate:                         relayerFeeShare,
		DefaultHourlyFundingRateCap:                 defaultHourlyFundingRateCap,
		DefaultHourlyInterestRate:                   defaultHourlyInterestRate,
		MaxDerivativeOrderSideCount:                 maxDerivativeSideOrderCount,
		InjRewardStakedRequirementThreshold:         injRewardStakedRequirementThreshold,
		TradingRewardsVestingDuration:               tradingRewardsVestingDuration,
		LiquidatorRewardShareRate:                   liquidatorRewardShareRate,
		BinaryOptionsMarketInstantListingFee:        binaryOptionsMarketInstantListingFee,
		AtomicMarketOrderAccessLevel:                atomicMarketOrderAccessLevel,
		SpotAtomicMarketOrderFeeMultiplier:          spotAtomicMarketOrderFeeMultiplier,
		DerivativeAtomicMarketOrderFeeMultiplier:    derivativeAtomicMarketOrderFeeMultiplier,
		BinaryOptionsAtomicMarketOrderFeeMultiplier: binaryOptionsAtomicMarketOrderFeeMultiplier,
		MinimalProtocolFeeRate:                      minimalProtocolFeeRate,
		IsInstantDerivativeMarketLaunchEnabled:      false,
		PostOnlyModeHeightThreshold:                 postOnlyModeHeightThreshold,
	}
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeySpotMarketInstantListingFee, &p.SpotMarketInstantListingFee, validateSpotMarketInstantListingFee),
		paramtypes.NewParamSetPair(KeyDerivativeMarketInstantListingFee, &p.DerivativeMarketInstantListingFee, validateDerivativeMarketInstantListingFee),
		paramtypes.NewParamSetPair(KeyDefaultSpotMakerFeeRate, &p.DefaultSpotMakerFeeRate, ValidateMakerFee),
		paramtypes.NewParamSetPair(KeyDefaultSpotTakerFeeRate, &p.DefaultSpotTakerFeeRate, ValidateFee),
		paramtypes.NewParamSetPair(KeyDefaultDerivativeMakerFeeRate, &p.DefaultDerivativeMakerFeeRate, ValidateMakerFee),
		paramtypes.NewParamSetPair(KeyDefaultDerivativeTakerFeeRate, &p.DefaultDerivativeTakerFeeRate, ValidateFee),
		paramtypes.NewParamSetPair(KeyDefaultInitialMarginRatio, &p.DefaultInitialMarginRatio, ValidateMarginRatio),
		paramtypes.NewParamSetPair(KeyDefaultMaintenanceMarginRatio, &p.DefaultMaintenanceMarginRatio, ValidateMarginRatio),
		paramtypes.NewParamSetPair(KeyDefaultFundingInterval, &p.DefaultFundingInterval, validateFundingInterval),
		paramtypes.NewParamSetPair(KeyFundingMultiple, &p.FundingMultiple, validateFundingMultiple),
		paramtypes.NewParamSetPair(KeyRelayerFeeShareRate, &p.RelayerFeeShareRate, ValidateFee),
		paramtypes.NewParamSetPair(KeyDefaultHourlyFundingRateCap, &p.DefaultHourlyFundingRateCap, ValidateFee),
		paramtypes.NewParamSetPair(KeyDefaultHourlyInterestRate, &p.DefaultHourlyInterestRate, ValidateFee),
		paramtypes.NewParamSetPair(KeyMaxDerivativeOrderSideCount, &p.MaxDerivativeOrderSideCount, validateDerivativeOrderSideCount),
		paramtypes.NewParamSetPair(KeyInjRewardStakedRequirementThreshold, &p.InjRewardStakedRequirementThreshold, validateInjRewardStakedRequirementThreshold),
		paramtypes.NewParamSetPair(KeyTradingRewardsVestingDuration, &p.TradingRewardsVestingDuration, validateTradingRewardsVestingDuration),
		paramtypes.NewParamSetPair(KeyLiquidatorRewardShareRate, &p.LiquidatorRewardShareRate, validateLiquidatorRewardShareRate),
		paramtypes.NewParamSetPair(KeyBinaryOptionsMarketInstantListingFee, &p.BinaryOptionsMarketInstantListingFee, validateBinaryOptionsMarketInstantListingFee),
		paramtypes.NewParamSetPair(KeyAtomicMarketOrderAccessLevel, &p.AtomicMarketOrderAccessLevel, validateAtomicMarketOrderAccessLevel),
		paramtypes.NewParamSetPair(KeySpotAtomicMarketOrderFeeMultiplier, &p.SpotAtomicMarketOrderFeeMultiplier, validateAtomicMarketOrderFeeMultiplier),
		paramtypes.NewParamSetPair(KeyDerivativeAtomicMarketOrderFeeMultiplier, &p.DerivativeAtomicMarketOrderFeeMultiplier, validateAtomicMarketOrderFeeMultiplier),
		paramtypes.NewParamSetPair(KeyBinaryOptionsAtomicMarketOrderFeeMultiplier, &p.BinaryOptionsAtomicMarketOrderFeeMultiplier, validateAtomicMarketOrderFeeMultiplier),
		paramtypes.NewParamSetPair(KeyMinimalProtocolFeeRate, &p.MinimalProtocolFeeRate, ValidateFee),
		paramtypes.NewParamSetPair(KeyIsInstantDerivativeMarketLaunchEnabled, &p.IsInstantDerivativeMarketLaunchEnabled, validateBool),
		paramtypes.NewParamSetPair(KeyPostOnlyModeHeightThreshold, &p.PostOnlyModeHeightThreshold, validatePostOnlyModeHeightThreshold),
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return Params{
		SpotMarketInstantListingFee:                  sdk.NewCoin("inj", math.NewIntWithDecimal(SpotMarketInstantListingFee, 18)),
		DerivativeMarketInstantListingFee:            sdk.NewCoin("inj", math.NewIntWithDecimal(DerivativeMarketInstantListingFee, 18)),
		DefaultSpotMakerFeeRate:                      math.LegacyNewDecWithPrec(-1, 4), // default -0.01% maker fees
		DefaultSpotTakerFeeRate:                      math.LegacyNewDecWithPrec(1, 3),  // default 0.1% taker fees
		DefaultDerivativeMakerFeeRate:                math.LegacyNewDecWithPrec(-1, 4), // default -0.01% maker fees
		DefaultDerivativeTakerFeeRate:                math.LegacyNewDecWithPrec(1, 3),  // default 0.1% taker fees
		DefaultInitialMarginRatio:                    math.LegacyNewDecWithPrec(5, 2),  // default 5% initial margin ratio
		DefaultMaintenanceMarginRatio:                math.LegacyNewDecWithPrec(2, 2),  // default 2% maintenance margin ratio
		DefaultFundingInterval:                       DefaultFundingIntervalSeconds,
		FundingMultiple:                              DefaultFundingMultipleSeconds,
		RelayerFeeShareRate:                          math.LegacyNewDecWithPrec(40, 2),      // default 40% relayer fee share
		DefaultHourlyFundingRateCap:                  math.LegacyNewDecWithPrec(625, 6),     // default 0.0625% max hourly funding rate
		DefaultHourlyInterestRate:                    math.LegacyNewDecWithPrec(416666, 11), // 0.01% daily interest rate = 0.0001 / 24 = 0.00000416666
		MaxDerivativeOrderSideCount:                  MaxDerivativeOrderSideCount,
		InjRewardStakedRequirementThreshold:          math.NewIntWithDecimal(100, 18), // 100 INJ
		TradingRewardsVestingDuration:                604800,                          // 7 days
		LiquidatorRewardShareRate:                    math.LegacyNewDecWithPrec(5, 2), // 5% liquidator reward
		BinaryOptionsMarketInstantListingFee:         sdk.NewCoin("inj", math.NewIntWithDecimal(BinaryOptionsMarketInstantListingFee, 18)),
		AtomicMarketOrderAccessLevel:                 AtomicMarketOrderAccessLevel_SmartContractsOnly,
		SpotAtomicMarketOrderFeeMultiplier:           math.LegacyNewDecWithPrec(25, 1),        // default 2.5 multiplier
		DerivativeAtomicMarketOrderFeeMultiplier:     math.LegacyNewDecWithPrec(25, 1),        // default 2.5 multiplier
		BinaryOptionsAtomicMarketOrderFeeMultiplier:  math.LegacyNewDecWithPrec(25, 1),        // default 2.5 multiplier
		MinimalProtocolFeeRate:                       math.LegacyMustNewDecFromStr("0.00005"), // default 0.005% minimal fee rate
		IsInstantDerivativeMarketLaunchEnabled:       false,
		PostOnlyModeHeightThreshold:                  0,
		MarginDecreasePriceTimestampThresholdSeconds: 60,
		ExchangeAdmins:                               []string{},
		InjAuctionMaxCap:                             DefaultInjAuctionMaxCap,
	}
}

// Validate performs basic validation on exchange parameters.
func (p Params) Validate() error {
	if err := validateSpotMarketInstantListingFee(p.SpotMarketInstantListingFee); err != nil {
		return fmt.Errorf("spot_market_instant_listing_fee is incorrect: %w", err)
	}
	if err := validateDerivativeMarketInstantListingFee(p.DerivativeMarketInstantListingFee); err != nil {
		return fmt.Errorf("derivative_market_instant_listing_fee is incorrect: %w", err)
	}
	if err := ValidateMakerFee(p.DefaultSpotMakerFeeRate); err != nil {
		return fmt.Errorf("default_spot_maker_fee_rate is incorrect: %w", err)
	}
	if err := ValidateFee(p.DefaultSpotTakerFeeRate); err != nil {
		return fmt.Errorf("default_spot_taker_fee_rate is incorrect: %w", err)
	}
	if err := ValidateMakerFee(p.DefaultDerivativeMakerFeeRate); err != nil {
		return fmt.Errorf("default_derivative_maker_fee_rate is incorrect: %w", err)
	}
	if err := ValidateFee(p.DefaultDerivativeTakerFeeRate); err != nil {
		return fmt.Errorf("default_derivative_taker_fee_rate is incorrect: %w", err)
	}
	if err := ValidateMarginRatio(p.DefaultInitialMarginRatio); err != nil {
		return fmt.Errorf("default_initial_margin_ratio is incorrect: %w", err)
	}
	if err := ValidateMarginRatio(p.DefaultMaintenanceMarginRatio); err != nil {
		return fmt.Errorf("default_maintenance_margin_ratio is incorrect: %w", err)
	}
	if err := validateFundingInterval(p.DefaultFundingInterval); err != nil {
		return fmt.Errorf("default_funding_interval is incorrect: %w", err)
	}
	if err := validateFundingMultiple(p.FundingMultiple); err != nil {
		return fmt.Errorf("funding_multiple is incorrect: %w", err)
	}
	if err := ValidateFee(p.RelayerFeeShareRate); err != nil {
		return fmt.Errorf("relayer_fee_share_rate is incorrect: %w", err)
	}
	if err := ValidateFee(p.DefaultHourlyFundingRateCap); err != nil {
		return fmt.Errorf("default_hourly_funding_rate_cap is incorrect: %w", err)
	}
	if err := ValidateFee(p.DefaultHourlyInterestRate); err != nil {
		return fmt.Errorf("default_hourly_interest_rate is incorrect: %w", err)
	}
	if err := validateDerivativeOrderSideCount(p.MaxDerivativeOrderSideCount); err != nil {
		return fmt.Errorf("max_derivative_order_side_count is incorrect: %w", err)
	}
	if err := validateInjRewardStakedRequirementThreshold(p.InjRewardStakedRequirementThreshold); err != nil {
		return fmt.Errorf("inj_reward_staked_requirement_threshold is incorrect: %w", err)
	}
	if err := validateLiquidatorRewardShareRate(p.LiquidatorRewardShareRate); err != nil {
		return fmt.Errorf("liquidator_reward_share_rate is incorrect: %w", err)
	}
	if err := validateBinaryOptionsMarketInstantListingFee(p.BinaryOptionsMarketInstantListingFee); err != nil {
		return fmt.Errorf("binary_options_market_instant_listing_fee is incorrect: %w", err)
	}
	if err := validateAtomicMarketOrderAccessLevel(p.AtomicMarketOrderAccessLevel); err != nil {
		return fmt.Errorf("atomic_market_order_access_level is incorrect: %w", err)
	}
	if err := validateAtomicMarketOrderFeeMultiplier(p.SpotAtomicMarketOrderFeeMultiplier); err != nil {
		return fmt.Errorf("spot_atomic_market_order_fee_multiplier is incorrect: %w", err)
	}
	if err := validateAtomicMarketOrderFeeMultiplier(p.DerivativeAtomicMarketOrderFeeMultiplier); err != nil {
		return fmt.Errorf("derivative_atomic_market_order_fee_multiplier is incorrect: %w", err)
	}
	if err := validateAtomicMarketOrderFeeMultiplier(p.BinaryOptionsAtomicMarketOrderFeeMultiplier); err != nil {
		return fmt.Errorf("binary_options_atomic_market_order_fee_multiplier is incorrect: %w", err)
	}
	if err := ValidateFee(p.MinimalProtocolFeeRate); err != nil {
		return fmt.Errorf("minimal_protocol_fee_rate is incorrect: %w", err)
	}
	if err := validatePostOnlyModeHeightThreshold(p.PostOnlyModeHeightThreshold); err != nil {
		return fmt.Errorf("post_only_mode_height_threshold is incorrect: %w", err)
	}
	if err := validateAdmins(p.ExchangeAdmins); err != nil {
		return fmt.Errorf("ExchangeAdmins is incorrect: %w", err)
	}
	return nil
}

func validateSpotMarketInstantListingFee(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !v.IsValid() || !v.Amount.IsPositive() {
		return fmt.Errorf("invalid SpotMarketInstantListingFee: %T", i)
	}

	return nil
}

func validateDerivativeMarketInstantListingFee(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !v.IsValid() || !v.Amount.IsPositive() {
		return fmt.Errorf("invalid DerivativeMarketInstantListingFee: %T", i)
	}

	return nil
}

func validateBinaryOptionsMarketInstantListingFee(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !v.IsValid() || !v.Amount.IsPositive() {
		return fmt.Errorf("invalid BinaryOptionsMarketInstantListingFee: %T", i)
	}

	return nil
}

func ValidateFee(i interface{}) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("exchange fee cannot be nil: %s", v)
	}

	if v.IsNegative() {
		return fmt.Errorf("exchange fee cannot be negative: %s", v)
	}
	if v.GT(math.LegacyOneDec()) {
		return fmt.Errorf("exchange fee cannot be greater than 1: %s", v)
	}

	return nil
}

func ValidateMakerFee(i interface{}) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("exchange fee cannot be nil: %s", v)
	}

	if v.GT(math.LegacyOneDec()) {
		return fmt.Errorf("exchange fee cannot be greater than 1: %s", v)
	}

	if v.LT(math.LegacyOneDec().Neg()) {
		return fmt.Errorf("exchange fee cannot be less than -1: %s", v)
	}

	return nil
}

func ValidateHourlyFundingRateCap(i interface{}) error {
	v, ok := i.(math.LegacyDec)

	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("hourly funding rate cap cannot be nil: %s", v)
	}

	if v.IsNegative() {
		return fmt.Errorf("hourly funding rate cap cannot be negative: %s", v)
	}

	if v.IsZero() {
		return fmt.Errorf("hourly funding rate cap cannot be zero: %s", v)
	}

	if v.GT(math.LegacyNewDecWithPrec(3, 2)) {
		return fmt.Errorf("hourly funding rate cap cannot be larger than 3 percent: %s", v)
	}

	return nil
}

func ValidateHourlyInterestRate(i interface{}) error {
	v, ok := i.(math.LegacyDec)

	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("hourly interest rate cannot be nil: %s", v)
	}

	if v.IsNegative() {
		return fmt.Errorf("hourly interest rate cannot be negative: %s", v)
	}

	if v.GT(math.LegacyNewDecWithPrec(1, 2)) {
		return fmt.Errorf("hourly interest rate cannot be larger than 1 percent: %s", v)
	}

	return nil
}

func ValidateTickSize(i interface{}) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("tick size cannot be nil: %s", v)
	}

	if v.IsNegative() {
		return fmt.Errorf("tick size cannot be negative: %s", v)
	}

	if v.IsZero() {
		return fmt.Errorf("tick size cannot be zero: %s", v)
	}

	if v.GT(MaxOrderPrice) {
		return fmt.Errorf("unsupported tick size amount")
	}

	// 1e18 scaleFactor
	scaleFactor := math.LegacyNewDec(1000000000000000000)
	// v can be a decimal (e.g. 1e-18) so we scale by 1e18
	scaledValue := v.Mul(scaleFactor)

	power := math.LegacyNewDec(1)
	ten := math.LegacyNewDec(10)

	// determine whether scaledValue is a power of 10
	for power.LT(scaledValue) {
		power = power.Mul(ten)
	}

	if !power.Equal(scaledValue) {
		return fmt.Errorf("unsupported tick size")
	}

	return nil
}

func ValidateMinNotional(i interface{}) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("min notional cannot be nil")
	}

	if v.IsNegative() {
		return fmt.Errorf("min notional cannot be negative: %s", v)
	}

	return nil
}

func ValidateMarginRatio(i interface{}) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("margin ratio cannot be nil: %s", v)
	}
	if v.LT(minMarginRatio) {
		return fmt.Errorf("margin ratio cannot be less than minimum: %s", v)
	}
	if v.GTE(math.LegacyOneDec()) {
		return fmt.Errorf("margin ratio cannot be greater than or equal to 1: %s", v)
	}

	return nil
}

func validateFundingInterval(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("fundingInterval must be positive: %d", v)
	}

	return nil
}

func validatePostOnlyModeHeightThreshold(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < 0 {
		return fmt.Errorf("postOnlyModeHeightThreshold must be non-negative: %d", v)
	}

	return nil
}

func validateAdmins(i interface{}) error {
	v, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	admins := make(map[string]struct{})

	for _, admin := range v {
		adminAddr, err := sdk.AccAddressFromBech32(admin)
		if err != nil {
			return fmt.Errorf("invalid admin address: %s", admin)
		}

		if _, found := admins[adminAddr.String()]; found {
			return fmt.Errorf("duplicate admin: %s", admin)
		}
		admins[adminAddr.String()] = struct{}{}
	}

	return nil
}

func validateFundingMultiple(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("fundingMultiple must be positive: %d", v)
	}

	return nil
}

func validateDerivativeOrderSideCount(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("DerivativeOrderSideCount must be positive: %d", v)
	}

	return nil
}

func validateInjRewardStakedRequirementThreshold(i interface{}) error {
	v, ok := i.(math.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsZero() {
		return fmt.Errorf("InjRewardStakedRequirementThreshold cannot be zero: %d", v)
	}

	if v.IsNegative() {
		return fmt.Errorf("InjRewardStakedRequirementThreshold cannot be negative: %d", v)
	}

	return nil
}

func validateTradingRewardsVestingDuration(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < 0 {
		return fmt.Errorf("trading rewards vesting duration must be non-negative: %d", v)
	}

	return nil
}

func validateLiquidatorRewardShareRate(i interface{}) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("reward ratio cannot be nil: %s", v)
	}
	if v.IsNegative() {
		return fmt.Errorf("reward ratio cannot be negative: %s", v)
	}
	if v.GT(math.LegacyOneDec()) {
		return fmt.Errorf("reward ratio cannot be greater than 1: %s", v)
	}

	return nil
}

func validateAtomicMarketOrderAccessLevel(i interface{}) error {
	v, ok := i.(AtomicMarketOrderAccessLevel)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if !v.IsValid() {
		return fmt.Errorf("invalid AtomicMarketOrderAccessLevel value: %v", v)
	}
	return nil
}

func validateAtomicMarketOrderFeeMultiplier(i interface{}) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("atomicMarketOrderFeeMultiplier cannot be nil: %s", v)
	}
	if v.LT(math.LegacyOneDec()) {
		return fmt.Errorf("atomicMarketOrderFeeMultiplier cannot be less than one: %s", v)
	}
	if v.GT(MaxFeeMultiplier) {
		return fmt.Errorf("atomicMarketOrderFeeMultiplier cannot be bigger than %v: %v", v, MaxFeeMultiplier)
	}
	return nil
}

func validateBool(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}
