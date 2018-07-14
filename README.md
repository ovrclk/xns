# xns

xns is a simple and fast wildcard DNS service for any IP address, written in go and inspired by [xip.io](http://xip.io). xns allows you to map any IP Address in the following DNS wildcard entries:

```
          10.0.0.1.xns.akash.network  resolves to  10.0.0.1
      www.10.0.0.1.xns.akash.network  resolves to  10.0.0.1
   mysite.10.0.0.1.xns.akash.network  resolves to  10.0.0.1
  foo.bar.10.0.0.1.xns.akash.network  resolves to  10.0.0.1
          10-0-0-1.xns.akash.network  resolves to  10.0.0.1
      www.10-0-0-1.xns.akash.network  resolves to  10.0.0.1
   mysite.10-0-0-1.xns.akash.network  resolves to  10.0.0.1
  foo.bar.10-0-0-1.xns.akash.network  resolves to  10.0.0.1
```

## Usage

To start the server, simply run `sudo xns`, and use dig to query:

```
$ dig +short foo.127.0.0.1.xns.akash.network @localhost

127.0.0.1
```

xns responds to `xns.akash.network` by default. To specify a domain, run `xns -r [domain]`, for example `sudo xns -r example.com`

```
$ xns -h                                                                                                                                                                                                              
xns is a simple and fast wildcard DNS service for any IP addressm,
written in go and inspired by xip.io. xns powers xns.akash.network.

Usage:
  xns [flags]

Flags:
  -b, --bind string   address to bind (udp) (default "0.0.0.0:53")
  -h, --help          help for xns
  -r, --root string   root domain (example: xns.akash.network)
```

### Docker

Start the docker container on port `53` and use dig to query.

```
$ sudo docker run --rm -p 53:53/udp -t quay.io/ovrclk/xns
```

In a seperate window run:

```
$ dig +short foo.127.0.0.1.xns.akash.network @localhost

127.0.0.1
```

### Kubernetes

Create a config map and specify the root domain and port by setting `xns.root` and `xns.port` variables respectively.

```
$ kubectl create configmap xns-config --from-literal=xns.root=xns.akash.network --from-literal=xns.port=53
$ kubectl create -f https://raw.githubusercontent.com/ovrclk/xns/master/contrib/k8s.yml
```

## Development

### Building

Fetch the source and place it under `$GOPATH/src/github.com/ovrclk/xns`

```
$ mkdir -p $GOPATH/src/github.com/ovrclk
$ git clone https://github.com/ovrclk/xns.git $GOPATH/src/github.com/ovrclk
$ cd $GOPATH/src/github.com/ovrclk/xns
$ make build
```

### Docker image

```
$ make image
```
