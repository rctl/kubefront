<template>
  <div class="container view">
    <div class="card white view">
      <NodeList></NodeList>
    </div>
  </div>
</template>

<script>
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
    this.$upstream.subscribe("NODES");
    this.$nodes.refresh()
  },
  updated(){

  },
  destroyed(){
    this.$upstream.unsubscribe("NODES");
  }
};
</script>

<style scoped lang="scss">
  .view{
    padding: 30px;
  }
</style>
