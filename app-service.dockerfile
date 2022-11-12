FROM alpine:latest

RUN mkdir /app

# COPY --from=builder /app/brokerApp /app
COPY taskApp /app

CMD [ "/app/taskApp" ]