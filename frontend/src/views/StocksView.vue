<template>
  <div class="min-h-screen bg-gray-100 px-4 py-8">
    <div class="max-w-7xl mx-auto">
      <h1 class="text-3xl font-bold text-blue-700 mb-6">Acciones del día</h1>

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
            <tr
              v-for="stock in stocks"
              :key="stock.ticker"
              class="border-b hover:bg-gray-50 transition"
            >
              <td class="px-6 py-4 font-semibold text-gray-800">{{ stock.ticker }}</td>
              <td class="px-6 py-4 text-gray-700">{{ stock.empresa }}</td>
              <td class="px-6 py-4">
                <span class="font-medium text-gray-600">{{ stock.objetivo_desde }}</span>
                →
                <span
                  :class="{
                    'text-green-600 font-semibold': isSubiendo(stock),
                    'text-red-600 font-semibold': isBajando(stock),
                    'text-gray-500': isIgual(stock)
                  }"
                >
                  {{ stock.objetivo_hasta }}
                </span>
              </td>
              <td class="px-6 py-4">
                <span class="text-gray-700">{{ stock.calificacion_desde }}</span>
                →
                <span class="text-blue-600 font-medium">{{ stock.calificacion_hasta }}</span>
              </td>
              <td class="px-6 py-4 flex items-center gap-2">
                <font-awesome-icon
                  :icon="getTendenciaIcon(stock)"
                  :class="getTendenciaColor(stock)"
                  class="text-lg"
                />
                <span
                  :class="getTendenciaColor(stock)"
                >
                  {{ getTendenciaLabel(stock) }}
                </span>
              </td>
              <td class="px-6 py-4 text-gray-600">
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
import axios from 'axios'

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

function parse(obj: string): number {
  return parseFloat(obj.replace(',', '.').replace('$', ''))
}

function getTendenciaIcon(stock: Stock): any {
  if (isSubiendo(stock)) return 'up-long'
  if (isBajando(stock)) return 'down-long'
  return 'pause'
}

function getTendenciaLabel(stock: Stock): string {
  if (isSubiendo(stock)) return 'Subiendo'
  if (isBajando(stock)) return 'Bajando'
  return 'Sin cambio'
}

function getTendenciaColor(stock: Stock): string {
  if (isSubiendo(stock)) return 'text-green-600 font-semibold'
  if (isBajando(stock)) return 'text-red-600 font-semibold'
  return 'text-gray-500'
}

function isSubiendo(stock: Stock): boolean {
  return parse(stock.objetivo_hasta) > parse(stock.objetivo_desde)
}

function isBajando(stock: Stock): boolean {
  return parse(stock.objetivo_hasta) < parse(stock.objetivo_desde)
}

function isIgual(stock: Stock): boolean {
  return parse(stock.objetivo_hasta) === parse(stock.objetivo_desde)
}

onMounted(async () => {
  try {
    const { data } = await axios.get(`${import.meta.env.VITE_API_URL}/stocks`)
    stocks.value = data
  } catch (err) {
    error.value = 'No se pudieron cargar los datos'
  } finally {
    loading.value = false
  }
})
</script>
