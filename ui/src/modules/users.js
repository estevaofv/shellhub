// import Vue from 'vue';
// import * as apiUser from '@/api/users';
import postUser from '@/api/users';

export default {
  namespaced: true,

  state: {
    user: '',
  },

  getters: {
    user: (state) => state.user,
  },

  mutations: {
    // setUser(user) {
    //   Vue.set(state, 'user', user);
    // },
  },

  actions: {
    setUser(context, user) {
      console.log(user);
      postUser(user);
    },
  },
};
