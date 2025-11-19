FROM alpine:latest

RUN addgroup -g 1000 app && \
    adduser -D -u 1000 -G app app

RUN apk --no-cache add curl

WORKDIR /app

COPY ./server server

COPY ./.conf .conf

RUN chown -R app:app /app

USER app

EXPOSE 3000

RUN chmod +x /app/server

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:3000/health || exit 1

CMD ["/app/server"]