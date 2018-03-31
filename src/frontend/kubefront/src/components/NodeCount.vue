<template>
    <div class="wrapper">
        <div class="center">Node count <span class="right"><a href="#" class="grey-text tooltipped" :data-tooltip="lastUpdate|moment('from', 'now')"><i class="material-icons">info</i></a></span></div>
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
       lastUpdate: new Date(),
    }
  },
  methods: {
    
  },
  mounted() {
    this.$nodes.count().then(d => this.count = d)
    this.$bus.$on(this.$nodes.broadcasts.ADDED, _ => {
      this.$nodes.count().then(d => this.count = d)
      this.lastUpdate = new Date()
    })
    M.Tooltip.init(this.$el.querySelector('.tooltipped'), {});
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
