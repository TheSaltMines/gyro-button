package main

import (
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/mdlayher/arp"
)

const gyroURL = "http://saltmines.us/gyroup.php"
const name = "Some button pusher"

const dashMac = "74:c2:46:81:f2:ac"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: arp <iface>")
	}

	iface, err := net.InterfaceByName(os.Args[1])
	if err != nil {
		log.Fatal("InterfaceByName: ", err)
	}

	mac, err := net.ParseMAC(dashMac)
	if err != nil {
		log.Fatal("ParseMAC: ", err)
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

		if p.Operation == arp.OperationRequest &&
			p.SenderIP.Equal(net.IPv4zero) &&
			p.SenderHardwareAddr.String() == mac.String() {

			log.Printf("Gyro button pressed at %s!", time.Now())
			resp, err := http.PostForm(gyroURL, url.Values{"user_name": {name}})
			if err != nil {
				log.Printf("Error updating count: %s", err)
			}
			resp.Body.Close()
		}
	}
}
