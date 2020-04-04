<template>
  <div class="article-item">
      <p name="author">@{{article.author}}</p>
      <p>{{article.pdate}}</p>
      <p name="title">{{article.title}}</p>
      <p name="content">{{article.content}}</p>
  </div>
</template>

<script>
export default {
  name: 'ArticleItem',
  props: ["article"],
  methods:{
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
div[class = "article-item"]{
  background: rgb(92, 92, 92);
  height: 30%;
  width: 60%;
  color: #909C67;
  border-radius: 24px;
  margin: auto;
  border-left: 8px #DAA520 solid;
}
p[name="author"]{
  padding-top: 10px;
  padding-left: 10px;
  text-align: start;
}
p[name="content"]{
  padding-bottom: 2%;
  text-align: center;
}
p[name="title"]{
  padding-top: 0;
  padding-bottom: 0;
  padding-left: 0;
  padding-right: 0;
  text-align: center;
  font-size: 17px;
  font-weight: 500;
}
</style>
