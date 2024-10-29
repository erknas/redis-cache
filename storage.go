package main

import (
	"context"
	"fmt"
)

type Storage struct {
	data  map[string]string
	cache Cacher
}

func NewStorage(c Cacher) *Storage {
	data := map[string]string{
		"1": "something",
		"2": "storage data",
		"3": "test",
	}

	return &Storage{
		data:  data,
		cache: c,
	}
}

func (s *Storage) Get(ctx context.Context, key string) (string, error) {
	val, ok := s.cache.Get(ctx, key)
	if ok {
		if err := s.cache.Remove(ctx, key); err != nil {
			fmt.Println(err)
		}
		fmt.Println("returning from cache")
		return val, nil
	}

	val, ok = s.data[key]
	if !ok {
		return "", fmt.Errorf("nothing found for (%s)", key)
	}

	if err := s.cache.Set(ctx, key, val); err != nil {
		return "", err
	}

	fmt.Println("returning from storage")

	return val, nil
}
