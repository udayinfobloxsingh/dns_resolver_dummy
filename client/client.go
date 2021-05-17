package main

import (
	dnspb "dns_resolver_dummy/dnspb"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":50551", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	defer conn.Close()

	c := dnspb.NewDnsServiceClient(conn)

	DNSauthenticate("google.com", c)
	DNSauthenticate("yahoo.com", c)
	DNSauthenticate("flipkart.com", c)
	DNSauthenticate("facebook.com", c)
	DNSauthenticate("infoblox.com", c)
	DNSauthenticate("twitter.com", c)
	DNSauthenticate("music.com", c)

	log.Println("END OF PROGRAM")

}

func DNSauthenticate(domainname string, c dnspb.DnsServiceClient) {
	response, err := c.DNSauthenticate(context.Background(), &dnspb.DomainReq{DomainName: domainname})
	if err != nil {
		log.Fatalf("Error when calling Server: %s", err)
	}
	log.Printf("Response from server %s for domain: %s", response, domainname)
}
