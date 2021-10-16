# HellPot
   
[![GoDoc](https://godoc.org/github.com/yunginnanet/HellPot?status.svg)](https://godoc.org/github.com/yunginnanet/HellPot) [![Go Report Card](https://goreportcard.com/badge/github.com/yunginnanet/HellPot)](https://goreportcard.com/report/github.com/yunginnanet/HellPot) [![IRC](https://img.shields.io/badge/ircd.chat-%23tcpdirect-blue.svg)](ircs://ircd.chat:6697/#tcpdirect)
  
### [v0.3 with fasthttp](https://github.com/yunginnanet/HellPot/releases/tag/v0.3) roughly [~75%](https://github.com/yunginnanet/HellPot/issues/2) faster!
  
HellPot is an endless honeypot that sends bots to hell. Based on [Heffalump](https://github.com/carlmjohnson/heffalump).   
  
  It finishes the work of Heffalump with a few improvements and the addition of a [toml configuration file](https://github.com/spf13/viper) and [JSON logging](https://github.com/rs/zerolog). It is built off of [CokePlate](https://git.tcp.direct/kayos/CokePlate).
    

The source of the honeypot data is [The Birth of Tragedy (Hellenism and Pessimism)](https://www.gutenberg.org/files/51356/51356-h/51356-h.htm) by Friedrich Nietzsche

![Exploding Heffalump](hellgif.gif)

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


## Example Program Config (toml) 
  
  If the configuration  file is missing, the default settings will automatically drop itself in the current working directory as `config.toml`.  
    
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
  
