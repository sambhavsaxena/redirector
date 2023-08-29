// Reverse proxy setup with SSL integration

package main

import (
    "crypto/tls"
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
)

type Proxy struct {
    target *url.URL
    proxy  *httputil.ReverseProxy
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    log.Println(r.URL)
    r.Host = p.target.Host
    p.proxy.ServeHTTP(w, r)
}

func main() {
    target, err := url.Parse("https://example.com")
    if err != nil {
        panic(err)
    }
    proxy := httputil.NewSingleHostReverseProxy(target)
    proxy.Transport = &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    p := &Proxy{target: target, proxy: proxy}
    err = http.ListenAndServe(":3000", p)
    if err != nil {
        panic(err)
    }
}
