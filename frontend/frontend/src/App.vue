<template>
  <div class="min-h-screen bg-gray-100">
    <nav class="bg-blue-600 text-white p-4 flex gap-4">
      <RouterLink to="/" class="hover:underline">Listado</RouterLink>
      <RouterLink to="/recommend" class="hover:underline">Recomendación</RouterLink>
    </nav>

    <main class="p-6">
      <RouterView />
    </main>
  </div>
</template>

<script setup lang="ts">
</script>
