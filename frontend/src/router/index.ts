import { createRouter, createWebHistory } from 'vue-router';
import LoginView from '@/views/client/LoginView.vue';
import RegisterView from '@/views/client/RegisterView.vue';
import HomeView from '@/views/client/HomeView.vue';

import DashboardView from '@/views/admin/DashboardView.vue';


const routes = [
    { name: 'login', path: '/login', component: LoginView },
    { name: 'register', path: '/register', component: RegisterView },
    { name: 'dashboard', path: '/admin', component: DashboardView},
    { name: 'home', path: '/', component: HomeView }
]

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes
})

export default router