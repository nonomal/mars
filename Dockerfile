FROM --platform=$BUILDPLATFORM node:21-alpine as web-build

WORKDIR /app

COPY ./frontend .

RUN yarn install --registry=https://registry.npm.taobao.org && \
    yarn build

FROM --platform=linux/amd64 golang:1.23 AS builder

ARG TARGETARCH
ARG TARGETOS

RUN apt update && \
  apt install -y ca-certificates tzdata git gcc-aarch64-linux-gnu xz-utils && \
  wget $(curl -s https://api.github.com/repos/upx/upx/releases/latest \
    | grep browser_download_url | grep amd64 | cut -d '"' -f 4) -O upx.tar.xz && \
  tar -xvf upx.tar.xz && \
  cd upx-*-amd64_linux && \
  mv upx /bin/upx

WORKDIR /app

COPY . .

COPY --from=web-build /app/build /app/frontend/build

RUN go mod download

RUN if [ "$TARGETARCH" = "arm64" ]; then CC=aarch64-linux-gnu-gcc && CC_FOR_TARGET=gcc-aarch64-linux-gnu && EXTRA_FLAGS='-extldflags "-static"'; fi && \
    VERSION_PATH=$(go list -m -f "{{.Path}}" | grep -v api)/internal/version && LDFLAGS="-w -s  \
     -X ${VERSION_PATH}.gitBranch=$(git rev-parse --abbrev-ref HEAD) \
     -X ${VERSION_PATH}.buildDate=$(date -u +'%Y-%m-%dT%H:%M:%SZ') \
     -X ${VERSION_PATH}.gitCommit=$(git rev-parse --short HEAD) \
     -X ${VERSION_PATH}.gitTag=$(git describe --exact-match --tags HEAD 2> /dev/null || echo '') \
     -X ${VERSION_PATH}.kubectlVersion=$(go list -m -f '{{.Path}} {{.Version}}' all | grep k8s.io/client-go | cut -d ' ' -f2) \
     -X ${VERSION_PATH}.helmVersion=$(go list -m -f '{{.Path}} {{.Version}}' all | grep helm.sh/helm/v3 | cut -d ' ' -f2)" \
    && CGO_ENABLED=1 CC=$CC CC_FOR_TARGET=$CC_FOR_TARGET GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags="$LDFLAGS $EXTRA_FLAGS" -o /bin/app main.go \
    && upx -9 /bin/app

FROM --platform=$TARGETPLATFORM gcr.io/distroless/base-debian12

WORKDIR /

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /bin/app /bin/app

CMD ["app"]
