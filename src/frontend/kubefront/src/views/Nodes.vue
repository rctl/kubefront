<template>
  <div class="container view">
    <div class="card white view">
      <NodeList></NodeList>
    </div>
  </div>
</template>

<script>
//Components
import NodeList from '../components/NodeList'
  
export default {
  name: "Nodes",
  components: {
    NodeList
  },
  data(){
    return{
    }
  },
  methods: {
    
  },
  mounted() {
    //Subcribe to topics needed for this view and fetch initial data
    this.$upstream.subscribe("NODES");
    this.$nodes.refresh()
  },
  updated(){

  },
  destroyed(){
    //Unsubscribe to topics
    this.$upstream.unsubscribe("NODES");
  }
};
</script>

<style scoped lang="scss">
  .view{
    padding: 30px;
  }
</style>
