FROM alpine:latest

COPY main server

ENV PORT=80
# ENV GO_ENV="production"
# ENV GIN_MODE="release"

EXPOSE 80
CMD [ "/server" ]