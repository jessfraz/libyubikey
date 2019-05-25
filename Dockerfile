FROM alpine:edge
MAINTAINER Jessica Frazelle <jess@linux.com>

ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go

RUN	apk add --no-cache \
	asciidoc \
	autoconf \
	automake \
	bash \
	build-base \
	ca-certificates \
	git \
	gcc \
	go \
	libc-dev \
	libgcc \
	libtool

RUN go get github.com/xlab/cgogen \
	&& go get golang.org/x/lint/golint

ENV YUBICOC_VERSION 1.13
RUN git clone --depth 1 -b "libyubikey-${YUBICOC_VERSION}" \
	https://github.com/Yubico/yubico-c.git /usr/src/yubico-c \
	&& ( \
		cd /usr/src/yubico-c \
		&& autoreconf --install \
		&& ./configure --prefix=/usr \
		&& make \
		&& make install \
	)

ENV PKG_CONFIG_PATH /usr/lib/pkgconfig:/usr/src/yubico-c
COPY . /go/src/github.com/jessfraz/libyubikey
WORKDIR /go/src/github.com/jessfraz/libyubikey
