<template>
  <div class="p-6">
    <h2 class="text-2xl font-bold mb-4">Recomendación de inversión</h2>

    <div v-if="loading" class="text-gray-600">Cargando recomendación...</div>
    <div v-else-if="error" class="text-red-500">Error: {{ error }}</div>
    <div v-else class="bg-white p-4 rounded shadow">
      <p><strong>Empresa:</strong> {{ stock?.empresa }}</p>
      <p><strong>Ticker:</strong> {{ stock?.ticker }}</p>
      <p><strong>Objetivo:</strong> {{ stock?.objetivo_desde }} → {{ stock?.objetivo_hasta }}</p>
      <p><strong>Calificación:</strong> {{ stock?.calificacion_desde }} → {{ stock?.calificacion_hasta }}</p>
      <p><strong>Fecha:</strong> {{ stock?.tiempo ? new Date(stock.tiempo).toLocaleDateString() : '' }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

interface Stock {
  ticker: string
  empresa: string
  objetivo_desde: string
  objetivo_hasta: string
  calificacion_desde: string
  calificacion_hasta: string
  tiempo: string
}

const stock = ref<Stock | null>(null)
const loading = ref(true)
const error = ref<string | null>(null)

onMounted(async () => {
  try {
    const res = await fetch('http://localhost:3000/stocks/recommend')
    const data = await res.json()
    stock.value = data
  } catch (_) {
    error.value = 'No se pudo obtener la recomendación'
  } finally {
    loading.value = false
  }
})
</script>
