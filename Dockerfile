FROM fholzer/nginx-brotli

RUN apk add --no-cache curl

COPY ./public /usr/share/nginx/html
COPY ./conf /etc/nginx/conf.d

