FROM --platform=$BUILDPLATFORM golang:1.21.3-alpine AS builder

#RUN apk update && apk add --no-cache git
#
#WORKDIR /app
#
#COPY . .
#
#RUN go mod tidy
#
#RUN go build -o binary
#
#ENTRYPOINT ["/app/binary"]

WORKDIR /app

ENV CGO_ENABLED 0
ENV GOPATH /go
ENV GOCACHE /go-build

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod/cache \
    go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod/cache \
    --mount=type=cache,target=/go-build \
    go build -o bin/product-management main.go

CMD ["/app/bin/product-management"]

FROM scratch AS app-release
COPY --from=builder /app/bin/product-management /usr/local/bin/go-product-management/apps
COPY --from=builder /app/.env .

CMD ["/usr/local/bin/go-product-management/apps"]