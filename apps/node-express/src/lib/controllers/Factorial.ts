import express, { Request, Response } from 'express';

const router = express.Router();

router.get('/factorial', (req: Request, res: Response) => {
  const factorial = computeFactorial(50);

  res.send(factorial);
});

function computeFactorial(n: number): number {
  if (n === 0) {
    return 1;
  }
  return n * computeFactorial(n - 1);
}

export default router;
