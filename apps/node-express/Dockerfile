FROM docker.io/node:lts-alpine

WORKDIR /app

RUN addgroup --system node-express && \
  adduser --system -G node-express node-express

COPY dist/apps/node-express node-express
RUN chown -R node-express:node-express .

RUN npm --prefix node-express --omit=dev -f install

CMD [ "node", "node-express" ]
