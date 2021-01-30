<template>
  <div class="bg-gray-800">
    <Sidebar />
    <Editor />
  </div>
</template>

<script lang="ts">
import { ref, defineComponent } from "vue";
import Editor from "@/components/Editor.vue"; // @ is an alias to /src
import Sidebar from "components/Sidebar.vue";

export default defineComponent({
  name: "Home",
  components: {
    Editor,
    Sidebar
  },
  setup() {
    const message = ref("Click the Icon");

    const getMessage = () => {
      window.backend.basic().then(result => {
        message.value = result;
      });
    };

    return { message: message, getMessage: getMessage };
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
