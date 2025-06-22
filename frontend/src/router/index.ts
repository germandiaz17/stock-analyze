import { createRouter, createWebHistory } from 'vue-router'
import StocksView from '../views/StocksView.vue'
import RecommendView from '../views/RecommendView.vue'

const routes = [
  { path: '/', name: 'stocks', component: StocksView },
  { path: '/recommend', name: 'recommend', component: RecommendView },
]

export const router = createRouter({
  history: createWebHistory(),
  routes,
})
