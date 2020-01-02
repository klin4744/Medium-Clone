module.exports = {
   entry: './static/index.js',
   module: {
      rules: [
         {
            test: /\.(js|jsx)$/,
            exclude: /node_modules/,
            use: ['babel-loader'],
         },
         {
            test: /\.css$/i,
            use: ['style-loader', 'css-loader'],
         },
      ],
   },
   resolve: {
      extensions: ['*', '.js', '.jsx'],
   },
   output: {
      path: __dirname + '/static/dist',
      publicPath: '/',
      filename: 'bundle.js',
   },
   devServer: {
      contentBase: './dist',
      historyApiFallback: true,
   },
};
