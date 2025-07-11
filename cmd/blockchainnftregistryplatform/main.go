// cmd/blockchainnftregistryplatform/main.go
package main

import (
"flag"
"log"
"os"

"blockchainnftregistryplatform/internal/blockchainnftregistryplatform"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := blockchainnftregistryplatform.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
