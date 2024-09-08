import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: [
      { find: "@images", replacement: '/src/assets/images'},
      { find: "@components", replacement: '/src/components'},
      { find: "@interface", replacement: '/src/interface'},
      { find: "@layout", replacement: '/src/layout'},
      { find: "@pages", replacement: '/src/pages'},
      { find: "@router", replacement: '/src/router'},
      { find: "@store", replacement: '/src/store'}, 
      { find: "@utils", replacement: '/src/utils'},
      { find: "@ui", replacement: '/src/components/ui'},
    ]
  }
})