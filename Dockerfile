FROM --platform=$BUILDPLATFORM  golang:1.16-alpine AS builder

ARG TARGETPLATFORM

WORKDIR /app

COPY . . 
RUN go mod download

RUN go build -o /victim

# ---------

FROM scratch

# COPY --from=builder /etc/passwd /etc/passwd
# COPY --from=builder /etc/group /etc/group
COPY --from=builder /victim /victim

# USER guest

# EXPOSE 8080

CMD [ "/victim" ]
