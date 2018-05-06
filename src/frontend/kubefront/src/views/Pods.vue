<template>
  <div class="container view">
    <div class="card white view">
      <PodList extended></PodList>
    </div>
  </div>
</template>

<script>
//Components
import PodList from '../components/PodList'
  
export default {
  name: "Nodes",
  components: {
    PodList
  },
  data(){
    return{
    }
  },
  methods: {
    
  },
  mounted() {
    //Subcribe to topics needed for this view and fetch initial data
    this.$upstream.subscribe("PODS");
    this.$pods.refresh()
  },
  updated(){

  },
  destroyed(){
    //Unsubscribe to topics
    this.$upstream.unsubscribe("PODS");
  }
};
</script>

<style scoped lang="scss">
  .view{
    padding: 30px;
  }
</style>
