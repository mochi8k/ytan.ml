var webpack = require('webpack');

module.exports = {
    entry: `${__dirname}/src/main.js`,
    output: {
        path: `${__dirname}/build`,
        publicPath: `${__dirname}/build`,
        filename: 'build.js'
    },
    module: {
        rules: [
            {
                test: /\.js$/,
                exclude: /node_modules/,
                use: 'babel-loader'
            },
            {
                test: /\.vue$/,
                use: 'vue'
            }
        ]
    },
    devtool: 'inline-source-map',
    plugins: [
        new webpack.HotModuleReplacementPlugin()
    ]
};
