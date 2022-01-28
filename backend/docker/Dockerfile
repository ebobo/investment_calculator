FROM golang:alpine AS builder
RUN apk add --no-cache git
RUN apk add --no-cache build-base
WORKDIR /build
COPY . .
RUN make init
RUN make build

FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /hbb
COPY --from=builder /build/bin/ic_server /bin/ic_server
CMD ./ic_server