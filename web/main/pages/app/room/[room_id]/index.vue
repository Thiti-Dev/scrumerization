<template>
    <div>
        <section class="py-1 bg-gray-100 h-screen">
            <div class="w-full xl:w-8/12 mb-12 xl:mb-0 px-4 mx-auto mt-24">
            <div class="relative flex flex-col min-w-0 break-words bg-white w-full mb-6 shadow-lg rounded ">
                <div class="rounded-t mb-0 px-4 py-3 border-0">
                    <div class="flex flex-wrap items-center">
                        <div class="relative w-full px-4 max-w-full flex-grow flex-1">
                            <nav class="flex" aria-label="Breadcrumb">
                                <ol class="inline-flex items-center space-x-1 md:space-x-2 rtl:space-x-reverse">
                                    <li class="inline-flex items-center">
                                    <a href="/app" class="inline-flex items-center text-sm font-medium text-gray-700 hover:text-blue-600 dark:text-gray-400 dark:hover:text-white">
                                        <svg class="w-3 h-3 me-2.5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 20 20">
                                        <path d="m19.707 9.293-2-2-7-7a1 1 0 0 0-1.414 0l-7 7-2 2a1 1 0 0 0 1.414 1.414L2 10.414V18a2 2 0 0 0 2 2h3a1 1 0 0 0 1-1v-4a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v4a1 1 0 0 0 1 1h3a2 2 0 0 0 2-2v-7.586l.293.293a1 1 0 0 0 1.414-1.414Z"/>
                                        </svg>
                                        App
                                    </a>
                                    </li>
                                    <li aria-current="page">
                                    <div class="flex items-center">
                                        <svg class="rtl:rotate-180 w-3 h-3 text-gray-400 mx-1" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 6 10">
                                        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 9 4-4-4-4"/>
                                        </svg>
                                        <span class="ms-1 text-sm font-medium text-gray-500 md:ms-2 dark:text-gray-400 select-none">Rooms</span>
                                    </div>
                                    </li>
                                </ol>
                            </nav>
                        </div>
                        <div :class="{flex: isCreateTopicPanelOpened}" class="w-full px-4 max-w-full flex-grow flex-1 text-right">
                            <div v-if="isCreateTopicPanelOpened" class="relative flex w-full max-w-[24rem]">
                                <div class="relative h-10 w-full min-w-[200px]">
                                    <input v-model="topicCreationName" type="text"
                                    class="peer h-full w-full rounded-[7px] border border-blue-gray-200 border-t-transparent bg-transparent px-3 py-2.5 pr-20 font-sans text-sm font-normal text-blue-gray-700 outline outline-0 transition-all placeholder-shown:border placeholder-shown:border-blue-gray-200 placeholder-shown:border-t-blue-gray-200 focus:border-2 focus:border-gray-900 focus:border-t-transparent focus:outline-0 disabled:border-0 disabled:bg-blue-gray-50"
                                    placeholder=" " />
                                    <label
                                    class="before:content[' '] after:content[' '] pointer-events-none absolute left-0 -top-1.5 flex h-full w-full select-none !overflow-visible truncate text-[11px] font-normal leading-tight text-gray-500 transition-all before:pointer-events-none before:mt-[6.5px] before:mr-1 before:box-border before:block before:h-1.5 before:w-2.5 before:rounded-tl-md before:border-t before:border-l before:border-blue-gray-200 before:transition-all after:pointer-events-none after:mt-[6.5px] after:ml-1 after:box-border after:block after:h-1.5 after:w-2.5 after:flex-grow after:rounded-tr-md after:border-t after:border-r after:border-blue-gray-200 after:transition-all peer-placeholder-shown:text-sm peer-placeholder-shown:leading-[3.75] peer-placeholder-shown:text-blue-gray-500 peer-placeholder-shown:before:border-transparent peer-placeholder-shown:after:border-transparent peer-focus:text-[11px] peer-focus:leading-tight peer-focus:text-gray-900 peer-focus:before:border-t-2 peer-focus:before:border-l-2 peer-focus:before:!border-gray-900 peer-focus:after:border-t-2 peer-focus:after:border-r-2 peer-focus:after:!border-gray-900 peer-disabled:text-transparent peer-disabled:before:border-transparent peer-disabled:after:border-transparent peer-disabled:peer-placeholder-shown:text-blue-gray-500">
                                    Topic name
                                    </label>
                                </div>
                                <button
                                    @click="onTopicCreate"
                                    class="!absolute right-1 top-1 rounded bg-blue-gray-500 py-2 px-4 text-center align-middle font-sans text-xs font-bold uppercase text-red-500 shadow-md shadow-blue-gray-500/20 transition-all hover:shadow-lg hover:shadow-blue-gray-500/40 focus:opacity-[0.85] focus:shadow-none active:opacity-[0.85] active:shadow-none disabled:pointer-events-none disabled:opacity-50 disabled:shadow-none"
                                    type="button">
                                    Create
                                </button>
                            </div>  
                            <button v-if="!isCreateTopicPanelOpened"  @click="isCreateTopicPanelOpened=true" :class="{hidden: authStore.userID !== result?.rooms.data[0].creator_id}" class="bg-indigo-500 text-white active:bg-indigo-600 text-xs font-bold uppercase px-3 py-1 rounded outline-none focus:outline-none mr-1 mb-1 ease-linear transition-all duration-150" type="button">Create Topic</button>
                        </div>
                    </div>
                    <p class="float-right text-red-500 pr-16">{{ errorMessage }}</p>
                </div>

                <div class="block w-full overflow-x-auto overflow-y-auto max-h-[535px]">
                <table class="items-center bg-transparent w-full border-collapse">
                    <thead>
                    <tr>
                        <th class="px-6 bg-blueGray-50 text-blueGray-500 align-middle border border-solid border-blueGray-100 py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-left">
                                Name
                        </th>
                        <th class="px-6 bg-blueGray-50 text-blueGray-500 align-middle border border-solid border-blueGray-100 py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-left">
                                Average Score
                        </th>
                        <th class="px-6 bg-blueGray-50 text-blueGray-500 align-middle border border-solid border-blueGray-100 py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-left">
                                Is on-going
                        </th>
                        <th class="px-6 bg-blueGray-50 text-blueGray-500 align-middle border border-solid border-blueGray-100 py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-left">
                                Created Date
                        </th>
                    </tr>
                    </thead>
                    <tbody>
                        <tr v-for="topic of result?.topics" class="table-auto border-b">
                            <th @click="navigateTo(`/app/room/${route.params.room_id}/topic/${topic.id}`,{external:true})" class="cursor-pointer border-t-0 px-6 align-middle border-l-0 border-r-0 text-xs whitespace-nowrap p-4 text-left text-blueGray-700 ">
                                {{ topic.name }}
                            </th>
                            <td class="border-t-0 px-6 align-middle border-l-0 border-r-0 text-xs whitespace-nowrap p-4 ">
                                {{ topic.avgScore }}
                            </td>
                            <td class="border-t-0 px-6 align-center border-l-0 border-r-0 text-xs whitespace-nowrap p-4">
                                {{ topic.isActive ? "Yes" : "No" }}
                            </td>
                            <td class="border-t-0 px-6 align-middle border-l-0 border-r-0 text-xs whitespace-nowrap p-4">
                            <i class="fas fa-arrow-up text-emerald-500 mr-4"></i>
                                {{ new Date(topic.createdAt).toISOString().substring(0, 10) }}
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

    const route = useRoute()

    const authStore = useAuthStore()

    const isCreateTopicPanelOpened = ref(false)
    const topicCreationName = ref('')
    const errorMessage = ref('')

    async function onTopicCreate(){
        if(!topicCreationName.value) return
        try{
            const res = await createTopic({input:{name:topicCreationName.value,roomID: route.params.room_id}})
            restart() // triggering refresh
            topicCreationName.value = ''
            isCreateTopicPanelOpened.value = false // close the panel
        }
        catch(e:any){
            // error occur
            if(e.hasOwnProperty('message')){
                errorMessage.value = e.message
            }
        }
    }

    const {result,restart} = useQuery(ROOM_INFO_ALONG_WITH_TOPICS_QUERY,{
    "roomWhere": {
        "id": route.params.room_id
    },
    "roomID": route.params.room_id
    },{...createApolloContextAuthorization(),fetchPolicy:'no-cache'})

    const {mutate: createTopic} = useMutation(CREATE_TOPIC_MUTATION,createApolloContextAuthorization)
</script>

<!-- GRAPHQL -->
<script lang="ts">
    import {graphql} from "../../../../.gen/gql"

    const CREATE_TOPIC_MUTATION = graphql(`
        mutation createTopic($input: CreateTopicInput!){
            createTopic(input: $input){
                id
                roomID
                name
                createdAt
                updatedAt
            }
        }
    `)

    const ROOM_INFO_ALONG_WITH_TOPICS_QUERY = graphql(`
        query roomInfoAlongWithTopics($roomWhere: RoomWhereClause!,$roomID: UUID!){
            rooms(where: $roomWhere){
                data{
                id
                room_name
                creator_id
                }
            }
            topics(roomID: $roomID){
                    id
                    roomID
                    name
                    avgScore
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
</script>