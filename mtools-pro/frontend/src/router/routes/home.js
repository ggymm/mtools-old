const Layout = () => import('@/layout/index.vue')

export default {
  name: 'Root',
  path: '/',
  component: Layout,
  redirect: '/home',
  children: [
    {
      name: 'Home',
      path: 'home',
      component: () => import('@/views/home/index.vue'),
      meta: {
        title: '工作台',
        icon: 'ic:baseline-menu',
        order: 0
      }
    }
  ]
}
