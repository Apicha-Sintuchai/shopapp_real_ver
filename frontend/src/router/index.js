import { createRouter, createWebHistory } from 'vue-router'
import MenuItem from '@/views/MenuItem.vue'
import LoginPage from '@/views/LoginPage.vue'
import OrderPage from '@/views/OrderPage.vue'
import MenuPage from '@/views/MenuPage.vue'
import RecievedPage from '@/views/RecievedPage.vue'
import PaymentPage from '@/views/PaymentPage.vue'
import QrcodePage from '@/views/QrcodePage.vue'
import Customer_Menu from '@/views/customer/Customer_Menu.vue'


const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),

  routes: [
    {
      path: '/',
      name: 'Login',
      component: LoginPage,
    },

    {
      path: '/MenuItems',
      name: 'MenuItems',
      component: MenuItem,

    },
    {
      path: '/Orderpage',
      name: 'Orderpage',
      component: OrderPage,

    },
    {
      path: '/Menupage',
      name: 'Menupage',
      component: MenuPage,

    },
    {
      path: '/Recivepage',
      name: 'Recivepage',
      component: RecievedPage,
    },
    {
      path: '/Paymentpage',
      name: 'Paymentpage',
      component: PaymentPage,
    },
    {
      path: '/Qrcodepage',
      name: 'Qrcodepage',
      component: QrcodePage,
    },
    {
      path: '/customer_menu',
      name: 'customer_menu',
      component: Customer_Menu,
    }
  ],
})





router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token');


  if (to.name === 'customer_menu') {
    return next();
  }

  if (!token && to.name !== 'Login') {
    return next({ name: 'Login' });
  }

  next();
});
export default router
