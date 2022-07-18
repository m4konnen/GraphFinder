package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup

	concurrency = 50
	maxSize     = int64(1024000)
)

func check(e error) {
	if e != nil {
		fmt.Println("Error")
		return
	}
}

func main() {

	f := flag.String("f", "", "Input File name.")
	o := flag.String("o", "", "Output File name.")
	flag.Parse()

	filename := *f
	output := *o

	banner()

	stat, _ := os.Stdin.Stat()
	if (stat.Mode()&os.ModeCharDevice) != 0 && filename == "" {
		flag.PrintDefaults()
		return
	}

	var fileLines []string

	if filename != "" {
		file, err := os.Open(filename)
		check(err)

		fileScanner := bufio.NewScanner(file)

		fileScanner.Split(bufio.ScanLines)

		for fileScanner.Scan() {
			fileLines = append(fileLines, fileScanner.Text())
		}

		file.Close()

		foundList := Scan(fileLines, output)
		scanIntrospect(foundList, output)
		return
	}

	var input io.Reader
	input = os.Stdin

	sc := bufio.NewScanner(input)

	if sc != nil {

		for sc.Scan() {
			fileLines = append(fileLines, sc.Text())
		}

		foundList := Scan(fileLines, output)
		scanIntrospect(foundList, output)

	}

	return

}

func Scan(urllist []string, output string) []string {

	found := make([]string, 0)

	proxyString := "http://localhost:8080"
	proxyURL, _ := url.Parse(proxyString)

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true,
			MinVersion: tls.VersionTLS11,
			MaxVersion: tls.VersionTLS11},
		MaxIdleConns:        concurrency,
		MaxIdleConnsPerHost: concurrency,
		MaxConnsPerHost:     concurrency,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   3 * time.Second}

	semaphore := make(chan bool, concurrency)

	for _, baseUrl := range urllist {

		enpoints := []string{
			"graphql", "qql", "ql", "console/graphql", "graphiql", "api/graphql",
		}

		for _, e := range enpoints {

			wg.Add(1)
			semaphore <- true

			rawURL := baseUrl + e

			go func(rawURL string) {
				defer wg.Done()

				finalURL, _ := url.Parse(rawURL)

				request, err := http.NewRequest("GET", finalURL.String(), nil)
				if err != nil {
					return
				}

				resp, err := client.Do(request)
				if err != nil {
					return
				}

				if resp.StatusCode != 404 {

					var data = []byte(`{"operationName":null,
					"variables":{},
					"query":"{\n  xpto: xpto\n}\n"
					}`)

					request, err := http.NewRequest("POST", finalURL.String(), bytes.NewBuffer(data))
					request.Header.Add("Content-Type", "application/json")
					check(err)

					response, err := client.Do(request)
					if err != nil {
						return
					}

					body, err := io.ReadAll(response.Body)
					if err != nil {
						return
					}

					if strings.Contains(string(body), "Cannot query field") {
						fmt.Println("[+] [FOUND] -", finalURL.String())
						found = append(found, finalURL.String())
						toOutfile(finalURL.String(), output)
					}

					defer response.Body.Close()

				}

			}(rawURL)
			<-semaphore

		}

		wg.Wait()

	}

	return found

}

func toOutfile(url, out string) {

	if out == "" {
		return
	}

	outfile, err := os.OpenFile(out, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer outfile.Close()

	if url != "" {
		_, err := outfile.WriteString(url + "\n")
		if err != nil {
			log.Fatal(err)
		}

	}

}

func scanIntrospect(list []string, output string) {

	proxyString := "http://localhost:8080"
	proxyURL, _ := url.Parse(proxyString)

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true,
			MinVersion: tls.VersionTLS11,
			MaxVersion: tls.VersionTLS11},
		MaxIdleConns:        concurrency,
		MaxIdleConnsPerHost: concurrency,
		MaxConnsPerHost:     concurrency,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   3 * time.Second}

	for _, line := range list {

		var data = []byte(`{"operationName":"IntrospectionQuery","variables":{},"query":"query IntrospectionQuery {\n  __schema {\n    queryType {\n      name\n    }\n    mutationType {\n      name\n    }\n    subscriptionType {\n      name\n    }\n    types {\n      ...FullType\n    }\n    directives {\n      name\n      description\n      locations\n      args {\n        ...InputValue\n      }\n    }\n  }\n}\n\nfragment FullType on __Type {\n  kind\n  name\n  description\n  fields(includeDeprecated: true) {\n    name\n    description\n    args {\n      ...InputValue\n    }\n    type {\n      ...TypeRef\n    }\n    isDeprecated\n    deprecationReason\n  }\n  inputFields {\n    ...InputValue\n  }\n  interfaces {\n    ...TypeRef\n  }\n  enumValues(includeDeprecated: true) {\n    name\n    description\n    isDeprecated\n    deprecationReason\n  }\n  possibleTypes {\n    ...TypeRef\n  }\n}\n\nfragment InputValue on __InputValue {\n  name\n  description\n  type {\n    ...TypeRef\n  }\n  defaultValue\n}\n\nfragment TypeRef on __Type {\n  kind\n  name\n  ofType {\n    kind\n    name\n    ofType {\n      kind\n      name\n      ofType {\n        kind\n        name\n        ofType {\n          kind\n          name\n          ofType {\n            kind\n            name\n            ofType {\n              kind\n              name\n              ofType {\n                kind\n                name\n              }\n            }\n          }\n        }\n      }\n    }\n  }\n}\n"}`)

		urlFinal, _ := url.Parse(line)

		req, err := http.NewRequest("POST", urlFinal.String(), bytes.NewBuffer(data))
		req.Header.Add("Content-Type", "application/json")

		if err != nil {
			continue
		}

		res, err := client.Do(req)

		if err != nil {
			continue
		}

		resStr, err := io.ReadAll(res.Body)

		if res.StatusCode != 404 && strings.Contains(string(resStr), "data") {

			fmt.Println("[+] [INSTROSPECTION ENABLED] - " + urlFinal.String())
			toOutfile(urlFinal.String()+" - [+] [INSTROSPECTION ENABLED] ", "instrospect-"+output)

			defer res.Body.Close()
		}

	}
}

func banner() {
	fmt.Println(`
	
   ____     ____        _       ____    _   _    _____              _   _    ____  U _____ u   ____     
U /"___|uU |  _"\ u U  /"\  u U|  _"\ u|'| |'|  |" ___|    ___     | \ |"|  |  _"\ \| ___"|/U |  _"\ u  
\| |  _ / \| |_) |/  \/ _ \/  \| |_) |/| |_| |\U| |_  u   |_"_|   <|  \| |>/| | | | |  _|"   \| |_) |/  
 | |_| |   |  _ <    / ___ \   |  __/ U|  _  |u\|  _|/     | |    U| |\  |uU| |_| |\| |___    |  _ <    
  \____|   |_| \_\  /_/   \_\  |_|     |_| |_|  |_|      U/| |\u   |_| \_|  |____/ u|_____|   |_| \_\   
  _)(|_    //   \\_  \\    >>  ||>>_   //   \\  )(\\,-.-,_|___|_,-.||   \\,-.|||_   <<   >>   //   \\_  
 (__)__)  (__)  (__)(__)  (__)(__)__) (_") ("_)(__)(_/ \_)-' '-(_/ (_")  (_/(__)_) (__) (__) (__)  (__) 

 USAGE: ./GraphFinder -f inputfile.txt -o outputfile.txt
													
	`)
}
