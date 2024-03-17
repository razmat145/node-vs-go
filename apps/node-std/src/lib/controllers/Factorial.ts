import { IncomingMessage, ServerResponse } from 'http';

export function factorialHandler(_: IncomingMessage, res: ServerResponse) {
  const factorial = computeFactorial(100);

  res.writeHead(200, { 'Content-Type': 'text/plain' });
  res.end(factorial.toString());
}

function computeFactorial(n: number): number {
  if (n === 0) {
    return 1;
  }
  return n * computeFactorial(n - 1);
}
