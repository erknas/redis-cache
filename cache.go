package main

import "context"

type Cacher interface {
	Get(context.Context, string) (string, bool)
	Set(context.Context, string, string) error
	Remove(context.Context, string) error
}

type Cache struct{}

func (c Cache) Get(context.Context, string) (string, bool) {
	return "", false
}

func (c Cache) Set(context.Context, string, string) error {
	return nil
}

func (c Cache) Remove(context.Context, string) error {
	return nil
}
