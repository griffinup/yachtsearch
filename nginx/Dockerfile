FROM nginx:1.13.12

COPY nginx.conf /etc/nginx/nginx.conf
COPY frontend /etc/nginx/www/

CMD ["nginx", "-g", "daemon off;"]
