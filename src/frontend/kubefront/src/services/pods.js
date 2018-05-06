import bus from "../bus"
import api from "../api"
import router from '../router'

let service = {

    //Properties
    lastUpdate: new Date(),
    pods: [],

    //Broadcasts
    broadcasts: {
        UPDATED: "POD_UPDATED",
        ADDED: "POD_ADDED"
    },

    //Methods
    /**
     * Force a state sync
     */
    refresh() {
        return new Promise((resolve, reject) => {
            api.get("/pods/")
            .then(r => {
                this.lastUpdate = new Date()
                this.pods = r.data
                bus.$emit(this.broadcasts.ADDED)
                bus.$emit(this.broadcasts.UPDATED)
            }).catch(reject);
        })
    },
    /**
     * Returns a Promise yeilding the number of pods in existance
     */
    count() {
        return new Promise((resolve, reject) => {
           resolve(this.pods.length)
        })
    },
    /**
     * Returns a Promise yeilding a list with pods in existance
     */
    list() {
        return new Promise((resolve, reject) => {
           resolve(this.pods)
        })
    },
    /**
     * Start a deletion job. Returns a request Promise
     * @param  {string} namespace - The namespace of the resource to be deleted
     * @param  {string} name - The name of the resource to be deleted
     */
    delete(namespace, name) {
        return api.delete("/pods/" + namespace + "/" + name)
    }
}

//Handle to events informing that a pod has been changed
bus.$on("POD_CHANGED", (entityID, data) => {
    if(data.status.containerStatuses){
        if(data.status.containerStatuses.every(x => x.state.terminated)){
            service.pods = service.pods.filter(x => x.metadata.name != entityID)
            bus.$emit(service.broadcasts.ADDED)
            return
        }
    }
    let i = service.pods.findIndex(x => x.metadata.name == entityID)
    data.lastUpdate = new Date()
    if(i != -1){
        service.pods[i] = data
        bus.$emit(service.broadcasts.UPDATED)
    }else{
        service.pods.push(data)
        bus.$emit(service.broadcasts.ADDED)
    }
})

//Export
export default service