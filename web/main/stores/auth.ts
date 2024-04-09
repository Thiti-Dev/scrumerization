import { defineStore } from 'pinia'

export const useAuthStore = defineStore('authStore', {
    state: () => ({
        userID: '',
        username: '',
        isAuthenticated: false
    }),
    actions: {
        setAuthenticationStatus(bool:boolean){
            if(!bool){
                // attempting to logout
                 localStorage.clear() // for now localstorage only contains 2 values
                 this.isAuthenticated = false
                 navigateTo("/login",{external:true}) // external = true solved the case when token found invalid after refresh and still stuck at the same page ignoring the navigateTo somehow
            }


            this.isAuthenticated = bool
        },

    }
  })