<template>
  <div>
      Pods
  </div>
</template>

<script>

export default {
  name: "pods",
  components: {
  },
  data(){
    return{
        lastUpdate: new Date()
    }
  },
  mounted(){
    this.$upstream.subscribe("PODS");
    /*this.$bus.$on(this.$pods.UPDATED, _ => {
        this.lastUpdate = new Date()
    })*/
  },
  destroyed(){
      this.$upstream.unsubscribe("PODS");
  }
};
</script>

<style scoped lang="scss">
  
</style>
