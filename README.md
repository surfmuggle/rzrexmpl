# Intro

For go i want to use the [gorazor template engine](https://github.com/sipin/gorazor/tree/v1.2.2). 
I used code from gorazor repository and from [go template benchmark](https://github.com/slinso/goTemplateBenchmark/blob/master/gorazor/tpl/layout/base.gohtml) and ran `gorazor tpl tpl` and got `panic: Can't find layout: rzrexmpl/tpl/layout/base [Home]`:

```bash
gorazor tpl tpl
gorazor processing dir: tpl -> tpl
panic: Can't find layout: rzrexmpl/tpl/layout/base [Home]
``` 
The project `rzrexmpl` looks like this

```bash
rzrexmpl
├─── main.go
├───models
|     └─── user.go
└───tpl
    ├───helper
    |     ├─── footer.gohtml
    |     ├─── header.gohtml
    |     └─── msg.gohtml
    ├───layout
    |       └─── base.gohtml
    └─── home.gohtml
```

See [this question at stackoverflow `gorazor-cant-find-layout`](https://stackoverflow.com/questions/79808592/gorazor-cant-find-layout)
