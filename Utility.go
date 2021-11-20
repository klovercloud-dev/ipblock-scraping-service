package main

import (
	"bytes"
	"math"
	"net"
	"strconv"
	"strings"
)

func AddIpRange(ipBlocks []IpBlock)  {
	var temp IpRange
	for i:=0; i<len(ipBlocks); i++{
		temp = GetIpRange(ipBlocks[i].Cidr)
		ipBlocks[i].FirstHost = temp.First
		ipBlocks[i].LastHost = temp.Last
	}
}

func Sort(ipBlock []IpBlock){
	var temp IpBlock
	var i int
	var j int
	length := len(ipBlock)
	for i =0; i<length; i++{
		for j = 0; j< length-1 ; j++{
			if whichBlockBig(ipBlock[j], ipBlock[j+1]) == "first"{
				temp = ipBlock[j]
				ipBlock[j] = ipBlock[j+1]
				ipBlock[j+1] = temp
			}
		}
	}
}

func whichBlockBig(a IpBlock, b IpBlock) string {
	if bytes.Compare(a.LastHost, b.FirstHost) == 1{
		return "first"
	}else {
		return "second"
	}
}


func AdjustLength(bin string) string{
	count := len(bin)
	temp := ""

	for i:= 0; i< (8-count); i++{
		temp += "0"
	}
	return temp + bin
}

func GetHighestRange(block uint8, index int) uint64 {
	bin := strconv.FormatUint(uint64(block), 2)
	bin = AdjustLength(bin)
	if(len(bin)< 8){

	}
	var modifiedBin string

	// fmt.Println("bin: ",bin)
	for i := 0; i<8; i++{
		if i >= 8 - index {
			modifiedBin += "1"
		}else{
			modifiedBin += string([]rune(bin)[i])
		}
	}

	// fmt.Println("Modified Bin: ", modifiedBin)

	out, _ := strconv.ParseUint(modifiedBin, 2, 8)

	return out
}

func dupIP(ip net.IP) net.IP {
	dup := make(net.IP, len(ip))
	copy(dup, ip)
	return dup
}

func GetIpRange(cidr string) IpRange{
	var ipRange IpRange

	cidrParts := strings.Split(cidr, "/")

	ip := net.ParseIP(cidrParts[0])
	ipRange.First = dupIP(ip)

	prefix, _ := strconv.Atoi(cidrParts[1])

	numberOfBlocksToManipulate := math.Ceil(float64(32 - prefix)/float64(8))

	numberOfBit := (32 - prefix) % 8
	condition := 15 - int(numberOfBlocksToManipulate) + 1

	for i :=  0; i <int(numberOfBlocksToManipulate); i++ {
		temp := 15 - i
		if temp == condition{
			ip[temp] = byte(GetHighestRange(ip[temp], numberOfBit))
		}else{
			ip[temp] = 255
		}
	}
	ipRange.Last = ip
	//fmt.Println(ipRange.first)
	return ipRange
}

