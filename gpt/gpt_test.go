package gpt

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func TestGetDescriptionIntegration(t *testing.T) {
	t.Parallel()

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		t.Fatalf("bruh")
	}

	client := NewGPT(apiKey)

	market := StrategyRequest{
		StrategyID:  "ethereum:0xeab7b7c8353ba1cb4b29cf7ae9c166efdc57835f",
		Name:        "PT stETH 25JUN2026",
		Description: "PT stETH 25JUN2026 on pendle — fixed-income principal token (annual_rate=4.44, tvl_usd=2888489.99)",
		YieldType:   "fixed",
		RateFormat:  "apy",
		CategoryIDs: []string{"fixed-income"},
		ChainIDs:    []string{"ethereum"},
		InputTokenIDs: []string{
			"ethereum:0x57e114B691Db790C35207b2e685D4A43181e6061", // ENA
			"ethereum:0x8bE3460A480c80728a8C4D7a5D5303c85ba7B3b9", // sENA
		},
		OutputTokenIDs: []string{
			"ethereum:0x4552c668eb8dedeac53e00cfd05d873f11a80204", // PT token
		},
		Status: "active",
		Limits: Limits{
			DepositsPaused:    false,
			WithdrawalsPaused: false,
		},
		RoutingEngine: "portals",

		Fees: Fees{
			TotalFee:   decimal.RequireFromString("0.3"),
			YIFIFee:    decimal.RequireFromString("0.3"),
			PartnerFee: decimal.RequireFromString("0"),
		},

		MarketsIDs: []string{
			"ethereum:0xeab7b7c8353ba1cb4b29cf7ae9c166efdc57835f",
		},

		Risk: "medium",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	desc, err := client.GetDescription(ctx, market)
	if err != nil {
		t.Fatalf("GetDescription returned error: %v", err)
	}
	if desc == nil {
		t.Fatal("GetDescription returned nil description")
	}

	t.Logf("Assistant description: %s", *desc)

	if len(*desc) == 0 {
		t.Fatal("assistant returned empty description")
	}
}
