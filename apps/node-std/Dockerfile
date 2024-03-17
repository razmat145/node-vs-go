FROM docker.io/node:lts-alpine

WORKDIR /app

RUN addgroup --system node-std && \
  adduser --system -G node-std node-std

COPY dist/apps/node-std node-std
RUN chown -R node-std:node-std .

RUN npm --prefix node-std --omit=dev -f install

CMD [ "node", "node-std" ]
