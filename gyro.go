package main

import (
	"log"
	"os"
	"os/exec"
	"time"

	"go.universe.tf/netboot/dhcp4"
)

const gyroURL = "http://saltmines.us/gyroup.php"
const name = "Some button pusher"
const dashMAC = "74:c2:46:81:f2:ac"
const debounceThreshold = 60 * time.Second

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: gyro-button")
	}

	log.Printf("Starting gyro at %s", time.Now().Format(time.RFC3339))

	dhcp, err := dhcp4.NewSnooperConn("0.0.0.0:67")
	if err != nil {
		log.Fatalf("Cannot start DHCP server: %s", err)
	}

	var debounce time.Time

	for {
		pkt, _, err := dhcp.RecvDHCP()
		if err != nil {
			log.Printf("RecvDHCP: %s", err)
			time.Sleep(3 * time.Second)
			continue
		}

		if pkt.HardwareAddr.String() != dashMAC {
			continue
		}

		if pkt.Type != dhcp4.MsgDiscover {
			continue
		}

		now := time.Now()
		if now.Before(debounce) {
			continue
		}
		debounce = now.Add(debounceThreshold)

		log.Printf("Gyro button pressed at %s", time.Now().Format(time.RFC3339))

		/*
						resp, err := http.PostForm(gyroURL, url.Values{"user_name": {name}})
						if err != nil {
							log.Printf("Error updating count: %s", err)
						}
			                        resp.Body.Close()
		*/

		if err := exec.Command("systemctl", "restart", "lightdm").Run(); err != nil {
			log.Printf("Error restarting lightdm: %s", err)
		}
	}
}
