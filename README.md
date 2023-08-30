# URL Shortener

For learning purposes I tried to build a URL Shortener. It is very close to Alan Bardales' solution [here](https://alan-g-bardales.medium.com/simple-url-shortener-with-go-gin-and-mongodb-87f5e13dbbae). It will deviate in the future since I want to extend on his solution.

## Next steps

- [x] Add Swagger UI
  - Followed article [here](https://santoshk.dev/posts/2022/how-to-integrate-swagger-ui-in-go-backend-gin-edition/)
  - Run swag init always you change the docs for the api. It uses comments.
  - [x] Add response structs (some maybe in common?) --> Isolate health and root endpoints for swagger doc!
    - No idea if this is good or just over engineered
  - [x] Redirect from /docs to /docs/index.html
- [ ] Build simple frontend on top (maybe Vue3?)
- [ ] Dockerize application
- [ ] Build a browser extension that resolves short links before you click on it
- [ ] Add SSL support
