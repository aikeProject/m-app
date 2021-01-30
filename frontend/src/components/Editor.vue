<template>
  <div class="mx-auto p-10">
    <p @click="openDir">{{ config.outDir }}</p>
    <p>{{ config.target }}</p>
    <input
      type="file"
      accept="image/jpeg, image/png, image/jpg, image/webp"
      multiple
      @change="processFileInput"
    />
    <label for="target">Target</label>
    <select class="text-black" name="target" id="target" @change="selectTarget">
      <option value="webp" selected>WebP</option>
      <option value="jpg">JPG</option>
      <option value="png">PNG</option>
    </select>
    <button
      class="flex ml-auto text-white bg-indigo-500 border-0 py-2 px-2 px-6 focus:outline-none hover:bg-indigo-600 rounded"
      @click="selectOutDir"
    >
      选择文件夹
    </button>
    <button
      class="flex ml-auto text-white bg-indigo-500 border-0 py-2 px-2 px-6 focus:outline-none hover:bg-indigo-600 rounded"
      :disabled="!canConvert"
      @click="convert"
    >
      转换
    </button>
    <button
      class="flex ml-auto text-white bg-indigo-500 border-0 py-2 px-2 px-6 focus:outline-none hover:bg-indigo-600 rounded"
      @click="clear"
    >
      清空
    </button>
    <div v-if="files.length > 0" class="table-wrapper">
      <table class="table-auto w-full text-left whitespace-nowrap">
        <thead>
          <tr>
            <th
              class="px-4 py-3 title-font tracking-wider font-medium text-gray-900 text-sm bg-gray-200 rounded-tl rounded-bl"
            >
              文件名
            </th>
            <th
              class="px-4 py-3 title-font tracking-wider font-medium text-gray-900 text-sm bg-gray-200 rounded-tl rounded-bl"
            >
              大小
            </th>
            <th
              class="px-4 py-3 title-font tracking-wider font-medium text-gray-900 text-sm bg-gray-200 rounded-tl rounded-bl"
            >
              转换后大小
            </th>
            <th
              class="px-4 py-3 title-font tracking-wider font-medium text-gray-900 text-sm bg-gray-200 rounded-tl rounded-bl"
            >
              转化率
            </th>
            <th
              class="px-4 py-3 title-font tracking-wider font-medium text-gray-900 text-sm bg-gray-200 rounded-tl rounded-bl"
            >
              结果
            </th>
            <th
              class="px-4 py-3 title-font tracking-wider font-medium text-gray-900 text-sm bg-gray-200"
            >
              完成
            </th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(file, i) in files" :key="`${i}-${file.name}`">
            <td class="px-4 py-3">{{ file.filename }}</td>
            <td class="font-mono px-4 py-3">{{ getPrettySize(file.size) }}</td>
            <td class="font-mono px-4 py-3">
              {{ getPrettySize(file.convertedSize) }}
            </td>
            <td class="px-4 py-3">{{ getSavings(file) }}</td>
            <td class="px-4 py-3" @click="openFile(file)">
              {{ file.convertedPath }}
            </td>
            <td class="px-4 py-3">{{ file.isConverted }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { fExt, fName, fSize } from "lib/filw";
import { CFile } from "components/Editor";
import Wails from "@wailsapp/runtime";

export default defineComponent({
  name: "Editor",
  data() {
    return {
      files: [] as CFile[]
    };
  },
  computed: {
    canConvert(): boolean {
      if (this.files.length === 0) return false;
      return this.files.some(f => {
        return !f.isConverted;
      });
    },
    /**
     * 从状态管理中获取config配置参数
     */
    config() {
      return this.$store.getters.config;
    }
  },
  methods: {
    /**
     * 图片格式转换
     * eg: .jpg -> .webp
     */
    convert() {
      window.backend.FileManager.Convert()
        .then(result => {
          console.log(result);
        })
        .catch(error => console.error(error));
    },
    /**
     * 根据文件名和大小创建文件id
     * @param name
     * @param size
     */
    createFileId(name: string, size: number): string {
      return name + size.toString();
    },
    /**
     * 通过文件名查找文件
     * @param id
     */
    getFileById(id: string) {
      if (this.files.length === 0) return;
      return this.files.find(f => f.id === id);
    },
    /**
     * 转换文件大小的格式
     * @param size
     */
    getPrettySize(size: number) {
      if (size === 0) return "";
      return fSize(size);
    },
    /**
     * 转换率
     * => 转换后大小/原文件大小
     * @param file {CFile}
     */
    getSavings(file: CFile) {
      if (file.convertedSize === 0) return "";
      const p = Math.floor(100 - (file.convertedSize / file.size) * 100);
      return `${p.toString()}%`;
    },
    /**
     * 文件是否存在
     * @param id
     */
    hasFile(id: string): boolean {
      if (this.files.length === 0) return false;
      return this.files.some(f => f.id === id);
    },
    /**
     * 检查文件类型
     * @param type
     */
    isValidType(type: string): boolean {
      const v = [
        "image/png",
        "image/.jpg",
        "image/jpg",
        "image/jpeg",
        "image/webp"
      ];
      return v.indexOf(type) >= 0;
    },
    processFileInput(e: InputEvent) {
      const target = e.target as HTMLInputElement;
      const files: File[] = target.files ? [].slice.apply(target.files) : [];
      files.forEach(f => {
        const name = fName(f.name);
        const size = f.size;
        const id = this.createFileId(name, size);
        if (!this.isValidType(f.type) || this.hasFile(name)) return;
        this.processFile(f, id);
        this.files.push({
          id,
          name,
          size,
          filename: f.name,
          isConverted: false,
          convertedSize: 0,
          convertedPath: ""
        });
      });
    },
    processFile(file: File, id: string) {
      const reader = new FileReader();
      reader.onload = () => {
        if (reader.result) {
          const name = file.name;
          window.backend.FileManager.HandleFile(
            JSON.stringify({
              id,
              data: (reader.result as string).split(",")[1],
              ext: fExt(name),
              name: fName(name),
              type: file.type
            })
          );
        }
      };
      reader.readAsDataURL(file);
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
    selectTarget(e: Event) {
      const target = e.target as HTMLOptionElement;
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
    },
    /**
     * 打开文件
     */
    openFile(file: CFile) {
      window.backend.FileManager.OpenFile(file.convertedPath)
        .then(res => console.log(res))
        .catch(err => console.error(err));
    },
    /**
     * 清空选择的所有文件
     */
    clear() {
      this.files = [];
      window.backend.Config.Clear()
        .then(console.log)
        .catch(console.error);
    }
  },
  mounted() {
    Wails.Events.On("conversion:complete", e => {
      const f = this.getFileById(e.id);
      if (!f) return;
      f.isConverted = true;
      f.convertedSize = e.size;
      f.convertedPath = e.path;
    });
  }
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="stylus">
.table-wrapper
  min-height calc(100vh / 5)
  max-height calc(100vh / 2)
  overflow auto
</style>
