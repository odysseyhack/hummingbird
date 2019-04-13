import {
  version
} from '../../package.json'

import {
  getters
} from './getters.js'

import {
  mutations
} from './mutations.js'

import {
  actions
} from './actions.js'

const state = {
  version: version,
  counter: 0,
  status: {
    system: 'connecting...',
    mqtt: false
  },
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
  strict: true
}