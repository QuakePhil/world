Yet another 2d canvas engine

Powered by gorilla websocket by way of melody

# Usage

```
go run . <world>
```

Then, point a browser to `localhost:8081`!

`<world>` is a folder with `.js` and `.go` code that renders the frontend and processes the backend

It listens on `localhost:8081` by default, if that doesn't work, try changing the value of `config.address` in `init.go`

# Examples

## Cursor
Barebones example that draws a circle under the cursor
```
go run . cursor
```

## Bouncyball
Launch particles and watch them bounce
```
go run . bouncyball
```
