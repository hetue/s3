ARG BINARY=s3

FROM ccr.ccs.tencentyun.com/storezhang/alpine:3.20.1 AS builder

# 复制执行程序
ARG BINARY
ARG TARGETPLATFORM=linux/amd64
COPY dist/${TARGETPLATFORM}/${BINARY} /docker/usr/local/bin/${BINARY}



FROM ccr.ccs.tencentyun.com/storezhang/alpine:3.20.1


LABEL author="storezhang<华寅>" \
    email="storezhang@gmail.com" \
    qq="160290688" \
    wechat="storezhang" \
    description="Drone持续集成Git插件，增加标签功能以及Github加速功能。同时支持推拉模式"


COPY --from=builder /docker /


# 执行命令
ENTRYPOINT ["/usr/local/bin/s3", "run"]
