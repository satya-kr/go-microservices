# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

# COPY --from=builder /app/authApp /app
COPY authApp /app

CMD [ "/app/authApp" ]