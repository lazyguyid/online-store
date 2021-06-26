FROM golang:1.15.13-alpine as builder

WORKDIR /app

COPY . .

RUN CG_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags -ldflags="-w -s" .

FROM busybox

WORKDIR /app

COPY --from=builder /app/race-conditions .
COPY --from=builder /app/.env .

EXPOSE 8085

CMD [ "./online-store" ]