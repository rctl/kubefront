import bus from "../bus"
import api from "../api"

let service = {

    //Properties
    jobs: [], //Stores a list of current jobs

    //Broadcasts
    STARTED: "STARTED",
    COMPLETED: "COMPLETED", 
    FAILED: "FAILED",

    /**
     * Returns all current jobs
     */
    get(){
        return new Promise((resolve, reject) => {
            resolve(this.jobs)
        })
    },
    /**
     * Returns all jobs by the id of the entity it is attached to (legacy)
     * @param  {string} id - The entity id
     */
    getByEntity(id){
        return new Promise((resolve, reject) => {
            resolve(this.jobs.filter(x => x.entity == id))
        })
    },
    /**
     * Returns all jobs by the fully qualified id of the entity it is attached to
     * @param  {} kind - The entity kind
     * @param  {} namespace - The entity namespace
     * @param  {} name - The entity name
     */
    get(kind, namespace, name){
        return new Promise((resolve, reject) => {
            resolve(this.jobs.filter(x => x.kind == kind && x.namespace == namespace && x.name == name))
        })
    },
    /**
     * Force a state sync with server
     */
    update(){
        return api.get("/workers/")
        .then(r => {
            r.data.forEach(x => {
                if(x.entity.includes("/")){
                    this.jobs.push({
                        id: x.id,
                        kind: x.entity.split("/")[0],
                        namespace: x.entity.split("/")[1],
                        name: x.entity.split("/")[2],
                        entity: x.entity,
                    })
                }else{
                    service.jobs.push({
                        id: x.id,
                        entity: x.entity,
                    })
                }
            });
        });
    }

}

//Event is triggered when a job is started on the server
bus.$on("JOB_STARTED", (id, data) => {
    if(data.includes("/")){
        //V2 (Fully Qualified ID)
        service.jobs.push({
            id: id,
            kind: data.split("/")[0],
            namespace: data.split("/")[1],
            name: data.split("/")[2],
            entity: data,
        })
    }else{
        service.jobs.push({
            id: id,
            entity: data,
        })
    }
    //Pass along the message that a job has started
    bus.$emit(service.STARTED, data)
})

//Event is triggered when a job is completed on the server
bus.$on("JOB_COMPLETED", (id, data) => {
    service.jobs = service.jobs.filter(x => x.id != id)
    //Pass along the message that a job has completed
    bus.$emit(service.COMPLETED, data)
})

//Event is triggered when a job fails on the server
bus.$on("JOB_FAILED", (id, data) => {
    service.jobs = service.jobs.filter(x => x.id != id)
    //Pass along the message that a job has failed
    bus.$emit(service.FAILED, data)
})

//Export
export default service