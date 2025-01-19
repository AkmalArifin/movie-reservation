<template>
    <div class="body">
        <div class="input">
            <input type="text" name="name" id="name" v-model="userInput.name">
            <input type="text" name="email" id="email" v-model="userInput.email">
            <input type="password" name="password" id="password" v-model="userInput.password">
            <input type="password" name="confirmPassword" id="confirmPassword" v-model="userInput.confirmPassword">
            <button @click="handleRegister">Sign Up</button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue'
const userInput = reactive({
    name: "",
    email: "",
    password: "",
    confirmPassword: ""
})

async function handleRegister(event:Event) {
    event.preventDefault();

    console.log(userInput);
}
</script>

<script lang="ts">
export default {
    name: "RegisterView"
}
</script>

<style scoped>

</style>