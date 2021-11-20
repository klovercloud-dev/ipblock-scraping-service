package main

import "net"

type IpBlocks struct {
	Values []IpBlock `json:"values"`
}

type IpBlock struct {
	Cidr      string `json:"cidr"`
	Country   string `json:"country"`
	FirstHost net.IP `json:"first_host"`
	LastHost  net.IP `json:"last_host"`
}

type IpRange struct {
	First net.IP
	Last  net.IP
}
