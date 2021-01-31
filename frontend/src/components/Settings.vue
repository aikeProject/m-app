<template>
  <section>
    <h1>Options</h1>
    <p @click="closeView">Close</p>
    <p @click="openDir">{{ config.outDir }}</p>
    <label for="target">Target</label>
    <select class="text-black" name="target" id="target" @change="selectTarget">
      <option value="webp" selected>WebP</option>
      <option value="jpg">JPG</option>
      <option value="png">PNG</option>
    </select>
    <button
      class="flex ml-auto bg-purple-default border-0 py-2 px-6 focus:outline-none hover:bg-indigo-600 rounded-full text-gray-900"
      @click="selectOutDir"
    >
      输出目录
    </button>
  </section>
</template>

<script>
export default {
  name: "Settings",
  computed: {
    /**
     * 从状态管理中获取config配置参数
     */
    config() {
      return this.$store.getters.config;
    }
  },
  methods: {
    closeView() {
      this.$emit("close-view");
    },
    // 选择输出目录
    selectOutDir() {
      window.backend.Config.SetOutDir()
        .then(result => {
          console.log(result);
          this.$store.dispatch("getConfig");
        })
        .catch(err => console.error(err));
    },
    /**
     * 选择转换格式
     */
    selectTarget(e) {
      const target = e.target;
      console.log(target);
      window.backend.Config.SetTarget(target.value)
        .then(res => {
          console.log(res);
          this.$store.dispatch("getConfig");
        })
        .catch(err => {
          console.error(err);
        });
    },
    // 打开文件输出目录
    openDir() {
      window.backend.Config.OpenOutputDir()
        .then(res => console.log(res))
        .catch(err => console.error(err));
    }
  }
};
</script>

<style scoped></style>
