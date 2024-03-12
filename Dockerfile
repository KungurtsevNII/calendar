FROM golang:1.22-alpine3.19 as builder

ENV CGO_ENABLED 0
ENV GOOS linux

RUN apk update --no-cache \
    && apk add --no-cache tzdata \
    && apk add --no-cache make \
    && apk add --no-cache curl

WORKDIR /calendar
COPY . .

RUN make build

FROM scratch

EXPOSE 5001

COPY --from=builder /calendar/bin/calendar-service /bin/calendar-service
COPY --from=builder /calendar/config /bin/config

ENTRYPOINT ["/bin/calendar-service"]