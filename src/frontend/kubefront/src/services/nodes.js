import bus from "../bus"
import api from "../api"
import router from '../router'
let service = {

    //Properties
    lastUpdate: new Date(),
    nodes: [],

    //Broadcasts
    broadcasts: {
        UPDATED: "NODE_UPDATED",
        ADDED: "NODE_ADDED"
    },

    //Methods
    /**
     * Force state sync
     */
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
    /**
     * Returns a Promise yeilding the count of nodes in existance
     */
    count() {
        return new Promise((resolve, reject) => {
           resolve(this.nodes.length)
        })
    },
    /**
     * Returns a Promise yeilding a list with nodes in existance
     */
    list() {
        return new Promise((resolve, reject) => {
           resolve(this.nodes)
        })
    }
}

//Handle to events informing that a node has been changed
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