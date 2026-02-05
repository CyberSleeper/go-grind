## Problem 1: The "No Floats" Rule (Fintech Basics)

### 📌 Concept
In financial software (Fintech), **never use floating-point numbers (`float64`, `float32`) for monetary calculations**. Computers store floats using binary fractions (IEEE 754), which leads to precision errors (e.g., `0.1 + 0.2 != 0.3`).

### 📝 Task
Write a function `CalculateTotal` that accurately computes the final price using **integer math** (representing cents/minor units).

**Formula:** `Total = Amount * Quantity * (1 - DiscountRate)`

**Inputs:**
* `amount`: Rp 100,000
* `qty`: 3
* `discount`: 10%

### 🚫 Starter Code (The Buggy Version)
```go
func CalculateTotal(amount float64, qty int, discount float64) float64 {
    // BUG: Precision loss happens here
    return amount * float64(qty) * (1 - discount)
}
```