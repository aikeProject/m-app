<template>
  <div class="bg-gray-900 flex h-full">
    <Sidebar v-on:select-view="handleViewSelect" />
    <component :is="currentView" v-on:close-view="handleViewClose" />
  </div>
</template>

<script lang="ts">
import { ref, defineComponent, onMounted } from "vue";
import { useStore } from "vuex";
import Editor from "@/components/Editor.vue";
import Sidebar from "components/Sidebar.vue";
import Settings from "components/Settings.vue";
import About from "components/About.vue";

export default defineComponent({
  name: "Home",
  components: {
    Editor,
    Sidebar,
    Settings,
    About
  },
  setup() {
    const store = useStore();
    const currentView = ref("Editor");

    const handleViewClose = () => {
      currentView.value = "Editor";
    };

    const handleViewSelect = (e: string) => {
      currentView.value = e;
    };

    onMounted(() => {
      store.dispatch("getConfig");
    });

    return { currentView, handleViewClose, handleViewSelect };
  }
});
</script>

<style scoped lang="stylus">
.logo
  width 16em
</style>
