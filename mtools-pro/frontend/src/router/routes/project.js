const Layout = () => import('@/layout/index.vue')

export default {
  name: 'Project',
  path: '/project',
  component: Layout,
  meta: {
    title: '项目管理',
    icon: 'ic:baseline-menu'
  },
  children: [
    {
      name: 'ProjectIndex',
      path: 'index',
      component: () => import('@/views/project/index/index.vue'),
      meta: {
        title: '项目列表',
        icon: 'ic:baseline-menu'
      }
    },
    {
      name: 'DemandIndex',
      path: 'demand',
      component: () => import('@/views/project/demand/index.vue'),
      meta: {
        title: '需求列表',
        icon: 'ic:baseline-menu'
      }
    },
    {
      name: 'FunctionIndex',
      path: 'function',
      component: () => import('@/views/project/function/index.vue'),
      meta: {
        title: '功能列表',
        icon: 'ic:baseline-menu'
      }
    }
  ]
}
