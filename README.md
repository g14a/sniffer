
## Packet Sniffing and Decoding

### Goal

#### To capture live with an open wireless interface and decode the various layers in each packet displaying all its headers.

### Implementation

#### Using a packet processing library, [go-packet](https://github.com/google/gopacket) and its submodules in Golang.

### Local Execution

<h4>Install Golang in Ubuntu</h4>

``` sudo apt install golang-go ```

<h4> Execute the following in home directory</h4>

``` cd go```
<br />
``` git clone https://github.com/gowtham-munukutla/nsd```

<h4>Install libraries</h4>

``` go get github.com/google/gopacket```
<br />
``` go get github.com/fatih/color ```

<h4> Navigate to </h4> 

``` cd go/src/github.com/nsd ```
<br />
``` go build```

<h4> Run the build </h4>

``` ./nsd wlp2s0 ```

<h4> The second parameter should be your wireless interface. In my case it is </h4>

``` wlp2s0 ```