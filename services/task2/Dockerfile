ARG BASE
ARG COMMON
FROM $COMMON as builder

ENV GOPATH=/go
ENV GOBIN=/go/bin
ENV PATH=/go/bin:$PATH
ENV GO111MODULE=on
WORKDIR /go/src/svc
COPY go.mod /go/src/svc/go.mod
COPY go.sum /go/src/svc/go.sum
RUN go mod download
COPY . /go/src/svc
# `skaffold debug` sets SKAFFOLD_GO_GCFLAGS to disable compiler optimizations
ARG SKAFFOLD_GO_GCFLAGS
RUN go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -o /app .

FROM $BASE
COPY --from=builder /app .