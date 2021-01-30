<template>
  <Sidebar v-on:select-view="handleViewSelect" />
  <component :is="currentView" v-on:close-view="handleViewClose" />
</template>

<script lang="ts">
import { ref, defineComponent } from "vue";
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
    const currentView = ref("Editor");

    const handleViewClose = () => {
      currentView.value = "Editor";
    };

    const handleViewSelect = (e: string) => {
      currentView.value = e;
    };

    return { currentView, handleViewClose, handleViewSelect };
  },
  mounted() {
    this.$store.dispatch("getConfig");
  }
});
</script>

<style scoped lang="stylus">
.logo
  width 16em
</style>
