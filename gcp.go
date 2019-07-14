package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/dns/v1"
)

func updateGCP(oldIP string, newIP string) {
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, dns.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	dnsService, err := dns.New(c)
	if err != nil {
		log.Fatal(err)
	}

	rb := &dns.Change{
		Additions: []*dns.ResourceRecordSet{
			&dns.ResourceRecordSet{
				Name: recordName,
				Type: "A",
				Ttl:  120,
				Rrdatas: []string{
					newIP,
				},
			},
		},
		Deletions: []*dns.ResourceRecordSet{
			&dns.ResourceRecordSet{
				Name: recordName,
				Type: "A",
				Ttl:  120,
				Rrdatas: []string{
					oldIP,
				},
			},
		},
	}

	resp, err := dnsService.Changes.Create(project, managedZone, rb).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Change code below to process the `resp` object:
	fmt.Printf("%#v\n", resp)
}

func getCurrentIP() (string, error) {
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, dns.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	dnsService, err := dns.New(c)
	if err != nil {
		log.Fatal(err)
	}

	lastKnown := ""

	req := dnsService.Changes.List(project, managedZone)
	if err := req.Pages(ctx, func(page *dns.ChangesListResponse) error {
		for _, change := range page.Changes {
			for _, addition := range change.Additions {
				if addition.Name == recordName && lastKnown == "" {
					lastKnown = addition.Rrdatas[0]
				}
			}
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	return lastKnown, err
}
