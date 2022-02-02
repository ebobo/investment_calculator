FROM golang:alpine AS builder
# After building run "docker image prune --filter label=stage=builder" to remove builder image
LABEL stage=builder
RUN apk add --no-cache git
RUN apk add --no-cache build-base
WORKDIR /build
COPY . ./
RUN make init
RUN make gen
RUN make build

FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /record-ms
# Copy binary
COPY --from=builder /build/bin/record ./
EXPOSE 9094
CMD ./record