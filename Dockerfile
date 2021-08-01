FROM golang:1.16-alpine AS builder

WORKDIR /app

COPY . . 
RUN go mod download

RUN CGO_ENABLED=0 GOARCH=amd64 go build -o /victim_amd64
RUN GOARCH=arm GOARM=7 go build -o /victim_armv7
RUN GOARCH=arm64 go build -o /victim_arm64

# ---------

# FROM scratch

# COPY --from=builder /etc/passwd /etc/passwd
# COPY --from=builder /etc/group /etc/group
# COPY --from=builder /victim /victim

# USER guest

# EXPOSE 8080

CMD [ "/victim_amd64" ]
