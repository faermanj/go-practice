# docker build -f Containerfile --progress=plain --no-cache -t go-practice .
# docker run --network=host go-practice

FROM jetpackio/devbox:latest AS builder

USER devbox
WORKDIR /devbox
COPY --chown=devbox . .
RUN eval "$(devbox shell --print-env)" && \
    cd app && \
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o go-practice ./cmd/main.go

FROM registry.access.redhat.com/ubi9/ubi-minimal:latest

COPY --from=builder --chown=nobody /devbox/app/go-practice /app/go-practice
USER nobody
EXPOSE 8080
ENTRYPOINT ["/app/go-practice"]