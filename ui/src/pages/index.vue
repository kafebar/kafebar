<script setup lang="ts" generic="T extends any, O extends any">
defineOptions({
  name: 'IndexPage',
})

// const name = ref('')

// const router = useRouter()

const hi = () => {
  console.log('hi')
}
</script>

<template>
  <div @click="hi()">
   hi
  </div>
</template>
