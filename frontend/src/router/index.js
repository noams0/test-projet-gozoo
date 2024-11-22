import { createRouter, createWebHistory } from 'vue-router'
import HomeView from "@/views/HomeView.vue";

const routes = [
  { path: '/', name: 'Home', component: HomeView },
  // { path: '/simulation', name: 'Simulation', component: () => import('../views/Simulation.vue') }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
