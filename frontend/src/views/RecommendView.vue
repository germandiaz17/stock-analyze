<template>
  <div class="p-6 min-h-screen bg-gray-100">
    <div class="max-w-xl mx-auto">
      <h2 class="text-3xl font-bold mb-6 text-blue-700 flex items-center gap-2">
        Recomendación de inversión
      </h2>

      <div v-if="loading" class="text-gray-600"> Cargando recomendación...</div>
      <div v-else-if="error" class="text-red-500"> Error: {{ error }}</div>

      <div
        v-else
        class="bg-white p-6 rounded-lg shadow-lg border border-gray-200 space-y-4 transition hover:shadow-xl"
      >
        <p class="text-lg">
          <strong class="text-gray-700"><font-awesome-icon icon="fa-building" /> Empresa: </strong>
          <span class="text-gray-900">{{ stock?.empresa }}</span>
        </p>

        <p class="text-lg">
          <strong class="text-gray-700"><font-awesome-icon icon="fa-circle-check" /> Ticker: </strong>
          <span class="text-gray-900">{{ stock?.ticker }}</span>
        </p>

        <p class="text-lg">
          <strong class="text-gray-700"><font-awesome-icon icon="fa-circle-dot" /> Objetivo: </strong>
          <span :class="objetivoClass(stock)">
            {{ stock?.objetivo_desde }} → {{ stock?.objetivo_hasta }}
          </span>
        </p>

        <p class="text-lg">
          <strong class="text-gray-700"><font-awesome-icon icon="fa-chart-column" /> Calificación: </strong>
          <span class="text-blue-700"> {{ stock?.calificacion_desde }} → {{ stock?.calificacion_hasta }}
          </span>
        </p>

        <p class="text-lg text-gray-600">
          <strong><font-awesome-icon icon="fa-calendar" /> Fecha:</strong>
          {{ stock?.tiempo ? new Date(stock.tiempo).toLocaleDateString() : '' }}
        </p>
      </div>
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

function objetivoClass(s: Stock | null): string {
  if (!s) return ''
  const from = parseFloat(s.objetivo_desde.replace(',', '.').replace('$', ''))
  const to = parseFloat(s.objetivo_hasta.replace(',', '.').replace('$', ''))
  if (to > from) return 'text-green-600 font-semibold'
  if (to < from) return 'text-red-600 font-semibold'
  return 'text-gray-700'
}

onMounted(async () => {
  try {
    const res = await fetch(`${import.meta.env.VITE_API_URL}/stocks/recommend`)
    const data = await res.json()
    stock.value = data
  } catch (_) {
    error.value = 'No se pudo obtener la recomendación'
  } finally {
    loading.value = false
  }
})
</script>
