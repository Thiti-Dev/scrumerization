<template>
    <div class="font-workbench">
        <div :class="{'opacity-50': isCreationPanelOpened}" class="px-24 h-screen_ flex flex-col justify-center">
            <div class="mt-2 flex flex-row justify-end">
                <button @click="isCreationPanelOpened = true" type="button" class="text-white bg-gradient-to-r from-purple-500 via-purple-600 to-purple-700 hover:bg-gradient-to-br focus:ring-4 focus:outline-none focus:ring-purple-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center me-2 mb-2">Create room</button>
            </div>
            <div class="relative overflow-x-auto shadow-md">
                <table class="w-full text-sm text-left rtl:text-right text-gray-500">
                    <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700">
                        <tr>
                            <th scope="col" class="px-6 py-3">
                                Room
                            </th>
                            <!-- <th scope="col" class="px-6 py-3">
                                Password
                            </th> -->
                            <th scope="col" class="px-6 py-3">
                                Created Date
                            </th>
                            <th scope="col" class="px-6 py-3">
                                Action
                            </th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="room of result?.rooms.data" class="bg-white border-b hover:bg-gray-300 cursor-pointer">
                            <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap">
                                {{ room.room_name }}
                            </th>
                            <!-- <td class="px-6 py-4">
                                {{ room.password ?? 'No-password' }}
                            </td> -->
                            <td class="px-6 py-4">
                                {{ new Date(room.createdAt).toISOString().substring(0, 10) }}
                            </td>
                            <td>
                                <a @click="enterRoom(room.id)" class="select-none cursor-pointer font-medium text-blue-600 hover:underline">Enter</a>
                                <a v-if="justCopiedRoomID !== room.id" @click="copyLink(room.id)" class="select-none cursor-pointer pl-5 font-medium text-blue-600 hover:underline">Copy link</a>
                                <a v-else class="select-none pl-5 font-medium text-blue-600 hover:underline">Copied</a>
                            </td>
                        </tr>
                        <tr v-for="_ of Array.from(Array(10-(result?.rooms.count || 0)))" class="border-b bg-white h-12">
                            <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap">
                            </th>
                            <td class="px-6 py-4">
                            </td>
                            <td class="px-6 py-4">
                            </td>
                            <td>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
            <div class="flex flex-col items-center bg-black rounded-lg rounded-t-none select-none">
                <!-- Help text -->
                <span class="text-sm text-white">
                    Showing <span class="font-semibold text-blue-500">{{ (currentPage * pageLimit)+1  }}</span> to <span class="font-semibold text-blue-500">{{ (currentPage * pageLimit) + result?.rooms.count! }}</span> of <span class="font-semibold text-blue-500">{{ result?.rooms.totalCount }}</span> Entries
                </span>
                <!-- Buttons -->
                <div class="inline-flex mt-2 xs:mt-0">
                    <button @click="pageChange('decrement')" class="flex items-center justify-center px-4 h-10 text-base font-medium text-white bg-gray-800 rounded-s hover:bg-gray-900">
                        Prev
                    </button>
                    <button @click="pageChange('increment')" class="flex items-center justify-center px-4 h-10 text-base font-medium text-white bg-gray-800 border-0 border-s border-gray-700 rounded-e hover:bg-gray-900">
                        Next
                    </button>
                </div>
            </div>
        </div>
        <div id="static-modal" :class="{hidden: !isCreationPanelOpened}" data-modal-backdrop="static" tabindex="-1" aria-hidden="true" class="flex overflow-y-auto overflow-x-hidden fixed top-0 right-0 left-0 z-50 justify-center items-center w-screen h-screen md:inset-0">
            <div class="relative p-4 w-full max-w-2xl max-h-full">
                <!-- Modal content -->
                <div class="relative bg-white rounded-lg shadow">
                    <!-- Modal header -->
                    <div class="flex items-center justify-between p-4 md:p-5 border-b rounded-t">
                        <h3 class="text-xl font-semibold text-gray-900">
                            Room creation
                        </h3>
                        <button @click="isCreationPanelOpened = false" type="button" class="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 ms-auto inline-flex justify-center items-center" data-modal-hide="static-modal">
                            <svg class="w-3 h-3" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 14">
                                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"/>
                            </svg>
                            <span class="sr-only">Close modal</span>
                        </button>
                    </div>
                    <!-- Modal body -->
                    <form class="max-w-sm mx-auto" @submit.prevent="onCreateRoomSubmit">
                        <div class="p-4 md:p-5 space-y-4">
                                <div class="mb-5">
                                    <label for="room-name" class="block mb-2 text-sm font-medium text-gray-900"><span class="text-red-500">*</span>Room name</label>
                                    <input v-model="roomFormData.name" type="text" id="room-name" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5" placeholder="eg. SprintPlanning 1" required />
                                </div>
                                <!-- <div class="mb-5">
                                    <label for="room-password" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Room password</label>
                                    <input v-model="roomFormData.password" type="password" id="room-password" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="*********"/>
                                </div> -->
                        </div>
                        <!-- Modal footer -->
                        <div class="flex items-center p-4 md:p-5 border-t border-gray-200 rounded-b dark:border-gray-600">
                            <button type="submit" data-modal-hide="static-modal" class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center">Create room</button>
                            <button @click="isCreationPanelOpened = false" data-modal-hide="static-modal" type="button" class="py-2.5 px-5 ms-3 text-sm font-medium text-gray-900 focus:outline-none bg-white rounded-lg border border-gray-200 hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-4 focus:ring-gray-100">Cancel</button>
                        </div>
                    </form>

                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
    definePageMeta({
        middleware: ["auth"]
    })

    import { useQuery } from '@vue/apollo-composable';
    import {graphql} from "../../.gen/gql"
    import { computed, watchEffect } from 'vue';
    import { useAuthStore } from '~/stores/auth';

    const justCopiedRoomID = ref<string | null>(null)
    let copiedResetTimeout: NodeJS.Timeout | null = null


    function copyLink(link:string){
        justCopiedRoomID.value = link
        navigator.clipboard.writeText(window.location.origin+"/app/room/"+ link);

        copiedResetTimeout && clearTimeout(copiedResetTimeout)
        copiedResetTimeout = setTimeout(() => {
            justCopiedRoomID.value = null
        }, 2000);
        
    }

    function enterRoom(roomID: string){
        navigateTo("/app/room/"+ roomID)
    }

    function pageChange(type: 'increment'|'decrement'){
        if(!result.value) return
        if(type === 'increment'){
            // const maxPage = result.value.rooms.totalCount / pageLimit
            // if(currentPage.value > maxPage - 2) return // -2 because currentPage start at 0
            if ((currentPage.value * pageLimit) + result.value.rooms.count >= result.value.rooms.totalCount) return
            currentPage.value++;
        }else{
            if(currentPage.value === 0) return // basically === 0
            currentPage.value--;
        }
    }

    async function onCreateRoomSubmit(){
        const {name,password} = roomFormData.value
        try{
            await createRoom({input:{name,password: password || undefined }})
            // creation succeed
            currentPage.value = 0; // triggering side-effect of watch # it got cached so have to call restart right after
            restart()
            isCreationPanelOpened.value = false // closing modal
        }
        catch(e:any){
            console.log(JSON.stringify(e))
        }
    }

    const pageLimit = 10
    const currentPage = ref(0)
    const isCreationPanelOpened = ref(false)
    const roomFormData = ref({
        name: '',
        password: ''
    })

    const authStore = useAuthStore()
    
    const {result,error,refetch,restart,onError} = useQuery(
        graphql(`
            query listCreatedRooms($where: RoomWhereClause!,$paginate: PaginationInput){
                rooms(where: $where,paginate:$paginate){
                    data {
                        id
                    room_name
                    password
                    is_active
                    createdAt
                    updatedAt
                        }
                    totalCount
                    count
                }
            }
        `),{where:{creator_id:authStore.userID},paginate:{limit:pageLimit,offset: currentPage.value * pageLimit}}
    ,{fetchPolicy:'no-cache',context:{headers: {
        "Authorization": localStorage.getItem("token")
    }}})

    const {mutate: createRoom} = useMutation(
        graphql(`
            mutation createRoom($input: RoomCreationInput!){
                createRoom(input: $input){
                id
                room_name
                password
                is_active
            }
            }        
        `),{context:{
            headers:{
                "Authorization": localStorage.getItem("token")
            }
        }}
    )

    watch(currentPage, (newPage,_) => {
        refetch({where:{creator_id:authStore.userID},paginate:{limit:pageLimit,offset: newPage * pageLimit}})
    })

    // watchEffect(() => {
    //     if(error.value?.graphQLErrors[0].extensions.code === 'Authentication'){
    //         navigateTo("/login")
    //     }
    // })

    onError((error) => {
        for(const err of error.graphQLErrors){
            if(err.extensions?.code === 'Authentication'){
                authStore.setAuthenticationStatus(false)
                break
            }
        }
    })
</script>