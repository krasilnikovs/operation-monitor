FROM golang:1.24.1-alpine as base
WORKDIR /app
COPY . .

FROM base AS dev
RUN go install github.com/air-verse/air@latest
CMD ["air", "-c", ".air.toml"]