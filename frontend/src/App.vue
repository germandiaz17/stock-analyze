<template>
  <div class="min-h-screen bg-gray-100 flex flex-col ">
    <nav class="bg-gray-800 text-white p-4 flex gap-4">
      <RouterLink to="/" :class="getLinkClass('/')">Listado</RouterLink>
      <RouterLink to="/recommend" :class="getLinkClass('/recommend')">Recomendación</RouterLink>
    </nav>

    <main class="p-6 flex-grow">
      <RouterView />
    </main>

    <footer class="bg-gray-800 text-white text-center py-4">
      Creado con <font-awesome-icon icon="fa-heart" /> por German Diaz
    </footer>
  </div>
</template>

<script setup>
import { useRoute } from 'vue-router'

const route = useRoute()

const getLinkClass = (path) => {
  const base = 'rounded-md px-3 py-2 text-sm font-medium'
  const inactive = 'text-gray-300 hover:bg-gray-700 hover:text-white'
  const active = 'bg-gray-700 text-white'

  return `${base} ${route.path === path ? active : inactive}`
}
</script>