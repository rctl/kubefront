<template>
    <div class="wrapper">
        <div class="center">All Pods <span class="right"><a href="#" class="grey-text tooltipped" :data-tooltip="lastUpdate|moment('from', 'now')"><i class="material-icons">info</i></a></span></div>
        <div class="wrapper scroll">
            <div class="collection-item" v-for="(l, _) in pods" :key="_">
                <b>{{l.namespace}}</b>
                <div class="group" v-for="(g, _) in l.groups" :key="_">
                    <ul class="collection full-width">
                        <li class="collection-item" v-for="(p, _) in g.pods" :key="_">
                            <a class='dropdown-trigger left grey-text' href='#' v-if="p.editable" :data-target='p.metadata.name'>
                                <i class="material-icons">more_vert</i>
                            </a>
                            <div v-if="!p.editable" class="preloader-wrapper small active left tooltipped" data-tooltip="There is a job running on this pod" style="height:20px; width: 20px; margin-right:10px;">
                                <div class="spinner-layer spinner-green-only">
                                <div class="circle-clipper left">
                                    <div class="circle"></div>
                                </div><div class="gap-patch">
                                    <div class="circle"></div>
                                </div><div class="circle-clipper right">
                                    <div class="circle"></div>
                                </div>
                                </div>
                            </div>
                            {{p.metadata.name}}
                            <a href="#" class="right tooltipped"
                                :data-tooltip="describe(p)">
                                <i v-if="p.warnings.length == 0 && p.errors.length == 0" class="material-icons green-text">brightness_1</i>
                                <i v-if="p.warnings.length > 0 && p.errors.length == 0" class="material-icons yellow-text">brightness_1</i>
                                <i v-if="p.errors.length > 0" class="material-icons red-text">brightness_1</i>
                            </a>
                            <ul :id='p.metadata.name' class='dropdown-content dropper'>
                                <li><a href="#!" class="red-text" @click="deletePod(p)"><i class="material-icons">delete</i> Delete Pod</a></li>
                            </ul>
                        </li>
                    </ul>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
  name: 'PodList',
  data() {
    return { 
       pods: [],
       lastUpdate: new Date()
    }
  },
  methods: {
      deletePod(p){
          p.editable = false;
          this.$pods.delete(p.metadata.namespace, p.metadata.name).then(this.refresh).catch(() => {
              p.editable = true;
              M.toast({html: 'Pod could not be deleted'})
          })
      },
      describe(p){
          if(p.errors.length > 0){
              return p.errors.join(', ') + "."
          }
          if(p.warnings.length > 0){
              return p.warnings.join(', ') + "."
          }
          return "Pod is running"
      },
      refresh(){
        this.$pods.list().then(d =>{
            this.pods = []
            let namespaces = {}
            d.forEach(x => {
                if(namespaces[x.metadata.namespace] === undefined){
                    namespaces[x.metadata.namespace] = {}
                }
                if(namespaces[x.metadata.namespace][x.metadata.generateName] === undefined){
                    namespaces[x.metadata.namespace][x.metadata.generateName] = []
                }
                namespaces[x.metadata.namespace][x.metadata.generateName].push(x);
                x.warnings = []
                x.errors = []
                this.$workers.getByEntity(x.metadata.name).then(w => {
                    x.editable = w.length == 0;
                })
                if(x.status.phase == "Pending"){
                    x.warnings.push("changes to this pod are pending")
                }
                if(x.status.containerStatuses){
                    x.status.containerStatuses.forEach(s => {
                        if(s.state.terminated){
                            x.errors.push("pod has been terminated")
                        }
                    })
                }
            })
            Object.keys(namespaces).forEach(x => {
                this.pods.push({
                    namespace: x,
                    groups: Object.keys(namespaces[x]).map(i =>{
                        return {
                            group: i,
                            pods: namespaces[x][i]
                        }
                    })
                })
            });
        }).then(_ => {
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
    this.$el.querySelectorAll('.dropdown-trigger').forEach(e => {
        if(!e.getAttribute("init")){
            e.setAttribute("init", true)
            M.Dropdown.init(e, {
                constrainWidth: false,
            });
        }
    });
  },
  mounted() {
    this.refresh()
    this.$bus.$on(this.$pods.broadcasts.UPDATED, _ => {
      this.refresh()
    })
    this.$bus.$on("JOB_STARTED", (id, data) => {
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
        padding-bottom: 15px;
    }
    .scroll{
        overflow:auto;
    }
    .content{
        margin: auto;
    }
    .full-width{
        width: 100%;
        overflow: visible;
    }
    .indicator{
        margin-top: 8px;
    }
    .group{
        padding-right: 20px;
        overflow: visible;
    }
    .dropper{
        overflow: visible;
    }
</style>

