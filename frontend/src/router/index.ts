import { createRouter, createWebHistory } from 'vue-router';
import { parseJwt } from '@/utils/jwt';
import { isAuth, isAdmin } from './middleware';

import LoginView from '@/views/client/LoginView.vue';
import RegisterView from '@/views/client/RegisterView.vue';
import HomeView from '@/views/client/HomeView.vue';

import DashboardView from '@/views/admin/DashboardView.vue';
import MovieMasterView from '@/views/admin/MovieMasterView.vue';


const routes = [
    { name: 'login', path: '/login', component: LoginView },
    { name: 'register', path: '/register', component: RegisterView },

    { name: 'dashboard', path: '/admin', component: DashboardView, meta: { middleware: [isAuth, isAdmin] } },
    { name: 'movie-master', path: '/admin/movies', component: MovieMasterView, meta: { middleware: [isAuth, isAdmin] } },

    { name: 'home', path: '/', component: HomeView },
]

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes
})


// TODO: handle middleware
// router.beforeEach((to, from, next) => {
//     var middlewares = to.meta.middleware as Function[] || undefined

//     if (middlewares && middlewares.length > 0) {

//     }
// })

export default router