FROM golang:1.14.3-alpine3.11 AS builder

RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

WORKDIR /PurdoobahDiscordBot
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/bot ./...

FROM alpine:3.11.6

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

WORKDIR /PurdoobahDiscordBot
COPY --from=builder /PurdoobahDiscordBot/assets assets
COPY --from=builder /PurdoobahDiscordBot/bin/bot bot

CMD /PurdoobahDiscordBot/bot -token=$PURDOOBAH_DISCORD_BOT_TOKEN
