FROM golang:latest
ENV XNS_ROOT=$GOPATH/src/github.com/ovrclk/xns
RUN mkdir -p $XNS_ROOT
ADD . $XNS_ROOT
WORKDIR $XNS_ROOT
RUN make deps && make build && go install .
CMD ["xns"]
