package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	client := InitDb()
	var ipblocks IpBlocks
	val := GetIpBlocks()
	ipblocks.Values = val
	AddIpRange(ipblocks.Values)
	Sort(ipblocks.Values)

	resp, _ := json.Marshal(ipblocks)
	fmt.Println(string(resp))

	SetIpBlocks(client, "ipblocks_v2", string(resp))
}
