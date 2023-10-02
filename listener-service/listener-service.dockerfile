# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

# COPY --from=builder /app/listenerApp /app
COPY listenerApp /app

CMD [ "/app/listenerApp" ]