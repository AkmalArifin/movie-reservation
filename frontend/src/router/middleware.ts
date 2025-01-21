import axios from 'axios';
import { parseJwt } from "@/utils/jwt";
import type { RouteLocationNormalized, NavigationGuardNext } from 'vue-router';

export const isAuth = async (
    to: RouteLocationNormalized, 
    from: RouteLocationNormalized, 
    next: NavigationGuardNext
): Promise<void> => {
    const NOW = Math.floor(Date.now() / 1000);
    const parsedToken = parseJwt(sessionStorage.getItem("token"));
    
    console.log("auth");
    if (parsedToken.exp > NOW) {
        return next();
    }


    try {
        const response = await axios.post("http://localhost:8080/refresh", {}, { withCredentials: true })

        const data : {
            message: string,
            token: string
        } = response.data;

        sessionStorage.setItem("token", data.token);

        return next();

    } catch (error) {
        if (axios.isAxiosError(error)) {
            console.error(error.response);
        } else {
            console.error(error);
        }

        return next({name: 'login'})
    }
}

export const isAdmin = async (
    to: RouteLocationNormalized,
    from: RouteLocationNormalized,
    next: NavigationGuardNext
): Promise<void> => {
    const NOW = Math.floor(Date.now() / 1000);
    const parsedToken = parseJwt(sessionStorage.getItem('token'));

    console.log("admin");

    if (parsedToken.exp > NOW && parsedToken.role === 'admin') {
        return next();
    }

    try {
        const response = await axios.post("http://localhost:8080/refresh", {}, { withCredentials: true });

        const data : {
            message: string,
            token: string
        } = response.data;

        sessionStorage.setItem("token", data.token);

        const newParsedToken = parseJwt(data.token);
        if (newParsedToken.role === 'admin') {
            return next()
        } else {
            return next({name: 'login'})
        };
    } catch(error) {
        if (axios.isAxiosError(error)) {
            console.log(error.response);
        } else {
            console.log(error);
        }

        return next({name: 'login'})
    }
}