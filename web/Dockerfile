FROM node:16

WORKDIR /gva_web/
COPY . .

# 通过npm下载cnpm
RUN npm install -g cnpm --registry=https://registry.npm.taobao.org

# 使用cnpm进行安装依赖
RUN cnpm install || npm install

# 使用npm run build命令打包web项目
RUN npm run build

FROM nginx:alpine
LABEL MAINTAINER="Jack@kesilent@live.com"

COPY .docker-compose/nginx/conf.d/my.conf /etc/nginx/conf.d/my.conf
COPY --from=0 /gva_web/dist /usr/share/nginx/html
RUN cat /etc/nginx/nginx.conf
RUN cat /etc/nginx/conf.d/my.conf
RUN ls -al /usr/share/nginx/html
