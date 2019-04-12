<template>
  <div id="app-home">
    <div class="logo">
      <img src="../assets/logo.png">
    </div>
    <div class="settings">
      <a v-on:click="goToSettings">
        <ons-icon style="color: whitesmoke;" size="25px" icon="fa-bars"></ons-icon>
      </a>
    </div>
  </div>
</template>

<script>
export default {
  name: "Home",
  components: {},
  data() {
    return {};
  },
  computed: {},
  mounted: function() {},
  mqtt: {},
  methods: {
    goToSettings: function() {
      console.log(" -- Go To Settings");
      this.$router.push("/settings");
    }
  }
};
</script>

<style>
.loader .vue-simple-progress,
.loader .vue-simple-progress-bar {
  border-radius: 5px !important;
}
</style>

<style scoped>
.loader .vue-simple-progress,
.loader .vue-simple-progress-bar {
  border-radius: 5px !important;
}

#app-home {
  position: relative;
  top: 0;
  text-align: center;
}

.loader {
  position: relative;
  top: 170px;
  width: 190px;
  margin: 0 auto;
}

.loader-status {
  position: relative;
  top: 160px;
  font-weight: bold;
  color: white;
}

.menu {
  position: relative;
  top: 120px;
  display: block;
}

.logo {
  position: relative;
  top: 120px;
  display: block;
  opacity: 0;
  -webkit-animation: fadein 0.7s forwards;
  animation: fadein 0.7s forwards;
  animation-delay: 0s;
}

@keyframes fadein {
  0% {
    top: 115px;
    opacity: 0;
  }
  100% {
    top: 120px;
    opacity: 1;
  }
}

@keyframes fadeinsimple {
  0% {
    opacity: 0;
  }
  100% {
    opacity: 1;
  }
}

.settings {
  position: absolute;
  top: 20px;
  right: 30px;
  display: block;
  -webkit-animation: fadeinsimple 0.4s forwards;
  animation: fadeinsimple 0.4s forwards;
  animation-delay: 0s;
}

.settings img {
  width: 35px;
  height: auto;
}
</style>
