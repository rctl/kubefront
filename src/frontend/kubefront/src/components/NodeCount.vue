<template>
    <div class="wrapper">
        <div class="center">Node count</div>
        <div class="wrapper valign-wrapper">
            <h3 class="center content">{{count}}</h3>
        </div>
    </div>
</template>

<script>
export default {
  name: 'NodeCount',
  data() {
    return { 
       count: 0,
    }
  },
  methods: {
    
  },
  mounted() {
    this.$nodes.count().then(d => this.count = d)
    this.$bus.$on(this.$nodes.broadcasts.UPDATED, _ => {
      this.$nodes.count().then(d => this.count = d)
    })
  }
}
</script>

<style scoped lang="scss">
    .wrapper{
        width: 100%;
        height: 100%;
    }
    .content{
        margin: auto;
    }
</style>
