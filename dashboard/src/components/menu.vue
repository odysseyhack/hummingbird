<template>
  <div class="button-wrapper">
    <a class="mano-button" @click="logOut" v-bind:class="{ active: isLoggedIn }">Logout</a>
  </div>
</template>

<script>
export default {
  name: "headerMenu",
  components: {},
  data() {
    return {};
  },
  computed: {},
  methods: {
    logOut: function() {
      console.log(" -- Logout triggered");
    }
  }
};
</script>

<style scoped>
.button-wrapper {
  display: block;
  position: absolute;
  background-color: transparent;
  padding: 0;
  height: 80px;
  left: 50%;
  opacity: 0;
  -webkit-animation: fadein 0.6s forwards;
  animation: fadein 0.6s forwards;
  animation-delay: 0.3s;
  font-size: 15px;
  text-decoration: none !important;
}
@keyframes fadein {
  0% {
    opacity: 0;
    transform: translateX(-50%) translateY(110%);
  }
  100% {
    opacity: 1;
    transform: translateX(-50%) translateY(110%);
  }
}
.active {
  opacity: 0 !important;
}
</style>
