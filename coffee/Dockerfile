FROM golang:1.23.2-alpine3.20 as builder
WORKDIR /app
COPY ./src .
RUN go build -o coffee . 
RUN mkdir -p /rootfs/usr/local/etc/containers
RUN mkdir -p /rootfs/usr/local/lib/containers/coffee
RUN mv coffee /rootfs/usr/local/lib/containers/coffee/coffee
COPY manifest.yaml /

FROM scratch
COPY --from=builder /rootfs/usr/local /rootfs/usr/local
COPY --from=builder /manifest.yaml /
COPY coffee.yaml /rootfs/usr/local/etc/containers/coffee.yaml
WORKDIR /rootfs/usr/local/lib/containers/coffee
COPY ./src/tmpl ./tmpl
COPY ./src/static ./static
ENTRYPOINT ["./coffee"]
