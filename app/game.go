package app

import (
	"context"
	"sync"
	"time"
)

type Game struct {
	LuckyStack      *LuckyStack
	Balance         float64
	LastGame        time.Time
	CleanupInterval time.Duration
	lock            sync.Mutex
}

func NewGame(initialBalance float64, limit int, cleanupInterval time.Duration) *Game {
	return &Game{
		LuckyStack:      &LuckyStack{&Stack{}, limit},
		Balance:         initialBalance,
		LastGame:        time.Time{},
		CleanupInterval: cleanupInterval,
		lock:            sync.Mutex{},
	}
}

func (g *Game) Play(userPair Pair, userFee float64) (float64, GameStatus) {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.LastGame = time.Now()

	pair := g.LuckyStack.Pop()

	g.LuckyStack.Push(userPair)
	g.Balance += userFee

	switch {
	case pair == nil:
		return 0, FAIL_GAME
	case *pair == userPair:
		balance := g.Balance
		if balance > 0 {
			g.Balance = 0
			return balance, WIN_GAME
		}
		return 0, WIN_FREE_GAME
	}
	return 0, FAIL_GAME
}

func (g *Game) Cleanup(ctx context.Context) {
	ticker := time.NewTicker(g.CleanupInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case now := <-ticker.C:
			g.lock.Lock()
			if now.Sub(g.LastGame) > g.CleanupInterval {
				pair := g.LuckyStack.Pop()
				GetLogger().Infow("cleanup", "now", now, "pair", pair)
			}
			g.lock.Unlock()
		}
	}
}
