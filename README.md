# Ipblock Scraping Service
This service webscrap cidr of ipv4 address of all the country of the globe from [ipverse](https://github.com/ipverse/rir-ip) github repo. 
Then it calculates the first and last range of the [cidr](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing). 
After that it sort the ip addresses based on cidr and store it to redis. 

The ips' stored in database like below:

```
{
  "values": [
    {
      "cidr": "1.0.0.0/24",
      "country": "AU",
      "first_host": "1.0.0.0",
      "last_host": "1.0.0.0"
    },
    {
      "cidr": "1.0.1.0/24",
      "country": "CN",
      "first_host": "1.0.1.0",
      "last_host": "1.0.1.0"
    },
    {
      "cidr": "1.0.2.0/23",
      "country": "CN",
      "first_host": "1.0.2.0",
      "last_host": "1.0.3.255"
    },
    --------------------------
    --------------------------
    --------------------------
  ]
}
```
