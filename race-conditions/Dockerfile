FROM golang:1.15.13-alpine as builder

WORKDIR /app

COPY . .

RUN apk add gcc musl-dev
RUN CG_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags -ldflags="-w -s" .

FROM busybox

WORKDIR /app

COPY --from=builder /app/online-store .
COPY --from=builder /app/.env .

EXPOSE 8085

CMD [ "./online-store" ]