import { createRouter, createWebHistory } from 'vue-router'
// import App from '@/App.vue'
import SomeThing from '@/views/SomeThing.vue'
import MenuItem from '@/views/MenuItem.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'MenuItems',
      component: MenuItem,
    },
    {
      path: '/about',
      name: 'about',
      component: SomeThing,

    },
    // {
    //   path: '/Menuitems',
    //   name: 'MenuItems',
    //   component: MenuItem,

    // },
  ],
})

export default router
