<template>
  <div class="home">
    <img
      @click="getMessage"
      alt="Vue logo"
      src="../assets/appicon.png"
      :style="{ height: '400px' }"
    />
    <div
      class="p-6 max-w-sm mx-auto bg-white rounded-xl shadow-md flex items-center space-x-4"
    >
      <div class="flex-shrink-0">
        <img
          class="h-12 w-12"
          src="../assets/appicon.png"
          alt="ChitChat Logo"
        />
      </div>
      <div>
        <div class="text-xl font-medium text-black">ChitChat</div>
        <p class="text-gray-500">You have a new message!</p>
      </div>
    </div>
    <div class="h-4"></div>
    <div
      class="py-8 px-8 max-w-sm mx-auto bg-white rounded-xl shadow-md space-y-2 sm:py-4 sm:flex sm:items-center sm:space-y-0 sm:space-x-6"
    >
      <img
        class="block mx-auto h-24 rounded-full sm:mx-0 sm:flex-shrink-0"
        src="../assets/appicon.png"
        alt="Woman's Face"
      />
      <div class="text-center space-y-2 sm:text-left">
        <div class="space-y-0.5">
          <p class="text-lg text-black font-semibold">
            Erin Lindford
          </p>
          <p class="text-gray-500 font-medium">
            Product Engineer
          </p>
        </div>
        <button
          class="px-4 py-1 text-sm text-purple-600 font-semibold rounded-full border border-purple-200 hover:text-white hover:bg-purple-600 hover:border-transparent focus:outline-none focus:ring-2 focus:ring-purple-600 focus:ring-offset-2"
        >
          Message
        </button>
      </div>
    </div>
    <button
      class="py-2 px-4 font-semibold rounded-lg shadow-md text-white bg-green-500 hover:bg-green-700"
    >
      Click Me
    </button>
    <button class="btn btn-green">
      Button
    </button>
    <div
      class="max-w-md mx-auto bg-white rounded-xl shadow-md overflow-hidden md:max-w-2xl"
    >
      <div class="md:flex">
        <div class="md:flex-shrink-0">
          <img
            class="h-48 w-full object-cover md:w-48"
            src="../assets/appicon.png"
            alt=""
          />
        </div>
        <div class="p-8">
          <div
            class="uppercase tracking-wide text-sm text-indigo-500 font-semibold"
          >
            Case study
          </div>
          <a
            href="#"
            class="block mt-1 text-lg leading-tight font-medium text-black hover:underline"
          >
            床前明月光，疑似地上霜。举头望明月，低头思故乡
          </a>
          <p class="mt-2 text-gray-500">
            Getting a new business off the ground is a lot of hard work. Here
            are five ideas you can use to find your first customers.
          </p>
        </div>
      </div>
    </div>
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
.btn
  @apply py-2 px-4 font-semibold rounded-lg shadow-md

.btn-green
  @apply text-white bg-green-500
  &:hover
    @apply bg-green-700
</style>
