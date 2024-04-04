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
                 navigateTo("/login")
            }


            this.isAuthenticated = bool
        },

    }
  })