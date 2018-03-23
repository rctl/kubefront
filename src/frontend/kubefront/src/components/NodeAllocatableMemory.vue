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
  name: 'NodeAllocatableMemory',
  data() {
    return { 
       allocatable: "0Gi",
    }
  },
  methods: {
    calculateAllocatableMemory(nodes){
        let sum = 0
        nodes.forEach(n => {
            sum += this.parseDataSize(n.status.allocatable.memory.string)
        });
        this.allocatable = Math.round(sum / (1024*1024)) + " GiB";
    },
    parseDataSize(size){
        if(size.includes("Ki")){
            return parseInt(size.replace("Ki", ""));
        }
        if(size.includes("Mi")){
            return parseInt(size.replace("Mi", ""))*1024;
        }
        if(size.includes("Gi")){
            return parseInt(size.replace("Gi", ""))*1024*1024;
        }
        if(size.includes("Ti")){
            return parseInt(size.replace("Gi", ""))*1024*1024*1024;
        }
        if(size.includes("Pi")){
            return parseInt(size.replace("Gi", ""))*1024*1024*1024*1024;
        }
        return parseInt(size)
    }
  },
  mounted() {
    this.$nodes.list().then(d => this.calculateAllocatableMemory(d))
    this.$bus.$on(this.$nodes.broadcasts.UPDATED, _ => {
      this.$nodes.list().then(d => this.calculateAllocatableMemory(d))
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
