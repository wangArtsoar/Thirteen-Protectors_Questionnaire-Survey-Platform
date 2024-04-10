package main

import (
	"context"
	"time"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func BookList(pageNum, pageSize int) []*Book {
	context.WithDeadline(context.Background(), time.Now())
	return nil
}
