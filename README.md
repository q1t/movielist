#### Personal Movies Watch List ![travis build label](https://magnum.travis-ci.com/sysashi/mymovieslist.svg?token=AbFANsv5Xi6mwLCZBaop&branch=master)

*Built with:*
- [Go](https://golang.org/)
- [Vue.js](http://vuejs.org)
- [Webpack](https://webpack.github.io)
- [Postgresql](http://www.postgresql.org)

##### *Build*
You have to install at least Go to serve the static files along with the database - sqlite or postgresql.

Set the variables at config/app.toml
```
[Database]
User = your postgresql user
Name = your postgresql database
```

To build assets, Run
```
npm i
./node_modules/webpack/bin/webpack.js
```

Run ```go build && ./crud```
