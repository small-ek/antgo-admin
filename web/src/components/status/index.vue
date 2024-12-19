<script setup>
import {defineEmits, defineProps} from 'vue';

const props = defineProps({
  row: {type: Object, required: true}
});

const emit = defineEmits(['onClick'])

const onClick = (status) => emit('onClick', status, props.row)
</script>

<template>
  <a-tag bordered v-if="props.row.status===1" color="magenta" size="small" class="hand" @click="onClick(2)">已禁用
  </a-tag>
  <a-tag bordered v-if="props.row.status===2" size="small" class="hand" @click="onClick(1)">已启用</a-tag>
</template>
