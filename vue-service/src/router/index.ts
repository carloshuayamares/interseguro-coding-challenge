import { createRouter, createWebHistory } from 'vue-router';
import MatrixCalculatorView from '../views/MatrixCalculatorView.vue';

export const router = createRouter({
  history: createWebHistory(),
  routes: [{ path: '/', name: 'matrix-calculator', component: MatrixCalculatorView }]
});
