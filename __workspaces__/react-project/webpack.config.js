const path = require("path");

module.exports = {

    resolve : {
        extensions : ['.js', '.jsx']
    },
    entry:{ 
        app : ['./src/index.jsx'], 
    },
    output:{
        path: path.join(__dirname, 'dist'),
        filename : '[name].bundle.js'  
    },
    module : { 
        rules : [{
            test: /\.jsx?/,
            loader: 'babel-loader',
            options: {
                presets: ['@babel/preset-env', ["@babel/preset-react", {"runtime": "automatic"}]],
            },
        }],
    },
    devServer: {
        static: {
            directory: path.join(__dirname, 'public'),
        },
        compress: true,
        port: 9000,
    },
};