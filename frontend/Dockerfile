FROM nginx:latest

RUN rm -f /etc/nginx/conf.d/default.conf
COPY site.conf /etc/nginx/conf.d/site.conf

RUN mkdir -p /var/www/html
COPY index.html /var/www/html
