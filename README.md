# Intro

For go i want to use the [gorazor template engine][1]. 
I used code from gorazor repository and from [go template benchmark][2] and ran `gorazor tpl tpl` and got `panic: Can't find layout: rzrexmpl/tpl/layout/base [Home]`:

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
