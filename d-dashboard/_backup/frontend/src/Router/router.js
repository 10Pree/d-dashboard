import { createRouter, createWebHistory } from 'vue-router'
import Dashboard from '../components/Dashboard/dashboard.vue'
import Admin from '../components/Admin/admin.vue'
import Login from '../components/Login/login.vue'
import Register from '../components/Register/register.vue'


const routes = [
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: Dashboard
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/admin',
    name: 'Admin',
    component: Admin
  },
  {
    path: '/register',
    name: 'Register',
    component: Register
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
