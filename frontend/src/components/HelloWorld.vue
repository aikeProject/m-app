<template>
  <div class="hello container">
    <h1>{{ msg }}</h1>
    <el-button @click="convert">Run Covert</el-button>
    <input type="file" multiple @change="processFileInput" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { fExt, fName } from "lib/filw";

export default defineComponent({
  name: "HelloWorld",
  props: {
    msg: String
  },
  methods: {
    convert() {
      window.backend.FileManager.Convert()
        .then(result => {
          console.log(result);
        })
        .catch(error => console.error(error));
    },
    processFileInput(e: InputEvent) {
      const target = e.target as HTMLInputElement;
      const files = target.files ? [].slice.apply(target.files) : [];
      files.forEach(f => this.processFile(f));
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
    }
  }
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="stylus">
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
