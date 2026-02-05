package main

import "fmt"

const (
	// BPSScale represents 100.00%
	// 10000 bps = 100%
	BPSScale = 10000
)

// CalculateTotal accepts ONLY integers.
// amount: stored in minor units (cents). Rp 100.000 -> 100000
// discountBPS: 10% -> 1000 bps (0.10 * 10000)
func CalculateTotal(amount int64, qty int, discountBPS int64) int64 {
	// 1. Calculate Gross (All integer math)
	gross := amount * int64(qty)

	// 2. Calculate Discount Amount
	// Formula: (Gross * BPS) / 10000
	discountAmount := (gross * discountBPS) / BPSScale

	// 3. Return Net (Integer)
	return gross - discountAmount
}

func main() {
	// Scenario: Rp 100.000, 3 items, 10% discount

	// Input is defined as INT from the start
	amount := int64(100000)
	qty := 3

	// 10% converted to BPS manually (10 * 100) or received as 1000 from frontend
	discountBPS := int64(1000)

	result := CalculateTotal(amount, qty, discountBPS)

	fmt.Printf("Total: %d\n", result) // Output: 270000
}
