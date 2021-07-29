package best_time_to_buy_and_sell_stock

import "testing"

func testCase(t *testing.T, prices []int, output int) {
	result := MaxProfit(prices)
	if result != output {
		t.Errorf("want: %v, got: %v\n", output, result)
	}
}

func TestMaxProfit(t *testing.T) {
	testCase(t, []int { 7,1,5,3,6,4 }, 5)
	testCase(t, []int { 7,6,4,3,1 }, 0)
}
