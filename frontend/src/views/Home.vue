<template>
  <div class="bg-gray-900 flex h-full">
    <Sidebar :active="currentView" v-on:select-view="handleViewSelect" />
    <keep-alive>
      <component :is="currentView" v-on:close-view="handleViewClose" />
    </keep-alive>
    <Notification />
  </div>
</template>

<script lang="ts">
import { ref, defineComponent, onMounted } from "vue";
import { useStore } from "vuex";
import Editor from "@/components/Editor.vue";
import Sidebar from "components/Sidebar.vue";
import Settings from "components/Settings.vue";
import About from "components/About.vue";
import Notification from "components/Notification.vue";

export default defineComponent({
  name: "Home",
  components: {
    Editor,
    Sidebar,
    Settings,
    About,
    Notification
  },
  setup() {
    const store = useStore();
    const currentView = ref("Editor");

    const handleViewClose = () => {
      currentView.value = "Editor";
    };

    const handleViewSelect = (e: string) => {
      if (currentView.value === e) {
        currentView.value = "Editor";
      } else {
        currentView.value = e;
      }
    };

    onMounted(() => {
      store.dispatch("getConfig");
      store.dispatch("getStats");
    });

    return { currentView, handleViewClose, handleViewSelect };
  }
});
</script>

<style scoped lang="stylus">
.logo
  width 16em
</style>
