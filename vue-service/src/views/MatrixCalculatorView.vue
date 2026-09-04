<script setup lang="ts">
import { computed, ref } from 'vue';
import { useMatrixStore } from '../stores/matrix.store';
import type { Matrix } from '../types/matrix';
import MatrixInput from '../components/MatrixInput.vue';
import MatrixTable from '../components/MatrixTable.vue';
import StatisticsCards from '../components/StatisticsCards.vue';
import ApiFlow from '../components/ApiFlow.vue';
import ErrorAlert from '../components/ErrorAlert.vue';

const store = useMatrixStore();
const matrix = ref<Matrix>([]);
const precision = ref(4);
const hideInsignificantZeros = ref(false);

const canCalculate = computed(() => matrix.value.length > 0 && matrix.value.every((row) =>
  row.length > 0 && row.length === matrix.value[0].length && row.every((value) => Number.isFinite(value))
));

async function calculate(): Promise<void> {
  if (!canCalculate.value) return;
  await store.calculate(matrix.value);
}

function clearAll(): void {
  matrix.value = [];
  store.clearResult();
  store.clearError();
}
</script>

<template>
  <div class="app-shell">
    <header class="hero">
      <div class="hero__topline"><span class="brand-mark">M / QR</span><span class="status-dot"><i /> Pipeline online</span></div>
      <div class="hero__content">
        <div><span class="eyebrow hero-eyebrow">Linear algebra workstation</span><h1>Matrix QR<br /><em>Decomposition</em></h1><p>Explore orthogonal structure, inspect the resulting matrices, and send the full calculation through the service pipeline.</p></div>
        <div class="hero__badge"><v-icon icon="mdi-vector-triangle" size="40" /><span>Q · R<br /><small>A = QR</small></span></div>
      </div>
    </header>

    <main class="content-wrap">
      <div class="intro-row"><div><span class="eyebrow">Interactive console</span><h2>Build a matrix. Read the structure.</h2></div><div class="precision-control"><span>Precision</span><v-select v-model="precision" :items="[2, 3, 4, 5, 6]" density="compact" variant="outlined" hide-details /></div></div>
      <ErrorAlert v-if="store.error" :message="store.error" @dismiss="store.clearError" />
      <MatrixInput v-model="matrix" />
      <div class="calculate-bar"><div><v-icon icon="mdi-shield-check-outline" /><span>Validated before transmission</span></div><v-btn color="secondary" size="x-large" :loading="store.loading" :disabled="store.loading || !canCalculate" prepend-icon="mdi-calculator-variant" @click="calculate">{{ store.loading ? 'Calculating...' : 'Calculate QR' }}</v-btn><v-btn variant="text" :disabled="store.loading" @click="clearAll">Reset workspace</v-btn></div>
      <div v-if="store.qMatrix.length" class="results-area">
        <div class="results-heading"><div><span class="eyebrow">Section 02 / 03</span><h2>Decomposition Results</h2></div><v-switch v-model="hideInsignificantZeros" color="primary" label="Hide insignificant zeros" hide-details /></div>
        <div class="original-result"><MatrixTable :matrix="store.matrix" title="Original Matrix" :precision="precision" :hide-insignificant-zeros="hideInsignificantZeros" /></div>
        <div class="qr-grid"><MatrixTable :matrix="store.qMatrix" title="Q Matrix" :precision="precision" :hide-insignificant-zeros="hideInsignificantZeros" /><MatrixTable :matrix="store.rMatrix" title="R Matrix" :precision="precision" :hide-insignificant-zeros="hideInsignificantZeros" /></div>
        <StatisticsCards v-if="store.statistics" :statistics="store.statistics" :precision="precision" />
        <ApiFlow />
      </div>
      <div v-else class="empty-state"><v-icon icon="mdi-table-plus" size="34" /><span>Your result matrices will appear here</span><small>Generate a matrix above and run the calculation</small></div>
    </main>
    <footer>INTERSEGuro / MATRIX LAB <span>·</span> Go + Fiber · Node.js + Express · Vue.js</footer>
  </div>
</template>
