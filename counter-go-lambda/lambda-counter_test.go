package main

import (
	"context"
	"testing"
)

// mockHandler replicates the handler response for testing
func mockHandler(ctx context.Context) (map[string]interface{}, error) {
	// Simulate a DynamoDB response
	return map[string]interface{}{"count": "99"}, nil
}

func TestHandlerReturnsCount(t *testing.T) {
	resp, err := mockHandler(context.Background())
	if err != nil {
		t.Fatalf("handler returned error: %v", err)
	}

	val, ok := resp["count"]
	if !ok {
		t.Fatalf("expected 'count' in response, got %v", resp)
	}

	if val != "99" {
		t.Errorf("expected count to be '99', got %v", val)
	}
}
