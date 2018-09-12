Jaffa is a DNS server/forwarder/dispatcher written in Go.

## Features

+ Full IPv6 support
+ Multiple DNS upstream
    + Via UDP/TCP with custom port
    + Via SOCKS5 proxy (TCP only)
    + With EDNS Client Subnet (ECS) [RFC7871](https://tools.ietf.org/html/rfc7871)
+ Dispatcher
    + IPv6 record (AAAA) redirection
    + Custom domain
    + Custom IP network
+ Minimum TTL modification
+ Hosts (Random order if there are multiple answers)
+ Cache with ECS
