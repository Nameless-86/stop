#STAGE ONE
FROM golang:1.23-alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o ./userapi

#STAGE TWO

FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /build/userapi ./userapi

CMD [ "/app/userapi" ]
