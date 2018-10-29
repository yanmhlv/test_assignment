package app

type LuckyStack struct {
	*Stack
	Limit int
}

func (l *LuckyStack) Push(pair Pair) bool {
	if l.Stack.count >= l.Limit {
		return false
	}
	l.Stack.Push(pair)
	return true
}
