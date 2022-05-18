package greedy

import (
	"fmt"
	"testing"
)

// 分糖果测试
func TestCandy(t *testing.T) {
	ratings := []int{1, 2, 3, 2, 1}
	fmt.Println(candy(ratings))

	ratings = []int{1, 2, 87, 87, 87, 2, 1}
	fmt.Println(candy(ratings))
}

// 买卖股票
func TestMaxProfit(t *testing.T) {
	prices := []int{1, 2, 3, 4, 5}
	fmt.Println(MaxProfit(prices))
}

func TestRemoveKdigits(t *testing.T) {
	fmt.Println(removeKdigits("3210", 2))
}
