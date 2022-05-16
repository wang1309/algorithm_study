package greedy

import (
	"fmt"
	"testing"
)

// 分糖果测试
func TestCandy(t *testing.T)  {
	ratings := []int{1,2,3,2,1}
	fmt.Println(candy(ratings))

	ratings = []int{1,2,87,87,87,2,1}
	fmt.Println(candy(ratings))
}
