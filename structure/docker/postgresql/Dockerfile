FROM postgres:13.3

# args from docker-compose
ARG POSTGRES_USER
ARG POSTGRES_PASSWORD
ARG POSTGRES_DB

COPY init.sh /docker-entrypoint-initdb.d/

USER postgres

CMD ["postgres"]

EXPOSE 5432
