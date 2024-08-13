const path = require('path');
module.exports = {
    entry: {
        count: './src/count.tsx'
    },
    output: {
        filename: "[name].js",
        path: path.resolve(__dirname, 'dist')
    },
    resolve: {
        extensions: ['.ts', '.tsx', '.js']
    },
    module: {
        rules: [
            {
                test: /\.(ts|tsx)$/,
                exclude: /node_modules/,
                use: "babel-loader"
            }
        ]
    },
    devServer: {
        devMiddleware: {
            publicPath: '/dist/',
            writeToDisk: true
        }
    }
}