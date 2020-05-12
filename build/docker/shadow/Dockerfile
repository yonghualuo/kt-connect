FROM golang:1.13
ENV GO111MODULE on
#RUN go env -w GO111MODULE=on
#RUN go env -w GOPROXY=https://goproxy.cn,direct
# Install go debugger
#RUN GO111MODULE=off go get -u github.com/go-delve/delve/cmd/dlv

FROM damonlyh/shadow-base:v0.2
COPY artifacts/shadow/shadow-linux-amd64 /usr/sbin/shadow-linux-amd64
#COPY --from=0 /go/bin/dlv /usr/sbin/dlv
#RUN apt-get install -y net-tools
ADD build/docker/shadow/run.sh /run.sh
RUN chmod 755 /run.sh

ENTRYPOINT ["/run.sh"]