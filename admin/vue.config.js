module.exports = {
    chainWebpack: config => {
        config.module.rule('three').test(/\.js$/)
            .use('babel')
            .loader('babel-loader')
            .end()
        // GraphQL Loader
        config.module.rule('js').exclude.add(/\.worker\.js$/)
        config.module
            .rule('worker-loader')
            .test(/\.worker\.js$/)
            .use('worker-loader')
            .loader('worker-loader')
            .options({
                inline: 'fallback', chunkFilename: 'workerName.[chunkhash].js'
              })
            .end()
        // config.module.rule('styl').options({
        //     import:[
        //         path.join(__dirname, "./src/main.styl")
        //     ]
        // })
    },
    pages:{
        // forgotpwd: {          
        //     entry: 'src/mima.js',          
        //     template: 'public/mima.html',        
        // },        
        index: {          
            entry: 'src/main.js',          
            template: 'public/index.html',        
        },  
    },//配置多入口
    publicPath: '/admin',
    integrity: true,
    configureWebpack: config => {
        config.entry.app = ["babel-polyfill", "./src/main.js"];
     }
}
