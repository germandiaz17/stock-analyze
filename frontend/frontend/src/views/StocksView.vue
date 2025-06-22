<template>
  <div class="min-h-screen bg-gray-100 px-4 py-8">
    <div class="bg-red-500">
      <h1>pruebaa</h1>
    </div>
    <div class="max-w-7xl mx-auto">
      <h1 class="text-3xl font-bold text-blue-700 mb-6">📈 Acciones del día</h1>

      <div v-if="loading" class="text-gray-500">Cargando acciones...</div>
      <div v-else-if="error" class="text-red-500">Error: {{ error }}</div>

      <div v-else class="relative overflow-x-auto shadow-md sm:rounded-lg">
        <table class="min-w-full text-sm text-left">
          <thead class="bg-blue-100 text-blue-800 text-xs uppercase tracking-wider">
            <tr>
              <th class="px-6 py-3">Ticker</th>
              <th class="px-6 py-3">Empresa</th>
              <th class="px-6 py-3">Objetivo</th>
              <th class="px-6 py-3">Calificación</th>
              <th class="px-6 py-3">Tendencia</th>
              <th class="px-6 py-3">Fecha</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="stock in stocks" :key="stock.ticker" class="border-b hover:bg-gray-50">
              <td class="px-6 py-4 font-semibold">{{ stock.ticker }}</td>
              <td class="px-6 py-4">{{ stock.empresa }}</td>
              <td class="px-6 py-4">
                <span class="text-green-600">{{ stock.objetivo_desde }}</span>
                →
                <span class="text-blue-600">{{ stock.objetivo_hasta }}</span>
              </td>
              <td class="px-6 py-4">
                {{ stock.calificacion_desde }} → {{ stock.calificacion_hasta }}
              </td>
              <td class="px-6 py-4">
                {{ getTendencia(stock) }}
              </td>
              <td class="px-6 py-4">
                {{ new Date(stock.tiempo).toLocaleDateString() }}
              </td>
            </tr>
          </tbody>
        </table>
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

const stocks = ref<Stock[]>([])
const loading = ref(true)
const error = ref<string | null>(null)

function getTendencia(stock: Stock): string {
  const from = parseFloat(stock.objetivo_desde.replace(',', '.').replace('$', ''))
  const to = parseFloat(stock.objetivo_hasta.replace(',', '.').replace('$', ''))

  if (to > from) return '📈 Subiendo'
  if (to < from) return '📉 Bajando'
  return '⏸️ Igual'
}


onMounted(async () => {
  try {
    const res = await fetch('http://localhost:3000/stocks')
    const data = await res.json()
    stocks.value = data
  } catch (err) {
    error.value = 'No se pudieron cargar los datos'
  } finally {
    loading.value = false
  }
})
</script>
