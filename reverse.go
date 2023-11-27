package main

import (
	"context"

	"github.com/ServiceWeaver/weaver"
)

type Reverser interface {
  Reverse(context.Context, string) (string, error)
}


type reverser struct {
 weaver.Implements[Reverser]
  prefixer weaver.Ref[Prefixer]
}

func (r *reverser) Reverse(ctx context.Context, s string) (string, error) {
  runes := []rune(s)
  n := len(runes)

  for i := 0; i <  n/2; i++ {
    runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
  }

  var p Prefixer = r.prefixer.Get()

  result, _ := p.Prefix(ctx, string(runes))

  return result, nil
}

