<template>
  <div class="container view">
    <div class="">
      <DeploymentList extended></DeploymentList>
    </div>
  </div>
</template>

<script>
//Components
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
    //Subcribe to topics needed for this view and fetch initial data
    this.$upstream.subscribe("DEPLOYMENTS");
    this.$upstream.subscribe("PODS");
    this.$deployments.refresh()
    this.$pods.refresh()
  },
  updated(){

  },
  destroyed(){
    //Unsubscribe to topics
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
