FROM golang:alpine3.15 as builder

ARG swagger_version="latest"

RUN apk update \
 && apk add --no-cache \
        bash \
        binutils \
        gcc \
        git \
        make \
        musl-dev

RUN go install github.com/swaggo/swag/cmd/swag@${swagger_version}

WORKDIR /tmp/go
ADD . /tmp/go
ADD swagger /tmp/go
COPY config.yaml /tmp/go
ARG goreleaser_flags
RUN make build

# Export stage used by docker BuildKit#RUN apk update \
# && apk add --no-cache \
#        bash \
#        binutils \
#        gcc \
#        git \
#        make \
#        musl-dev
FROM scratch AS export
COPY --from=builder /tmp/go/dist .

FROM alpine:3.15
COPY --from=builder /tmp/go/onlyCloud /tmp/go/onlyCloud
COPY --from=builder /tmp/go/config.yaml .
COPY --from=builder /tmp/go/swagger/doc.json /app/swagger/doc.json
RUN apk update
WORKDIR .
CMD ["/tmp/go/onlyCloud"]
EXPOSE 8888