FROM golang:1.11.2-alpine

ENV GO111MODULE=on

RUN apk add --update git bzr

WORKDIR /ctx/
ADD ./ /ctx/

RUN CGO_ENABLED=0 go build -o /tautulli-metrics ./cmd/tautulli_metrics.go

FROM scratch

COPY --from=0 /tautulli-metrics /tautulli-metrics

ENTRYPOINT ["/tautulli-metrics"]
