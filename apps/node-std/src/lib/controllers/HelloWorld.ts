import { IncomingMessage, ServerResponse } from 'http';

export function helloWorldHandler(_: IncomingMessage, res: ServerResponse) {
  res.writeHead(200, { 'Content-Type': 'text/plain' });
  res.end('Hello, World!');
}
