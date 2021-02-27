# Go embed directive

With Go 1.16 comes the `//go:embed` directive which allows for embedding additional assets within the binary. Before now, tools like [go-bindata](https://github.com/jteeuwen/go-bindata) and [go-bindata-assetfs](https://github.com/elazarl/go-bindata-assetfs) were required.

This snippet shows how to embed a SPA in your binary. For additional use cases, checkout the [official docs](https://golang.org/pkg/embed/)

```go
//go:embed website/*
var website embed.FS

func loadEmbeddedContent () http.Handler {
    handler, err := fs.Sub(website, "website")
    if err != nil {
        return http.NotFoundHandler()
    }

    return http.FileServer(http.FS(handler))
}
```

## Why is fs.Sub required?
If we were to just pass `website` variable to the FileServer, we would get a directory listing instead of the desired index.html.
