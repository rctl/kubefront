<template>
    <div class="wrapper">
        <div class="center">All Nodes <span class="right"><a href="#" class="grey-text tooltipped" :data-tooltip="lastUpdate|moment('from', 'now')"><i class="material-icons">info</i></a></span></div>
        <div class="wrapper scroll">
            <ul class="collection full-width">
                <li class="collection-item" v-for="(l, _) in nodes" :key="_" >
                    {{l.metadata.name}}
                    <a href="#" class="right" v-if="l.warning && l.warning.length > 0"><i 
                    class="right material-icons red-text tooltipped"
                    :data-tooltip="l.warning"
                    >warning</i></a>
                    <a href="#" class="right" v-if="!l.warning"><i 
                    class="right material-icons green-text tooltipped"
                    data-tooltip="Node is ready"
                    >done</i></a>
                    <a href="#" class="right" v-if="l.master"><i 
                    class="right material-icons grey-text tooltipped"
                    data-tooltip="This is a master node"
                    >grade</i></a>
                </li>
            </ul>
        </div>
    </div>
</template>

<script>
export default {
  name: 'NodeList',
  data() {
    return { 
       nodes: [],
       lastUpdate: new Date()
    }
  },
  methods: {
      refresh(){
        this.$nodes.list().then(d => this.nodes = d).then(_ => {
            this.nodes.forEach(x => {
                x.master = Object.keys(x.metadata.labels).some(i => i == "node-role.kubernetes.io/master")
                x.status.conditions.forEach(c => {
                    switch (c.type) {
                        case "MemoryPressure":
                            if(c.status == "True"){
                                x.warning = "Nodes memory is under pressure!"
                            }
                            break;
                        case "OutOfDisk":
                            if(c.status == "True"){
                                x.warning = "Node is out of storage space!"
                            }
                            break;
                        case "DiskPressure":
                            if(c.status == "True"){
                                x.warning = "Node disk is under pressure!"
                            }
                            break;
                        case "Ready":
                            let heartbeat = (new Date() - new Date(c.lastHeartbeatTime))/1000;
                            if(c.status == "True" && heartbeat < 30){
                                x.warning = ""
                            }else{
                                x.warning = "Node is not ready!";
                            }
                            break;
                        default:
                            x.warning = "Node status is unknown!";
                            break;
                    }
                })
            });
            this.lastUpdate = new Date()
        })
      }
  },
  updated() {
    this.$el.querySelectorAll('.tooltipped').forEach(e => {
        if(!e.getAttribute("init")){
            e.setAttribute("init", true)
            M.Tooltip.init(e, {});
        }
    });
  },
  mounted() {
    this.refresh()
    this.$bus.$on(this.$nodes.broadcasts.UPDATED, _ => {
      this.refresh()
    })
    this.$el.querySelectorAll('.tooltipped').forEach(e => {
        M.Tooltip.init(e, {});
    });
  }
}
</script>

<style scoped lang="scss">
    .wrapper{
        width: 100%;
        height: 100%;
    }
    .scroll{
        overflow:auto;
    }
    .content{
        margin: auto;
    }
    .full-width{
        width: 100%;
    }
</style>
