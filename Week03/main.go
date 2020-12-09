package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
)

func main() {
	servers := []http.Server{
		{Addr: "8090"},
		{Addr: "8091"},
		{Addr: "8092"},
	}

	group, _ := errgroup.WithContext(context.Background())

	stopChan := make(chan struct{}, 1)
	for i := 0; i < len(servers); i++ {
		i := i
		group.Go(func() (err error) {
			go func() {
				<-stopChan
				_ = servers[i].Shutdown(context.Background())
			}()
			err = servers[i].ListenAndServe()
			if err != nil {
				return
			}

			return nil
		})
	}

	go func() {
		group.Wait()
		close(stopChan)
	}()

	signalChan := make(chan os.Signal, 1)

	select {
	case <-signalChan:
		stopChan <- struct{}{}
		return

	}

}
