var autoprefixer = require('autoprefixer')
var precss = require('precss')
var palette = require('postcss-color-palette')
var colorFunction = require('postcss-color-function')
var lost = require('lost')
var postcss_variables = require('postcss-advanced-variables')
var LiveReloadPlugin = require('webpack-livereload-plugin')

module.exports = {
  entry: './javascript/app.js',
    devtool: 'inline-source-map',
  output: {
    path: './public/js',
    filename: 'bundle.js'
  },
  babel: {
    presets: ['es2015', 'stage-0'],
    // plugins: ['transform-runtime']
    plugins: []
  },
  module: {
    loaders: [
      {
        test: /\.json$/, loader: 'json'
      },
      {
        test: /\.css$/,
        loader: 'style-loader!css-loader!postcss-loader'
      },
      {
        test: /\.js$/,
        exclude: /(node_modules|bower_components)/,
        loader: 'babel'
      },
      {
        test: /\.vue$/, loader: 'vue'
      },
      {
        test: /\.md$/,
        loader: 'html!markdown!emojione-loader'
      },
      {
        test: /\.svg$/,
        loader: 'svg-inline'
      }
    ]
  },
  postcss: function () {
    return {
        defaults: [
          autoprefixer,
          precss,
          palette({ palette: 'mrmrs' }),
          colorFunction,
          lost
        ]
    }
  },
  plugins: [
    new LiveReloadPlugin()
  ]
}
