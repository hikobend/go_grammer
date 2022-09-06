package foo

// スコープ

const (
	// 先頭が大文字 : 参照可能
	Max = 100
	// 先頭が小文字 : 参照不可
	min = 1
)

func ReturnMin() int {
	return min
}
