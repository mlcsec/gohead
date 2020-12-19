# gohead

Check for missing HTTP security headers, verbose server headers, and sensitive software headers concurrently.

`[+]` indicates a present server/software header (green)<br> 
`[-]` indicates a missing security header (red)

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
go get github.com/mlcsec/gohead
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
$ assetfinder boozt.com | httprobe | gohead
[>] http://ia.boozt.com
[+] Server: AmazonS3
[-] Strict-Transport-Security
[-] Content-Security-Policy
[-] X-Frame-Options
[-] X-XSS-Protection
[-] X-Content-Type
[>] https://mail.boozt.com
[+] Server: cloudflare
[-] Strict-Transport-Security
[-] Content-Security-Policy
[-] X-Frame-Options
[-] X-XSS-Protection
[-] X-Content-Type
[>] http://mail.boozt.com
[+] Server: cloudflare
[-] Strict-Transport-Security
[-] Content-Security-Policy
[-] X-Frame-Options
[-] X-XSS-Protection
[-] X-Content-Type
[>] http://ib.boozt.com
[+] Server: cloudflare
[-] Strict-Transport-Security
[-] Content-Security-Policy
[-] X-Frame-Options
[-] X-XSS-Protection
[-] X-Content-Type
...

$ cd out/
$ ls
boozt.com_http                delivery-time.boozt.com_https  ia.boozt.com_http            mail.boozt.com_https        rel.boozt.com_http        sp-dev.boozt.com_https
boozt.com_https               drive.boozt.com_http           ia.boozt.com_https           m.boozt.com_http            rel.boozt.com_https       t.boozt.com_http
calendar.boozt.com_http       drive.boozt.com_https          ib.boozt.com_http            m.boozt.com_https           sendgrid.boozt.com_http   t.boozt.com_https
calendar.boozt.com_https      gcp.boozt.com_http             ib.boozt.com_https           nw.boozt.com_http           sendgrid.boozt.com_https  view.email.boozt.com_http
click.email.boozt.com_http    gcp.boozt.com_https            image.email.boozt.com_http   nw.boozt.com_https          sp.boozt.com_http         view.email.boozt.com_https
click.email.boozt.com_https   groups.boozt.com_http          image.email.boozt.com_https  parcel-api.boozt.com_http   sp.boozt.com_https        www.boozt.com_http
delivery-time.boozt.com_http  groups.boozt.com_https         mail.boozt.com_http          parcel-api.boozt.com_https  sp-dev.boozt.com_http     www.boozt.com_https

# Can then grep for certain headers
$ grep -iRl X-Powered
```
