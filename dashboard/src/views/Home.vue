<template>
  <div id="app-home">
    <div class="logo">
      <img src="../assets/logo.png">
    </div>
    <div class="settings">
      <span class="title">Hummingbird</span>
      <a v-on:click="goToSettings">
        <img src="../assets/settings.svg">
      </a>
      <a v-on:click="logOut">
        <img src="../assets/logout.svg">
      </a>
    </div>
    <div class="sidebar"></div>
    <div class="map">
      <GmapMap
        :center="{lat:10, lng:10}"
        :zoom="7"
        map-type-id="terrain"
        style="width: 100%; height: 100%"
        :options="{
          zoomControl: false,
          mapTypeControl: false,
          scaleControl: false,
          streetViewControl: false,
          rotateControl: false,
          fullscreenControl: false
        }"
      >
        <GmapMarker
          :key="index"
          v-for="(m, index) in markers"
          :position="m.position"
          :clickable="true"
          :draggable="false"
          @click="center=m.position"
        />
      </GmapMap>
    </div>
  </div>
</template>

<script>
import GmapCustomMarker from "vue2-gmap-custom-marker";

export default {
  name: "Home",
  components: {
    GmapCustomMarker
  },
  data() {
    return {
      markers: {}
    };
  },
  computed: {},
  mounted: function() {},
  mqtt: {},
  methods: {
    goToSettings: function() {
      console.log(" -- Go To Settings");
      this.$router.push("/settings");
    },
    logOut: function() {
      console.log(" -- Log out?");
      this.$ons.notification
        .confirm({
          title: "Confirm",
          message: "Are you sure you want to exit Hummingbird?"
        })
        .then(arg => {
          if (arg == 1) {
            // yes
            console.log(" -- Shutting down");
          }
        });
    }
  }
};
</script>

<style>
.loader .vue-simple-progress,
.loader .vue-simple-progress-bar {
  border-radius: 5px !important;
}
@font-face {
  font-family: "Skycoin";
  font-style: normal;
  font-weight: 700;
  src: url("../assets/fonts/skycoin/Skycoin-Bold.woff2") format("woff2"),
    url("../assets/fonts/skycoin/Skycoin-Bold.woff") format("woff");
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

.sidebar {
  display: block;
  position: absolute;
  float: left;
  top: 50px;
  width: 120px;
  height: 100vh;
  background-color: whitesmoke;
}

.map {
  display: block;
  position: absolute;
  right: 0;
  top: 50px;
  width: calc(100vw - 120px);
  height: 100vh;
  background-color: black;
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
  float: left;
  display: block;
  opacity: 0;
  left: 20px;
  margin-top: 10px;
  -webkit-animation: fadein 0.7s forwards;
  animation: fadein 0.7s forwards;
  animation-delay: 0s;
}

.logo img {
  width: 30px;
}

@keyframes fadein {
  0% {
    top: 0;
    opacity: 0;
  }
  100% {
    top: 0;
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
  height: 50px;
  line-height: 50px;
  right: 30px;
  display: block;
  -webkit-animation: fadeinsimple 0.4s forwards;
  animation: fadeinsimple 0.4s forwards;
  animation-delay: 0s;
  float: left;
}

.settings img {
  position: relative;
  width: 20px;
  top: 3px;
  height: auto;
  padding-left: 20px;
}

.settings .title {
  position: relative;
  /* top: -2px; */
  font-family: "Skycoin";
  font-size: 20px;
  color: whitesmoke;
  padding-right: 20px;
}
</style>
