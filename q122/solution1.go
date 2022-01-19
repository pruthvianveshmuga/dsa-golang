package q122

func Solution1(prices []int) int {
    totalProfit := 0
    for i := 0; i < len(prices)-1; i++ {
        diff := prices[i+1] - prices[i]
        if diff > 0 {
            totalProfit += diff
        }
    }
    return totalProfit
}