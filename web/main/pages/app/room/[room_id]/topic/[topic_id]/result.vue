<template>
    <div class="font-workbench">
        <section v-if="votes" class="py-1 bg-blueGray-50">
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
                                    <li>
                                    <div class="flex items-center">
                                        <svg class="rtl:rotate-180 w-3 h-3 text-gray-400 mx-1" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 6 10">
                                        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 9 4-4-4-4"/>
                                        </svg>
                                        <a :href="`/app/room/${route.params.room_id}`" class="ms-1 text-sm font-medium text-gray-700 hover:text-blue-600 md:ms-2 dark:text-gray-400 dark:hover:text-white">Room</a>
                                    </div>
                                    </li>
                                    <li aria-current="page">
                                    <div class="flex items-center">
                                        <svg class="rtl:rotate-180 w-3 h-3 text-gray-400 mx-1" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 6 10">
                                        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 9 4-4-4-4"/>
                                        </svg>
                                        <span class="ms-1 text-sm font-medium text-gray-500 md:ms-2 dark:text-gray-400 select-none">Topic-Result</span>
                                    </div>
                                    </li>
                                    <li aria-current="page">
                                        <div class="flex items-center">
                                            <svg class="rtl:rotate-180 w-3 h-3 text-gray-400 mx-1" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 6 10">
                                            <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 9 4-4-4-4"/>
                                            </svg>
                                            <span class="ms-1 text-sm font-medium text-gray-500 md:ms-2 dark:text-gray-400 select-none">{{ result?.topic?.name }}</span>
                                        </div>
                                    </li>
                                </ol>
                            </nav>
                            <p class="font-semibold text-base text-blueGray-700 float-right"><span class="bg-blue-100 text-blue-800 text-sm font-medium me-2 px-2.5 py-0.5 rounded dark:bg-blue-900 dark:text-blue-300">Participants</span><span class="bg-red-100 text-red-800 text-xs font-medium me-2 px-2.5 py-0.5 rounded dark:bg-gray-700 dark:text-red-400 border border-red-400">{{ votes.length }}</span></p>
                        </div>
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
                                Voted
                        </th>
                    </tr>
                    </thead>
                    <tbody>
                        <tr class="table-auto" v-for="vote of votes">
                            <th class="border-t-0 px-6 align-middle border-l-0 border-r-0 text-xs whitespace-nowrap p-4 text-left text-blueGray-700 ">
                                {{ vote.user?.name }}
                            </th>
                            <td class="border-t-0 px-6 align-center border-l-0 border-r-0 text-xs whitespace-nowrap p-4">
                               {{ vote.votedDesc || '-' }}
                            </td>
                            <td class="border-t-0 px-6 align-middle border-l-0 border-r-0 text-xs whitespace-nowrap p-4">
                                <div class="w-full">
                                    <PointingCard view-only :view-only-point="vote.voted"/>
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

    const route = useRoute()
    const {result} = useQuery(GET_TOPIC_VOTES_QUERY,{where:{topicID: route.params.topic_id},topicID: route.params.topic_id},createApolloContextAuthorization())

    const votes = computed(() => result.value?.topicVotes)
</script>

<script lang="ts">
    import { useQuery } from '@vue/apollo-composable';
    import {graphql} from '../../../../../../.gen/gql'
    import { useRoute } from 'vue-router';

    const GET_TOPIC_VOTES_QUERY = graphql(`
        query topicVote($where:TopicVoteQueryWhereClause,$topicID: UUID!){
            topicVotes(where: $where){
                user: User{
                id
                username
                name
                }
                voted
                votedDesc
                createdAt
                updatedAt
            }
                topic(topicID:$topicID){
                    name
                    avgScore
                }
            }
    `)
</script>