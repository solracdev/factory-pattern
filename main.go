package main

import (
	"context"
	"fmt"
)

type Recorder interface {
	Record(ctx context.Context) string
}

type Media func(ctx context.Context) string

func (m Media) Record(ctx context.Context) string {
	return m(ctx)
}

func MP3(fileName string) Recorder {
	return Media(func(ctx context.Context) string {
		return fileName + ctx.Value("extension").(string)
	})
}

type Service struct {
	r Recorder
}

func NewService(r Recorder) *Service {
	return &Service{
		r: r,
	}
}

func main() {
	ctx := context.WithValue(context.TODO(), "extension", ".mp3")
	service := NewService(MP3("billi_jean"))
	fmt.Println(service.r.Record(ctx))
}
