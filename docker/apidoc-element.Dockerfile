FROM nginx:1.27.0
COPY ./docker/resource/api-doc.html /usr/share/nginx/html/index.html
COPY ./api/api.yaml /usr/share/nginx/html/openapi.yaml
