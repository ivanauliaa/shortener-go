# Readme

URL shortener built with GDM (Go, Docker, mongoDB) tech stack and multi stage builds

- Gin for http server and router management
- shortid for short URL generations
- mongo express for mongoDB monitoring
- Multi stage docker builds for creatine tiny Go images

## POST Request Example

```curl
curl -X POST -d '{"longURL": "https://ivanauliaa.netlify.app"}' localhost:5100/shorten
```

Response:
```json
{
  "data": {
    "expires": "2022-03-20 13:27:50",
    "id": "623094563e026bd6e8e0ba8a",
    "new_url": "http://localhost:5100/VJ7ofvPnR"
  },
  "message": "Success create new short URL"
}
```
