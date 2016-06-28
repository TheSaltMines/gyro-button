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
const dashMAC = "74:c2:46:81:f2:ac"
const debounceThreshold = 60 * time.Second

func initClient(ifaceName string) (*arp.Client, error) {
	iface, err := net.InterfaceByName(ifaceName)
	if err != nil {
		return nil, err
	}

	return arp.NewClient(iface)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: arp <iface>")
	}

	log.Printf("Starting gyro at %s", time.Now())

	mac, err := net.ParseMAC(dashMAC)
	if err != nil {
		log.Fatal("ParseMAC: ", err)
	}

	var c *arp.Client
	for {
		var err error
		c, err = initClient(os.Args[1])
		if err == nil {
			break
		}

		log.Printf("initClient: %s", err)
		time.Sleep(3 * time.Second)
	}

	var debounce time.Time

	for {
		p, _, err := c.Read()
		if err != nil {
			log.Printf("c.Read: %s", err)
			time.Sleep(3 * time.Second)
			continue
		}

		if p.Operation == arp.OperationRequest &&
			p.SenderIP.Equal(net.IPv4zero) &&
			p.SenderHardwareAddr.String() == mac.String() {

			now := time.Now()
			if now.Before(debounce) {
				log.Printf("Discarding button press within debounce window")
				continue
			}
			debounce = now.Add(debounceThreshold)

			log.Printf("Gyro button pressed at %s!", time.Now())
			resp, err := http.PostForm(gyroURL, url.Values{"user_name": {name}})
			if err != nil {
				log.Printf("Error updating count: %s", err)
			}
			resp.Body.Close()
		}
	}
}
