package main

import (
)

func main() {
  router := CreateServer()
  router.Run(":8082")
}
