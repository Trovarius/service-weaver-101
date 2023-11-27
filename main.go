package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ServiceWeaver/weaver"
)

func main()  {
  if err := weaver.Run(context.Background(), serve); err != nil {
    log.Fatal(err)
  }
}

type app struct {
  weaver.Implements[weaver.Main]
  reverser weaver.Ref[Reverser]
  hello weaver.Listener 
}

func serve(ctx context.Context, app *app) error {
  // var r Reverser = app.reverser.Get()
  // original := "Trovarius"
  // reversed, err := r.Reverse(ctx, original)
  //
  // if err != nil {
  //   return err
  // }
  //
  // fmt.Println(original)
  // fmt.Println(reversed)
  // return nil

  fmt.Printf("Hello listener %v\n", app.hello)

  http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
    name := r.URL.Query().Get("name")

    if name == "" {
      name = "Trovarius"
    }

    reversed, err := app.reverser.Get().Reverse(ctx, name)

    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    fmt.Fprintf(w, "Hello %s!\n", reversed)
  })

  return http.Serve(app.hello, nil)
}
