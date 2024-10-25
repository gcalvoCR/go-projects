# image to run the app
FROM alpine:latest

RUN mkdir /app

COPY loggerServiceApp /app

CMD ["/app/loggerServiceApp"]