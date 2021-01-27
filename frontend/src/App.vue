<template>
  <div id="app" class="bg-red-800">
    <div id="nav">
      <router-link to="/">
        <button>Home</button>
      </router-link>
      <router-link to="/about">
        <el-button>About</el-button>
      </router-link>
      <router-link to="/todo">
        <el-button>Todo</el-button>
      </router-link>
    </div>
    <router-view />
  </div>
</template>

<style lang="stylus"></style>
