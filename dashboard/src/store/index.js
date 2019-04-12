import {
  version
} from '../../package.json'

const state = {
  version: version,
  counter: 0,
  status: {
    system: 'connecting...',
    mqtt: false
  },
}

// mutations
const mutations = {
  SET_SYSTEM_STATUS: (state, status) => {
    state.status.system = status
  }
}

// actions
const actions = {
  updateSystemStatus: (context, status) => {
    context.commit('SET_SYSTEM_STATUS', status)
  },
}

const getters = {
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

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
  strict: true
}