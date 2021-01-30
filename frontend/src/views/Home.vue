<template>
  <div class="home">
    <img
      class="logo"
      @click="getMessage"
      alt="Vue logo"
      src="../assets/appicon.png"
    />
    <HelloWorld :msg="message" />
  </div>
</template>

<script lang="ts">
import { ref, defineComponent } from "vue";
import HelloWorld from "@/components/HelloWorld.vue"; // @ is an alias to /src

export default defineComponent({
  name: "Home",
  components: {
    HelloWorld
  },
  setup() {
    const message = ref("Click the Icon");

    const getMessage = () => {
      window.backend.basic().then(result => {
        message.value = result;
      });
    };

    return { message: message, getMessage: getMessage };
  }
});
</script>

<style scoped lang="stylus">
.logo
  width 16em
</style>
