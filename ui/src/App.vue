<script setup lang="ts">

</script>

<template>
  <div class="max-w-screen-lg mx-auto p-4 pb-48">
    <RouterView />
  </div>
  <div class="fixed bottom-0 w-full bg-gray-800">
    <div class="flex max-w-screen-lg mx-auto">
      <RouterLink to="/" class="router-link">Counter</RouterLink>
      <RouterLink to="/orders" class="router-link">Orders</RouterLink>
      <RouterLink to="/products" class="router-link">Products</RouterLink>
    </div>
  </div>
</template>

<style scoped>
.router-link {
  @apply grow text-center p-4 text-xl
}

.router-link.router-link-active {
 @apply bg-gray-900
}
</style>