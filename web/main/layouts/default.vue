<template>
    <div class="font-workbench  sticky inset-0 z-50 border-b border-slate-100 bg-white/80 backdrop-blur-lg h-15 ">
        <nav class="mx-auto flex max-w-6xl gap-8 px-6 transition-all duration-200 ease-in-out lg:px-12 py-4">
            <div class="relative flex items-center gap-3">
                <a href="/">
                    <img src="https://www.svgrepo.com/show/499831/target.svg" loading="lazy" style="color:transparent" width="32" height="32">
                </a>
                <p class="text-2xl">Scrumerization</p>
            </div>
            <ul class="hidden items-center justify-center gap-6 md:flex">
                <NuxtLink to="/app">
                    <li class="pt-1.5 font-dm text-sm font-medium text-slate-700">
                        <p>My Rooms</p>
                    </li>
                </NuxtLink>
                <!-- <li class="pt-1.5 font-dm text-sm font-medium text-slate-700">
                    <a href="#">Blog</a>
                </li>
                <li class="pt-1.5 font-dm text-sm font-medium text-slate-700">
                    <a href="#">Docs</a>
                </li> -->
            </ul>
            <div class="flex-grow"></div>
                <div className="flex flex-wrap justify-center">
                    <p className="relative cursor-pointer" @click="authStore.setAuthenticationStatus(false)">
                        <span className="absolute top-0 left-0 mt-1 ml-1 h-full w-full rounded bg-black"></span>
                        <span className="active:bg-gray-700 select-none fold-bold relative inline-block h-full w-full rounded border-2 border-black bg-white px-3 py-1 text-base font-bold text-black transition duration-100 hover:bg-red-400 hover:text-gray-900 min-w-[150px] text-center">Logout</span>
                    </p>
                </div>
            <!-- <div class="relative flex items-center justify-center md:hidden">
                <button type="button">
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" aria-hidden="true" class="h-6 w-auto text-slate-900"><path stroke-linecap="round" stroke-linejoin="round" d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5"></path></svg>
                </button>
            </div> -->
        </nav>
    </div>
    <slot />
</template>

<script setup lang="ts">
    const authStore = useAuthStore()
</script>