<template>
  <Teleport to="body">
        <Transition name="modal">
            <div v-if="show" class="fixed flex justify-center items-center transition-opacity z-40 top-0 left-0 w-full h-full bg-black bg-opacity-60" @click="$emit('close')">
                <div class="modal-container max-w-96 w-full mx-4 px-6 py-4 bg-gray-700 rounded-sm shadow-md transition-all" @click.stop>
                    <div class="flex justify-between gap-4 mb-4">
                        <h3 class="text-xl">{{ name }}</h3>
                        <button @click="$emit('close')">X</button>
                    </div>
                    <slot></slot>
                </div>
            </div>
        </Transition>
    </Teleport>
</template>

<script setup lang="ts">
import { Transition, Teleport } from 'vue';

defineProps({
  show: {type: Boolean, default: false},
  name: {type: String}
})

defineEmits(['close'])

</script>

<style>

.modal-enter-from {
  opacity: 0;
}

.modal-leave-to {
  opacity: 0;
}

.modal-enter-from .modal-container,
.modal-leave-to .modal-container {
  -webkit-transform: scale(1.1);
  transform: scale(1.1);
}
</style>