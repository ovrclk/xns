IMAGE: quay.io/ovrclk/xns

test:
	go test ./...

clean:
	rm -rf _build

build: 
	mkdir -p _build
	go build -o _build/xns .

deps:
	go get github.com/miekg/dns
	go get github.com/spf13/cobra

image:
	docker build . -t quay.io/ovrclk/xns

install:
	go install .
