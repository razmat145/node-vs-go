import express from 'express';

import Prom from './lib/Prom';

import HelloWorldRouter from './lib/controllers/HelloWorld';
import FactorialRouter from './lib/controllers/Factorial';

const app = express();
const port = parseInt(process.env.API_PORT || '8989');

Prom.applyPromMetrics(app);

app.use(HelloWorldRouter);
app.use(FactorialRouter);

app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});
