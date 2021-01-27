<template>
  <div class="hello container">
    <h1>{{ msg }}</h1>
    <input type="file" @change="processFile" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "HelloWorld",
  props: {
    msg: String
  },
  methods: {
    processFile(e: InputEvent) {
      const target = e.target as HTMLInputElement;
      const file = target.files ? target.files[0] : null;
      if (file) {
        const reader = new FileReader();
        reader.onload = () => {
          // base64
          console.log(reader.result);
          if (reader.result) {
            window.backend.HandleFile(
              JSON.stringify({
                data: (reader.result as string).split(",")[1],
                name: file.name,
                type: file.type
              })
            );
          }
        };
        // reader.readAsBinaryString(file);
        reader.readAsDataURL(file);
        // reader.readAsArrayBuffer(file);
      }
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
