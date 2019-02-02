module.exports = {
    chainWebpack: webpackConfig => {
      webpackConfig.module
       .rule('wasm')
         .test(/.wasm$/)
         .use('wasm-loader')
         .loader('wasm-loader')
    }
  }
  