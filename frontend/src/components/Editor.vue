<template>
  <div class="w-full p-10">
    <header class="border-b-2 border-gray-800 flex">
      <div class="w-1/2">
        <div
          class="bg-gray-800 flex flex-col items-center justify-center py-10"
          ref="dropZone"
        >
          <svg
            version="1.1"
            id="dropZone-plus"
            xmlns="http://www.w3.org/2000/svg"
            xmlns:xlink="http://www.w3.org/1999/xlink"
            x="0px"
            y="0px"
            viewBox="0 0 48 48"
            enable-background="new 0 0 48 48"
            width="48px"
            height="48px"
            xml:space="preserve"
          >
            <path
              fill="#808080"
              d="M47,20.6H28.4c-0.6,0-1-0.4-1-1V1c0-0.6-0.4-1-1-1h-4.9c-0.6,0-1,0.4-1,1v18.6c0,0.6-0.4,1-1,1H1
                c-0.6,0-1,0.4-1,1v4.9c0,0.6,0.4,1,1,1h18.6c0.6,0,1,0.4,1,1V47c0,0.6,0.4,1,1,1h4.9c0.6,0,1-0.4,1-1V28.4c0-0.6,0.4-1,1-1H47
                c0.6,0,1-0.4,1-1v-4.9C48,21,47.6,20.6,47,20.6z"
            />
          </svg>
          <p class="mt-6 text-gray-200">Drag and drop or select images</p>
        </div>
        <section class="flex justify-between py-6 w-full">
          <button
            class="bg-purple-default border-0 flex py-2 px-8 focus:outline-none hover:bg-green-default rounded-full text-gray-900"
            @click="convert"
            :disabled="!canConvert"
          >
            转换格式
          </button>
          <button
            class="bg-purple-default border-0 flex py-2 px-8 focus:outline-none hover:bg-green-default rounded-full text-gray-900"
            @click="clear"
          >
            清除
          </button>
        </section>
      </div>
      <div class="w-1/2">
        <div class="flex h-full items-center justify-center">
          <h2 class="leading-none text-4xl text-center text-green-default">
            Add image files<br />to get started
          </h2>
        </div>
      </div>
    </header>
    <p @click="openDir">{{ config.outDir }}</p>
    <p>{{ config.target }}</p>
    <input
      type="file"
      accept="image/jpeg, image/png, image/jpg, image/webp"
      multiple
      @input="processFileInput"
      ref="fileInput"
    />
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
      保存目录
    </button>
    <div v-if="files.length > 0" class="table-wrapper">
      <table class="table-auto w-full text-left whitespace-nowrap">
        <thead>
          <tr>
            <th
              class="font-medium pl-3 pt-6 text-gray-400 text-left text-sm tracking-wider uppercase"
            >
              文件名
            </th>
            <th
              class="font-medium pl-3 pt-6 text-gray-400 text-left text-sm tracking-wider uppercase"
            >
              大小
            </th>
            <th
              class="font-medium pl-3 pt-6 text-gray-400 text-left text-sm tracking-wider uppercase"
            >
              转换后大小
            </th>
            <th
              class="font-medium pl-3 pt-6 text-gray-400 text-left text-sm tracking-wider uppercase"
            >
              转化率
            </th>
            <th
              class="font-medium pl-3 pt-6 text-gray-400 text-left text-sm tracking-wider uppercase"
            >
              结果
            </th>
            <th
              class="font-medium pl-3 pt-6 text-gray-400 text-left text-sm tracking-wider uppercase"
            >
              完成
            </th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(file, i) in files" :key="`${i}-${file.name}`">
            <td class="cell-l">{{ file.filename }}</td>
            <td>{{ getPrettySize(file.size) }}</td>
            <td>
              {{ getPrettySize(file.convertedSize) }}
            </td>
            <td>{{ getSavings(file) }}</td>
            <td @click="openFile(file)">
              {{ file.convertedPath }}
            </td>
            <td>
              <p
                v-if="file.isConverted"
                class="flex items-center justify-center"
              >
                <svg
                  version="1.1"
                  :id="`${i}-check`"
                  xmlns="http://www.w3.org/2000/svg"
                  xmlns:xlink="http://www.w3.org/1999/xlink"
                  x="0px"
                  y="0px"
                  viewBox="0 0 20 20"
                  enable-background="new 0 0 20 20"
                  width="20px"
                  height="20px"
                  xml:space="preserve"
                >
                  <path
                    fill="#07FDBC"
                    d="M10,0C4.5,0,0,4.5,0,10s4.5,10,10,10s10-4.5,10-10S15.5,0,10,0z M8,14.4l-3.7-3.7l1.4-1.4L8,11.6l5.3-5.3
                            l1.4,1.4L8,14.4z"
                  />
                </svg>
              </p>
              <p v-else></p>
            </td>
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
     * 兼容ie的写法 ......
     */
    getFileType(type: string, ext: string) {
      if (this.isValidType(type)) return type;
      if (this.isValidExt(ext)) return `image/${ext}`;
      return "";
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
     * isValidExt returns true if the given file extension is of an accepted
     * set of extensions.
     * @param {string} ext - A file extension
     * @returns {boolean}
     */
    isValidExt(ext: string) {
      const v = ["jpg", "jpeg", "png", "webp"];
      return v.indexOf(ext) >= 0;
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
        const ext = fExt(f.name);
        const type = this.getFileType(f.type, ext);
        const id = this.createFileId(name, size);
        if (!type || this.hasFile(name)) return;
        this.processFile(f, id, type);
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
    processFile(file: File, id: string, type: string) {
      const reader = new FileReader();
      reader.onload = () => {
        if (reader.result) {
          const name = file.name;
          window.backend.FileManager.HandleFile(
            JSON.stringify({
              id,
              type,
              ext: fExt(name),
              name: fName(name),
              data: (reader.result as string).split(",")[1]
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
      // eslint-disable-next-line
      this.$refs["fileInput"].value = "";
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
  min-height 80px
  /*max-height: calc(100vh / 2);*/
  overflow auto
  height calc(100vh - 344px)

table tr:nth-child(odd) p
  @apply bg-gray-800


td
  margin 0
  padding 0

td p
  @apply my-1 pl-3 py-2;
  min-height 40px


td p.cell-l
  border-top-left-radius 6px
  border-bottom-left-radius 6px


td p.cell-r
  border-top-right-radius 6px
  border-bottom-right-radius 6px
</style>
