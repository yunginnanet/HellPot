# HellPot
   
   [![GoDoc](https://godoc.org/github.com/yunginnanet/HellPot?status.svg)](https://godoc.org/github.com/yunginnanet/HellPot) [![Go Report Card](https://goreportcard.com/badge/github.com/yunginnanet/HellPot)](https://goreportcard.com/report/github.com/yunginnanet/HellPot) [![IRC](https://img.shields.io/badge/ircd.chat-%23tcpdirect-blue.svg)](ircs://ircd.chat:6697/#tcpdirect) [![Mentioned in Awesome Honeypots](https://awesome.re/mentioned-badge.svg)](https://github.com/paralax/awesome-honeypots)
   
## Summary  
  
HellPot is an endless honeypot based on [Heffalump](https://github.com/carlmjohnson/heffalump) that sends unruly HTTP bots to hell.

Notably it implements a [toml configuration file](https://github.com/spf13/viper), has [JSON logging](https://github.com/rs/zerolog), and comes with significant performance gains.
  
 ![Exploding Heffalump](hellgif.gif)

## Grave Consequences

Clients (hopefully bots) that disregard `robots.txt` and connect to your instance of HellPot will **suffer eternal consequences**. 

HellPot will send an infinite stream of data that is *just close enough* to being a real website that they might just stick around until their soul is ripped apart and they cease to exist.

Under the hood of this eternal suffering is a markov engine that chucks bits and pieces of [The Birth of Tragedy (Hellenism and Pessimism)](https://www.gutenberg.org/files/51356/51356-h/51356-h.htm) by Friedrich Nietzsche at the client using [fasthttp](https://github.com/valyala/fasthttp).

## Compilation

HellPot should probably be built with Go version 1.17 or higher.

HellPot uses [go modules](https://go.dev/blog/using-go-modules). This should make it dead simple to build with a stock Go installation.

1 ) `git clone https://github.com/yunginnanet/HellPot`

2 ) `cd HellPot`

4 ) `go build cmd/HellPot/HellPot.go`

5 ) *Consider the potential grave consequences of your actions.*

## Usage

### YOLO Method:

In the event of a missing configuration file, HellPot will attempt to place it's default config in **$HOME/.config/HellPot/config.toml**. This allows irresponsible souls to begin raining hellfire with ease, ***immediately***:

1 ) Download a [compiled release](https://github.com/yunginnanet/HellPot/releases/latest) 

2 ) Run binary and immedidately begin sending clients directly to hell.

---

### Reasonable Method:

1 ) Configure webserver as reverse proxy (see below)

2 ) `./HellPot --genconfig `

3 ) Edit your newly generated `config.toml` as desired.

4 ) Ponder your ~~existence~~ server's ability to handle your chosen performance values.

5 ) ./HellPot -c config.toml

666 ) 𝙏͘͝𝙝̓̓͛𝙚͑̈́̀ 𝙨͆͠͝𝙠͑̾͌𝙮̽͌͆ 𝙞̓̔̔𝙨͒͐͝ 𝙛͑̈́̚𝙖͛͒𝙡͑͆̽𝙡̾̚̚𝙞͋̒̒𝙣̾͛͝𝙜͒̒̀.́̔͝​

## Example Config (toml) 
   
```  
[http]
  # TCP Listener (default)
  bind_addr = "127.0.0.1"
  bind_port = "8080"
  paths = ["wp-login.php","wp-login"]

  # Unix Socket Listener (will override default)
  use_unix_socket = false
  unix_socket = "/var/run/hellpot"

[logger]
  debug = true
  directory = "/home/kayos/.config/HellPot/logs/"
  nocolor = false
  use_date_filename = true

[performance]
  # max_workers is only valid if restrict_concurrency is true
  restrict_concurrency = false
  max_workers = 256
  
[deception]
  # Used as "Server: " header (if not proxied)
  server_name = "nginx"

```
  

## Example Web Server Config (nginx)  
    
```          
location '/robots.txt' {
	proxy_set_header Host $host;
	proxy_set_header X-Real-IP $remote_addr;
	proxy_pass http://127.0.0.1:8080$request_uri;
}  

location '/wp-login.php' {
	proxy_set_header Host $host;
	proxy_set_header X-Real-IP $remote_addr;
	proxy_pass http://127.0.0.1:8080$request_uri;
}
```
