<script setup lang="ts">
import type { Matrix } from '../types/matrix';

withDefaults(
defineProps<{
  matrix: Matrix;
  title: string;
  precision?: number;
  hideInsignificantZeros?: boolean;
}>(), {
  precision: 4,
  hideInsignificantZeros: false
});

function formatValue(value: number, precision: number, hideZeros: boolean): string {
  if (hideZeros && Math.abs(value) < 10 ** -precision) return '0';
  return value.toFixed(precision);
}
</script>

<template>
  <section class="matrix-panel">
    <div class="matrix-panel__header">
      <span class="eyebrow">Output</span>
      <h3>{{ title }}</h3>
    </div>
    <div class="matrix-scroll">
      <table class="matrix-table" :aria-label="title">
        <tbody>
          <tr v-for="(row, rowIndex) in matrix" :key="rowIndex">
            <td v-for="(value, columnIndex) in row" :key="columnIndex">
              {{ formatValue(value, precision, hideInsignificantZeros) }}
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </section>
</template>
