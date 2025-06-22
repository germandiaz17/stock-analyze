<template>
  <div class="min-h-screen bg-gray-100 flex flex-col">
    <nav class="bg-blue-600 text-white p-4 flex gap-4">
      <RouterLink to="/" class="hover:underline">Listado</RouterLink>
      <RouterLink to="/recommend" class="hover:underline">Recomendación</RouterLink>
    </nav>

    <main class="p-6 flex-grow">
      <RouterView />
    </main>

    <footer class="bg-blue-700 text-white text-center py-4">
      Creado con <font-awesome-icon icon="fa-heart" /> por German Diaz
    </footer>
  </div>
</template>
