# Newton's Method Fractal Generation

Code to draw the Newton's method fractal for the solutions to z³ - 1 = 0. I've
written solutions several times in different languages/environments. It's meaty
enough to require some actual work but still tractable, particularly if don't
need to roll your own complex point structs/arithmetic.

Here, it offered a chance to learn go, it's tools and how golang packages are
typically structured.

It runs surprisingly slowly and I think it got even slower with a recent golang
update. About 1600ms on my MacBookAir. I made some limited attempts to profile
but got no further than learning that malloc seemed to be the culprit. Can't
tell which new's are the worst.
