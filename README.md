# gohead

Check for missing HTTP security headers, verbose server headers, and sensitive software headers concurrently.

`[+]` indicates a present server/software header<br>
`[-]` indicates a missing security header

Outputs file for each host to `out` directory in current working folder for easy navigation of results.

`gohead` tests for the following HTTP headers:
* Server
* X-Powered-By
* X-AspNetMvc-Version
* X-Asp-Version
* X-Version
* X-Runtime
* Strict-Transport-Security
* Content-Security-Policy
* X-Frame-Options
* X-XSS-Protection
* X-Content-Type

## Install 
```
$ go get github.com/mlcsec/gohead
```

## Help
```
Usage of gohead:
  -c int
        concurrency level (default 20)
  -t int
        timeout (milliseconds) (default 10000)
```

## Example
Takes input via stdin - `cat` a range of hosts returned from httprobe or `echo` a single host to `gohead`.
```
$ cat urls | gohead
[>] http://scanme.nmap.org
[+] Server: Apache/2.4.7 (Ubuntu)
[-] Strict-Transport-Security
[-] Content-Security-Policy
[-] X-Frame-Options
[-] X-XSS-Protection
[-] X-Content-Type
[>] https://mlcsec.com
[+] Server: Netlify
[-] Content-Security-Policy
[-] X-Frame-Options
[-] X-XSS-Protection
[-] X-Content-Type
$ cd out/
$ ls
mlcsec.com_https  scanme.nmap.org_http
```
