FROM alpine:latest
RUN apk --no-cache add make git go gcc libtool musl-dev

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

COPY . /
RUN make

RUN apk add ca-certificates && \
    rm -rf /var/cache/apk/* /tmp/* && \
    [ ! -e /etc/nsswitch.conf ] && echo 'hosts: files dns' > /etc/nsswitch.conf

# Increase UDP receive buffer size for max performance of QUIC
RUN sysctl -w net.core.rmem_max=2500000

ENTRYPOINT ["/hari"]
