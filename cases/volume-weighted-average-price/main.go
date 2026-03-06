package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/shopspring/decimal"
	"golang.org/x/sync/errgroup"
)

// OrderBook represents the top-of-book liquidity from a single exchange.
type OrderBook struct {
	ExchangeID string
	Price      string // e.g., "15000.50"
	Volume     string // e.g., "10.5"
}

type VWAP struct {
	PriceAggr decimal.Decimal
	Volume    decimal.Decimal
	mu        sync.RWMutex
}

func NewVWAP() *VWAP {
	return &VWAP{
		PriceAggr: decimal.Zero,
		Volume:    decimal.Zero,
	}
}

func (v *VWAP) Insert(o *OrderBook) error {
	price, err := decimal.NewFromString(o.Price)
	if err != nil {
		return err
	}
	volume, err := decimal.NewFromString(o.Volume)
	if err != nil {
		return err
	}

	v.mu.Lock()
	v.PriceAggr = v.PriceAggr.Add(price.Mul(volume))
	v.Volume = v.Volume.Add(volume)
	v.mu.Unlock()

	return nil
}

func (v *VWAP) Calc() string {
	v.mu.RLock()
	defer v.mu.RUnlock()

	if v.Volume.IsZero() {
		return decimal.Zero.StringFixed(8)
	}

	return v.PriceAggr.Div(v.Volume).StringFixed(8)
}

// ExchangeClient simulates an API client for a specific exchange.
type ExchangeClient interface {
	// FetchOrderBook makes a network call to get the latest liquidity.
	FetchOrderBook(ctx context.Context) (OrderBook, error)
}

func CalculateGlobalVWAP(clients []ExchangeClient) (string, error) {
	vwap := NewVWAP()

	var eg errgroup.Group
	eg.SetLimit(50)

	for _, client := range clients {
		eg.Go(func() error {
			ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
			defer cancel()

			orderBook, err := client.FetchOrderBook(ctx)
			if err != nil {
				log.Printf("Error getting from exchange %s: %v", orderBook.ExchangeID, err)
				return nil
			}
			if err := vwap.Insert(&orderBook); err != nil {
				log.Printf("Error parsing data from %s: %v", orderBook.ExchangeID, err)

			}
			return nil
		})
	}

	eg.Wait()

	return vwap.Calc(), nil
}
