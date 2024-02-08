import { acceptHMRUpdate, defineStore } from 'pinia'

export const useUserStore = defineStore('users', () => {
  const test = ref('hello')

  return {
    test,
  }
})

if (import.meta.hot)
  import.meta.hot.accept(acceptHMRUpdate(useUserStore as any, import.meta.hot))
