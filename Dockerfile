FROM --platform=${BUILDPLATFORM} golang:1.22-alpine AS builder
ARG TARGETOS
ARG TARGETARCH

WORKDIR /src
COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /main ./cmd/main.go

FROM python:3.12-bookworm AS final

COPY ./scripts/entrypoint.sh /entrypoint.sh
COPY --from=builder /main /main
EXPOSE 8194

ENTRYPOINT ["bash", "/entrypoint.sh"]
