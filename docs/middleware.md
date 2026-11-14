# Middleware

Constellation ships a small set of middleware in `config/middleware/`. Most are
wired automatically during route registration; you rarely call them directly.

## CORS

`CorsSetup()` registers [gin-contrib/cors](https://github.com/gin-contrib/cors):

- **Origins** — `ALLOWED_ORIGINS` (comma-separated) or the defaults
  `http://localhost:3000` and `http://localhost:5173`.
- **Methods** — `GET, POST, PATCH, PUT, DELETE, OPTIONS`.
- **Headers** — `Origin, Content-Type, Authorization, X-API-Key`.
- **Exposed** — `Content-Length`.
- Credentials are allowed and preflight results are cached for 12 hours.

It also installs a response-header middleware that sets `Developer: MPOP Reverse II`
on every response.

## Bot protection

`BotProtection()` runs on every registered route. It inspects the `User-Agent`
and rejects requests that match a list of known scraper agents (`scrapy`,
`python-requests`, `httpx`) with `403 Access denied`.

> This is a heuristic, not a security boundary. It matches the raw `User-Agent`
> against the blocked list, so truncated or spoofed agents can slip through.
> Harden it (case-insensitive token matching, rate limiting, IP rules) before
> relying on it in production.

## Cookie access control

`RequestHandlerCookie()` guards routes with `Permission: PERMISSION_COOKIE`. It
requires a `temporary` cookie whose value is at least 30 characters; otherwise it
returns `403 Unauthorized`.

Because `Register` defaults `POST` routes to `COOKIE` permission, form submissions
are protected by default. Set `Permission: utils.PERMISSION_ALL` to opt a `POST`
route out.

## Request logging

`StartServer` adds a logger that prints the method, path, status, and a timestamp
for each request:

```
[2026-11-14 02:07:31] [GET] / -> 200
```

`gin.Default()` also logs requests. See the ordering note in
[architecture.md](architecture.md) — middleware added after routes do not affect
them.

## Custom 404

`StartServer` installs a `NoRoute` handler returning an HTML 404 page.

## Adding your own middleware

Define a `gin.HandlerFunc` and either:

- register it globally **before** routes with `app.Use(...)`, or
- wrap a handler at registration time.

For a per-route concern, the cleanest option is to extend `Register` with a new
permission or route option rather than calling `app.Use` after the fact.
