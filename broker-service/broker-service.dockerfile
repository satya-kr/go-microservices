# #base go image
# FROM golang:1.20-alpine as builder

# #creare a dir app
# RUN mkdir /app

# # we use . fro copy everything from current folder
# # now copy everything from current folder to app
# COPY . /app

# WORKDIR /app

# # we not gona use any "c/c++" library so disable CGO
# RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

# # give permission or make sure its executable
# RUN chmod +x /app/brokerApp

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

# COPY --from=builder /app/brokerApp /app
COPY brokerApp /app

CMD [ "/app/brokerApp" ]