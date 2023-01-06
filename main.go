package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
)

type DnsRecordEntity struct {
	Data string `json:"data"`
	TTL  int    `json:"ttl"`
}

func putDNS(domainName string, domainType, domainValue string, domainRecord string, shopperID string, apiKey string, apiSecret string) (int, error) {
	var dnsRecord = DnsRecordEntity{
		Data: domainValue,
		TTL:  1800,
	}
	var putData = [1]DnsRecordEntity{dnsRecord}
	var putDataJson, _ = json.Marshal(putData)

	domainUrl := fmt.Sprintf("https://api.godaddy.com/v1/domains/%s/records/%s/%s", domainName, domainType, domainRecord)
	req, err := http.NewRequest("PUT", domainUrl, bytes.NewBuffer(putDataJson))
	if err != nil {
		return -1, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Shopper-Id", shopperID)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("sso-key %s:%s", apiKey, apiSecret))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return resp.StatusCode, err
	}
	return resp.StatusCode, nil
}

func getIP(ipv6 bool) (string, error) {
	url := "https://4.ipw.cn"
	if ipv6 {
		url = "https://6.ipw.cn"
	}
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	//Convert the body to type string
	var ip = string(body)

	return ip, nil
}

func main() {
	h := flag.Bool("help", false, "--help")
	flagDomain := flag.String("domain", "", "your domain")
	flagType := flag.String("type", "A", "default 'A'")
	flagName := flag.String("name", "", "default nil")
	flagShopperID := flag.String("shopperid", "", "shopper id")
	flagKey := flag.String("key", "", "api key")
	flagSecret := flag.String("secret", "", "api secret")
	flag.CommandLine.SortFlags = false
	flag.Parse()

	if *h {
		flag.Usage()
		return
	}

	bytesWriter := &bytes.Buffer{}
	stdoutWriter := os.Stdout
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(io.MultiWriter(bytesWriter, stdoutWriter))
	log.SetLevel(log.InfoLevel)

	ipv6 := false

	if *flagType == "AAAA" {
		ipv6 = true
	}

	ip, err := getIP(ipv6)
	if err != nil {
		log.Fatalln(err)
	}
	log.WithField("IP", ip).Info("get ip success")

	statusCode, err := putDNS(*flagDomain, *flagType, ip, *flagName, *flagShopperID, *flagKey, *flagSecret)
	if err != nil {
		log.Fatalln(err)
	}
	log.WithField("statusCode", statusCode).Info("put request end")

}
