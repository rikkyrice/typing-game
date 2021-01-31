module.exports = {
  publicPath: process.env.VUE_APP_PUBLIC_PATH,
  css: {
    loaderOptions: {
      sass: {
        prependData: `@import "@/style.scss";`,
      },
    },
  },
  devServer: {
    proxy: process.env.VUE_APP_API_BASE_URL,
  },
};
