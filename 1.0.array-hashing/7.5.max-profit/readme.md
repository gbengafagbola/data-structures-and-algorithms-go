Sure! Here's a beginner-friendly, visually enhanced **README** for the Go solution to the *Best Time to Buy and Sell Stock* problem:

---

# 📈 Best Time to Buy and Sell Stock

## 🧩 Problem Statement

You're given a list of **daily stock prices** (e.g., `prices[i]` is the stock price on day `i`).
Your task is to **find the best single day to buy** and **a later day to sell** to achieve the **maximum profit**.

⚠️ You **can only buy and sell once**.

---

## 🧪 Example

```go
Input:  prices = [7, 1, 5, 3, 6, 4]
Output: Buy on day 1 and sell on day 4 for a profit of 5
```

**Explanation**:
Buy at `1` (day 1) and sell at `6` (day 4): `6 - 1 = 5` profit.
This is the best possible profit in one transaction.

---

## 🧠 Approach

We traverse the prices once, tracking:

* The **minimum price** so far (`minPrice`) — a candidate for buying
* The **maximum profit** so far, based on selling at current day and buying at `minPrice`

### 🧮 Key Idea

At each day `i`, we:

1. Check if `prices[i] - minPrice` gives us more profit than previously seen
2. If so, update `maxProfit`, `buyDay`, and `sellDay`
3. If `prices[i]` is lower than `minPrice`, we update `minPrice` and its index

---

## 🔁 Algorithm Flow (Illustration)

Let’s walk through:
`prices = [7, 1, 5, 3, 6, 4]`

```
Index (Day):    0   1   2   3   4   5
Price:          7   1   5   3   6   4
                 ↓   ↑   ↑   ↑   ↑   ↑
minPrice        7   1   1   1   1   1
maxProfit       0   0   4   4   5   5
buyDay          -   1   1   1   1   1
sellDay         -   -   2   2   4   4
```

* **Day 1**: New minimum found (`1`)
* **Day 2**: Potential profit `5 - 1 = 4` → update profit and days
* **Day 4**: Profit `6 - 1 = 5` → best profit so far

✅ Final answer: Buy on **day 1** (price = `1`), sell on **day 4** (price = `6`), **profit = 5**

---

## 🧾 Code (Go)

```go
func MaxProfit(prices []int) (buyDay, sellDay, maxProfit int) {
	if len(prices) < 2 {
		return -1, -1, 0 // Not enough data
	}

	minPrice := prices[0]
	minDay := 0
	maxProfit = 0

	for i := 1; i < len(prices); i++ {
		profit := prices[i] - minPrice
		if profit > maxProfit {
			maxProfit = profit
			buyDay = minDay
			sellDay = i
		}
		if prices[i] < minPrice {
			minPrice = prices[i]
			minDay = i
		}
	}

	if maxProfit == 0 {
		return -1, -1, 0 // No profit possible
	}
	return buyDay, sellDay, maxProfit
}
```

---

## 🧰 Time and Space Complexity

| Metric | Value |
| ------ | ----- |
| Time   | O(n)  |
| Space  | O(1)  |

---

## ✅ Edge Cases

* `[1, 2, 3, 4, 5]` → Profit = 4 (Buy day 0, sell day 4)
* `[5, 4, 3, 2, 1]` → No profit possible → returns `-1, -1, 0`
* `[]` or `[7]` → Not enough data to make a transaction

---

## ✨ Optional Improvement

If you want to return **prices instead of days**:

```go
return prices[buyDay], prices[sellDay], maxProfit
```

---

Let me know if you want this in a `.md` file or need a version with ASCII graphs!
