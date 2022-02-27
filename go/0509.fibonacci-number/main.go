package main

func fib(n int) int {
	if n == 0 {
		return 0
	}
	dp_i_1, dp_i_2 := 1, 0
	for i := 2; i <= n; i++ {
		dp_i := dp_i_1 + dp_i_2
		dp_i_2, dp_i_1 = dp_i_1, dp_i
	}
	return dp_i_1
}
