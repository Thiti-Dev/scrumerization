// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: ['@nuxtjs/tailwindcss', '@nuxtjs/apollo', ["@pinia/nuxt",{autoImports: ["defineStore", "acceptHMRUpdate"],}],
  ['@nuxtjs/google-fonts', {
    families: {
      Workbench: true,
    }
  }]
],
  apollo: {
    clients: {
      default: {
        httpEndpoint: process.env.NUXT_PUBLIC_GRAPHQL_ENDPOINT || "http://localhost:8080/query",
        wsEndpoint: process.env.NUXT_PUBLIC_WEB_SOCKET_ENDPOINT || 'ws://localhost:8080/query',
        // httpLinkOptions:{
        //   //# TODO: refresh token
        // }
      }
    },
  },
  routeRules: {
    '/app': {ssr: false},
    '/app/room/**' : {ssr:false}
  },
  runtimeConfig:{
    public:{
      version: 0 // Initial version, will be overrided by .env
    }
  }
})