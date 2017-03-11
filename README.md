#go-iris-reactjs-webpack

Starting template for websites with Golang backend, Iris and ReactJS with webpack and hot reload. In addition to that, go-bindata is used in the production mode.

Borrowed a lot of ideas from [https://github.com/codeskyblue/go-reactjs-example](https://github.com/codeskyblue/go-reactjs-example)

### Development
To setup the development environment

    npm install -g webpack
    npm install
    make dev

Now open VSCode and start debugging  
Open in your browser the following url: [http://127.0.0.1:8080](http://127.0.0.1:8080)vv
Edit file `app/containers/app.jsx` and the site should refresh and you should see your changes immediately.

## Production
To prepare the site for production, just run the following command

    make prod

Now you can start the production ready site

    go run main.go

## Thanks

[https://github.com/codeskyblue/go-reactjs-example](https://github.com/codeskyblue/go-reactjs-example)
