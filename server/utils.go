package server

import (
	"context"
	dnspb "dns_resolver_dummy/dnspb"
	"errors"
	"log"
)

type DnsServiceServer struct {
}

var actionCloud2Local = [3]dnspb.DomainRes_Action{
	dnspb.DomainRes_DROP,
	dnspb.DomainRes_BLOCK,
	dnspb.DomainRes_ALLOW,
}
var (
	errNoValue   = errors.New("no value")
	errNoAction  = errors.New("no DNS action found")
	errInvAction = errors.New("invalid DNS action found")
)

var dnsrecords = map[string]int{

	"google.com":    2,
	"amazon.com":    1,
	"infoblox.com":  2,
	"cricbuzz.com":  1,
	"flipkart.com":  1,
	"uber.com":      1,
	"instagram.com": 1,
	"facebook.com":  1,
	"twitter.com":   1,
	"LinkedIn.com":  2,
}

func (s *DnsServiceServer) DNSauthenticate(ctx context.Context, req *dnspb.DomainReq) (*dnspb.DomainRes, error) {
	resp := &dnspb.DomainRes{}
	var err error
	domainname := req.GetDomainName()
	log.Printf("Value of result is %s", domainname)
	result, ok := dnsrecords[domainname]
	log.Printf("Value of result is %d", result)
	if !ok {
		return &dnspb.DomainRes{}, nil
	}
	if resp.Action, err = getAction(result); err != nil {
		return nil, err
	}
	/*
	   	response := &dnspb.DomainRes{
	         DomainRes_Action: getAction(result)
	   	}
	*/
	return resp, nil
}

func getAction(result int) (dnspb.DomainRes_Action, error) {

	localAction := actionCloud2Local[result]
	log.Printf("Value of localAction is %d", localAction)

	if localAction < 0 {
		return 0, errInvAction
	}
	return localAction, nil
}
