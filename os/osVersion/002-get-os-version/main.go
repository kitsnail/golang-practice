package main

import (
   "fmt"
   "github.com/shirou/gopsutil/host"
)

func main() {
   platform, family, pver, _ := host.PlatformInformation()
   fmt.Println("Platform:",platform)
   fmt.Println("Famliy:",family)
   fmt.Println("ProductVersion:",pver)
  // fmt.Println("Version:",version)
}

