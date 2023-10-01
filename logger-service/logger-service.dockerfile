# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

# COPY --from=builder /app/loggerServiceApp /app
COPY loggerServiceApp /app

CMD [ "/app/loggerServiceApp" ]