FROM golang:1.23.6-bookworm AS builder

# static build
ENV CGO_ENABLED=0

RUN --mount=type=bind,source=tools/replacetool,target=. go build -o /replacetool

FROM ubuntu:24.04 AS runtime

ENV DEBIAN_FRONTEND=noninteractive
RUN --mount=type=cache,target=/var/lib/apt,sharing=locked \
    --mount=type=cache,target=/var/cache/apt,sharing=locked \
    apt-get update -yqq && \
    apt-get install -y --no-install-recommends \
            devscripts \
            build-essential \
            cdbs

COPY template /template
COPY tools/replacetool/template /replacetool_template
COPY entrypoint.sh /usr/local/bin/
COPY --from=builder /replacetool /usr/local/bin/

ENTRYPOINT ["entrypoint.sh"]
