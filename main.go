package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	domains := strings.Split(os.Getenv("DOMAINS"), ",")

	url := "https://api.ipify.org?format=text"
	fmt.Printf("Getting IP address from  ipify ...\n")
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic(resp.Status)
	}
	ip, err := io.ReadAll(resp.Body)
	_ = ip
	if err != nil {
		panic(err)
	}

	api, err := cloudflare.NewWithAPIToken(os.Getenv("CLOUDFLARE_API_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	zones, err := api.ListZones(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, z := range zones {
		records, _, err := api.ListDNSRecords(context.Background(), cloudflare.ZoneIdentifier(z.ID), cloudflare.ListDNSRecordsParams{Type: "A"})
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, r := range records {
			for _, d := range domains {
				if r.Name == d {
					if r.Content != string(ip) {
						fmt.Printf("Updating %s to %s\n", r.Name, string(ip))
						_, err = api.UpdateDNSRecord(context.Background(), cloudflare.ZoneIdentifier(z.ID), cloudflare.UpdateDNSRecordParams{Type: "A", Name: d, ID: r.ID, Content: string(ip)})
						if err != nil {
							fmt.Println(err)
							return
						}
					}
				}
			}
		}

	}
}
