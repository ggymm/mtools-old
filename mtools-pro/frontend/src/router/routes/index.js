import home from '@/router/routes/home.js'

import project from '@/router/routes/project.js'
import wiki from '@/router/routes/wiki.js'

import tools from '@/router/routes/tools.js'

import video from '@/router/routes/video.js'
import magnet from '@/router/routes/magnet.js'

export const basicRoutes = [
  {
    name: '404',
    path: '/404',
    component: () => import('@/views/error-page/404.vue'),
    isHidden: true
  }
]

const asyncRoutes = []
asyncRoutes.push(home)

asyncRoutes.push(tools)
asyncRoutes.push(project)
asyncRoutes.push(wiki)

asyncRoutes.push(video)
asyncRoutes.push(magnet)

export { asyncRoutes }
