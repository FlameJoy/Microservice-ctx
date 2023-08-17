package main

import (
	"context"
	"fmt"
	"time"
)

type LoggingService struct {
	next Service
}

func NewLoggingService(next Service) Service {
	return &LoggingService{
		next: next,
	}
}

func (s *LoggingService) GetCatFact(ctx context.Context) (fact *CatFact, err error) {
	defer func(start time.Time) {
		fmt.Printf("fact=%s err=%s took=%v\n", fact.Fact, err, time.Since(start))
	}(time.Now())
	// timeout check
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()

	respCh := make(chan *CatFact)

	go func() {
		fact, err = s.next.GetCatFact(ctx)
		respCh <- fact
	}()

	for {
		select {
		case <-ctx.Done():
			return &CatFact{Fact: "Can't get fact"}, fmt.Errorf("request is too slow (more than 200 ms)")
		case resp := <-respCh:
			return resp, nil
		}
	}
}
