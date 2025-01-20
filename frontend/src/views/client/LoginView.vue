<template>
    <div class="body">
        <h1>Login View</h1>
        <div class="input">
            <input type="text" v-model="userInput.email">
            <input type="password" v-model="userInput.password">
            <button @click="handleLogin">Login</button>
        </div>

    </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue';
import { useRouter } from 'vue-router';
import axios from 'axios';

import { parseJwt } from '../../utils/jwt'

const router = useRouter();

const userInput = reactive({
    email: "",
    password: ""
})

async function handleLogin(event:Event) {
    event.preventDefault();

    const data = {
        email: userInput.email,
        password: userInput.password
    };

    await axios.post("http://localhost:8080/login", data)
        .then(response => {
            const data : {
                message: string,
                token: string,
                refreshToken: string
            } = response.data
            
            sessionStorage.setItem("token", data.token)
        }).catch(error => {
            console.log(error)
        })

    const parsedToken = parseJwt(sessionStorage.getItem('token'));
    if (parsedToken.role === 'admin') {
        router.push({name: "dashboard"});
    } else if (parsedToken.role === 'user') {
        router.push({name: "home"});
    }
}

</script>

<script lang="ts">
export default {
    name: "LoginView"
}
</script>

<style scoped>

</style>