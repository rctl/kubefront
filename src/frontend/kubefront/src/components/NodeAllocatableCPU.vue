<template>
    <div class="wrapper">
        <div class="center">Allocatable CPU</div>
        <div class="wrapper valign-wrapper">
            <h3 class="center content">{{allocatable}}</h3>
        </div>
    </div>
</template>

<script>
export default {
  name: 'NodeAllocatableCPU',
  data() {
    return { 
       allocatable: 0,
    }
  },
  methods: {
    calculateAllocatableCPU(nodes){
        this.allocatable = 0
        nodes.forEach(n => {
            this.allocatable += parseInt(n.status.allocatable.cpu.string)
        });
    }
  },
  mounted() {
    this.$nodes.list().then(d => this.calculateAllocatableCPU(d))
    this.$bus.$on(this.$nodes.broadcasts.UPDATED, _ => {
      this.$nodes.list().then(d => this.calculateAllocatableCPU(d))
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
