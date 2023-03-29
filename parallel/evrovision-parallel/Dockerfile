FROM python:3.9.16 AS init

WORKDIR /init
COPY ./ ./
RUN ./etl.py

FROM eclipse-temurin:17-jdk-jammy

WORKDIR /app
COPY --from=init /init/sqlite.db .

CMD [ "ls -l" ]