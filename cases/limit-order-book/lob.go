package main

import (
	"fmt"

	"github.com/HuKeping/rbtree"
)

// Note: It is assumed that all incoming orders are already serialized by a ring buffer
// Will attempt to implement the Ring Buffer next, it won't be that hard, right?

type Order struct {
	ID       int64
	Quantity int64
	Prev     *Order // Doubly Linked List pointers
	Next     *Order
	Level    *PriceLevel // The price level this order belongs to
}

// PriceLevel represents a specific price (e.g., 10,500 IDR)
// and holds the queue of orders at that price.
type PriceLevel struct {
	Price       int64
	TotalVolume int64 // Sum of all order quantities here

	Head *Order // First order in line (oldest)
	Tail *Order // Last order in line (newest)
}

// OrderBook is the master struct for a single trading pair (e.g., BBCA-IDR)
type OrderBook struct {
	// Fast lookup for cancellations
	ActiveOrders map[int64]*Order

	// Sorted Price Levels (using a generic Red-Black tree library)
	// Bids = Buyers (sorted descending, highest price first)
	Bids *rbtree.Rbtree

	// Asks = Sellers (sorted ascending, lowest price first)
	Asks *rbtree.Rbtree
}

// Note that this func reduce the total volume as it literally removes the whole Order
func (pl *PriceLevel) PopHead() error {
	if pl.Head == nil {
		return fmt.Errorf("Head in nil")
	}
	pl.TotalVolume -= pl.Head.Quantity
	if pl.Head.Next != nil {
		pl.Head.Next.Prev = nil
	}
	pl.Head = pl.Head.Next
	return nil
}

func (pl *PriceLevel) ReduceHead(qty int64) error {
	if qty > pl.Head.Quantity {
		return fmt.Errorf("Qty in ReduceHead is more than qty available")
	}

	pl.Head.Quantity -= qty
	pl.TotalVolume -= qty

	if pl.Head.Quantity == 0 {
		if err := pl.PopHead(); err != nil {
			return err
		}
	}
	return nil
}

// GetBestAsk returns the PriceLevel with the LOWEST selling price.
// If there are no sellers, it returns nil.
// If a PriceLevel is emptied, calling this again will return the NEXT lowest price.
func (ob *OrderBook) GetBestAsk() *PriceLevel {
	// Assume this is fully implemented and O(1)
	return nil
}

// RemoveBestAsk completely deletes the current best ask level from the order book trees.
func (ob *OrderBook) RemoveBestAsk() {
	// Assume this is fully implemented and O(log N)
}

// ExecuteMarketBuy attempts to buy 'quantity' shares at the best available prices.
// It returns the total cost of the purchase, and the remaining un-bought quantity
// (which should be 0 if fully filled, or >0 if the order book ran out of sellers).
func (ob *OrderBook) ExecuteMarketBuy(quantity int64) (totalCost int64, unfulfilled int64) {

	for ob.GetBestAsk() != nil && quantity > 0 {
		curPriceLevel := ob.GetBestAsk()
		for curPriceLevel.TotalVolume > 0 && quantity > 0 {
			curOrderToBuy := curPriceLevel.Head
			qtyToBuy := min(curOrderToBuy.Quantity, quantity)

			quantity -= qtyToBuy
			curPriceLevel.ReduceHead(qtyToBuy)

			totalCost += qtyToBuy * curPriceLevel.Price
		}
		if curPriceLevel.TotalVolume == 0 {
			ob.RemoveBestAsk()
		}
	}

	return totalCost, quantity
}

// CancelOrder removes an order from the order book in O(1) time
// (excluding potential O(log N) tree cleanup if the price level empties).
func (ob *OrderBook) CancelOrder(orderID int64) error {
	targetOrder, exists := ob.ActiveOrders[orderID]
	if !exists {
		return fmt.Errorf("Order does not exists")
	}

	prevOrder := targetOrder.Prev
	nextOrder := targetOrder.Next
	priceLevel := targetOrder.Level

	if prevOrder != nil {
		prevOrder.Next = nextOrder
	} else {
		priceLevel.Head = nextOrder
	}
	if nextOrder != nil {
		nextOrder.Prev = prevOrder
	} else {
		priceLevel.Tail = prevOrder
	}

	priceLevel.TotalVolume -= targetOrder.Quantity
	delete(ob.ActiveOrders, targetOrder.ID)

	return nil
}
