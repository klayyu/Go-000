package main

import(
	"fmt"
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"
	"net/http"
	"golang.org/x/sync/errgroup"
	"time"
)

func main()  {
	fmt.Println("hello")

	var stopSignal chan struct{}
	g,ctx := errgroup.WithContext(context.BackGround())

	g.Go(func() error{
		server := http.Server{
			Addr:":8080",
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second			
		}
		go func(){
			<-ctx.Done()
			fmt.Println("http server 8080 ctx done")
			if err := server.Shutdown(context.BackGround());err != nil{
				fmt.Println("http server 8080 shutdown")
			}
			stopSignal<- struct{}{}
		}()
		return server.ListenAndServe()
	})

	g.Go(func() error{
		server := http.Server{
			Addr:":8081",
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second			
		}
		go func(){
			<-ctx.Done()
			fmt.Println("http server 8081 ctx done")
			if err := server.Shutdown(context.BackGround());err != nil{
				fmt.Println("http server 8081 shutdown")
			}
			stopSignal<- struct{}{}
		}()
		return server.ListenAndServe()
	})

	g.Go(func() error{
		signals := make(chan os.Signal,1)
		signal.Notify(signals,syscall.SIGINT,syscall.SIGTERM)
		select{
		case <-signals:
			fmt.Println("receive stop signal")
			return errors.New("receive stop signal")
		case <-ctx.Done():
			fmt.Println("ctx done")
			return ctx.Err()
		}
	})

	fmt.Println("main thread")
	if err := g.Wait(); err != nil{
		fmt.Println("errgroup wati ",err.Error())
	}

	<-stopSignal
}