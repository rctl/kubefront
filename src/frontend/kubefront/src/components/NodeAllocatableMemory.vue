<template>
    <div class="wrapper">
        <div class="center">Allocatable Memory<span class="right"><a href="#" class="grey-text tooltipped" :data-tooltip="lastUpdate|moment('from', 'now')"><i class="material-icons">info</i></a></span></div>
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
       lastUpdate: new Date(),
    }
  },
  methods: {
    /**
     * Calculate allocatable memory and set component state
     * @param {Array} nodes - Array containing nodes
     */
    calculateAllocatableMemory(nodes){
        let sum = 0
        nodes.forEach(n => {
            //Make all nodes allocatable memory in common format
            sum += this.parseDataSize(n.status.allocatable.memory.string)
        });
        this.allocatable = Math.round(sum / (1024*1024)) + " GiB";
    },
    /**
     * Convert size string to integer
     * @param {string} size - Size described in string
     */
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
    //Fetch data from service
    this.$nodes.list().then(d => this.calculateAllocatableMemory(d))
    //Listen for data updates
    this.$bus.$on(this.$nodes.broadcasts.UPDATED, _ => {
      this.$nodes.list().then(d => this.calculateAllocatableMemory(d))
      this.lastUpdate = new Date()
    })
    //Initialize lastUpdated tooltip
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
