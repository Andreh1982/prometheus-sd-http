FROM golang:1.17.1 as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

ENV GOOS linux
ENV GOARCH ${GOARCH:-amd64}
ENV CGO_ENABLED=0

RUN go build -v -o prom-http-sd main.go

RUN apt-get update && apt-get install -y xz-utils && rm -rf /var/lib/apt/lists/*
ADD https://github.com/upx/upx/releases/download/v3.95/upx-3.95-amd64_linux.tar.xz /usr/local
RUN xz -d -c /usr/local/upx-3.95-amd64_linux.tar.xz | tar -xOf - upx-3.95-amd64_linux/upx > /bin/upx && chmod a+x /bin/upx

RUN upx --brute prom-http-sd

##########################
## user creator
##########################
FROM alpine:latest as user

ENV APP_HOME /build
ENV APP_USER olist
ENV APP_GROUP olist

RUN addgroup -S ${APP_GROUP} && adduser -S ${APP_USER} -G ${APP_GROUP}  --no-create-home
RUN apk --no-cache add ca-certificates \
    && update-ca-certificates

COPY --from=builder ${APP_HOME}/prom-http-sd ${APP_HOME}/prom-http-sd


RUN sed 's/ash/nologin/g' /etc/passwd

ENV GIN_MODE release

# ################################
# ## generate clean, final image
# ################################ 
# FROM scratch

# ENV APP_HOME /build
# ENV APP_USER olist

# ARG VERSION
# ENV APP_VERSION $VERSION

# COPY --from=user ${APP_HOME}/prom-http-sd ${APP_HOME}/prom-http-sd
# COPY --from=user /etc/passwd /etc/passwd
# COPY --from=user /etc/group /etc/group
# COPY --from=user /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# USER $APP_USER
# WORKDIR $APP_HOME

EXPOSE 9990

CMD ["/build/prom-http-sd"]
