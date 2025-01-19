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

const userInput = reactive({
    email: "",
    password: ""
})

async function handleLogin(event:Event) {
    event.preventDefault()

    console.log(userInput)
}

</script>

<script lang="ts">
export default {
    name: "LoginView"
}
</script>

<style scoped>

</style>