<template>
    <div>
        <section class="bg-gray-50 dark:bg-gray-900">
            <div class="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
                <a href="#" class="flex items-center mb-6 text-2xl font-semibold text-gray-900 dark:text-white">
                    <img class="w-8 h-8 mr-2" src="https://www.svgrepo.com/show/530552/mouthwash-cup.svg" alt="logo">
                    Scrumerization    
                </a>
                <div class="w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700">
                    <div class="p-6 space-y-4 md:space-y-6 sm:p-8">
                    <form @submit.prevent="submitForm" class="space-y-4 md:space-y-6" action="#">
                        <div>
                            <label for="username" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Username</label>
                            <input v-model="formData.username" :class="{'dark:border-rose-600 border-rose-600': false}" type="text" name="username" id="username" class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="john.doe" required>
                            <p class="italic text-red-500 mt-3" v-if="validationErrors.username">{{ validationErrors.username }}</p>
                        </div>
                        <div>
                            <label for="password" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Password</label>
                            <input v-model="formData.password" type="password" name="password" id="password" placeholder="••••••••" class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" required>
                            <p class="italic text-red-500 mt-3" v-if="validationErrors.password">{{ validationErrors.password }}</p>
                        </div>
                        <button type="submit" class="w-full text-white bg-primary-600 hover:bg-primary-700 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800">Login</button>
                        <p class="italic text-red-500 mt-3" v-if="generalError">{{ generalError }}</p>
                        <p class="text-sm font-light text-gray-500 dark:text-gray-400">
                            Doesn't have accoint? <a href="#" class="font-medium text-primary-600 hover:underline dark:text-primary-500">Register here</a>
                        </p>
                    </form>
                </div>
            </div>
        </div>
        </section>
    </div>
</template>

<script setup lang="ts">
    definePageMeta({
        layout: false,
    })

    import { ref, watch, watchEffect } from 'vue';
    import { useMutation } from '@vue/apollo-composable'
    import {graphql} from "../../.gen/gql"

    const { mutate: loginUser } = useMutation(
        graphql(`
            mutation loginUser($input: LoginUserInput!){
                loginUser(input: $input){
                    refreshToken
                    token
                }
            }
        `)
    )

    const formData = ref({
      username: '',
      password: ''
    });
    const validationErrors = ref({
        username: '',
        password: ''
    })
    const generalError = ref('')
    
    async function submitForm(){
        generalError.value = '' // reset

        try{
            const res = await loginUser({input:{username: formData.value.username,password:formData.value.password}})
            const {token,refreshToken} = res!.data!.loginUser
            localStorage.setItem('token',token)
            localStorage.setItem('refreshToken',refreshToken)

            navigateTo("/app")
        }
        catch(e:any){
            if(e.message === 'validation Error'){
                validationErrorExtract(e, validationErrors)
            }
            else {
                generalError.value = e.message
            }
        }
    }
</script>