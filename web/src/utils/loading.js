import { ref } from "vue";
import EventBus from "@/utils/eventBus.js";

const loading = ref(false);

EventBus.on('setLoading', (data) => {
    loading.value = data;
});

export { loading };