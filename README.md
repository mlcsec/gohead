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
$ assetfinder -subs-only bitso.com | httprobe | gohead
[>] http://blog.bitso.com
[+] Server: openresty
[+] X-Powered-By: Medium
[-] Strict-Transport-Security
[-] X-Content-Type
[>] https://blog.bitso.com
[+] Server: openresty
[+] X-Powered-By: Medium
[-] Strict-Transport-Security
[-] X-Content-Type
[>] http://lri.bitso.com
[+] Server: openresty/1.15.8.1
[-] Content-Security-Policy
[-] X-Frame-Options
[-] X-XSS-Protection
[-] X-Content-Type
[>] http://status.bitso.com
[+] X-Runtime: 0.164759
[-] Content-Security-Policy
[-] X-Frame-Options
[-] X-Content-Type
...

$ cd out/
$ ls
api.bitso.com_http          click.email.boozt.com_https     help.bitso.com_http          mail.boozt.com_https        stageassets.bitso.com_http
api.bitso.com_https         delivery-time.boozt.com_http    help.bitso.com_https         m.boozt.com_http            stageassets.bitso.com_https
api-dev.bitso.com_http      delivery-time.boozt.com_https   ia.boozt.com_http            m.boozt.com_https           status.bitso.com_http
api-dev.bitso.com_https     devassets.bitso.com_http        ia.boozt.com_https           nw.boozt.com_http           status.bitso.com_https
assets.bitso.com_http       devassets.bitso.com_https       ib.boozt.com_http            nw.boozt.com_https          t.boozt.com_http
assets.bitso.com_https      dev.bitso.com_http              ib.boozt.com_https           parcel-api.boozt.com_http   t.boozt.com_https
bitso.com_http              dev.bitso.com_https             image.email.boozt.com_http   parcel-api.boozt.com_https  view.email.boozt.com_http
bitso.com_https             devmaltaassets.bitso.com_http   image.email.boozt.com_https  pos.bitso.com_http          view.email.boozt.com_https
blog.bitso.com_http         devmaltaassets.bitso.com_https  landing.bitso.com_http       pos.bitso.com_https         ws.bitso.com_http
blog.bitso.com_https        drive.boozt.com_http            landing.bitso.com_https      rel.boozt.com_http          ws.bitso.com_https
blog-en.bitso.com_http      drive.boozt.com_https           lra.bitso.com_http           rel.boozt.com_https         ws-dev.bitso.com_http
blog-en.bitso.com_https     edu.bitso.com_http              lra.bitso.com_https          sendgrid.boozt.com_http     ws-dev.bitso.com_https
boozt.com_http              edu.bitso.com_https             lrd.bitso.com_http           sendgrid.boozt.com_https    www.bitso.com_http
boozt.com_https             gcp.boozt.com_http              lrd.bitso.com_https          sp.boozt.com_http           www.bitso.com_https
calendar.boozt.com_http     gcp.boozt.com_https             lri.bitso.com_http           sp.boozt.com_https          www.boozt.com_http
calendar.boozt.com_https    groups.boozt.com_http           lri.bitso.com_https          sp-dev.boozt.com_http       www.boozt.com_https


# Can then grep for specific headers...
$ grep -iRl X-Powered
```
