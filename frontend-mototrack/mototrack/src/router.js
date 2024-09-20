import { createRouter, createWebHistory } from 'vue-router';
import Login from './components/Login.vue';
import Register from './components/Register.vue';
import Control from './components/Control.vue';

const routes = [
  { path: '/', component: Login },
  { path: '/register', component: Register },
  { path: '/login', component: Login },
  { 
    path: '/control', 
    component: Control,
    meta: { requiresAuth: true } // Adiciona meta para verificar autenticação
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

// Navigation guard
router.beforeEach((to, from, next) => {
  const userId = localStorage.getItem('userId');
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (!userId) {
      next({ path: '/login' }); // Redireciona para a página de login se não estiver logado
    } else {
      next(); // Permite a navegação se estiver logado
    }
  } else {
    next(); // Permite a navegação para rotas que não requerem autenticação
  }
});

export default router;
