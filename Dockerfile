FROM --platform=$BUILDPLATFORM  golang:1.16-alpine AS builder

# Parses TARGETPLATFORM and converts it to GOOS, GOARCH, and GOARM.
COPY --from=tonistiigi/xx:golang / /

ARG TARGETPLATFORM

WORKDIR /app

COPY . . 
RUN go mod download

# Build using GOOS, GOARCH, and GOARM
RUN CGO_ENABLED=0 go build -a -o /victim

# ---------

# FROM scratch

# COPY --from=builder /etc/passwd /etc/passwd
# COPY --from=builder /etc/group /etc/group
# COPY --from=builder /victim /victim

# USER guest

# EXPOSE 8080

CMD [ "/victim" ]
