import { jwtDecode } from "jwt-decode";

type UserPayload = {
  exp: number;
  expired_at: string;
  id: string;
  issued_at: string;
  role: string;
  username: string;
  uuid: string;
}

const AUTHORIZATION_RELATED_ROUTES = ["/login", "/registration"]

const NO_AUTH_REQUIRED_ROUTES = [
  ...AUTHORIZATION_RELATED_ROUTES,
  "/"
]

export default defineNuxtRouteMiddleware(to => {
    // skip middleware on server
    if (process.server) return
    const token = localStorage.getItem("token")
    if(!token && NO_AUTH_REQUIRED_ROUTES.includes(to.fullPath)) return // doesn't have token in storage and also navigating to the no auth required route
    if (!token) return navigateTo("/login")
    if(AUTHORIZATION_RELATED_ROUTES.includes(to.fullPath)) return navigateTo("/app") // if the token invalid, it will be directed back here with the token already removed
    try{
      const decoded = jwtDecode(token) as UserPayload;
      if(decoded.exp <= Date.now()/1000){
        // try for extending token live by calling refresh-token API <later>
        return navigateTo("/login")
      }
      const authStore = useAuthStore()
      authStore.userID = decoded.uuid
      authStore.username = decoded.username
      authStore.setAuthenticationStatus(true)
    }catch{
      // failed to decode
      return navigateTo("/login")
    }

})