const path = require('path');
const MiniCssExtractPlugin = require("mini-css-extract-plugin");

module.exports = {
    mode: 'production',
    entry: {
        sass: './app/sass/src/index.js',
        js: './app/js/src/index.js'
    },
    output: {        
        filename: 'app/[name]/dist/bundle.js',
        path: path.resolve(__dirname, ''),
    },
    devtool: 'source-map',
    watchOptions: {
        ignored: /node_modules/,
    },
    module: {
        rules: [
        {
            test: /\.s[ac]ss$/i,
            use: [
                MiniCssExtractPlugin.loader,
                {
                    loader: 'css-loader',
                    options: {
                        url: false,
                        sourceMap: true
                    }
                },
                {
                    // Compiles Sass to CSS
                    loader: "sass-loader",
                    options: {
                        sourceMap: true
                    }
                }
            ],
        },
        {
            test: /\.(png|svg|jpg|gif|woff2)$/i,
            type: 'asset/resource',
        },
        ],
    },
    resolve: {
        extensions: ['.ts', '.js'],
    },
    plugins: [
        new MiniCssExtractPlugin({            
            filename: "./app/sass/dist/bundle.css",
            chunkFilename: "[id].css",
        }),
    ]
};