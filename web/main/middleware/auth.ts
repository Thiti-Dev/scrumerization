import { jwtDecode } from "jwt-decode";
export default defineNuxtRouteMiddleware(to => {
    // skip middleware on server
    if (process.server) return
    const token = localStorage.getItem("token")
    if (!token) return navigateTo("/login")
    type UserPayload = {
        exp: number;
        expired_at: string;
        id: string;
        issued_at: string;
        role: string;
        username: string;
        uuid: string;
      }
    try{
      const decoded = jwtDecode(token) as UserPayload;
      console.log(decoded.exp)

      if(decoded.exp <= Date.now()/1000){
        // try for extending token live by calling refresh-token API <later>
        return navigateTo("/login")
      }
      const authStore = useAuthStore()
      authStore.userID = decoded.uuid
      authStore.username = decoded.username
    }catch{
      // failed to decode
      return navigateTo("/login")
    }

})