FROM node:20-alpine
WORKDIR /app

# Instala deps primeiro (melhor cache)
COPY package*.json ./
RUN npm ci

EXPOSE 3000
CMD ["npm","run","dev"]