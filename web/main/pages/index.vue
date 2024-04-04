<template>
    <div>
        <!-- credit: https://github.com/codebushi/gatsby-starter-lander -->

        <div>

        <main class="text-gray-900 font-workbench">


            <section class="pt-20 md:pt-40">
                <div class="container mx-auto px-8 lg:flex">
                    <div class="text-center lg:text-left lg:w-1/2">
                        <h1 class="text-4xl lg:text-5xl xl:text-6xl font-bold leading-none">Scrumerization</h1>
                        <p class="text-xl lg:text-2xl mt-6 font-light">
                            Elevate teamwork with our real-time estimation tool. Collaborate effortlessly, make informed decisions, and boost productivity with ease.
                        </p>
                        <p class="mt-8 md:mt-12">
                            <NuxtLink to="/registration">
                                <button type="button"
                                class=" py-4 px-12 bg-teal-500 hover:bg-teal-600 rounded text-white">Get Started</button>
                            </NuxtLink>
                        </p>
                        <p class="mt-4 text-gray-600">Already have account? <a class="text-blue-500 text-lg" href="/login">Login</a></p>
                    </div>

                    <div class="lg:w-1/2">
                        <img src="~/assets/images/landing_promo.png" alt="promo" height="20" />
                    </div>
                </div>
            </section>


            <section id="features" class="py-20 lg:pb-40 lg:pt-48">
                <div class="container mx-auto text-center">
                    <h2 class="text-3xl lg:text-5xl font-semibold">Main Features</h2>
                    <div class="flex flex-col sm:flex-row sm:-mx-3 mt-12">
                        <div class="flex-1 px-3">
                            <div class="p-12 rounded-lg border border-solid border-gray-200 mb-8"
                                style="box-shadow:0 10px 28px rgba(0,0,0,.08)">
                                <p class="font-semibold text-xl">Creates room

                                </p>
                                <p class="mt-4">Experience seamless collaboration with our room system—create and join rooms effortlessly to engage with your team like never before..
                                </p>
                            </div>
                        </div>
                        <div class="flex-1 px-3">
                            <div class="p-12 rounded-lg border border-solid border-gray-200 mb-8"
                                style="box-shadow:0 10px 28px rgba(0,0,0,.08)">
                                <p class="font-semibold text-xl">Creates Topic/Story
                                </p>
                                <p class="mt-4">ADive deeper into discussions with our topics system—explore diverse topics within each room for focused and productive conversations.
                                </p>
                            </div>
                        </div>
                        <div class="flex-1 px-3">
                            <div class="p-12 rounded-lg border border-solid border-gray-200 mb-8"
                                style="box-shadow:0 10px 28px rgba(0,0,0,.08)">
                                <p class="font-semibold text-xl">Votes

                                </p>
                                <p class="mt-4">Empower your team with real-time decision-making using our voting system—make quick and informed choices with instant feedback from your colleagues.
                                </p>
                            </div>
                        </div>
                    </div>
                </div>
            </section>
        </main>
        </div>
    </div>
</template>

<script setup lang="ts">
   definePageMeta({
        middleware: ["auth"]
    })


    const authStore = useAuthStore()
</script>