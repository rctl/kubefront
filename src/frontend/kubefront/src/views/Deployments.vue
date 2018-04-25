<template>
  <div class="container view">
    <div class="">
      <DeploymentList extended></DeploymentList>
    </div>
  </div>
</template>

<script>
import DeploymentList from '../components/DeploymentList'
  
export default {
  name: "Nodes",
  components: {
    DeploymentList
  },
  data(){
    return{
    }
  },
  methods: {
    
  },
  mounted() {
    this.$upstream.subscribe("DEPLOYMENTS");
    this.$deployments.refresh()
    this.$upstream.subscribe("PODS");
    this.$pods.refresh()
  },
  updated(){

  },
  destroyed(){
    this.$upstream.unsubscribe("DEPLOYMENTS");
    this.$upstream.unsubscribe("PODS");
  }
};
</script>

<style scoped lang="scss">
  .view{
    padding: 30px;
  }
</style>
