var webpack = require('webpack');
var path = require('path');
var WebpackNotifierPlugin = require('webpack-notifier');

module.exports = {
    devtool: 'eval',
    // This will be our app's entry point (webpack will look for it in the 'src'
    // directory due to the modulesDirectory setting below). Feel free to change as
    // desired.
    entry: [
        // Add the react hot loader entry point - in reality, you only want this in your
        // dev Webpack config
        'react-hot-loader/patch',
        'webpack-dev-server/client?http://localhost:3000',
        'webpack/hot/only-dev-server',
        './public/entry.js'
    ],
    // Output the bundled JS to dist/app.js
    output: {
        filename: 'bundle.js',
        publicPath: '/bundle/',
        path: path.resolve('bundle')
    },
    resolve: {
        // Look for modules in .ts(x) files first, then .js(x)
        extensions: [
            '.js', '.jsx'
        ],
        // Add 'src' to our modulesDirectories, as all our app code will live in there,
        // so Webpack should look in there for modules
        modules: ['app', 'node_modules']
    },
    module: {
        loaders: [
            {test: /\.css$/, loader: "style-loader!css-loader"},
            {
                test: /\.jsx?$/,
                loaders: ['babel-loader']
            }
        ]
    },
    plugins: [
        // Add the Webpack HMR plugin so it will notify the browser when the app code
        // changes
        new webpack.HotModuleReplacementPlugin(),
        // Set up the notifier plugin - you can remove this (or set alwaysNotify false)
        // if desired
        new WebpackNotifierPlugin({alwaysNotify: true}),
        new webpack.ProvidePlugin({
			$: "jquery",
			jQuery: "jquery",
			"window.jQuery": "jquery"
		})
    ]
};