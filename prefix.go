
package main

import (
	"context"

	"github.com/ServiceWeaver/weaver"
)

type Prefixer interface {
  Prefix(context.Context, string) (string, error)
}

type prefixer struct {
 weaver.Implements[Prefixer]
}

func (r *prefixer) Prefix(_ context.Context, s string) (string, error) {
  return "PREFIX_" + s , nil
}

