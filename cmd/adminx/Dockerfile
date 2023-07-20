FROM golang:1.20-alpine AS builder

LABEL stage=gobuilder

#ARG DEPLOY_CRED_PSW
#ARG DEPLOY_CRED_USER=deploy

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk update --no-cache && apk add --no-cache ca-certificates git tzdata

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GIT_TERMINAL_PROMPT 0
ENV GOPROXY http://nexus.hycapital.hk/repository/goproxy/,direct
ENV GOPRIVATE *.qeelyn.com
ENV GOINSECURE *.qeelyn.com
ENV GOSUMDB off

#RUN echo -e "http://$DEPLOY_CRED_USER:$DEPLOY_CRED_PSW@t.qeelyn.com" > ~/.git-credentials && \
#    git config --global credential.helper store

WORKDIR /build/adminx

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .
COPY cmd/adminx/etc/app.yaml /app/etc/app.yaml
COPY cmd/adminx/etc/rbac_model.conf /app/etc/rbac_model.conf
RUN go build -tags=go-json -ldflags="-s -w" -o /app/serve cmd/adminx/main.go

FROM alpine

COPY --from=builder /usr/share/zoneinfo/UTC /usr/share/zoneinfo/UTC
COPY --from=builder /usr/share/zoneinfo/Asia/Hong_Kong /usr/share/zoneinfo/Asia/Hong_Kong
ENV TZ Asia/Hong_Kong

WORKDIR /app

COPY --from=builder /app/serve /app/adminx-serve
COPY --from=builder /app/etc /app/etc

CMD ["./adminx-serve"]
