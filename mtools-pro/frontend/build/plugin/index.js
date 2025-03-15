import vue from '@vitejs/plugin-vue'

import Unocss from 'unocss/vite'

import monaco from './monaco'
import unplugin from './unplugin'

export function createVitePlugins() {
  return [vue(), Unocss(), monaco(), ...unplugin]
}
