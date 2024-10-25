# image to run the app
FROM alpine:latest

RUN mkdir /app

COPY authApp /app

CMD ["/app/authApp"]