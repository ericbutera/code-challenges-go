package stockstreamer_test

import (
	"fmt"
	"testing"
	"time"
)

// Mocked stock price streamer
func mockStockPriceStreamer(prices chan<- float64, values []float64) {
	for _, price := range values {
		prices <- price // Send predefined prices
	}
	close(prices)
}

// Worker logic
func testWorker(id int, prices <-chan float64, alerts chan<- string) {
	for price := range prices {
		if price > 900 { // Example threshold
			alerts <- fmt.Sprintf("Worker %d: High price detected $%.2f", id, price)
		}
	}
}

func TestStockPriceStreamerAndWorkers(t *testing.T) {
	t.Parallel()

	stockPrices := []float64{500.0, 950.5, 875.3, 920.7, 899.9}
	expectedAlerts := []string{
		"Worker 1: High price detected $950.50",
		"Worker 1: High price detected $920.70",
	}

	prices := make(chan float64)
	alerts := make(chan string)

	// Start the worker
	go testWorker(1, prices, alerts)

	// Mock the stock price streamer
	go mockStockPriceStreamer(prices, stockPrices)

	// Collect alerts
	var actualAlerts []string
	done := make(chan struct{})
	go func() {
		for alert := range alerts {
			actualAlerts = append(actualAlerts, alert)
		}
		close(done)
	}()

	// Allow time for processing
	time.Sleep(time.Millisecond * 500)
	close(alerts)

	// Wait for the alert collection goroutine to finish
	<-done

	// Validate results
	if len(actualAlerts) != len(expectedAlerts) {
		t.Fatalf("expected %d alerts, got %d", len(expectedAlerts), len(actualAlerts))
	}

	for i, expected := range expectedAlerts {
		if actualAlerts[i] != expected {
			t.Errorf("expected alert '%s', got '%s'", expected, actualAlerts[i])
		}
	}
}
