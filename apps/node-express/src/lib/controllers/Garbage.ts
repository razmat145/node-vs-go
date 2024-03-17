import express, { Request, Response } from 'express';

const router = express.Router();

router.get('/garbage', (_: Request, res: Response) => {
  createGarbage();

  res.send('ok!');
});

function createGarbage() {
  for (let i = 0; i < 10_000; i++) {
    const garbage = {
      a: 'a',
      b: 'b',
      c: 'c',
      d: 'd',
      e: 'e',
      f: 'f',
      g: 'g',
      h: 'h',
      i: 'i',
      j: 'j',
      k: 'k',
      l: 'l',
      m: 'm',
      n: 'n',
      o: 'o',
      p: 'p',
      q: 'q',
      r: 'r',
      s: 's',
      t: 't',
      u: 'u',
      v: 'v',
      w: 'w',
      x: 'x',
      y: 'y',
      z: 'z',
    };
    const _ = garbage;
  }
}

export default router;
