import { RequestListener, IncomingMessage, ServerResponse } from 'http';

import {
  Registry,
  Histogram,
  collectDefaultMetrics,
  Counter,
} from 'prom-client';

type AsyncRequestListener = (
  req: IncomingMessage,
  res: ServerResponse
) => Promise<void>;

class Prom {
  private registry = new Registry();

  private timingHistogram: Histogram;
  private requestCounter: Counter;

  private isInitialised = false;

  public applyPromMetrics(
    requestListener: AsyncRequestListener
  ): RequestListener {
    if (!this.isInitialised) {
      this.initialise();
    }

    return this.applyMiddlewareRecordHook(requestListener);
  }

  private initialise() {
    this.timingHistogram = new Histogram({
      name: 'http_response_time_seconds',
      help: 'Duration of HTTP requests.',
      labelNames: ['method', 'path'],
      registers: [this.registry],
    });
    this.requestCounter = new Counter({
      name: 'http_requests_total',
      help: 'Total number of HTTP requests.',
      labelNames: ['method', 'path'],
      registers: [this.registry],
    });

    collectDefaultMetrics({ register: this.registry });

    this.isInitialised = true;
  }

  private applyMiddlewareRecordHook(requestListener: AsyncRequestListener) {
    return async (req: IncomingMessage, res: ServerResponse) => {
      if (req.url !== '/metrics') {
        const end = this.timingHistogram.startTimer();

        await requestListener(req, res);

        end({
          method: req.method,
          path: req.url,
        });

        this.requestCounter.inc({
          method: req.method,
          path: req.url,
        });
      }

      await this.metricsRouteHandler(req, res);
    };
  }

  private async metricsRouteHandler(req: IncomingMessage, res: ServerResponse) {
    if (req.url === '/metrics' && req.method === 'GET') {
      const metrics = await this.registry.metrics();

      res.writeHead(200, { 'Content-Type': this.registry.contentType });
      return res.end(metrics);
    }
  }
}

export default new Prom();
