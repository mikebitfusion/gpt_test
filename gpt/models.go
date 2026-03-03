package gpt

import (
	"sort"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

type Msg struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Choice struct {
	Index        int    `json:"index"`
	Messages     Msg    `json:"message"`
	FinishReason string `json:"finish_reason"`
}

type Response struct {
	Choices []Choice `json:"choices"`
}

type MetaAnswer struct {
	Description string `json:"description"`
}

type StrategyRequest struct {
	StrategyID     string   `json:"strategy_id"`
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	YieldType      string   `json:"yield_type"`
	RateFormat     string   `json:"rate_format"`
	CategoryIDs    []string `json:"category_ids"`
	ChainIDs       []string `json:"chain_ids"`
	InputTokenIDs  []string `json:"input_tokens_ids"`
	OutputTokenIDs []string `json:"output_tokens_ids"`
	Status         string   `json:"status"`
	Limits         Limits   `json:"limits"`
	RoutingEngine  string   `json:"routing_engine"`
	Fees           Fees     `json:"fee_breakdown"`
	MarketsIDs     []string `json:"market_ids"` // upd new
	Risk           string   `json:"risk"`       // upd new
	// ExecutionConfig ExecutionConfig `json:"execution_config"` // composed at pre repo call
	// Metrics         Metrics         `json:"metrics"` // gonna be added by the auto update script
	// CreatedAt int64 `json:"created_at"`
	// UpdatedAt int64 `json:"updated_at"`
}

func (s StrategyRequest) Slice() []string {
	sort.Strings(s.CategoryIDs)
	sort.Strings(s.ChainIDs)
	sort.Strings(s.InputTokenIDs)
	sort.Strings(s.OutputTokenIDs)
	sort.Strings(s.MarketsIDs)

	out := []string{
		s.StrategyID,
		s.Name,
		s.Description,
		string(s.YieldType),
		string(s.RateFormat),
		strings.Join(s.CategoryIDs, ","),
		strings.Join(s.ChainIDs, ","),
		strings.Join(s.InputTokenIDs, ","),
		strings.Join(s.OutputTokenIDs, ","),
		string(s.Status),
		s.Limits.String(),
		s.RoutingEngine,
		s.Fees.String(),
		strings.Join(s.MarketsIDs, ","),
		string(s.Risk),
	}

	return out
}

type Limits struct {
	DepositsPaused    bool `json:"deposits_paused"`
	WithdrawalsPaused bool `json:"withdrawals_paused"`
}

func (l Limits) String() string {
	return strings.Join([]string{
		strconv.FormatBool(l.DepositsPaused),
		strconv.FormatBool(l.WithdrawalsPaused),
	}, ",")
}

// check fee in param from whom it came
type Fees struct {
	TotalFee   decimal.Decimal `json:"total_fee"`
	YIFIFee    decimal.Decimal `json:"yifi_fee"`
	PartnerFee decimal.Decimal `json:"partner_fee"`
}

func (f Fees) String() string {
	return strings.Join([]string{
		f.TotalFee.String(),
		f.YIFIFee.String(),
		f.PartnerFee.String(),
	}, ",")
}
