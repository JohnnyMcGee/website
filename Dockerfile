FROM fholzer/nginx-brotli

RUN apt-get update && \
    apt-get install -y --no-install-recommends curl && \
    rm -rf /var/lib/apt/lists/*

COPY ./public /usr/share/nginx/html
COPY ./conf /etc/nginx/conf.d

