FROM golang:1.22 as build

WORKDIR /go/src/lab-4
COPY . .

RUN go test ./...
ENV CGO_ENABLED=0
RUN go install ./cmd/...

# ==== Final image ====
FROM alpine:latest
WORKDIR /opt/lab-4
COPY entry.sh /opt/lab-4/
COPY --from=build /go/bin/* /opt/lab-4
RUN ls /opt/lab-4
ENTRYPOINT ["/opt/lab-4/entry.sh"]
CMD ["server"]