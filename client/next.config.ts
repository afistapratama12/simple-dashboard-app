import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  /* config options here */
  // ignore lint for type any
  eslint: {
    ignoreDuringBuilds: true,
  },
};

export default nextConfig;
