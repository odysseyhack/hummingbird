export const actions = {
    updateSystemStatus: (context, status) => {
        context.commit('SET_SYSTEM_STATUS', status)
    },
}