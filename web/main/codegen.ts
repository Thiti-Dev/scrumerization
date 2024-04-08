
import type { CodegenConfig } from '@graphql-codegen/cli';

const config: CodegenConfig = {
  overwrite: true,
  schema: process.env.NUXT_PUBLIC_GRAPHQL_ENDPOINT || "http://localhost:8080/query",
  documents: "**/*.vue",
  generates: {
    ".gen/gql/": {
      preset: "client",
      plugins: []
    },
    "./graphql.schema.json": {
      plugins: ["introspection"]
    }
  }
};

export default config;
