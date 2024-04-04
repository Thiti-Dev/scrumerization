<template>
    <div class="font-workbench">
        <section v-if="topicData" class="py-1 bg-blueGray-50">
            <div class="w-full xl:w-8/12 mb-12 xl:mb-0 px-4 mx-auto mt-24">
            <div class="relative flex flex-col min-w-0 break-words bg-white w-full mb-6 shadow-lg rounded ">
                <div class="rounded-t mb-0 px-4 py-3 border-0">
                    <div class="flex flex-wrap items-center">
                        <div class="relative w-full px-4 max-w-full flex-grow flex-1">
                        <h3 class="font-semibold text-base text-blueGray-700">{{ topicData.name }}</h3>
                        </div>
                        <button @click="onConfirmHandler" class="bg-indigo-500 text-white active:bg-indigo-600 text-xs font-bold uppercase px-3 py-1 rounded outline-none focus:outline-none mr-1 mb-1 ease-linear transition-all duration-150" type="button">Confirm {{ (peers.filter((p) => p.isVoted).length + (isConfirmed ? 1 : 0)) }} / {{ peers.length+1 }}</button>
                    </div>
                </div>

                <div class="block w-full overflow-x-auto overflow-y-auto max-h-[450px]">
                <table class="items-center bg-transparent w-full border-collapse">
                    <thead>
                    <tr>
                        <th class="px-6 bg-blueGray-50 text-blueGray-500 align-middle border border-solid border-blueGray-100 py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-left">
                                Voter
                        </th>
                        <th class="px-6 bg-blueGray-50 text-blueGray-500 align-middle border border-solid border-blueGray-100 py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-left">
                                Description
                        </th>
                        <th class="text-center px-6 bg-blueGray-50 text-blueGray-500 align-middle border border-solid border-blueGray-100 py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-left">
                                Action
                        </th>
                    </tr>
                    </thead>
                    <tbody>
                        <tr class="table-auto">
                            <th class="border-t-0 px-6 align-middle border-l-0 border-r-0 text-xs whitespace-nowrap p-4 text-left text-blueGray-700 ">
                                Me
                            </th>
                            <td class="border-t-0 px-6 align-center border-l-0 border-r-0 text-xs whitespace-nowrap p-4">
                                <textarea :disabled="isConfirmed" id="message" v-model="givenDesc" rows="3" class="block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500" placeholder="Write your thoughts here..."></textarea>
                            </td>
                            <td class="border-t-0 px-6 align-middle border-l-0 border-r-0 text-xs whitespace-nowrap p-4">
                                <div class="w-full">
                                    <PointingCard v-if="!isConfirmed" @point-select="onPointSelect"/>
                                    <PointingCard v-else view-only :view-only-point="givenPoint!"/>
                                </div>  
                                <!-- <a class="font-medium text-blue-600 dark:text-blue-500 hover:underline">Enter</a> -->
                            </td>
                        </tr>       
                        <tr class="table-auto" v-for="peer of peers">
                            <th class="border-t-0 px-6 align-middle border-l-0 border-r-0 text-xs whitespace-nowrap p-4 text-left text-blueGray-700 ">
                                {{ peer.name }}
                            </th>
                            <td class="border-t-0 px-6 align-center border-l-0 border-r-0 text-xs whitespace-nowrap p-4">
                               <div v-if="!peer.isVoted" class="flex items-center">
                                    <div role="status">
                                        <svg aria-hidden="true" class="w-4 h-4 me-2 text-gray-200 animate-spin dark:text-gray-600 fill-blue-600" viewBox="0 0 100 101" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z" fill="currentColor"/><path d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z" fill="currentFill"/></svg>
                                        <span class="sr-only">Loading...</span>
                                    </div>
                                    Descripting . . .
                               </div>
                               <div v-else>
                                ✅
                               </div>
                            </td>
                            <td class="border-t-0 px-6 align-middle border-l-0 border-r-0 text-xs whitespace-nowrap p-4">
                                <div class="w-full flex justify-center">
                                    <button v-if="!peer.isVoted" disabled type="button" class="py-2.5 px-5 me-2 text-sm font-medium text-gray-900 bg-white rounded-lg border border-gray-200 hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-2 focus:ring-blue-700 focus:text-blue-700 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700 inline-flex items-center">
                                        <svg aria-hidden="true" role="status" class="inline w-4 h-4 me-3 text-gray-200 animate-spin dark:text-gray-600" viewBox="0 0 100 101" fill="none" xmlns="http://www.w3.org/2000/svg">
                                        <path d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z" fill="currentColor"/>
                                        <path d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z" fill="#1C64F2"/>
                                        </svg>
                                        Voting
                                    </button>
                                    <p v-else>✅</p>
                                </div>  
                                <!-- <a class="font-medium text-blue-600 dark:text-blue-500 hover:underline">Enter</a> -->
                            </td>
                        </tr>             
                    </tbody>

                </table>
                </div>
            </div>
            </div>
            <footer class="relative pt-8 pb-6 mt-16">
            <div class="container mx-auto px-4">
                <div class="flex flex-wrap items-center md:justify-between justify-center">
                <div class="w-full md:w-6/12 px-4 mx-auto text-center">
                    <div class="text-sm text-blueGray-500 font-semibold py-1">
                    Made with <a href="#" class="text-blueGray-500 hover:text-gray-800" target="_blank">Love</a> by <a href="#" class="text-blueGray-500 hover:text-blueGray-800" target="_blank"> Thiti-Dev</a>.
                    </div>
                </div>
                </div>
            </div>
            </footer>
            </section>
    </div>
</template>


<script setup lang="ts">
    definePageMeta({
        middleware: ["auth"]
    })

    const shouldConnectRoom = ref(false)
    const givenPoint = ref<number | null>(null)
    const givenDesc = ref('')
    const isConfirmed = ref(false)
    const route = useRoute()
    const topicID = route.params.topic_id
    const authStore = useAuthStore()
    const subscriptionVariables = reactive({
        roomID:route.params.room_id,
        token: localStorage.getItem("token")!
    })
    const {result:topicResult} = useQuery(FIND_TOPIC_QUERY,{topicID:topicID},createApolloContextAuthorization)
    const {result: subscriptionResult} = useSubscription(CONNECT_TO_ROOM_SUBSCRIPTION,subscriptionVariables,{enabled: shouldConnectRoom})
    const {mutate:emitVote} = useMutation(EMIT_VOTE_MUTATION,createApolloContextAuthorization())

    function onPointSelect(point: number){
        // emitVote({topicID,input:{point,description: undefined}})
        givenPoint.value = point
    }

    function onConfirmHandler(){
        if(givenPoint.value === null) return
        isConfirmed.value = true
        emitVote({topicID,input:{point:givenPoint.value,description: givenDesc.value || undefined}})
    }

    watch(() => subscriptionResult.value?.connectToRoom, (newValue,oldValue) => {
        if(!newValue) return
        if(newValue.onGoingTopic === null){
            // the room just has been terminated
            navigateTo(`/app/room/${route.params.room_id}/topic/${topicID}/result`,{replace:true})
        }
    },{deep:true})

    watch(() => topicResult.value, (newValue,oldValue) => {
        if(!newValue?.topic) return navigateTo('/app')
        if(!newValue.topic.isActive) return navigateTo(`/app/room/${route.params.room_id}/topic/${topicID}/result`,{replace:true})
        shouldConnectRoom.value = true
    })

    const topicData = computed(() => {
        if(topicResult.value?.topic === null) return null
        return topicResult.value?.topic
    })

    const peers = computed(() => {
        if(!subscriptionResult.value?.connectToRoom.clients) return []
        return subscriptionResult.value?.connectToRoom.clients.filter((c) => c.uuid != authStore.userID)
    })
</script>

<script lang="ts">
    import { useQuery } from '@vue/apollo-composable';
    import {graphql} from '../../../../../../.gen/gql'
    import { useRoute } from 'vue-router';

    const FIND_TOPIC_QUERY = graphql(`
    query listSigleTopic($topicID: UUID!){
        topic(topicID: $topicID){
            id
            roomID
            name
            createdAt
            updatedAt
            isActive
        }
    }    
    `)

    const CONNECT_TO_ROOM_SUBSCRIPTION = graphql(`
        subscription connectToRoom($token: String!,$roomID:UUID!){
            connectToRoom(roomID: $roomID,token: $token){
                clients{
                uuid
                name
                isVoted
                }
                active
                onGoingTopic{
                id
                name
                }
            }
        }    
    `)

    const EMIT_VOTE_MUTATION = graphql(`
        mutation emitVoting($topicID: UUID!,$input: VoteInput!){
            vote(topicID: $topicID, input: $input)
        }    
    `)
</script>