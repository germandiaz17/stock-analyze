<template>
  <div class="min-h-screen bg-gray-100 px-4 py-8">
    <div class="max-w-7xl mx-auto">
      <h1 class="text-3xl font-bold text-blue-700 mb-6">Acciones del día</h1>

      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-8">
        <!-- Gráfico de barras por tendencia -->
        <div class="bg-white p-4 rounded-lg shadow-sm">
          <h2 class="text-xl font-semibold text-gray-800 mb-4">Acciones por tendencia</h2>
          <canvas ref="tendenciaChart" class="max-h-80"></canvas>
        </div>
        
        <!-- Gráfico circular por empresa recomendada -->
        <div class="bg-white p-4 rounded-lg shadow-sm">
          <h2 class="text-xl font-semibold text-gray-800 mb-4">Empresas más recomendadas</h2>
          <canvas ref="empresaChart" class="max-h-80"></canvas>
        </div>
      </div>

      <!-- Filtros de búsqueda -->
      <div class="mb-6 bg-white p-4 rounded-lg shadow-sm">
        <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
          <!-- Filtro por ticker -->
          <div>
            <label for="tickerFilter" class="block text-sm font-medium text-gray-700 mb-1">Buscar por ticker</label>
            <input
              id="tickerFilter"
              v-model="tickerFilter"
              type="text"
              placeholder="Ej: AAPL, MSFT..."
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              @input="resetPagination"
            />
          </div>

          <div>
            <label for="nameFilter" class="block text-sm font-medium text-gray-700 mb-1">Buscar por nombre de empresa</label>
            <input
              id="nameFilter"
              v-model="nameFilter"
              type="text"
              placeholder="Ej: Apple, Truora..."
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              @input="resetPagination"
            />
          </div>
          
          <!-- Filtro por tendencia -->
          <div>
            <label for="tendenciaFilter" class="block text-sm font-medium text-gray-700 mb-1">Filtrar por tendencia</label>
            <select
              id="tendenciaFilter"
              v-model="tendenciaFilter"
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              @change="resetPagination"
            >
              <option value="all">Todas las tendencias</option>
              <option value="up">Subiendo</option>
              <option value="down">Bajando</option>
              <option value="neutral">Sin cambio</option>
            </select>
          </div>
          
          <!-- Contador de resultados -->
          <div class="flex items-end">
            <span class="text-sm text-gray-500">
              Mostrando {{ paginatedStocks.length }} de {{ filteredStocks.length }} acciones (Página {{ currentPage }} de {{ totalPages }})
            </span>
          </div>
        </div>
      </div>

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
              v-for="stock in paginatedStocks"
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
        
        <!-- Mensaje cuando no hay resultados -->
        <div v-if="filteredStocks.length === 0" class="p-8 text-center text-gray-500">
          No se encontraron acciones que coincidan con los filtros
        </div>
        
        <!-- Paginación -->
        <div v-if="filteredStocks.length > itemsPerPage" class="flex items-center justify-between px-4 py-3 bg-white border-t border-gray-200 sm:px-6">
          <div class="flex-1 flex justify-between sm:hidden">
            <button
              @click="prevPage"
              :disabled="currentPage === 1"
              class="relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50"
            >
              Anterior
            </button>
            <button
              @click="nextPage"
              :disabled="currentPage === totalPages"
              class="ml-3 relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50"
            >
              Siguiente
            </button>
          </div>
          <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
            <div>
              <p class="text-sm text-gray-700">
                Mostrando <span class="font-medium">{{ (currentPage - 1) * itemsPerPage + 1 }}</span>
                a <span class="font-medium">{{ Math.min(currentPage * itemsPerPage, filteredStocks.length) }}</span>
                de <span class="font-medium">{{ filteredStocks.length }}</span> resultados
              </p>
            </div>
            <div>
              <nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px" aria-label="Pagination">
                <button
                  @click="goToPage(1)"
                  :disabled="currentPage === 1"
                  class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50"
                >
                  <span class="sr-only">Primera</span>
                  <font-awesome-icon icon="fa-backward" />
                </button>
                <button
                  @click="prevPage"
                  :disabled="currentPage === 1"
                  class="relative inline-flex items-center px-2 py-2 border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50"
                >
                  <span class="sr-only">Anterior</span>
                  <font-awesome-icon icon="fa-arrow-left" />
                </button>
                
                <!-- Números de página -->
                <template v-for="page in visiblePages" :key="page">
                  <button
                    @click="goToPage(page)"
                    :class="{
                      'bg-blue-50 border-blue-500 text-blue-600 z-10': currentPage === page,
                      'bg-white border-gray-300 text-gray-500 hover:bg-gray-50': currentPage !== page
                    }"
                    class="relative inline-flex items-center px-4 py-2 border text-sm font-medium"
                  >
                    {{ page }}
                  </button>
                </template>
                
                <button
                  @click="nextPage"
                  :disabled="currentPage === totalPages"
                  class="relative inline-flex items-center px-2 py-2 border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50"
                >
                  <span class="sr-only">Siguiente</span>
                  <font-awesome-icon icon="fa-arrow-right" />
                </button>
                <button
                  @click="goToPage(totalPages)"
                  :disabled="currentPage === totalPages"
                  class="relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50"
                >
                  <span class="sr-only">Última</span>
                  <font-awesome-icon icon="fa-forward" />
                </button>
              </nav>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch, nextTick  } from 'vue'
import axios from 'axios'
import Chart from 'chart.js/auto'

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
const tickerFilter = ref('')
const nameFilter = ref('')
const tendenciaFilter = ref('all')
const currentPage = ref(1)
const itemsPerPage = ref(10)

const tendenciaChart = ref<HTMLCanvasElement | null>(null)
const empresaChart = ref<HTMLCanvasElement | null>(null)
let tendenciaChartInstance: Chart | null = null
let empresaChartInstance: Chart | null = null

// Computed property para filtrar las acciones
const filteredStocks = computed(() => {
  return stocks.value.filter(stock => {
    // Filtro por ticker (case insensitive)
    const matchesTicker = stock.ticker.toLowerCase().includes(tickerFilter.value.toLowerCase())
    const matchesName = stock.empresa.toLowerCase().includes(nameFilter.value.toLowerCase())
    
    // Filtro por tendencia
    let matchesTendencia = true
    if (tendenciaFilter.value === 'up') {
      matchesTendencia = isSubiendo(stock)
    } else if (tendenciaFilter.value === 'down') {
      matchesTendencia = isBajando(stock)
    } else if (tendenciaFilter.value === 'neutral') {
      matchesTendencia = isIgual(stock)
    }
    
    return matchesTicker && matchesTendencia && matchesName
  })
})

// Función para inicializar los gráficos
const initCharts = () => {
  // Destruir instancias anteriores si existen
  if (tendenciaChartInstance) tendenciaChartInstance.destroy()
  if (empresaChartInstance) empresaChartInstance.destroy()

  // Datos para el gráfico de tendencias
  const tendenciasData = {
    up: filteredStocks.value.filter(stock => isSubiendo(stock)).length,
    down: filteredStocks.value.filter(stock => isBajando(stock)).length,
    neutral: filteredStocks.value.filter(stock => isIgual(stock)).length
  }

  // Datos para el gráfico de empresas recomendadas (top 5)
  const empresasData = filteredStocks.value
    .filter(stock => stock.calificacion_hasta === 'Comprar') // Solo acciones con recomendación COMPRA
    .reduce((acc, stock) => {
      acc[stock.empresa] = (acc[stock.empresa] || 0) + 1
      return acc
    }, {} as Record<string, number>)

  const sortedEmpresas = Object.entries(empresasData)
    .sort((a, b) => b[1] - a[1])
    .slice(0, 20) // Top 5 empresas

  // Crear gráfico de barras para tendencias
  if (tendenciaChart.value) {
    tendenciaChartInstance = new Chart(tendenciaChart.value, {
      type: 'bar',
      data: {
        labels: ['Subiendo', 'Bajando', 'Sin cambio'],
        datasets: [{
          label: 'Cantidad de acciones',
          data: [tendenciasData.up, tendenciasData.down, tendenciasData.neutral],
          backgroundColor: [
            'rgba(16, 185, 129, 0.7)', // verde para subiendo
            'rgba(239, 68, 68, 0.7)',  // rojo para bajando
            'rgba(156, 163, 175, 0.7)' // gris para neutral
          ],
          borderColor: [
            'rgba(16, 185, 129, 1)',
            'rgba(239, 68, 68, 1)',
            'rgba(156, 163, 175, 1)'
          ],
          borderWidth: 1
        }]
      },
      options: {
        responsive: true,
        maintainAspectRatio: false,
        scales: {
          y: {
            beginAtZero: true,
            ticks: {
              stepSize: 1
            }
          }
        },
        plugins: {
          legend: {
            display: false
          }
        }
      }
    })
  }

  // Crear gráfico circular para empresas recomendadas
  if (empresaChart.value && sortedEmpresas.length > 0) {
    const backgroundColors = [
      'rgba(59, 130, 246, 0.7)',  // azul
      'rgba(16, 185, 129, 0.7)',  // verde
      'rgba(245, 158, 11, 0.7)',  // amarillo
      'rgba(139, 92, 246, 0.7)', // violeta
      'rgba(20, 184, 166, 0.7)',  // turquesa
    ]

    empresaChartInstance = new Chart(empresaChart.value, {
      type: 'pie',
      data: {
        labels: sortedEmpresas.map(item => item[0]),
        datasets: [{
          label: 'Recomendaciones COMPRA',
          data: sortedEmpresas.map(item => item[1]),
          backgroundColor: backgroundColors,
          borderColor: backgroundColors.map(color => color.replace('0.7', '1')),
          borderWidth: 1
        }]
      },
      options: {
        responsive: true,
        maintainAspectRatio: false,
        plugins: {
          legend: {
            position: 'right'
          },
          tooltip: {
            callbacks: {
              label: function(context) {
                const label = context.label || ''
                const value = Number(context.raw) || 0
                const total = context.dataset.data.reduce((a: number, b: number) => Number(a) + Number(b), 0)
                const percentage = Math.round((value / total) * 100)
                return `${label}: ${value} (${percentage}%)`
              }
            }
          }
        }
      }
    })
  }
}

// Actualizar gráficos cuando cambian los datos filtrados
watch(filteredStocks, () => {
  nextTick(() => {
    initCharts()
  })
})

// Inicializar gráficos cuando el componente se monta
onMounted(() => {
  nextTick(() => {
    initCharts()
  })
})

// Computed property para las acciones paginadas
const paginatedStocks = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage.value
  const end = start + itemsPerPage.value
  return filteredStocks.value.slice(start, end)
})

// Computed property para el total de páginas
const totalPages = computed(() => {
  return Math.ceil(filteredStocks.value.length / itemsPerPage.value)
})

// Computed property para las páginas visibles en el paginador
const visiblePages = computed(() => {
  const pages = []
  const maxVisible = 5 // Máximo de páginas visibles en el paginador
  let start = Math.max(1, currentPage.value - Math.floor(maxVisible / 2))
  let end = Math.min(totalPages.value, start + maxVisible - 1)
  
  // Ajustar si estamos cerca del final
  if (end - start + 1 < maxVisible) {
    start = Math.max(1, end - maxVisible + 1)
  }
  
  for (let i = start; i <= end; i++) {
    pages.push(i)
  }
  
  return pages
})

// Métodos para navegar entre páginas
function goToPage(page: number) {
  currentPage.value = page
}

function nextPage() {
  if (currentPage.value < totalPages.value) {
    currentPage.value++
  }
}

function prevPage() {
  if (currentPage.value > 1) {
    currentPage.value--
  }
}

function resetPagination() {
  currentPage.value = 1
}

// Resetear a página 1 cuando cambian los filtros
watch([tickerFilter, nameFilter, tendenciaFilter], () => {
  resetPagination()
})

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