# Newton's Method Fractal Generation

Code to draw the Newton's method fractal for the solutions to z³ - 1 = 0. I've
written solutions several times in different languages/environments. It's meaty
enough to require some actual work but still tractable, particularly if you
don't need to roll your own complex point structs/arithmetic.

Here, it offered a chance to learn go, its tools, and how golang packages are
typically structured.

It runs surprisingly slowly and I think it got a lot slower with a recent golang
update. About 1300ms on my MacBookAir. The most recent of my prior
implementations of this code was in JavaScript and it computed the same size
image in under 200ms on the same hardware.

The slowness on the older golang version
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

I'm surprised to see all the trig functions (and satan!) which I don't invoke
directly so they may all come from the `cmplx.Abs()`. Even if I managed to get
rid of all of that slowness, the runtime.(\*mcentral).grow would still make
this implementation slower than the JS version. According to my "research" (a
10 second perusal of Google search results) grow is memory/heap allocation.
The fact that there is no listing for `sweep` at least means that I don't need
to fight the GC. Still, 270ms for memory allocation is insane. Maybe I should
preallocate a fixed size array of `color.RGBA`s instead of creating
and returning them one at a time.
