# Newton's Method Fractal Generation

Code to draw the Newton's method fractal for the solutions to z³ - 1 = 0. I've
written solutions several times in different languages/environments. It's meaty
enough to require some actual work but still tractable, particularly if don't
need to roll your own complex point structs/arithmetic.

Here, it offered a chance to learn go, it's tools and how golang packages are
typically structured.

It runs surprisingly slowly and I think it got a lot slower with a recent golang
update. About 1600ms on my MacBookAir. The slowness on the older golang version
seemed to be memory allocation related. The new slowness seems more related
to FP math.

```
(pprof) top10
950ms of 1250ms total (76.00%)
Showing top 10 nodes out of 60 (cum >= 50ms)
      flat  flat%   sum%        cum   cum%
     270ms 21.60% 21.60%      270ms 21.60%  runtime.(*mcentral).grow
     150ms 12.00% 33.60%      150ms 12.00%  math.Sincos
     120ms  9.60% 43.20%      120ms  9.60%  math.Hypot
     100ms  8.00% 51.20%      250ms 20.00%  math.Pow
      70ms  5.60% 56.80%      190ms 15.20%  math.atan2
      60ms  4.80% 61.60%       60ms  4.80%  math.xatan
      50ms  4.00% 65.60%       50ms  4.00%  math.modf
      50ms  4.00% 69.60%      110ms  8.80%  math.satan
      40ms  3.20% 72.80%      770ms 61.60%  github.com/ghurley/newton.newton64
      40ms  3.20% 76.00%       50ms  4.00%  math.frexp
```
