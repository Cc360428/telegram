FROM repo.rdvalidasi.com/golang/docker:1-16-10 AS builder

WORKDIR /opt/Cc/SendTG

COPY . .
RUN CGO_ENABLED=0 go build -mod=vendor -ldflags "-w -s" -o SendTG

FROM alpine:3.5

RUN apk add --no-cache tzdata
ENV TZ Asia/Bangkok

WORKDIR /opt/SendTG
RUN apk --no-cache add ca-certificates && update-ca-certificates

COPY --from=builder /opt/Cc/SendTG/SendTG .
COPY --from=builder /opt/Cc/SendTG/conf ./conf
CMD ["./SendTG"]