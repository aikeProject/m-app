<template>
  <div class="mx-auto p-10">
    <input
      type="file"
      accept="image/jpeg, image/png, image/jpg"
      multiple
      @change="processFileInput"
    />
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
              class="px-4 py-3 title-font tracking-wider font-medium text-gray-900 text-sm bg-gray-200"
            >
              完成
            </th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(file, i) in files" :key="`${i}-${file.name}`">
            <td class="px-4 py-3">{{ file.filename }}</td>
            <td class="font-mono px-4 py-3">{{ file.size }}</td>
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
import { CFile } from "components/HelloWorld";
import Wails from "@wailsapp/runtime";

export default defineComponent({
  name: "HelloWorld",
  props: {
    msg: String
  },
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
    }
  },
  methods: {
    convert() {
      window.backend.FileManager.Convert()
        .then(result => {
          console.log(result);
        })
        .catch(error => console.error(error));
    },
    /**
     * 通过文件名查找文件
     * @param name
     */
    getFileByName(name: string) {
      if (this.files.length === 0) return;
      return this.files.find(f => f.name === name);
    },
    /**
     * 文件是否存在
     * @param name}
     */
    hasFile(name: string): boolean {
      if (this.files.length === 0) return false;
      return this.files.some(f => f.name === name);
    },
    /**
     * 检查文件类型
     * @param type
     */
    isValidType(type: string): boolean {
      const v = ["image/png", "image/jpeg", "image/jpeg"];
      return v.indexOf(type) >= 0;
    },
    processFileInput(e: InputEvent) {
      const target = e.target as HTMLInputElement;
      const files: File[] = target.files ? [].slice.apply(target.files) : [];
      files.forEach(f => {
        const name = fName(f.name);
        if (!this.isValidType(f.type) || this.hasFile(name)) return;
        this.processFile(f);
        this.files.push({
          filename: f.name,
          isConverted: false,
          name: fName(f.name),
          size: fSize(f.size)
        });
      });
    },
    processFile(file: File) {
      const reader = new FileReader();
      reader.onload = () => {
        if (reader.result) {
          const name = file.name;
          window.backend.FileManager.HandleFile(
            JSON.stringify({
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
    selectOutDir() {
      window.backend.FileManager.SetOutDir()
        .then(result => console.log(result))
        .catch(err => console.error(err));
    }
  },
  mounted() {
    Wails.Events.On("conversion:complete", name => {
      const f = this.getFileByName(name);
      if (!f) return;
      f.isConverted = true;
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
