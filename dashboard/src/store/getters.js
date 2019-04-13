export const getters = {
    getSystemStatus: state => {
        return state.status.system
    },
    getMQTTStatus: state => {
        return state.status.mqtt
    },
    getStatus: state => {
        return state.status
    },
    getVersion: state => {
        return state.version
    }
}