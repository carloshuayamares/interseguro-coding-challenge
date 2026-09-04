<script setup lang="ts">
import { computed, ref } from 'vue';
import type { Matrix } from '../types/matrix';

const props = defineProps<{ modelValue: Matrix }>();
const emit = defineEmits<{ 'update:modelValue': [value: Matrix] }>();

const rows = ref(Math.max(props.modelValue.length, 3));
const columns = ref(Math.max(props.modelValue[0]?.length ?? 2, 2));
const validationMessage = ref('');

const hasMatrix = computed(() => props.modelValue.length > 0);

function generateMatrix(): void {
  if (rows.value < 1 || columns.value < 1 || !Number.isInteger(rows.value) || !Number.isInteger(columns.value)) {
    validationMessage.value = 'Rows and columns must be whole numbers greater than or equal to 1.';
    return;
  }
  const current = props.modelValue;
  emit('update:modelValue', Array.from({ length: rows.value }, (_, row) =>
    Array.from({ length: columns.value }, (_, column) => current[row]?.[column] ?? 0)
  ));
  validationMessage.value = '';
}

function updateValue(rowIndex: number, columnIndex: number, rawValue: string): void {
  const value = Number(rawValue);
  const next = props.modelValue.map((row) => [...row]);
  next[rowIndex][columnIndex] = value;
  emit('update:modelValue', next);
}

function clear(): void {
  emit('update:modelValue', []);
  validationMessage.value = '';
}
</script>

<template>
  <section class="input-panel">
    <div class="section-heading">
      <div>
        <span class="eyebrow">Section 01</span>
        <h2>Matrix Input</h2>
      </div>
      <v-icon icon="mdi-grid" size="28" />
    </div>
    <div class="dimension-controls">
      <v-text-field v-model.number="rows" label="Rows" type="number" min="1" variant="outlined" density="comfortable" hide-details="auto" />
      <v-text-field v-model.number="columns" label="Columns" type="number" min="1" variant="outlined" density="comfortable" hide-details="auto" />
      <v-btn color="primary" variant="flat" size="large" prepend-icon="mdi-grid-plus" @click="generateMatrix">Generate Matrix</v-btn>
    </div>
    <p v-if="validationMessage" class="field-error">{{ validationMessage }}</p>
    <div v-if="hasMatrix" class="editable-matrix matrix-scroll">
      <div v-for="(row, rowIndex) in modelValue" :key="rowIndex" class="editable-row">
        <v-text-field
          v-for="(value, columnIndex) in row"
          :key="columnIndex"
          :model-value="value"
          type="number"
          variant="outlined"
          density="compact"
          hide-details
          aria-label="Matrix value"
          @update:model-value="updateValue(rowIndex, columnIndex, String($event))"
        />
      </div>
    </div>
    <div class="input-actions">
      <v-btn variant="text" prepend-icon="mdi-delete-outline" @click="clear">Clear</v-btn>
      <span class="matrix-size" v-if="hasMatrix">{{ modelValue.length }} × {{ modelValue[0]?.length ?? 0 }}</span>
    </div>
  </section>
</template>
