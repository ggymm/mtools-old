const Layout = () => import('@/layout/index.vue')

export default {
  name: 'Magnet',
  path: '/magnet',
  component: Layout,
  meta: {
    title: '磁力链接工具',
    icon: 'ic:baseline-menu'
  },
  children: [
    {
      name: 'Magnet_Search',
      path: 'search',
      component: () => import('@/views/magnet/search/index.vue'),
      meta: {
        title: '磁力搜索',
        icon: 'ic:baseline-menu'
      }
    },
    {
      name: 'Magnet_Config',
      path: 'config',
      component: () => import('@/views/magnet/config/index.vue'),
      meta: {
        title: '搜索配置',
        icon: 'ic:baseline-menu'
      }
    }
  ]
}
