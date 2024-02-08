<script setup lang="ts" generic="T extends any, O extends any">
import { useUserStore } from '~/stores/userStore'

defineOptions({
  name: 'IndexPage',
})
const userStore = useUserStore()
// const name = ref('')
const test = userStore.test
// const router = useRouter()
</script>

<template>
  <div>
    hi
    {{ test }}
    hi
  </div>
</template>
