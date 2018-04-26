<template>
  <div class="container view">
    <div class="card white view">
      <ServiceList extended></ServiceList>
    </div>
  </div>
</template>

<script>
import ServiceList from '../components/ServiceList'
  
export default {
  name: "Services",
  components: {
    ServiceList
  },
  data(){
    return{
    }
  },
  methods: {
    
  },
  mounted() {
    this.$upstream.subscribe("SERVICES");
    this.$upstream.subscribe("PODS");
    this.$services.refresh()
    this.$pods.refresh()
  },
  updated(){

  },
  destroyed(){
    this.$upstream.unsubscribe("SERVICES");
    this.$upstream.unsubscribe("PODS");
  }
};
</script>

<style scoped lang="scss">
  .view{
    padding: 30px;
  }
</style>
