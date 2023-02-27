FROM mysql:8.0

ENV MYSQL_ROOT_PASSWORD=root

RUN mkdir -p /try_go_clean_architecture/query_log

COPY ./container/db/conf.d /etc/mysql/conf.d

EXPOSE 3306
