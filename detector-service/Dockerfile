
ARG BASE
FROM golang:1.15 as builder
COPY . /go/src
# `skaffold debug` sets SKAFFOLD_GO_GCFLAGS to disable compiler optimizations
ARG SKAFFOLD_GO_GCFLAGS
ENV GOPATH=/go
ENV GOBIN=/go/bin
ENV PATH=/go/bin:$PATH
ENV GO111MODULE=on
WORKDIR /go/src
RUN go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -o /app .

FROM $BASE
COPY --from=builder /app .