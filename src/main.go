package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/mahiro72/todo-task-app/config"
	"golang.org/x/sync/errgroup"
)

func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	s := &http.Server{
		Addr: fmt.Sprintf(":%d",cfg.Port),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w,"Hello, %s!",r.URL.Path[1:])
		}),
	}

	eg, ctx := errgroup.WithContext(ctx)

	// 別ゴルーチンでサーバーを起動する
	eg.Go(func() error {
		if err:=s.ListenAndServe();err!=nil&&err!=http.ErrServerClosed{
			log.Printf("failed to close: %v",err)
			return err
		}
		return nil
	})

	// チャネルからの終了通知を待機する
	<-ctx.Done()
	if err:=s.Shutdown(context.Background());err!=nil{
		log.Printf("failed to shutdown: %v",err)
	}

	// goメソッドで起動した別ゴルーチンの終了を待つ
	return eg.Wait()
}

func main(){
	run(context.Background())
}