import bus from "../bus"
import api from "../api"
import router from '../router'
let service = {

    //Properties
    lastUpdate: new Date(),
    nodes: [],

    //Broadcasts
    broadcasts: {
        UPDATED: "UPDATED",
        ADDED: "ADDED"
    },

    //Methods
    refresh() {
        return new Promise((resolve, reject) => {
            api.get("/nodes/")
            .then(r => {
                this.lastUpdate = new Date()
                this.nodes = r.data
                bus.$emit(this.broadcasts.ADDED)
                bus.$emit(this.broadcasts.UPDATED)
            }).catch(reject);
        })
    },
    count() {
        return new Promise((resolve, reject) => {
           resolve(this.nodes.length)
        })
    },
    list() {
        return new Promise((resolve, reject) => {
           resolve(this.nodes)
        })
    }
}

//Init and listeners
bus.$on("NODE_CHANGED", (entityID, data) => {
    let i = service.nodes.findIndex(x => x.metadata.name == entityID)
    data.lastUpdate = new Date()
    if(i != -1){
        service.nodes[i] = data
        bus.$emit(service.broadcasts.UPDATED)
    }else{
        service.nodes.push(data)
        bus.$emit(service.broadcasts.ADDED)
    }
})

//Export
export default service