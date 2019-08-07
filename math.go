package tools

// http://cavaliercoder.com/blog/optimized-abs-for-int64-in-go.html
func INT64Abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}
