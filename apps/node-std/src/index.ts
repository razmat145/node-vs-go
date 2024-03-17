import http, { IncomingMessage, ServerResponse } from 'http';

import Prom from './lib/Prom';

import { helloWorldHandler } from './lib/controllers/HelloWorld';

const server = http.createServer(
  Prom.applyPromMetrics(async (req: IncomingMessage, res: ServerResponse) => {
    if (req.url === '/hello' && req.method === 'GET') {
      return helloWorldHandler(req, res);
    }

    res.writeHead(404, { 'Content-Type': 'text/plain' });
    res.end('Not Found');
  })
);
const port = parseInt(process.env.API_PORT || '8991');

server.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});
