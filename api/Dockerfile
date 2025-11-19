FROM alpine:latest

RUN addgroup -g 1000 app && \
    adduser -D -u 1000 -G app app

WORKDIR /app

COPY ./server server

RUN chmod +x /app/server

COPY ./.conf .conf

RUN chown -R app:app /app

USER app

EXPOSE 3000

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:3000/health || exit 1

CMD ["/app/server"]