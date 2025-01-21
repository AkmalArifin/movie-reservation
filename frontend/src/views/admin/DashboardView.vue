<template>
    <div>
        <h1>Dashboard View</h1>
        <router-link :to="{name: 'movie-master'}">Movie Table</router-link>
    </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import axios from 'axios';

import { parseJwt } from '../../utils/jwt';

const router = useRouter();

const parsedToken = parseJwt(sessionStorage.getItem('token'));
const now = Math.floor(Date.now() / 1000);

if (parsedToken.exp < now) {
    axios.post("http://localhost:8080/refresh", {}, { withCredentials: true })
        .then(response => {
            const data : {
                message: string,
                token: string
            } = response.data

            sessionStorage.setItem("token", data.token)
        }).catch(error => {
            console.error(error.response)
            alert("please login first")
            router.push({name: "login"})
        })
}

if (parsedToken.role !== 'admin') {
    alert("not authorized")
    router.push({name: "login"})
}

</script>

<script lang="ts">
export default {
    name: "DashboardView"
}
</script>

<style scoped>

</style>