package main

import "fmt"

func main() {
	var choices int
	ipList := []string{"192.168.1.11", "192.168.1.12", "192.168.1.13"}
	for i, ip := range ipList {
		fmt.Printf("[%d]: %s\n", i, ip)
	}

	for {

		fmt.Printf("Please select a ip route:")
		_, err := fmt.Scanln(&choices)
		if err != nil {
			fmt.Println("choices error:", err)
			continue
		}
		if choices >= len(ipList) {
			fmt.Println("your choices is over range!")
			continue
		}
		break
	}
	fmt.Printf("Your select ip is: %s\n", ipList[choices])
}
