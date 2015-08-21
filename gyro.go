package main

import (
	"log"
	"net"
	"net/http"
	"net/url"
	"os"

	"github.com/mdlayher/arp"
)

const gyroURL = "http://saltmines.us/gyroup.php"
const name = "Some button pusher"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: arp <iface>")
	}

	iface, err := net.InterfaceByName(os.Args[1])
	if err != nil {
		log.Fatal("InterfaceByName: ", err)
	}

	c, err := arp.NewClient(iface)
	if err != nil {
		log.Fatal("arp.NewClient: ", err)
	}

	for {
		p, _, err := c.Read()
		if err != nil {
			log.Fatal(err)
		}

		if p.Operation == arp.OperationRequest && p.SenderIP.Equal(net.IPv4zero) {
			resp, err := http.PostForm(gyroURL, url.Values{"user_name": {name}})
			if err != nil {
				log.Printf("Error updating count: %s", err)
			}
			resp.Body.Close()
		}
	}
}
