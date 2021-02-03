<template>
  <div class="w-full p-10">
    <header class="border-b-2 border-gray-800 flex flex-wrap">
      <div class="w-5/12">
        <div
          class="bg-gray-800 border-2 border-dashed cursor-pointer drop-zone flex flex-col items-center justify-center py-10 ta-slow rounded-sm"
          :class="isDragging ? 'border-green-default' : 'border-gray-400'"
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
          <p class="mt-6 text-gray-200">拖拽文件到此处或选择文件</p>
        </div>
      </div>
      <div class="pl-6 w-7/12">
        <transition name="fade-fast" mode="out-in">
          <div
            v-if="!stats.time"
            key="intro"
            class="flex h-full items-center justify-center"
          >
            <h2 class="leading-none text-4xl text-center text-green">
              Add image files<br />to get started
            </h2>
          </div>
          <div v-else key="stats" class="flex flex-wrap items-end h-full">
            <div class="px-3 w-5/12">
              <h2
                class="font-bold leading-none text-5xl text-green-default tracking-tight"
              >
                {{ getPrettySize(stats.savings) }}
              </h2>
              <p class="font-medium text-gray-300 tracking-wider uppercase">
                Saved
              </p>
            </div>
            <div class="px-3 w-3/12">
              <p class="font-bold text-2xl">{{ stats.count }}</p>
              <p
                class="font-medium text-gray-300 text-sm tracking-wider uppercase"
              >
                {{ stats.count > 1 ? "Images" : "Image" }}
              </p>
              <p class="font-bold text-2xl">
                {{ getPrettyTime(stats.time)[0] }}
              </p>
              <p
                class="font-medium text-gray-300 text-sm tracking-wider uppercase"
              >
                {{ getPrettyTime(stats.time)[1] }}
              </p>
            </div>
            <div class="px-3 w-4/12">
              <p class="font-bold text-2xl">
                {{ getPrettySize(totalStats.byteCount) }}
              </p>
              <p
                class="font-medium text-gray-300 text-sm tracking-wider uppercase"
              >
                All time Savings
              </p>
              <p class="font-bold text-2xl">{{ totalStats.imageCount }}</p>
              <p
                class="font-medium text-gray-300 text-sm tracking-wider uppercase"
              >
                All time Images
              </p>
            </div>
          </div>
        </transition>
      </div>
      <footer class="w-full">
        <section class="flex justify-between py-6 w-5/12">
          <button
            class="btn focus:outline-none ta-slow"
            :class="
              canConvert
                ? 'border-purple hover:bg-purple-default hover:text-gray-900 text-gray-200'
                : 'btn--disabled'
            "
            @click="convert"
            :disabled="!canConvert"
          >
            转换格式
          </button>
          <button
            class="btn focus:outline-none ta-slow"
            :class="
              files.length > 0
                ? 'border-gray-400 hover:bg-gray-400 hover:text-gray-900'
                : 'btn--disabled'
            "
            @click="clear"
            :disabled="files.length === 0"
          >
            清空
          </button>
        </section>
      </footer>
    </header>
    <input
      type="file"
      accept="image/jpeg, image/jpg, image/png, image/webp"
      multiple
      class="hidden"
      @input="handleFileInput"
      ref="fileInput"
    />
    <transition name="fade" mode="out-in">
      <div v-if="files.length > 0" class="table-wrapper">
        <table class="table-auto w-full text-left whitespace-nowrap">
          <thead>
            <tr>
              <th
                class="font-medium p-3 text-gray-400 text-left text-sm tracking-wider uppercase"
              >
                文件名
              </th>
              <th
                class="font-medium p-3 text-gray-400 text-left text-sm tracking-wider uppercase"
              >
                大小
              </th>
              <th
                class="font-medium p-3 text-gray-400 text-left text-sm tracking-wider uppercase"
              >
                转换后大小
              </th>
              <th
                class="font-medium p-3 text-gray-400 text-left text-sm tracking-wider uppercase"
              >
                转化率
              </th>
              <!--            <th-->
              <!--              class="font-medium pl-3 pt-6 text-gray-400 text-left text-sm tracking-wider uppercase"-->
              <!--            >-->
              <!--              结果-->
              <!--            </th>-->
              <th
                class="font-medium text-center pl-3 pt-6 text-gray-400 text-left text-sm tracking-wider uppercase"
              >
                状态
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
              <!--            <td @click="openFile(file)">-->
              <!--              {{ file.convertedPath }}-->
              <!--            </td>-->
              <td>
                <p
                  v-if="file.isConverted"
                  class="cell-r flex items-center justify-center"
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
                <p v-else class="cell-r"></p>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </transition>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { fExt, fName, fSize } from "lib/filw";
import { CFile } from "components/Editor";
import Wails from "@wailsapp/runtime";
import { prettyTime } from "lib/time";
import { EventBus } from "../lib/event-bus";

export default defineComponent({
  name: "Editor",
  data() {
    return {
      files: [] as CFile[],
      isDragging: false,
      stats: {
        count: 0,
        savings: 0,
        time: 0
      }
    };
  },
  computed: {
    canConvert(): boolean {
      if (this.files.length === 0) return false;
      return this.files.some(f => {
        return !f.isConverted;
      });
    },
    totalStats() {
      return this.$store.getters.stats;
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
     * handleFileInput handles the file submission via form input.
     * @param {Event} e
     */
    handleFileInput(e: InputEvent) {
      const target = e.target as HTMLInputElement;
      target.files && this.processFileInput(target.files);
    },
    /**
     * 时间戳格式化
     */
    getPrettyTime(ms: number) {
      return prettyTime(ms);
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
    processFileInput(files: FileList) {
      const fs: File[] = files ? [].slice.apply(files) : [];
      fs.forEach(f => {
        const name = fName(f.name);
        const size = f.size;
        const ext = fExt(f.name);
        const type = this.getFileType(f.type, ext);
        const id = this.createFileId(name, size);
        if (!type) {
          EventBus.emit("notify", {
            msg:
              "File type not supported. Valid file types include JPG, PNG, and WebP.",
            type: "warn"
          });
          return;
        }
        if (this.hasFile(name)) {
          EventBus.emit("notify", {
            msg: "Image file has already been added to the file list.",
            type: "warn"
          });
          return;
        }
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
      // eslint-disable-next-line
      // @ts-ignore
      this.$refs["fileInput"].value = null;
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
              size: file.size,
              data: (reader.result as string).split(",")[1]
            })
          );
        }
      };
      reader.readAsDataURL(file);
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
      // @ts-ignore
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

    Wails.Events.On("conversion:stat", e => {
      const c = e.count;
      const t = e.time;
      this.stats.count += c;
      this.stats.time += t;
      this.stats.savings += e.savings;
      this.$store.dispatch("getStats");
      EventBus.emit("notify", {
        msg: `Optimized ${c} ${c > 1 ? "images" : "image"} in ${
          prettyTime(t)[0]
        } ${prettyTime(t)[1].toLowerCase()}.`
      });
    });

    // eslint-disable-next-line
    // @ts-ignore
    const dz = this.$refs["dropZone"] as HTMLDivElement;

    dz.addEventListener(
      "click",
      () => {
        // eslint-disable-next-line
        // @ts-ignore
        this.$refs["fileInput"].click();
      },
      false
    );
    dz.addEventListener(
      "dragenter",
      e => {
        e.stopPropagation();
        e.preventDefault();
      },
      false
    );
    dz.addEventListener(
      "dragover",
      e => {
        e.stopPropagation();
        e.preventDefault();
        if (this.isDragging) return;
        this.isDragging = true;
      },
      false
    );
    dz.addEventListener(
      "dragend",
      e => {
        e.stopPropagation();
        e.preventDefault();
        this.isDragging = false;
      },
      false
    );
    dz.addEventListener(
      "dragleave",
      e => {
        e.stopPropagation();
        e.preventDefault();
        this.isDragging = false;
      },
      false
    );
    dz.addEventListener(
      "drop",
      e => {
        e.stopPropagation();
        e.preventDefault();
        const dt = e.dataTransfer;
        if (dt) {
          const f = dt.files;
          this.isDragging = false;
          this.processFileInput(f);
        }
      },
      false
    );
  }
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="stylus">
.table-wrapper
  min-height 80px
  height calc(100vh - 344px)
  max-height calc(100vh / 2)
  overflow auto

table thead tr
  @apply bg-gray-800

table tbody td
  @apply px-3

table tbody tr:nth-child(2n)
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

.drop-zone > svg path
  transition fill 600ms cubic-bezier(.07, .95, 0, 1)


.drop-zone:hover > svg path
  fill #07fdbc
</style>
