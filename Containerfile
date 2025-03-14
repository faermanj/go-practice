# docker build --debug -f Containerfile --progress=plain --no-cache -t go-practice .
# docker run --network=host go-practice
ARG UBI=registry.access.redhat.com/ubi9/ubi-minimal:latest

FROM jetpackio/devbox:latest AS builder

USER devbox
WORKDIR /devbox
COPY --chown=devbox . .
RUN devbox run build 
    
FROM $UBI

COPY --from=builder --chown=nobody /devbox/app/bin/go-practice /app/go-practice

USER nobody
EXPOSE 8080
ENTRYPOINT ["/app/go-practice"]