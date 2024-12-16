<script setup>
import { defineProps, useAttrs } from 'vue';
import MenuItem from "@/layout/componets/menu-item/index.vue";

const props = defineProps({
  item: {
    type: Object,
    required: true
  }
});

const attrs = useAttrs();
</script>

<template>
  <a-sub-menu v-if="props.item.children && props.item.children.length > 0" v-bind="attrs" :key="props.item.id">
    <template #icon>
      <font-awesome-icon v-if="props.item.icon !== ''" :icon="props.item.icon" />
      <font-awesome-icon v-if="props.item.icon === ''" icon="fa-solid fa-table-list" />
    </template>
    <template #title>
      <span>{{ props.item.title }}</span>
    </template>
    <template v-for="child in props.item.children" :key="child.id">
      <MenuItem :item="child" />
    </template>
  </a-sub-menu>
  <a-menu-item v-if="props.item.children && props.item.children.length === 0" v-bind="attrs" :key="props.item.id">
    <template #icon>
      <font-awesome-icon v-if="props.item.icon !== ''" :icon="props.item.icon" />
      <font-awesome-icon v-if="props.item.icon === ''" icon="fa-solid fa-table-list" />
    </template>
    {{ props.item.title }}
  </a-menu-item>
</template>