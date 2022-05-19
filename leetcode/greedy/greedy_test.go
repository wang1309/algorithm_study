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

// 移掉 K 位数字
func TestRemoveKdigits(t *testing.T) {
	fmt.Println(removeKdigits("3210", 2))
}

// 跳跃游戏
func TestCanJump(t *testing.T) {
	jumps := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(jumps))
}
