package main

import (
    "context"
    "fmt"
    "log"

    "github.com/ServiceWeaver/weaver"
)

func main() {
    if err := weaver.Run(context.Background(), serve); err != nil {
        log.Fatal(err)
    }
}

type app struct{
    weaver.Implements[weaver.Main]
    reverser weaver.Ref[Reverser]
}

func serve(ctx context.Context, app *app) error {
    // Call the Reverse method.
    var r Reverser = app.reverser.Get()
    reversed, err := r.Reverse(ctx, "!dlroW ,olleH")
    if err != nil {
        return err
    }
    fmt.Println(reversed)
    return nil
}
