<template>
    <div>
        <h1>Movie Master</h1>
    </div>
</template>

<script setup lang="ts">
import axios from 'axios';
import { useRouter } from 'vue-router'; 
import { parseJwt } from '@/utils/jwt';

const router = useRouter();

const parsedToken = parseJwt(sessionStorage.getItem("token"));
const NOW = Math.floor(Date.now() / 1000);

if (parsedToken.exp < NOW) {
    try {
        const data = axios.post('http://localhost:8080/refresh', {}, { withCredentials: true })
            .then(response => {
                const data: {
                    message: string,
                    token: string
                } = response.data;

                sessionStorage.setItem("token", data.token);
                const newParsedToken = parseJwt(data.token)
                if (newParsedToken.role !== 'admin') {
                    alert('not authorized');
                    router.push({ name: 'login' });
                }
            })

    } catch (error) {
        if (axios.isAxiosError(error)) {
            console.log(error.message);
        } else {
            console.log(error);
        }

        alert('Please login first');
        router.push({ name: 'login' });
    }
}

if (parsedToken.role !== 'admin') {
    alert('not authorized');
    router.push({ name: 'login' });
}

// TODO: get data for table
async function getData() {
    axios.get('http://localhost:8080/movies')
}

</script>

<script lang="ts">
export default {
    name: "MovieMasterVue"
}
</script>

<style scoped>

</style>