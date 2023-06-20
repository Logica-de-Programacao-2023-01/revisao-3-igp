package q3

func ClimbStairs(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {

		P1, P2 := 1, 1

		for i := 2; i <= n; i++ {
			P := P1 + P2
			P1 = P2
			P2 = P
		}
		return P2
	}
}
