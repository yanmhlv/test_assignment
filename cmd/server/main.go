package main

import (
	"context"
	"encoding/json"
	"flag"
	"net/http"
	"time"

	"github.com/chapsuk/grace"
	"github.com/chapsuk/wait"
	"github.com/julienschmidt/httprouter"

	"github.com/yanmhlv/test_assignment/app"
)

var (
	flagAddr            = flag.String("addr", ":3000", "")
	flagCleanupInterval = flag.Duration("cleanup-interval", 10*time.Second, "")
	flagStackSize       = flag.Int("stack-size", 100, "")
	flagInitialBalance  = flag.Float64("balance", 0, "")
)

func main() {
	flag.Parse()

	var (
		ctx  = grace.ShutdownContext(context.Background())
		game = app.NewGame(*flagInitialBalance, *flagStackSize, *flagCleanupInterval)
		wg   = wait.Group{}
	)

	wg.Add(func() { game.Cleanup(ctx) })

	router := httprouter.New()
	router.POST("/game", func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		p := app.GameRequest{}

		defer req.Body.Close()

		if err := json.NewDecoder(req.Body).Decode(&p); err != nil {
			http.Error(w, "", http.StatusBadRequest)
		}
		app.GetLogger().Infow("start game", "data", p)

		amount, gameStatus := game.Play(p.Pair, p.Fee)
		app.GetLogger().Infow("finish game", "result", gameStatus, "amount", amount)

		json.NewEncoder(w).Encode(&app.GameResponse{gameStatus, amount})
	})

	server := &http.Server{
		Handler: router,
		Addr:    *flagAddr,
	}

	wg.Add(func() {
		<-ctx.Done()
		server.Shutdown(ctx)
	})

	wg.Add(func() {
		app.GetLogger().Infow("server is starting")
		server.ListenAndServe()
	})

	wg.Wait()
}
