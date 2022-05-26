package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/cameronbrill/fig-issue/listener"
	"github.com/cameronbrill/fig-issue/pipeline"
)

func main() {
	source := []string{"FOO", "BAR", "BAX", "ABCDEFGHIJKLMNOPQRSTUVWXYZ"}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	readStream := listener.Producer(ctx, source)

	lowerStage := make(chan string)
	errorChannel := make(chan error)

	transformA := func(s string) (string, error) {
		return strings.ToLower(s), nil
	}

	go func() {
		pipeline.Step(ctx, readStream, lowerStage, errorChannel, transformA)
	}()

	type result struct {
		v string
		l int
	}

	titleStage := make(chan result)
	transformB := func(s string) (result, error) {
		if len(s) > 14 {
			return result{
				v: s,
				l: len(s),
			}, fmt.Errorf("invalid input string: %s", s)
		}

		return result{
			v: strings.Title(s),
			l: len(s),
		}, nil
	}

	go func() {
		pipeline.Step(ctx, lowerStage, titleStage, errorChannel, transformB)
	}()

	err := pipeline.Consumer(ctx, cancel, titleStage, errorChannel)
	if err != nil {
		fmt.Println(err)
	}
}
