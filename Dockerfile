FROM ccr.ccs.tencentyun.com/storezhang/alpine:3.20.1 AS builder

# 复制脚本程序
COPY docker /docker
# 复制执行程序
ARG TARGETPLATFORM=linux/amd64
COPY dist/${TARGETPLATFORM}/git /docker/usr/local/bin/



FROM ccr.ccs.tencentyun.com/storezhang/alpine:3.20.1

LABEL author="storezhang<华寅>" \
    email="storezhang@gmail.com" \
    qq="160290688" \
    wechat="storezhang" \
    description="Drone持续集成Git插件，增加标签功能以及Github加速功能。同时支持推拉模式"

COPY --from=builder /docker /

RUN set -ex \
    \
    \
    \
    && apk update \
    \
    # 安装工具
    && apk --no-cache add git openssh-client \
    \
    \
    \
    # 增加执行权限
    && chmod +x /usr/local/bin/* \
    \
    \
    \
    && rm -rf /var/cache/apk/*

# 执行命令
ENTRYPOINT /usr/local/bin/fastgit
