// Code generated from OpenSoha contracts. DO NOT EDIT.
package sohaapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const defaultUserAgent = "opensoha-contracts/0.1"

type Client struct {
	BaseURL    string
	Token      string
	HTTPClient *http.Client
	UserAgent  string
}

type ClientOption func(*Client)

func WithBearerToken(token string) ClientOption {
	return func(c *Client) {
		c.Token = strings.TrimSpace(token)
	}
}

func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) {
		c.HTTPClient = client
	}
}

func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) {
		c.UserAgent = strings.TrimSpace(userAgent)
	}
}

func NewClient(baseURL string, opts ...ClientOption) *Client {
	c := &Client{
		BaseURL:   strings.TrimSpace(baseURL),
		UserAgent: defaultUserAgent,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(c)
		}
	}
	return c
}

func (c *Client) ClaimExecutionTask(ctx context.Context, req ExecutionTaskClaimRequest) (ExecutionTask, error) {
	var out ExecutionTaskEnvelope
	if err := c.doJSON(ctx, http.MethodPost, "/delivery/execution-tasks/claim", true, req, &out); err != nil {
		return ExecutionTask{}, err
	}
	return out.Data, nil
}

func (c *Client) GetExecutionTaskRunnerStatus(ctx context.Context, taskID string) (ExecutionTask, error) {
	var out ExecutionTaskEnvelope
	path := "/delivery/execution-tasks/" + url.PathEscape(strings.TrimSpace(taskID)) + "/runner-status"
	if err := c.doJSON(ctx, http.MethodGet, path, true, nil, &out); err != nil {
		return ExecutionTask{}, err
	}
	return out.Data, nil
}

func (c *Client) RecordExecutionCallback(ctx context.Context, req ExecutionCallbackRequest) (ExecutionTask, error) {
	var out ExecutionTaskEnvelope
	if err := c.doJSON(ctx, http.MethodPost, "/delivery/execution-callbacks", false, req, &out); err != nil {
		return ExecutionTask{}, err
	}
	return out.Data, nil
}

func (c *Client) ClaimDockerOperation(ctx context.Context, req DockerOperationClaimRequest) (DockerOperation, error) {
	var out DockerOperationEnvelope
	if err := c.doJSON(ctx, http.MethodPost, "/docker/operations/claim", true, req, &out); err != nil {
		return DockerOperation{}, err
	}
	return out.Data, nil
}

func (c *Client) GetDockerOperationRunnerStatus(ctx context.Context, operationID string) (DockerOperation, error) {
	var out DockerOperationEnvelope
	path := "/docker/operations/" + url.PathEscape(strings.TrimSpace(operationID)) + "/runner-status"
	if err := c.doJSON(ctx, http.MethodGet, path, true, nil, &out); err != nil {
		return DockerOperation{}, err
	}
	return out.Data, nil
}

func (c *Client) RecordDockerOperationCallback(ctx context.Context, req DockerOperationCallbackRequest) (DockerOperation, error) {
	var out DockerOperationEnvelope
	if err := c.doJSON(ctx, http.MethodPost, "/docker/operation-callbacks", true, req, &out); err != nil {
		return DockerOperation{}, err
	}
	return out.Data, nil
}

func (c *Client) ClaimAgentRun(ctx context.Context, req AgentRunClaimRequest) (AgentRun, error) {
	var out AgentRunEnvelope
	if err := c.doJSON(ctx, http.MethodPost, "/copilot/agent-runs/claim", true, req, &out); err != nil {
		return AgentRun{}, err
	}
	return out.Data, nil
}

func (c *Client) RecordAgentRunCallback(ctx context.Context, req AgentRunCallbackRequest) (AgentRun, error) {
	var out AgentRunEnvelope
	if err := c.doJSON(ctx, http.MethodPost, "/copilot/agent-runs/callback", true, req, &out); err != nil {
		return AgentRun{}, err
	}
	return out.Data, nil
}

func (c *Client) RecordAgentRunToolCall(ctx context.Context, req AgentRunToolCallRequest) (AgentToolCallResult, error) {
	var out AgentToolCallResultEnvelope
	if err := c.doJSON(ctx, http.MethodPost, "/copilot/agent-runs/tool-call", true, req, &out); err != nil {
		return AgentToolCallResult{}, err
	}
	return out.Data, nil
}

func (c *Client) doJSON(ctx context.Context, method, path string, useToken bool, body any, out any) error {
	endpoint, err := c.endpoint(path)
	if err != nil {
		return err
	}
	var reader io.Reader
	if body != nil {
		raw, err := json.Marshal(body)
		if err != nil {
			return err
		}
		reader = bytes.NewReader(raw)
	}
	req, err := http.NewRequestWithContext(ctx, method, endpoint, reader)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if useToken && strings.TrimSpace(c.Token) != "" {
		req.Header.Set("Authorization", "Bearer "+strings.TrimSpace(c.Token))
	}
	client := c.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	raw, err := io.ReadAll(io.LimitReader(resp.Body, 10<<20))
	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		return fmt.Errorf("%s %s failed: %s: %s", method, path, resp.Status, responseErrorMessage(raw))
	}
	if out == nil || len(raw) == 0 {
		return nil
	}
	if err := json.Unmarshal(raw, out); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}
	return nil
}

func (c *Client) endpoint(path string) (string, error) {
	base := strings.TrimRight(strings.TrimSpace(c.BaseURL), "/")
	if base == "" {
		return "", fmt.Errorf("base URL is required")
	}
	if strings.HasSuffix(base, "/api/v1") {
		base = strings.TrimSuffix(base, "/api/v1")
	}
	path = "/" + strings.TrimLeft(path, "/")
	return base + "/api/v1" + path, nil
}

func responseErrorMessage(raw []byte) string {
	var wrapped ErrorEnvelope
	if json.Unmarshal(raw, &wrapped) == nil && wrapped.Error.Message != "" {
		if wrapped.Error.Code != "" {
			return wrapped.Error.Code + ": " + wrapped.Error.Message
		}
		return wrapped.Error.Message
	}
	return strings.TrimSpace(string(raw))
}
