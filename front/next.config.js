/** @type {import('next').NextConfig} */

const isProd = (process.env.NODE_ENV === 'production' || process.env.DETA_ENV === 'PROD')
/**
 * GO API SERVICE RUNS in / and redirects to app
 * NEXT APP SHOULD RUN in /app
 */
const nextConfig = {
  basePath: isProd ? "" : "/app",
  assetPrefix: "/app",
  output: "standalone",
};

module.exports = nextConfig
