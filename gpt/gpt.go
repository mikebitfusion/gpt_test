package gpt

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"io"
	"net/http"
)

type gpt struct {
	gptAPIKey string
}

func NewGPT(gptAPIKey string) *gpt {
	return &gpt{
		gptAPIKey: gptAPIKey,
	}
}

func (g *gpt) GetDescription(ctx context.Context, market StrategyRequest) (*string, error) {
	return g.makeRequest(ctx, String(market.Slice()))
}

func (g *gpt) makeRequest(ctx context.Context, msg string) (*string, error) {
	body := make(map[string]interface{})
	body["model"] = model
	content := fmt.Sprintf(FeedMetaAdsPrompt, msg)
	body["messages"] = []map[string]interface{}{
		{
			"role":    role,
			"content": content,
		},
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+g.gptAPIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	contents, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var gptResp Response
	err = json.Unmarshal(contents, &gptResp)
	if err != nil {
		return nil, err
	}

	unmarshalled := gptResp.Choices[0].Messages.Content
	// g.logger.Info(unmarshalled)
	var feed MetaAnswer
	err = json.Unmarshal([]byte(unmarshalled), &feed)
	if err != nil {
		return nil, err
	}
	// g.logger.Infof("clean: %+v", feed)
	return &feed.Description, nil
}
