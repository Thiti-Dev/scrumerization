<template>
    <div class="w-full flex bg-gray-100 justify-center rounded-sm py-2">
        <div v-for="point of [0,1,2,3,5,7,11,13]" :class="{'bg-green-200': point === currentPoint,'select-none': viewOnly,'cursor-pointer': !viewOnly}" @click="onCardSelect(point)" class="px-3 mx-1 w-6 h-8 bg-wheat-500 rounded-lg shadow-lg flex items-center justify-center">
            <span class="text-2 font-bold text-gray-800">{{ point }}</span>
        </div>
    </div>
</template>

<script setup lang="ts">
    const props = defineProps({
        viewOnly: Boolean,
        viewOnlyPoint: Number
    })


    const emit = defineEmits<{
    (e: 'point-select', point: number): void
    }>()
    const currentPoint = ref(props.viewOnlyPoint ?? 999)
    function onCardSelect(point:number){
        if(props.viewOnly) return
        // if(currentPoint.value !== 999) return
        currentPoint.value = point
        emit('point-select', point)
    }
</script>