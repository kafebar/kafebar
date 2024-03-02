<script setup lang="ts">

</script>

<template>
  <div class="relative">
    <RouterView />
    <div class="h-48"></div>
    <div class="fixed flex bottom-0 w-full bg-gray-800">
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
 @apply bg-red-500
}
</style>