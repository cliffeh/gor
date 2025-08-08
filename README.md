# gor

I wanted to put together a template golang web server project. Some of the
things I was aiming for:

* Using the Go standard library to the extent possible (no/minimal dependencies)
* A few example endpoints, including some that serve `application/json` content
  serialized from a struct
* A reasonably complete unit/integration test suite
* Github actions for running tests on PR/merge and publishing "production"
  container images

For a list of build targets take a look at the [Makefile](./Makefile) or just do
`make help`.

For some additional things I'd like to do in this repo see [TODO.md](./TODO.md).

**Attribution:** I drew (_heavy_) inspiration from this blog post:
<https://blog.arcjet.com/building-a-minimalist-web-server-using-the-go-standard-library-tailwind-css/>
