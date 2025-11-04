# Intro

For go i want to use the [gorazor template engine](https://github.com/sipin/gorazor/tree/v1.2.2). 
I used code from gorazor repository and from [go template benchmark](https://github.com/slinso/goTemplateBenchmark/blob/master/gorazor/tpl/layout/base.gohtml) and ran `gorazor tpl tpl` and got `panic: Can't find layout: rzrexmpl/tpl/layout/base [Home]`. To fix this you have to use `-prefix ...` like this `gorazor -prefix rzrexmpl ./tpl ./tpl`

```bash
gorazor -prefix rzrexmpl ./tpl ./tpl
gorazor processing dir: ./tpl -> ./tpl
tpl\helper\footer.gohtml -> C:\dev2\go\rzrexmpl\tpl\helper\footer.go
tpl\helper\header.gohtml -> C:\dev2\go\rzrexmpl\tpl\helper\header.go
tpl\layout\base.gohtml -> C:\dev2\go\rzrexmpl\tpl\layout\base.go
tpl\helper\msg.gohtml -> C:\dev2\go\rzrexmpl\tpl\helper\msg.go
tpl\home.gohtml -> C:\dev2\go\rzrexmpl\tpl\home.go
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
This repo has all the code to reproduce the error. 
It is unclear to me what i have to change either in the template files or in the imports. To be able to generate the code from the `gohtml` files. 
See [this question at stackoverflow `gorazor-cant-find-layout`](https://stackoverflow.com/questions/79808592/gorazor-cant-find-layout)
