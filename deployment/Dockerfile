FROM golang:1.17.7-alpine3.15 AS builder

# update alpine linux dependencies
RUN apk update
RUN apk add --no-cache git ca-certificates tzdata make
RUN update-ca-certificates

WORKDIR /purdoobahs-discord-bot

# generate binary
COPY . .
RUN make build

FROM alpine:3.15.0

# copy certs
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

WORKDIR /purdoobahs-discord-bot

# copy binary
COPY --from=builder /purdoobahs-discord-bot/bin/bot bot

# copy static assets
COPY --from=builder /purdoobahs-discord-bot/assets assets

# run bot
CMD ./bot -token=$PURDOOBAH_DISCORD_BOT_TOKEN
