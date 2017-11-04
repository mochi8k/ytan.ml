module.exports = {
    entry: './src/main.js',
    output: {
        path: `${__dirname}/build`,
        filename: 'build.js'
    },
    module: {
        loaders: [
            { test: /\.vue$/, loader: 'vue' },
        ]
    }
};
