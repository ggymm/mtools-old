import { defineConfig } from 'vite'

import { getSrcPath, getWailsPath } from './build/utils'
import { createVitePlugins } from './build/plugin'

// noinspection JSUnusedGlobalSymbols
export default defineConfig(() => {
  const srcPath = getSrcPath()
  const wailsPath = getWailsPath()

  return {
    server: {
      host: '0.0.0.0'
    },
    resolve: {
      alias: {
        '@': srcPath,
        '~': wailsPath
      }
    },
    plugins: createVitePlugins(),
    build: {
      outDir: 'dist',
      target: 'esnext',
      reportCompressedSize: false,
      chunkSizeWarningLimit: 1024
    }
  }
})
