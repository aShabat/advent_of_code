package util

func InBounds[T any](nums [][]T, row, col int) bool {
	return row >= 0 && row < len(nums) && col >= 0 && col < len(nums[row])
}
