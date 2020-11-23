package services

import (
	"context"
	"fmt"
	"time"
)

const delay = 500

type useless string

func NewUseless(s string) *useless {
	u := useless(s)
	return &u
}

func (u *useless) Start(ctx context.Context) error {
	fmt.Printf("U: Starting useless [%v]",u)
	time.Sleep(delay * time.Millisecond)
	fmt.Printf("U: Started useless [%v]",u)
	return nil
}
//
//func (u *useless) HealthCheck(ctx context.Context) error {
//	fmt.Printf("U: Healthchecking useless [%v]",u)
//	return nil
//}

func (u *useless) Stop(ctx context.Context) error {
	fmt.Printf("U: Stopping useless [%v]",u)
	time.Sleep(delay * time.Millisecond)
	fmt.Printf("U: Stopped useless [%v]",u)
	return nil
}