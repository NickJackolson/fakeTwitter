<template>
  <div class="articles">
      <div v-bind:key="article.id" v-for="article in articles">
          <ArticleItem v-bind:article="article"/> 
      </div>
  </div>
</template>

<script>

import ArticleItem from "./ArticleItem.vue";
export default {
  name: 'Articles',
  components:{
      ArticleItem
  },
  props: ["articles"]
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
