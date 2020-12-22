package main

import (
  "os"
  "strings"
  "flag"
  "sync"
  "net"
  "net/http"
  "time"
  "bufio"
  "github.com/fatih/color"
)

var concurrency int
var to int
var sec_headers = []string{"Strict-Transport-Security","Content-Security-Policy","X-Frame-Options","X-XSS-Protection","X-Content-Type"}
var software_headers = []string{"Server","X-Powered-By","X-AspNetMvc-Version","X-Asp-Version", "X-Version", "X-Runtime"}

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func runChecks (){
    timeout := time.Duration(to * 1000000)
    var tr = &http.Transport{
		MaxIdleConns:      30,
		IdleConnTimeout:   time.Second,
		DisableKeepAlives: true,
		DialContext: (&net.Dialer{
			Timeout:   timeout,
			KeepAlive: time.Second,
		}).DialContext,
	}
	client := &http.Client{
		Transport:     tr,
		Timeout:       timeout,
	}
    jobs := make(chan string)
    var wg sync.WaitGroup
    for i := 0; i < concurrency; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for domain := range jobs {
                url := domain
                req, err := client.Head(url)
                if err != nil {
                    continue
                }
                defer req.Body.Close()
                // Colors
                w := color.New(color.FgCyan)//.Add(color.Bold)
                r := color.New(color.FgRed)//.Add(color.Bold)
                g := color.New(color.FgGreen)//.Add(color.Bold)

                // Create file for each target
                domain_string := url
                if strings.Contains(domain_string, "http://") {
                    remHttp := strings.Replace(domain_string,"http://","",-1)
                    file_name_http := remHttp+"_http"

                    err := os.Mkdir("out", 0755)
		    check(err)
                    f, err := os.Create("out/"+file_name_http)
                    check(err)
                    defer f.Close()
                    httpw := bufio.NewWriter(f)

                    w.Println("[>] " + url)
                    _, err = w.Fprintf(httpw, "[>] " + url+"\n")
                    check(err)

                    // HTTP Software Headers
                    for _,softHeader := range software_headers {
                        softHeadercheck := req.Header.Get(softHeader)
                        if softHeadercheck != "" {
                            g.Println("[+] "+softHeader + ": " + softHeadercheck)
                            _, err = g.Fprintf(httpw,"[+] "+softHeader + ": " + softHeadercheck+"\n")
                            check(err)
                        }
                    }
                    // HTTP Security Headers
                    for _,secHeader := range sec_headers {
                        secHeadercheck := req.Header.Get(secHeader)
                        if secHeadercheck == "" {
                            r.Println("[-] " + secHeader)
                            _, err = r.Fprintf(httpw,"[-] " + secHeader+"\n")
                            check(err)
                        }
                    }
                    httpw.Flush()
                }   else {
                    remHttps := strings.Replace(domain_string,"https://","",-1)
                    file_name_https := remHttps+"_https"

                    f, err := os.Create("out/"+file_name_https)
                    check(err)
                    defer f.Close()
                    httpsw := bufio.NewWriter(f)

                    w.Println("[>] " + url)
                    _, err = w.Fprintf(httpsw, "[>] " + url+"\n")
                    check(err)

                    // HTTP Software Headers
                    for _,softHeader := range software_headers {
                        softHeadercheck := req.Header.Get(softHeader)
                        if softHeadercheck != "" {
                            g.Println("[+] "+softHeader + ": " + softHeadercheck)
                            _, err = g.Fprintf(httpsw,"[+] "+softHeader + ": " + softHeadercheck+"\n")
                            check(err)
                        }
                    }
                    // HTTP Security Headers
                    for _,secHeader := range sec_headers {
                        secHeadercheck := req.Header.Get(secHeader)
                        if secHeadercheck == "" {
                            r.Println("[-] " + secHeader)
                            _, err = r.Fprintf(httpsw,"[-] " + secHeader+"\n")
                            check(err)
                        }
                    }
                    httpsw.Flush()
                }
            }
        }()
    }
    sc := bufio.NewScanner(os.Stdin)
    for sc.Scan() {
        domain := sc.Text()
        jobs <- domain
    }
    close(jobs)
    wg.Wait()
}

func main() {
    flag.IntVar(&concurrency, "c", 20, "concurrency level")
    flag.IntVar(&to, "t", 10000, "timeout (milliseconds)")
    flag.Parse()
    runChecks()
}
