<script setup lang="ts">
const props = defineProps<{
  text: string
  active: boolean
}>()
const isActive = computed(() => {
  return props.active ? ' hover:-translate-y--2 hover:translate-x--2' : 'grayscale'
})
const isActiveBg = computed(() => {
  return props.active ? 'slicing' : 'grayscale'
})
const isActiveParent = computed(() => {
  return props.active ? 'parent' : ''
})
</script>

<template>
  <div class="relative inline-block" :class="isActiveParent">
    <div class="slicing absolute h-18 w-88 translate-x--1 translate-y-2 transform rounded-lg bg-#BF8B5E -bottom-1 -left-2" :class="isActiveBg" />
    <button class="relative z-10 h-15 w-full transform rounded-lg bg-#F0DDB7 px-4 py-2 text-xl text-#1C120D font-bold font-bold transition duration-150 ease-in-out ease-in-out" :class="isActive">
      {{ text }}
    </button>
  </div>
</template>

<style>
.slicing {
  clip-path: polygon(0 25%, 5% 0, 100% 84%, 96% 100%, 0 100%);
  transition: clip-path 0.1s ease-in-out;
}
.parent:hover .slicing {
  clip-path: polygon(0 25%, 2% 12%, 98% 90%, 96% 100%, 0 100%);
}
</style>
