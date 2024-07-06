package dice

type IMovementStrategy interface {
	GetMoves(vals []int) int
}

type SumStrategy struct{}

func (s *SumStrategy) GetMoves(vals []int) int {
	ans := 0
	for i := 0; i < len(vals); i++ {
		ans += vals[i]
	}
	return ans
}

type MinStrategy struct{}

func (m *MinStrategy) GetMoves(vals []int) int {
	ans := 7
	for i := 0; i < len(vals); i++ {
		if vals[i] < ans {
			ans = vals[i]
		}
	}
	return ans
}

type MaxStrategy struct{}

func (m *MaxStrategy) GetMoves(vals []int) int {
	ans := 0
	for i := 0; i < len(vals); i++ {
		if vals[i] > ans {
			ans = vals[i]
		}
	}
	return ans
}
