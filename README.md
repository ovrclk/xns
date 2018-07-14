# xns

xns is a simple and fast wildcard DNS service for any IP address, written in go and inspired by [xip.io](http://xip.io). xns allows you to map any IP Address in the following DNS wildcard entries:

          10.0.0.1.xns.akash.network  resolves to   10.0.0.1
      www.10.0.0.1.xns.akash.network  resolves to   10.0.0.1
   mysite.10.0.0.1.xns.akash.network  resolves to   10.0.0.1
  foo.bar.10.0.0.1.xns.akash.network  resolves to   10.0.0.1
          10-0-0-1.xns.akash.network  resolves to   10.0.0.1
      www.10-0-0-1.xns.akash.network  resolves to   10.0.0.1
   mysite.10-0-0-1.xns.akash.network  resolves to   10.0.0.1
  foo.bar.10-0-0-1.xns.akash.network  resolves to   10.0.0.1
