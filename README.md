# VPL Jail Server Reverse Proxy

This project is a Reverse Proxy server created to assist developers working on VPL Jail Server development. This server can accept both HTTP and WebSocket requests and forward them to a specific target server.

## Getting Started

<!-- TOC -->
* [VPL Jail Server Reverse Proxy](#vpl-jail-server-reverse-proxy)
  * [Getting Started](#getting-started)
    * [Requirements](#requirements)
    * [Installation](#installation)
    * [Usage](#usage)
    * [Contributing](#contributing)
    * [License](#license)
    * [Contact](#contact)
<!-- TOC -->

### Requirements

- Go (at least version 1.21 is required)
- VPL Jail Server (>= 3.0.0)

### Installation

Explain the steps on how to install and run the project. For example:

1. Clone the project:

   ```bash
   git clone https://github.com/yourusername/vpl-jail-reverse-proxy.git

2. Install the required packages:

   ```bash
   go get

3. Start the project:
   
   ```bash
   go run main.go

### Usage

Change the URL value for your own Jail Server

```go
  const URL = "http://192.168.71.2"
 ```

Example output

```
[2023-10-29 23:44:17] GET 192.168.18.7:57578/130166687421020/monitor 
Sec-Websocket-Key: [gBL+wu80fskfcrVqtVYLcg==]
Sec-Websocket-Extensions: [permessage-deflate; client_max_window_bits]
Connection: [Upgrade]
Upgrade: [websocket]
Sec-Websocket-Version: [13]
Accept-Encoding: [gzip, deflate]
Accept-Language: [en-US,en;q=0.9]
Pragma: [no-cache]
Cache-Control: [no-cache]
User-Agent: [Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36]
Origin: [http://localhost:9922]
{} 
[ws-response] message:compilation 
[ws-response] compilation: 
[ws-response] run:terminal 
```

### Contributing

* Fork the project and develop
* Send a pull request (PR)
* Report bugs
* Suggest new features or improvements

### License
* MIT

### Contact
ata.gunay@outlook.com


