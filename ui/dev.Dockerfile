FROM node:18-buster as build-frontend

WORKDIR /app
COPY package.json package-lock.json index.html tsconfig.json tsconfig.node.json vite.config.ts tailwind.config.js postcss.config.js ./

RUN npm ci

CMD [ "npm", "run", "dev" ]