package main

import (
  "fmt"
  "GoSha256/ssl"
)


func main(){
  data := "Decidable-Unsavory-Marmalade-Onward-Bazooka-Supply-Hardness-Boondocks-Cosmic-Improving"
  fmt.Printf("Input:\t%s\nOutput:\t%s\n", data, ssl.Sha256(data))
}

