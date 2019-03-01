FROM golang:1.11-alpine as builder
RUN adduser -S -D -g '' motn
WORKDIR $GOPATH/src/github.com/vickleford/motn
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s"  -o /go/bin/motn main.go

# ------------------
FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/motn /motn

USER motn
EXPOSE 8080
CMD [ "/motn" ]
