FROM --platform=$BUILDPLATFORM golang:1.22.1 as builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /pg-backuper .

FROM alpine:latest

RUN apk --no-cache add ca-certificates postgresql-client

COPY --from=builder /pg-backuper /pg-backuper

ENTRYPOINT ["/pg-backuper"]

CMD [ "start" ]

# docker buildx create --name multiarch --driver docker-container --use
# docker buildx build --push --platform linux/amd64,linux/arm64 -t ghcr.io/blesswinsamuel/infra-base/pg-backuper:latest .
