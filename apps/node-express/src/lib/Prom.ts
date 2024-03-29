import { Application, Request, Response, Next } from 'express';

import {
  Registry,
  Histogram,
  collectDefaultMetrics,
  Counter,
} from 'prom-client';

class Prom {
  private registry = new Registry();

  private timingHistogram: Histogram;
  private requestCounter: Counter;

  private isInitialised = false;

  public applyPromMetrics(app: Application) {
    if (!this.isInitialised) {
      this.initialise();
    }

    this.applyMiddlewareRecordHook(app);
    this.applyPromMetricsRoute(app);
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

  private applyMiddlewareRecordHook(app: Application) {
    app.use((req: Request, res: Response, next: Next) => {
      if (req.path !== '/metrics') {
        const end = this.timingHistogram.startTimer();

        res.on('finish', () => {
          end({
            method: req.method,
            path: req.path,
          });

          this.requestCounter.inc({
            method: req.method,
            path: req.path,
          });
        });
      }

      next();
    });
  }

  private applyPromMetricsRoute(app: Application) {
    app.get('/metrics', async (_: Request, res: Response) => {
      const metrics = await this.registry.metrics();

      res.set('Content-Type', this.registry.contentType);
      return res.end(metrics);
    });
  }
}

export default new Prom();
