package main

import (
		"fmt"
		"math/rand"
		"net"
		"time"
)

func scanDaInterwebz() {
		for {
				t := time.Second * 2
				ip := fmt.Sprintf("%d.%d.%d.%d:6379",
						rand.Intn(256), rand.Intn(256),
						rand.Intn(256), rand.Intn(256))
				c, err := net.DialTimeout("tcp", ip, t)
				if err == nil {
						c.SetDeadline(time.Now().Add(time.Second * 5))
						fmt.Fprintf(c, "RANDOMKEY\r\n")

						var l string
						_, err := fmt.Fscanln(c, &l)
						if err == nil {
								fmt.Println(ip)
								fmt.Println(l)
						}
				}
		}
}

func main() {
		for i := 0;i<300;i++ {
			go scanDaInterwebz()
		}
		fmt.Scanln()
}
