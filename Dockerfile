# FROM golang:latest

# WORKDIR /app

# COPY go.mod go.sum ./

# RUN go mod download

# COPY . .

# RUN go build -o main .

# EXPOSE 7000

# CMD ["./main"]

# --------------------------------------------------------------

FROM golang:alpine as builder

WORKDIR /go/src
    
COPY . .
    
RUN go get
    
RUN go build -o /go/bin/backend_api
    
FROM alpine
    
RUN apk update && apk add --no-cache tzdata
    
# COPY config.yaml /app/config.yaml
    
COPY --from=builder /go/bin/backend_api /backend_api
    
ENV TZ="Asia/Bangkok"
    
ENTRYPOINT [ "/backend_api" ]

# ต้นแบบ

# FROM golang:alpine as builder

# WORKDIR /go/src

# COPY . .

# RUN go get

# RUN go build -o /go/bin/app_api

# FROM alpine

# RUN apk update && apk add --no-cache tzdata

# COPY --from=builder /go/bin/app_api /app_api

# ENV TZ="Asia/Bangkok"

# ENTRYPOINT [ "/app_api" ]