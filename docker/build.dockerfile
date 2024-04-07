# Stage 1 - building client
FROM node:21-alpine AS client
WORKDIR /app
COPY ./client/*.json ./client/vite.config.ts ./client/index.html  ./
COPY ./client/postcss.config.js ./client/tailwind.config.js  ./
COPY ./client/public/ public/
COPY ./client/src/ src/
RUN npm install && npm run build

# Stage 2 - building backend
FROM golang:alpine3.19 AS builder

WORKDIR /src/backend

COPY ./backend/go.mod ./backend/go.sum ./
RUN go mod download

COPY ./backend .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o app ./cmd/main.go


# Stage 3 - compose all up
FROM alpine:3.19 AS deploy

WORKDIR /opt

COPY --from=client /app/build /opt/build
COPY --from=builder /src/backend/app .
COPY --from=builder /src/backend/config.yaml .
RUN apk --no-cache add ca-certificates tzdata

EXPOSE 8000
CMD ["/opt/app"]

