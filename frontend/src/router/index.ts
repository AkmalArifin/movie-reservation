import { createRouter, createWebHistory } from 'vue-router';
import LoginView from '@/views/client/LoginView.vue';
import RegisterView from '@/views/client/RegisterView.vue';

const routes = [
    { path: '/login', component: LoginView },
    { path: '/register', component: RegisterView }
]

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes
})

export default router