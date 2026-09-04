<script setup lang="ts">
import type { Statistics } from '../types/matrix';

defineProps<{ statistics: Statistics; precision: number }>();

function display(value: number, precision: number): string {
  return value.toFixed(precision);
}
</script>

<template>
  <section class="statistics-section">
    <div class="section-heading">
      <div>
        <span class="eyebrow">Section 04</span>
        <h2>Statistics</h2>
      </div>
      <v-icon icon="mdi-chart-box-outline" size="28" />
    </div>
    <div class="stats-grid">
      <article class="stat-card stat-card--warm"><v-icon icon="mdi-arrow-up-bold" /><span>Maximum</span><strong>{{ display(statistics.max, precision) }}</strong></article>
      <article class="stat-card stat-card--cool"><v-icon icon="mdi-arrow-down-bold" /><span>Minimum</span><strong>{{ display(statistics.min, precision) }}</strong></article>
      <article class="stat-card stat-card--lime"><v-icon icon="mdi-approximately-equal" /><span>Average</span><strong>{{ display(statistics.average, precision) }}</strong></article>
      <article class="stat-card stat-card--ink"><v-icon icon="mdi-sigma" /><span>Total Sum</span><strong>{{ display(statistics.sum, precision) }}</strong></article>
      <article class="stat-card stat-card--outline"><v-icon icon="mdi-axis-arrow" /><span>Diagonal Matrix</span><strong>{{ statistics.hasDiagonalMatrix ? 'Yes' : 'No' }}</strong></article>
    </div>
  </section>
</template>
